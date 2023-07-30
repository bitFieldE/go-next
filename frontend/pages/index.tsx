// import React, { useEffect, useState } from 'react'
// import useSWR from 'swr'
// import type { NextPage } from 'next'
// import styles from '../styles/Home.module.css'

// type Props = {
//   message: string
// }

// // ページコンポーネントを定義する
// const Home = (props: Props) => {
//   return (
//     <div>
//       <h1>{props.message}</h1>
//       <p>メッセージ</p>
//     </div>
//   );
// }

// // サーバサイドで実行する処理(getServerSideProps)を定義する
// export const getServerSideProps = async () => {
//   // APIやDBからのデータ取得処理などを記載
//   const res = await fetch(process.env.API_BASE_URL + '/api/v1/')
//   const data = await res.json()
//   const props = data

//   return {
//     props: props,
//   };
// }

// export default Home

// import Header from '../components/organisms/Header'
// import { useTranslation } from "next-i18next"

// export default function Index(): JSX.Element {
//   const { t } = useTranslation('common')
//   return (
//     <>
//       <Header />
//       <h1>{t('title')}</h1>
//     </>
//   )
// }

import { useState } from 'react';
import { useTranslation } from 'next-i18next'
import { serverSideTranslations } from 'next-i18next/serverSideTranslations'
import { GetStaticProps } from 'next'
import { Button } from '@chakra-ui/react'

interface IndexProps {
  locale: string
}

const Index = () => {
  const { t } = useTranslation('index')
  const [location, setLocation] = useState<string | null>(null);
  
  const getCurrentLocation = () => {
    if (navigator.geolocation) {
      navigator.geolocation.getCurrentPosition(
        (position) => {
          const latitude = position.coords.latitude;
          const longitude = position.coords.longitude;
          alert(`緯度: ${latitude}, 経度: ${longitude}`);
        },
        (error) => {
          alert(error);
        }
      );
    } else {
      alert('Geolocationがサポートされていません');
    }
  };

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

      <Button size='sm' onClick={getCurrentLocation} colorScheme="teal">
        クリックしてください
      </Button>
    </>
  )
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

export default Index
