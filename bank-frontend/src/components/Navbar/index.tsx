import Link from 'next/link';
import { BankAccount } from '../../model';
import classes from './Navbar.module.scss'

interface NavbarProps {
    bankAccount?: BankAccount
}
const Navbar: React.FC<NavbarProps> = (props) => {
    const { bankAccount } = props;

    return (
        <nav className={`navbar navbar-expand-lg ${classes.root} ${classes.bank001}`}>
            <div className={`container-fluid ${classes.navbarBody}`}>
                <Link href={`/bank-accounts`}>
                    <a className={`navbar-brand ${classes.navbarBrand}`} href="#">
                        <img src="/img/icon_banco.png" alt="" className={classes.logoBank} />
                        <div className={classes.bankName}>
                            <span>Cod - 001</span>
                            <h2>BBX</h2>
                        </div>
                    </a>
                </Link>
                {/* This block will only by rendered if we have received bankAccount */}
                {bankAccount && (
                    <div className={`collapse navbar-collapse ${classes.navbarRightRoot}`} id="navbarSupportedContent">
                        <ul className={`navbar-nav ml-auto ${classes.navbarRightBody}`}>
                            <li className={`nav-item ${classes.bankAccountInfo}`}>
                                <img src="/img/icon_user.png" alt="" className={`nav-item ${classes.iconUser}`} />
                                <p className={`nav-item ${classes.ownerName}`}>{bankAccount.owner_name} | C/C: {bankAccount.account_number}</p>
                            </li>
                        </ul>
                    </div>
                )}
            </div>
        </nav>
    );
};

export default Navbar;