import { NextPage, GetServerSideProps } from 'next'
import Link from 'next/link';
import { useRouter } from 'next/router';
import { useForm } from 'react-hook-form'
import Button from '../../../../../components/Button';
import Card from '../../../../../components/Card';
import FormButtonActions from '../../../../../components/FormButtonActions';
import Input from '../../../../../components/Input';
import Layout from '../../../../../components/Layout'
import Select from '../../../../../components/Select';
import Title from '../../../../../components/Title';
import { BankAccount } from '../../../../../model';
import { bankHttp } from '../../../../../util/http';
import Modal from '../../../../../util/modal';

interface Props {
    bankAccount: BankAccount
}

const TransactionRegister: NextPage<Props> = (props) => {
    
    const {bankAccount} = props;

    const {register, handleSubmit} = useForm();
    const {query: {id}, push} = useRouter();

    async function onSubmit(data) {

        try {

            await bankHttp.post(`/bank-accounts/${id}/transactions`, {
                ...data,
                amount: new Number(data.amount),
            });
            Modal.fire({
                title: `Transação realizada com sucesso!`,
                icon: 'success'
            });
            push(`/bank-accounts/${id}`);

        } catch (err) {
            console.error(err);
            Modal.fire({
                title: 'Ocorreu um erro. Verifique o console',
                icon: 'error'
            });
        };
        
    }

    return (
        <div>
            <Layout bankAccount={bankAccount}>
                <Title>Realizar transferência</Title>
                <Card>
                    <form onSubmit={handleSubmit(onSubmit)}>
                        <div className="row">
                            <div className="col-sm-4">
                                <Select labelText="Tipo" name="pix_key_type_to" ref={register}>
                                    <option value='cpf'>CPF</option>
                                    <option value='email'>E-mail</option>
                                </Select>
                            </div>
                            <div className="col-sm-8">
                                <Input name="pix_key_to" labelText='Chave' ref={register}/>
                            </div>
                        </div>
                        <Input
                            name='amount'
                            type='number'
                            step='.01'
                            labelText='Valor'
                            ref={register}
                            defaultValue='0.00'
                        />
                        <Input name='description' labelText='Descrição' ref={register}/>
                        <FormButtonActions>
                            <Link href={`/bank-accounts/${id}`}>
                                <Button variant='info' type='button'>Voltar</Button>
                            </Link>
                            <Button type='submit'>Cadastrar</Button>
                        </FormButtonActions>
                    </form>
                </Card>
            </Layout>
        </div>
    )
}

export const getServerSideProps: GetServerSideProps = async (ctx) => {
    const {
        query: { id },
      } = ctx;

      const { data: bankAccount } = await bankHttp.get(`bank-accounts/${id}`);

      return {
        props: {
          bankAccount,
        },
      };
}

export default TransactionRegister
