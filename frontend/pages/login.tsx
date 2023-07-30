import { serverSideTranslations } from 'next-i18next/serverSideTranslations'
import { GetStaticProps } from 'next'

interface LoginProps {
  locale: string
}

const Login = () => {
  return (
    <>
      <h1>ログイン</h1>
    </>
  )
}

export const getStaticProps: GetStaticProps<LoginProps> = async ({ locale }) => {
  const translations = await serverSideTranslations(locale as string, 'common');

  return {
    props: {
      ...translations,
      locale: locale as string,
    },
  }
}

export default Login
