import {BeforeInsert, Column, CreateDateColumn, Entity, JoinColumn, ManyToOne, PrimaryGeneratedColumn} from "typeorm";
import {v4 as uuidv4} from 'uuid';
import { BankAccount } from "./bank-account.model";

export enum PixKeyType {
    cpf = 'cpf',
    email = 'email',
}

@Entity({name: 'pix'})
export class Pix {
    @PrimaryGeneratedColumn('uuid')
    id: string;

    @Column()
    keyType: PixKeyType;

    @Column()
    key: string;

    @ManyToOne(() => BankAccount)
    @JoinColumn({name: 'bank_account_id'})
    bankAccount: BankAccount;

    @Column()
    bank_account_id: string;

    @CreateDateColumn({type: 'timestamp'})
    created_at: Date;

    
    @BeforeInsert()
    generateID(){
        if(this.id) {
            return;
        }
        this.id = uuidv4();
    }
}
