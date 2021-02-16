import { GetServerSideProps } from 'next';
import React from 'react';
import Card from '../../../../components/Card';
import Layout from '../../../../components/Layout';
import PixCard from '../../../../components/PixCard';
import Title from '../../../../components/Title';
import { Pix, PixKeyType } from '../../../../model';
import { bankHttp } from '../../../../util/http';

interface PixRegisterProps {
    pixKeys: Pix[];
}

const PixRegister: React.FC<PixRegisterProps> = (props) => {
    
    const {pixKeys} = props

    return (
        <Layout>
            <div className="row">
                <div className="col-sm-6">
                    <Title>Cadastrar chave pix</Title>
                    <Card>
                        <form action="">

                        </form>
                    </Card>
                </div>
                <div className="col-sm-4 offset-md-2">
                    <Title>Minhas chaves pix</Title>
                    {pixKeys.map((p, key) => (
                        <PixCard key={key} pix={p}></PixCard>
                    ))}
                </div>
            </div>
        </Layout>
    );
};

export default PixRegister;

export const getServerSideProps: GetServerSideProps = async (ctx) => {
    const {query: {id}} = ctx;
    const {data: pixKeys} = await bankHttp.get(`bank-accounts/${id}/pix-keys`)

    return {
        props: {
            pixKeys,
        }
    }
}
