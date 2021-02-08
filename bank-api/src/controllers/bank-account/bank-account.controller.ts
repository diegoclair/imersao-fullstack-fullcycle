import { Controller, Get, Param, ParseUUIDPipe } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { BankAccount } from 'src/models/bank-account.model';
import { Repository } from 'typeorm';

@Controller('bank-accounts')
export class BankAccountController {
    
    constructor(
        @InjectRepository(BankAccount)
        private bankAccountRepo: Repository<BankAccount>,
    ){}
    
    @Get()
    index(){
        return this.bankAccountRepo.find()
    }

    @Get(':bank_account_id')
    show(
        @Param('bank_account_id', new ParseUUIDPipe({version: '4'})
        ) bankAccountID: string
    ){
        return this.bankAccountRepo.findOneOrFail(bankAccountID)
    }
}
