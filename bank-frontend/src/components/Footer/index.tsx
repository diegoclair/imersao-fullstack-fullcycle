import classes from './Footer.module.scss';

const Footer: React.FC = (props) => {
    return (
        <footer className={classes.root}>
            <img src="/img/logo_pix.png" alt="Code Pix"/>
        </footer>
    );
};

export default Footer;