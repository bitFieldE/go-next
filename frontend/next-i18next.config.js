/** @type {import('next-i18next').UserConfig} */

const path = require('path')

module.exports = {
  reloadOnPrerender: true,
  i18n: {
    defaultLocale: 'ja',
    locales: ['en', 'ja'],
  },
  localePath: path.resolve('./locales')
}
