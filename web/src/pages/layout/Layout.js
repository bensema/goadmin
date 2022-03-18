import React from "react";
import Box from "@mui/material/Box";
import CssBaseline from "@mui/material/CssBaseline";
import {Outlet} from "react-router-dom"
import Toolbar from "@mui/material/Toolbar";
import {SnackbarProvider} from 'notistack';

import SideBar from "./SideBar"
import AppBar from "./AppBar";

function Layout() {
    const [sideBarOpen, setSideBarOpen] = React.useState(true);

    const handleToggleSideBar = () => {
        setSideBarOpen(!sideBarOpen);
    };

    return (
        <SnackbarProvider maxSnack={3}>
            <Box sx={{display: "flex"}}>
                <CssBaseline/>
                <AppBar sideBarOpen={sideBarOpen} handleToggleSideBar={handleToggleSideBar}/>
                <SideBar sideBarOpen={sideBarOpen} setSideBarOpen={setSideBarOpen}/>

                <Box
                    component="main"
                    sx={{
                        backgroundColor: (theme) =>
                            theme.palette.mode === 'light'
                                ? theme.palette.grey[100]
                                : theme.palette.grey[900],
                        flexGrow: 1,
                        height: '100vh',
                        overflow: 'auto',
                    }}
                >
                    <Toolbar/>
                    <Outlet/>
                </Box>
            </Box>
        </SnackbarProvider>
    )

}

export default Layout
