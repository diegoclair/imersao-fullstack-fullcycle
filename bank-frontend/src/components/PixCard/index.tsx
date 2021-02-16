import classes from './PixCard.module.scss';
import React from 'react';
import { Pix } from '../../model';

interface PixCardProps {
    pix: Pix
}

const pixTypes = {
    cpf: 'CPF',
    email: 'E-mail'
}

const PixCard: React.FC<PixCardProps> = (props) => {
    const {pix} = props
    return (
        <div className={`${classes.root} ${classes.bank001}`}>
            <p className={classes.type}>{pixTypes[pix.key_type]}</p>
            <span className={classes.key}>{pix.key}</span>
        </div>
    );
};

export default PixCard;