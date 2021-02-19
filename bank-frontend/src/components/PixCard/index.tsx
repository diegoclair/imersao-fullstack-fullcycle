import classes from './PixCard.module.scss';
import React, { useContext } from 'react';
import { Pix } from '../../model';
import BankContext from '../../context/BankContext';

interface PixCardProps {
    pix: Pix
}

const pixTypes = {
    cpf: 'CPF',
    email: 'E-mail'
}

const PixCard: React.FC<PixCardProps> = (props) => {
    const {pix} = props;
    
    const bank = useContext(BankContext);

    return (
        <div className={`${classes.root} ${classes[bank.cssCode]}`}>
            <p className={classes.type}>{pixTypes[pix.key_type]}</p>
            <span className={classes.key}>{pix.key}</span>
        </div>
    );
};

export default PixCard;