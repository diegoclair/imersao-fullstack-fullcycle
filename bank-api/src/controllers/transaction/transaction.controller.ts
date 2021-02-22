import { Body, Controller, Get, Inject, OnModuleDestroy, OnModuleInit, Param, ParseUUIDPipe, Post, ValidationPipe } from '@nestjs/common';
import { ClientKafka, MessagePattern, Payload } from '@nestjs/microservices';
import { Producer } from '@nestjs/microservices/external/kafka.interface';
import { InjectRepository } from '@nestjs/typeorm';
import { TransactionDto } from 'src/dto/transaction.dto';
import { BankAccount } from 'src/models/bank-account.model';
import { Pix, PixKeyType } from 'src/models/pix.model';
import { Transaction, TransactionOperation, TransactionStatus } from 'src/models/transaction.model';
import { Repository } from 'typeorm';

@Controller('bank-accounts/:bank_account_id/transactions')
export class TransactionController implements OnModuleInit, OnModuleDestroy{
    
    private kafkaProducer: Producer;

    constructor(
        @InjectRepository(BankAccount)
        private bankAccountRepo: Repository<BankAccount>,
        @InjectRepository(Transaction)
        private transactionRepo: Repository<Transaction>,
        @InjectRepository(Pix)
        private pixRepo: Repository<Pix>,

        @Inject('KAFKA_CODEPIX_TRANSACTION_SERVICE')
        private kafkaClient: ClientKafka,
    ){}

    async onModuleInit() {
        this.kafkaProducer = await this.kafkaClient.connect();
    }
    
    async onModuleDestroy() {
        await this.kafkaProducer.disconnect();
    }

    @Get('')
    index(
        @Param('bank_account_id', new ParseUUIDPipe({version: '4'})
        ) bankAccountID: string
    ){
        return this.transactionRepo.find({
            where: {
                bank_account_id: bankAccountID
            },
            order: {
                created_at: 'DESC'
            }
        })
    }
    
    @Post('')
    async store(
        @Param('bank_account_id', new ParseUUIDPipe({version: '4'})
        ) bankAccountID: string,
        @Body(new ValidationPipe({errorHttpStatusCode: 422})) //default is 400, for invalid body we use 422
        body: TransactionDto
    ){
        await this.bankAccountRepo.findOneOrFail(bankAccountID);

        let transaction = this.transactionRepo.create({
            ...body,
            amount: body.amount * -1,           //because it is a debit operation.. a new transfer to some pix key
            bank_account_id: bankAccountID,
            operation: TransactionOperation.debit
        });

        transaction = await this.transactionRepo.save(transaction);

        const sendDataToKafka = {
            id: transaction.external_id,
            accountID: bankAccountID,
            amount: body.amount,
            pixKeyTo: transaction.pix_key_to,
            pixKeyTypeTo: transaction.pix_key_type_to,
            description: transaction.description,
        }
        const topic = process.env.KAFKA_TRANSACTION_TOPIC

        await this.sendMessageToKafka(topic, sendDataToKafka)
        
        return transaction

    }

    async sendMessageToKafka(topic:string, message) {
        await this.kafkaProducer.send({
            topic: topic,
            messages: [
                {
                    key: topic,
                    value: JSON.stringify(message),
                }
            ]
        });
    }

    @MessagePattern(`bank${process.env.BANK_CODE}`) //MessagePattern means what topic we want to stay listening
    async doTransaction(@Payload() message){ //payload is to receive the message from kafka topic
        
        if(message.value.status === TransactionStatus.pending){
            await this.receivedTransaction(message.value);
        }

        if(message.value.status === TransactionStatus.confirmed){
            await this.confirmedTransaction(message.value);
        }
    }

    async receivedTransaction(data){
        const pix = await this.pixRepo.findOneOrFail({
            where: {
                key: data.pixKeyTo,
                key_type: data.pixKeyTypeTo,

            }
        });

        const transaction = this.transactionRepo.create({
            external_id: data.id,
            amount: data.amount,
            description: data.description,
            bank_account_id: pix.bank_account_id,
            bank_account_from_id: data.accountID,
            operation: TransactionOperation.credit,
            status: TransactionStatus.completed,
        });

        this.transactionRepo.save(transaction)

        const sendConfirmationDataToKafka = {
            ...data,
            status: TransactionStatus.confirmed,
        }
        const topic = process.env.KAFKA_TRANSACTION_CONFIRMATION_TOPIC
        this.sendMessageToKafka(topic, sendConfirmationDataToKafka)
    }

    async confirmedTransaction(data) {
        const transaction = await this.transactionRepo.findOneOrFail({
            where: {
                external_id: data.id,
            }
        });

        this.transactionRepo.update(
            {id: data.id},
            {
                status: TransactionStatus.completed,
            },
        );

        const sendDataToKafka = {
            id: data.id,
            accountID: transaction.bank_account_id,
            amount: Math.abs(transaction.amount),
            pixKeyTo: transaction.pix_key_to,
            pixKeyTypeTo: transaction.pix_key_type_to,
            description: transaction.description,
            status: TransactionStatus.completed,
        }
        const topic = process.env.KAFKA_TRANSACTION_CONFIRMATION_TOPIC

        this.sendMessageToKafka(topic, sendDataToKafka)
    }
}
