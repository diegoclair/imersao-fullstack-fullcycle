import { Controller, Get, Param, ParseUUIDPipe, Post } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { Pix } from 'src/models/pix.model';
import { Repository } from 'typeorm';

@Controller('bank-accounts/:bank_account_id/pix-keys')
export class PixController {

    constructor(
        @InjectRepository(Pix)
        private pixRepo: Repository<Pix>
    ){
        
    }

    @Get()
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
    
    @Post()
    store(){}

    @Get('exists')
    exists(){}
}
