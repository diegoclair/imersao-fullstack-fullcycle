import { Body, Controller, Get, Param, ParseUUIDPipe, Post, ValidationPipe } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { TransactionDto } from 'src/dto/transaction.dto';
import { BankAccount } from 'src/models/bank-account.model';
import { Transaction, TransactionOperation } from 'src/models/transaction.model';
import { Repository } from 'typeorm';

@Controller('bank_accounts/:bank_account_id/transactions')
export class TransactionController {
    
    constructor(
        @InjectRepository(BankAccount)
        private bankAccountRepo: Repository<BankAccount>,
        @InjectRepository(Transaction)
        private transactionRepo: Repository<Transaction>,
    ){
        
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
            amount: body.amount * -1,           //because it is a debit operation.. a new transfer to some pix_key
            bank_account_id: bankAccountID,
            operation: TransactionOperation.debit
        });

        transaction = await this.transactionRepo.save(transaction);

    }
}
