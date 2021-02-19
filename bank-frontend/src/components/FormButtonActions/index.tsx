import classes from "./FormButtonActions.module.scss";

const FormButtonActions: React.FC = (props) => {
  return <div className={classes.root}>{props.children}</div>;
};

export default FormButtonActions;