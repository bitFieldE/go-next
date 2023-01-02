import React, { useEffect, useState } from 'react'
// import useSWR from 'swr'
// import type { NextPage } from 'next'
// import styles from '../styles/Home.module.css'

type Props = {
  message: string
}

// ページコンポーネントを定義する
const Home = (props: Props) => {
  return (
    <div>
      <h1>{props.message}</h1>
    </div>
  );
}

// サーバサイドで実行する処理(getServerSideProps)を定義する
export const getServerSideProps = async () => {
  // APIやDBからのデータ取得処理などを記載
  const res = await fetch(process.env.API_BASE_URL)
  const data = await res.json()
  
  const props = data

  return {
    props: props,
  };
}

export default Home
