
import { Box, Flex, Link as Text } from "@chakra-ui/react"
import HamburgerBtn from '../atoms/HamburgerBtn'
import { useTranslation } from 'next-i18next'
import { useEffect, useState } from "react"

const Header = () => {
  const { t } = useTranslation('common')
  const [headerColor, setHeaderColor] = useState('green.300')

  useEffect(() => {
    const handleScroll = () => {
      const scrollPosition = window.pageYOffset

      if (scrollPosition > 150) {
        setHeaderColor('white')
      } else {
        setHeaderColor('green.300')
      }
    }

    window.addEventListener('scroll', handleScroll)

    return () => {
      window.removeEventListener('scroll', handleScroll)
    }
  }, [])

  return (
    <Box
      bg={headerColor}
      py={5}
      boxShadow="md"
      position="fixed"
      top={0}
      left={0}
      right={0}
      zIndex="sticky"
    >
      <Flex mx="auto" px={4} align="center" justify="space-between">
        <Text fontSize="xl" fontWeight="bold">
          {t('header.title')}
        </Text>
        <HamburgerBtn />
      </Flex>
    </Box>
  )
}

export default Header
