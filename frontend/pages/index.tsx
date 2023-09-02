import { useState } from 'react';
import { useTranslation } from 'next-i18next'
import { serverSideTranslations } from 'next-i18next/serverSideTranslations'
import { GetStaticProps } from 'next'
import { Button } from '@chakra-ui/react'

interface IndexProps {
  locale: string
}

export const getStaticProps: GetStaticProps<IndexProps> = async ({ locale }) => {
  const translations = await serverSideTranslations(locale as string, ['common', 'index']);

  return {
    props: {
      ...translations,
      locale: locale as string,
    },
  }
}

const Index = () => {
  const { t } = useTranslation('index')
  const [location, setLocation] = useState<LocationResponse | null>(null);
  const [errorMsg, setErrorMsg] = useState<ErrorResponse | null>(null);

  const getCurrentLocation = async () => {
    try {
      if (!navigator.geolocation) throw new Error;

      const position = await new Promise<GeolocationPosition>((resolve, reject) => {
        navigator.geolocation.getCurrentPosition(resolve, reject);
      });
      await getAddress(position);
    } catch {
      setErrorMsg({ error: t('index.geolocation_no_support_error') });
    }
  }

  const getAddress = async (position: GeolocationPosition) => {
    try {
      const res = await fetch(`${process.env.NEXT_PUBLIC_API_BASE_URL}/api/v1/locations/current_location?lon=${position.coords.longitude}&lat=${position.coords.latitude}`);
      const data = await res.json();
      if (!res.ok) throw new Error;
      setLocation(data);
    } catch {
      setErrorMsg({ error: t('index.geolocation_failure') });
    }
  }

  return (
    <>
      <h1>{t('index.title')}</h1>
      <p>{t('index.description')}</p>
      <h1>{t('index.title')}</h1>
      <p>{t('index.description')}</p>
      <h1>{t('index.title')}</h1>
      <p>{t('index.description')}</p>
      <h1>{t('index.title')}</h1>
      <p>{t('index.description')}</p>
      <h1>{t('index.title')}</h1>
      <p>{t('index.description')}</p>
      <h1>{t('index.title')}</h1>
      <p>{t('index.description')}</p>
      <h1>{t('index.title')}</h1>
      <p>{t('index.description')}</p>
      <h1>{t('index.title')}</h1>
      <p>{t('index.description')}</p>
      <h1>{t('index.title')}</h1>
      <p>{t('index.description')}</p>
      <h1>{t('index.title')}</h1>
      <p>{t('index.description')}</p>
      <h1>{t('index.title')}</h1>
      <p>{t('index.description')}</p>
      <h1>{t('index.title')}</h1>

      <Button size='sm' onClick={getCurrentLocation} colorScheme="teal">
        クリックしてください
      </Button>
      <h1></h1>
      <p>
        {location ? location.Feature[0].Property.Address : ''}
        {errorMsg ? errorMsg.error : ''}
      </p>
    </>
  )
}

export default Index
