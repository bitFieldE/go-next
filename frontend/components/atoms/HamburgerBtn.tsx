import { HamburgerIcon } from '@chakra-ui/icons';
import Link from "next/link"

const HamburgerBtn = () => {
  return (
    <Link href='https://chakra-ui.com'>
      <HamburgerIcon boxSize={8} color="white" />
    </Link>
  );
};

export default HamburgerBtn