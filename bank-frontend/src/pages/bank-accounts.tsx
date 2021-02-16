import React from 'react';
import BankAccountCard from '../components/BankAccountCard';
import Layout from '../components/Layout';
import Title from '../components/Title';

type Props = {

}

const BankAccountsList = (props: Props) => {
    return (
        <Layout>
            <Title>Contas bancárias</Title>
            <div className="row">
                <a className="col-12 col-sm-6 col-md4">
                    <BankAccountCard bankAccount={{id: 'teste1', owner_name: 'teste1', balance: 0, account_number: '11111'}}/>
                </a>
                <a className="col-12 col-sm-6 col-md4">
                    <BankAccountCard bankAccount={{id: 'teste2', owner_name: 'teste2', balance: 0, account_number: '22222'}}/>
                </a>
                <a className="col-12 col-sm-6 col-md4">
                    <BankAccountCard bankAccount={{id: 'teste3', owner_name: 'teste3', balance: 0, account_number: '33333'}}/>
                </a>
            </div>
        </Layout>
    );
};

export default BankAccountsList;