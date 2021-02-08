import { Body, Controller, Get, HttpCode, Inject, InternalServerErrorException, NotFoundException, Param, ParseUUIDPipe, Post, Query, UnprocessableEntityException, ValidationPipe } from '@nestjs/common';
import { ClientGrpc } from '@nestjs/microservices';
import { InjectRepository } from '@nestjs/typeorm';
import { PixExistsDto } from 'src/dto/pix-exists.dto';
import { PixDto } from 'src/dto/pix.dto';
import { PixService } from 'src/grpc-types/pix-service.grpc';
import { BankAccount } from 'src/models/bank-account.model';
import { Pix } from 'src/models/pix.model';
import { Repository } from 'typeorm';

@Controller('/bank_accounts/:bank_account_id/pix-keys')
export class PixController {

    constructor(
        @InjectRepository(Pix)
        private pixRepo: Repository<Pix>,
        @InjectRepository(BankAccount)
        private bankAccountRepo: Repository<BankAccount>,
        @Inject('GRPC_CODEPIX_PACKAGE') //the same we are using on app.module
        private client: ClientGrpc
    ){
        
    }

    @Get('')
    index(
        @Param('bank_account_id', new ParseUUIDPipe({version: '4'})
        ) bankAccountID: string
    ){
        return this.pixRepo.find({
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
        @Body(new ValidationPipe({errorHttpStatusCode: 422})) //default is 400, for invalid boddy we use 422
        body: PixDto
    ){
        await this.bankAccountRepo.findOneOrFail(bankAccountID);

        const pixService: PixService = this.client.getService('PixService'); //this pix service is the service on file pix.proto
        const pixExists = await this.checkPixExists(body);
        if(pixExists){
            throw new UnprocessableEntityException("Pix key already exists");
        }

        const createdPixKey = await pixService.addPixKey({
            key: body.key,
            keyType: body.key_type,
            accountID: bankAccountID
        }).toPromise();

        if (createdPixKey.error) {
            throw new InternalServerErrorException(createdPixKey.error)
        }

        const pix = this.pixRepo.create({
            id: createdPixKey.id,
            bank_account_id: bankAccountID,
            ...body
        })

        return await this.pixRepo.save(pix)
    }

    async checkPixExists(params: {key: string, key_type: string}) {
        const pixService: PixService = this.client.getService('PixService');
        try {
            await pixService.findPixKeyByID({key: params.key, keyType: params.key_type}).toPromise()
            return true;
        } catch (error) {
            if(error.details === "no key was found") {
                return false;
            }
            throw new InternalServerErrorException("Server not available");
        }
    }

    @Get('/exists')
    @HttpCode(204)
    async exists(
        @Query(new ValidationPipe({errorHttpStatusCode: 422})) //default is 400, for invalid body we use 422
        params: PixExistsDto
    ){

        const pixService: PixService = this.client.getService('PixService');
        try {
            await pixService.findPixKeyByID({key: params.key, keyType: params.key_type}).toPromise()
        } catch (e) {
            if(e.details === "no key was found") {
                throw new NotFoundException(e.details)
            }
            throw new InternalServerErrorException("Server not available");
        }
    
    }
}
