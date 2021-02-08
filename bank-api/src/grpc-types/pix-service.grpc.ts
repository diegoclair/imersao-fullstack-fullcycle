import { Observable } from "rxjs";
//this contract is similar that we have in pix.proto file
interface Account {
    accountID: string;
    accountNumber: string;
    bankID: string;
    bankName: string;
    ownerName: string;
    createdAt: string;
}

export interface FindPixKeyByIDResponse {
    id: string;
    keyType: string;
    key: string;
    account: Account;
    createdAt: string;
}

export interface PixService {
    addPixKey: (data: {
        keyType: string;
        key: string;
        accountID: string;
    }) => Observable<{ id: string; status: string; error: string }>;
    findPixKeyByID: (data: { keyType: string; key: string }) => Observable<FindPixKeyByIDResponse>;
}