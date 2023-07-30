/** @type {import('next').NextConfig} */

/**
 * @param {string} appEnv
 */

const { i18n } = require('./next-i18next.config')

module.exports = {
  i18n,
  reactStrictMode: true,
  swcMinify: true,
}

const loadEnv = (appEnv = 'local') => {
  const env = {
    ...require(`./env/env.${appEnv}`),
    NEXT_PUBLIC_APP_ENV: appEnv,
  }

  Object.entries(env).forEach(([key, value]) => {
    process.env[key] = value
  })
}

loadEnv(process.env.APP_ENV)
