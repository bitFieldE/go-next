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
  const [location, setLocation] = useState<LocationResponse | ErrorResponse | null>(null);

  const getCurrentLocation = async () => {
    try {
      if (!navigator.geolocation) throw new Error(t('index.geolocation_error'));

      const position = await new Promise<GeolocationPosition>((resolve, reject) => {
        navigator.geolocation.getCurrentPosition(resolve, reject);
      });

      const res = await getAddress(position);
      setLocation(res);
    } catch (error) {
      setLocation({ error: t('index.geolocation_failure') });
    }
  }

  const getAddress = async (position: GeolocationPosition): Promise<LocationResponse | ErrorResponse> => {
    const res = await fetch(`${process.env.NEXT_PUBLIC_API_BASE_URL}/api/v1/location/current_location?lon=${position.coords.longitude}&lat=${position.coords.latitude}`)
    const data = await res.json()

    return data
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
      <div>{typeof location === 'LocationResponse' ? <p>location?.Feature[0]?.Property?.Address</p> : null}</div>

      <Button size='sm' onClick={getCurrentLocation} colorScheme="teal">
        クリックしてください
      </Button>
    </>
  )
}

export default Index
