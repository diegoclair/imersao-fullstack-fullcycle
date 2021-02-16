import React from 'react';
import classes from './Card.module.scss';
import {GetServerSideProps} from 'next'


interface CardProps {
    
}

const Card: React.FC<CardProps> = (props) => {
    return (
        <div className={classes.root}>
            {props.children}
        </div>
    );
};

export default Card;

export const getServerSideProps: GetServerSideProps = async (ctx) => {
    return {
        props: {}
    }
}
