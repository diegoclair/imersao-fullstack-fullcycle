import '../../styles/sass/globals.scss'
import BankContext from '../context/BankContext'

function MyApp({ Component, pageProps }) {
  return (
    <BankContext.Provider value={{
      name: process.env.NEXT_PUBLIC_BANK_NAME,
      code: process.env.NEXT_PUBLIC_BANK_CODE,
      // cssCode: `bank${process.env.NEXT_PUBLIC_BANK_CODE}`,   //for cssCode is this code or the code bellow
      get cssCode() {
        return `bank${this.code}`;
      }
    }}>
      <Component {...pageProps} />
    </BankContext.Provider>
  )
}

export default MyApp
