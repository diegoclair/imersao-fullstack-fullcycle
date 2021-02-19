import * as React from 'react';
import Link, { LinkProps } from 'next/link';
import classes from "./BankAccountDashboard.module.scss";
import { BankAccountBalance } from '../../../components/BankAccountBalance';
import Layout from '../../../components/Layout';
import { BankAccount, Transaction } from '../../../model';
import { GetServerSideProps, NextPage } from 'next';
import { bankHttp } from '../../../util/http';
import format from "date-fns/format";
import parseISO from "date-fns/parseISO";

//This file is an example that we can create components and the page in the same file

//==========================================
// ActionLinkProps component 
//==========================================
interface ActionLinkProps extends LinkProps { }

const ActionLink: React.FunctionComponent<ActionLinkProps> = (props) => {
    const { children, ...rest } = props;

    return (
        <Link {...rest}>
            <a className={`${classes.actionLink} ${classes.bank001}`}>{children}</a>
        </Link>
    );
};

//==========================================
// Header component 
//==========================================
interface HeaderProps {
    bankAccount: BankAccount;
}

const Header: React.FC<HeaderProps> = (props) => {

    const { bankAccount } = props;

    return (
        <div className={`container ${classes.header}`}>
            <BankAccountBalance balance={bankAccount.balance} />
            <div className={classes.buttonActions}>
                <ActionLink href={`/bank-accounts/${bankAccount.id}/pix/transactions/register`}>
                    Realizar transferência
          </ActionLink>
                <ActionLink href={`/bank-accounts/${bankAccount.id}/pix/register`}>
                    Cadastrar chave pix
          </ActionLink>
            </div>
        </div>
    );
};

//==========================================
// BankAccountDashboard page 
//==========================================
interface BankAccountDashboardProps {
    bankAccount: BankAccount;
    transactions: Transaction[];
}

const BankAccountDashboard: NextPage<BankAccountDashboardProps> = (props) => {

    const { bankAccount, transactions } = props;

    return (
        <Layout bankAccount={bankAccount}>
            <Header bankAccount={bankAccount} />
            <div>
                <h1 className={classes.titleTable}>Últimos lançamentos</h1>
                <table className={`table table-borderless table-striped ${classes.tableTransactions}`}>
                    <thead>
                        <tr>
                            <th scope="col">Data</th>
                            <th scope="col" colSpan={3}>Descrição</th>
                            <th scope="col">Valor (R$)</th>
                        </tr>
                    </thead>
                    <tbody>
                        {transactions.map((t, key) => (
                            <tr key={key}>
                                <td>{format(parseISO(t.created_at), 'dd/MM')}</td>
                                <td colSpan={3}>{t.description}</td>
                                <td className='text-right'>{t.amount.toLocaleString('pt-BR')}</td>
                            </tr>
                        ))}
                    </tbody>
                </table>
            </div>
        </Layout>
    );
};

export default BankAccountDashboard;

export const getServerSideProps: GetServerSideProps = async (ctx) => {

    const { query: { id } } = ctx;
    const [{ data: bankAccount }, { data: transactions }] = await Promise.all([
        await bankHttp.get(`/bank-accounts/${id}`),
        await bankHttp.get(`/bank-accounts/${id}/transactions`),
    ]) 

    return {
        props: {
            bankAccount,
            transactions,
        },
    };
};