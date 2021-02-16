import * as React from 'react';
import Link, { LinkProps } from 'next/link';
import classes from "./BankAccountDashboard.module.scss";
import { BankAccountBalance } from '../../../components/BankAccountBalance';
import Layout from '../../../components/Layout';
import Title from '../../../components/Title';
import { BankAccount } from '../../../model';
import { GetServerSideProps, NextPage } from 'next';
import { bankHttp } from '../../../util/http';

interface ActionLinkProps extends LinkProps {

}

const ActionLink: React.FunctionComponent<ActionLinkProps> = (props) => {
  const { children, ...rest } = props;

  return (
    <Link {...rest}>
      <a className={`${classes.actionLink} ${classes.bank001}`}>{children}</a>
    </Link>
  );
};

interface HeaderProps {
  bankAccount: BankAccount;
}

const Header: React.FC<HeaderProps> = (props) => {

    const {bankAccount} = props;

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

interface BankAccountDashboardProps {
    bankAccount: BankAccount;
}

const BankAccountDashboard: NextPage<BankAccountDashboardProps> = (props) => {

    const {bankAccount} = props;

    return (
        <Layout>
            <Header bankAccount={bankAccount}/>
            <Title>Dashboard da conta</Title>
        </Layout>
    );
};

export default BankAccountDashboard;

export const getServerSideProps: GetServerSideProps = async (ctx) => {

    const {query: {id}} = ctx;    
    const {data: bankAccount} = await bankHttp.get(`/bank-accounts/${id}`);
    
    return {
        props: {
            bankAccount,
        },
    };
};