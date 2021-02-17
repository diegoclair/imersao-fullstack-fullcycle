import React from 'react';

interface ButtonProps extends React.DetailedHTMLProps<React.ButtonHTMLAttributes<HTMLButtonElement>, HTMLButtonElement> {
    variant?: 'primary' | 'info';
}

const buttonClasses = {
    primary: 'btn-primary',
    info: 'btn-info',
}

const Button = React.forwardRef<any, ButtonProps>((props, ref) => {
    const {variant = 'primary', ...rest} = props;

    const className = [
        'btn', 
        buttonClasses[variant], 
        'bank001',
        props.className
    ].join(' ').trim();

    return (
        <button className={className} {...rest} ref={ref}></button>
    );
});

export default Button;