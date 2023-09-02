import { ChakraProvider } from '@chakra-ui/react'
import { AppProps } from 'next/app'
import { appWithTranslation } from 'next-i18next'
import Header from '../components/organisms/Header'

const MyApp = ({ Component, pageProps }: AppProps) => {
  return (
    <ChakraProvider>
      <Header />
      <Component {...pageProps} />
    </ChakraProvider>
  )
}

export default appWithTranslation(MyApp)
