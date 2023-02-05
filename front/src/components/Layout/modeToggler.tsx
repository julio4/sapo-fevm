import { MoonIcon, SunIcon } from '@chakra-ui/icons';
import { Button, useColorMode, useColorModeValue } from '@chakra-ui/react';

export default function ModeToggler() {
    const { colorMode, toggleColorMode } = useColorMode();

    return (
    <Button 
        bg={useColorModeValue('green.50', 'whiteAlpha.200')}
        _hover={{
        bg: useColorModeValue('whiteAlpha.600', 'whiteAlpha.300'),
        }}
        onClick={toggleColorMode}>
        {colorMode === 'light' ? <MoonIcon /> : <SunIcon />}
    </Button>
    )
}