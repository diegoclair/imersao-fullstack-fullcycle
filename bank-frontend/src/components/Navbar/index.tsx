import Link from 'next/link';
import classes from './Navbar.module.scss'

const Navbar: React.FC = (props) => {
    return (
        <nav className={`navbar navbar-expand-lg ${classes.root} ${classes.bank001}`}>
            <div className={`container-fluid ${classes.navbarBody}`}>
                <Link href={`/bank-accounts`}>
                    <a className={`navbar-brand ${classes.navbarBrand}`} href="#">
                        <img src="/img/icon_banco.png" alt="" className={classes.logoBank}/>
                        <div className={classes.bankName}>
                            <span>Cod - 001</span>
                            <h2>BBX</h2>
                        </div>
                    </a>
                    </Link>
                <div className={`collapse navbar-collapse ${classes.navbarRightRoot}`} id="navbarSupportedContent">
                    <ul className={`navbar-nav ml-auto ${classes.navbarRightBody}`}>
                        <li className={`nav-item ${classes.bankAccountInfo}`}>
                            <img src="/img/icon_user.png" alt="" className={`nav-item ${classes.iconUser}`}/>
                            <p className={`nav-item ${classes.ownerName}`}>Proprietário | C/C: Número da conta</p>
                        </li>
                    </ul>
                </div>
            </div>
        </nav>
    );
};

export default Navbar;