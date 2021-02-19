import classes from './Title.module.scss'

const Title: React.FC = (props) => {
    return (
        <h1 className={classes.root}>
            {props.children}
        </h1>
    );
};

export default Title;