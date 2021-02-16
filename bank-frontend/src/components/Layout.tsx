import React from 'react';
import Footer from './Footer';
import MainContent from './MainContent';
import Navbar from './Navbar';


const Layout: React.FC = (props) => {
    return (
        <>
            <Navbar/>
            <MainContent>
                {props.children}
            </MainContent>
            <Footer/>
        </>
    );
};

export default Layout;