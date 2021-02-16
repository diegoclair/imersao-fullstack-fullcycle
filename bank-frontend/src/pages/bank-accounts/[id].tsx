import React from 'react';
import Layout from '../../components/Layout';
import Title from '../../components/Title';

interface BankAccountDashboardProps {
    
}

const BankAccountDashboard: React.FC<BankAccountDashboardProps> = () => {
    return (
        <Layout>
            <Title>Dashboard da conta</Title>
        </Layout>
    );
};

export default BankAccountDashboard;