import React from "react"
import MuiAppBar from "@mui/material/AppBar";
import {styled} from "@mui/material/styles";
import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";
import IconButton from "@mui/material/IconButton";
import MenuIcon from "@mui/icons-material/Menu";
import Toolbar from "@mui/material/Toolbar";

import Button from "@mui/material/Button";
import {useNavigate} from "react-router-dom"
import { useTheme } from '@mui/material/styles';
import Brightness4Icon from '@mui/icons-material/Brightness4';
import Brightness7Icon from '@mui/icons-material/Brightness7';
import {ColorModeContext} from "@/context"


const MyAppBar = styled(MuiAppBar, {
    shouldForwardProp: (prop) => prop !== "open"
})(({theme, open}) => ({
    zIndex: theme.zIndex.drawer + 1,
    transition: theme.transitions.create(['width', 'margin'], {
        easing: theme.transitions.easing.sharp,
        duration: theme.transitions.duration.leavingScreen,
    }),
}));

function AppBar(props){
    const theme = useTheme();
    const colorMode = React.useContext(ColorModeContext);

    const navigate = useNavigate();
    const {sideBarOpen, handleToggleSideBar} = props
    return(
        <MyAppBar position="fixed" open={sideBarOpen}>
            <Toolbar>
                <IconButton
                    color="inherit"
                    aria-label="menu"
                    onClick={handleToggleSideBar}
                    edge="start"
                    sx={{
                        marginRight: '36px',
                    }}
                >
                    <MenuIcon/>
                </IconButton>
                <Typography
                    variant="h6"
                    noWrap
                    component="div"
                    sx={{display: {xs: "none", sm: "block"}}}
                >
                    Go Admin
                </Typography>
                <Box sx={{flexGrow: 1}}/>
                <IconButton sx={{ ml: 1 }} onClick={colorMode.toggleColorMode} color="inherit">
                    {theme.palette.mode === 'dark' ? <Brightness7Icon /> : <Brightness4Icon />}
                </IconButton>
                <Button
                    color="inherit"
                    onClick={() => {
                        localStorage.setItem("login", "false");
                        navigate("/sign-in", {replace: true});
                    }}
                >
                    Logout
                </Button>
            </Toolbar>
        </MyAppBar>
    )
}


export default AppBar
