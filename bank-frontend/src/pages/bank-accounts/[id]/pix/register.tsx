import classes from './Register.module.scss';
import { GetServerSideProps } from 'next';
import React from 'react';
import Card from '../../../../components/Card';
import Input from '../../../../components/Input';
import Layout from '../../../../components/Layout';
import PixCard from '../../../../components/PixCard';
import Title from '../../../../components/Title';
import { Pix } from '../../../../model';
import { bankHttp } from '../../../../util/http';
import FormButtonActions from '../../../../components/FormButtonActions';
import Button from '../../../../components/Button';
import Link from 'next/link';
import {useRouter} from 'next/router'
import {useForm} from 'react-hook-form'
import Modal from '../../../../util/modal';

interface PixRegisterProps {
    pixKeys: Pix[];
}

const PixRegister: React.FC<PixRegisterProps> = (props) => {
    
    const {pixKeys} = props
    const {query: {id}, push} = useRouter();

    const {register, handleSubmit} = useForm();

    async function onSubmit(data) {
        try {
            await bankHttp.post(`/bank-accounts/${id}/pix-keys`, data)
            Modal.fire({
                title: `Chave ${data.key} cadastrada com sucesso`,
                icon: 'success'
            });
            push(`/bank-accounts/${id}`)
        } catch (err) {
            console.error(err);
            Modal.fire({
                title: 'Ocorreu um erro. Verifique o console',
                icon: 'error'
            });
        };
        
    }

    return (
        <Layout>
            <div className="row">
                <div className="col-sm-6">
                    <Title>Cadastrar chave pix</Title>
                    <Card>
                        <form onSubmit={handleSubmit(onSubmit)}>
                            <p className={classes.keyTypeInfo}>Escolha um tipo de chave</p>
                            <Input 
                                name="key_type" 
                                type="radio" 
                                labelText="CPF" 
                                value="cpf"
                                ref={register}
                            />
                            <Input 
                                name="key_type" 
                                type="radio" 
                                labelText="E-mail" 
                                value="email"
                                ref={register}
                            />
                            <Input 
                                name="key" 
                                labelText="Digite a chave"
                                ref={register}
                            />
                            <FormButtonActions>
                                <Link href={`/bank-accounts/${id}`}>
                                    <Button type="button" variant='info'>Voltar</Button>
                                </Link>
                                <Button type="submit">Cadastrar</Button>
                            </FormButtonActions>
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
