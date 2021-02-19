import React from 'react';
import { BankAccount } from '../model';
import Footer from './Footer';
import MainContent from './MainContent';
import Navbar from './Navbar';

interface LayoutProps{
    bankAccount?: BankAccount;
}

const Layout: React.FC<LayoutProps> = (props) => {
    const {bankAccount} = props;

    return (
        <>
            <Navbar bankAccount={bankAccount}/>
            <MainContent>
                {props.children}
            </MainContent>
            <Footer/>
        </>
    );
};

export default Layout;