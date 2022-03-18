import React from 'react';
import {createTheme, ThemeProvider} from '@mui/material/styles';
import routes from "./routes/routes";
import {useRoutes} from "react-router";
import {ColorModeContext} from "@/context"

export default function App() {
    const [mode, setMode] = React.useState('light');
    const colorMode = React.useMemo(
        () => ({
            toggleColorMode: () => {
                setMode((prevMode) => (prevMode === 'light' ? 'dark' : 'light'));
            },
        }),
        [],
    );

    const theme = React.useMemo(
        () =>
            createTheme({
                palette: {
                    mode,
                },
            }),
        [mode],
    );

    return (
        <ColorModeContext.Provider value={colorMode}>
            <ThemeProvider theme={theme}>
                {useRoutes(routes)}
            </ThemeProvider>
        </ColorModeContext.Provider>
    );
}
