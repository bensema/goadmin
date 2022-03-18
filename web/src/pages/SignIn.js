import * as React from "react";
import Avatar from "@mui/material/Avatar";
import Button from "@mui/material/Button";
import CssBaseline from "@mui/material/CssBaseline";
import TextField from "@mui/material/TextField";
import Link from "@mui/material/Link";
import Box from "@mui/material/Box";
import LockOutlinedIcon from "@mui/icons-material/LockOutlined";
import Typography from "@mui/material/Typography";
import Container from "@mui/material/Container";
import {createTheme, ThemeProvider} from "@mui/material/styles";
import {useNavigate} from "react-router-dom";

import {login} from "@/api/api";

function Copyright(props) {
    return (
        <Typography
            variant="body2"
            color="text.secondary"
            align="center"
            {...props}
        >
            {"Copyright © "}
            <Link color="inherit" href="https://mui.com/">
                Your Website
            </Link>{" "}
            {new Date().getFullYear()}
            {"."}
        </Typography>
    );
}

const theme = createTheme();

export default function SignIn() {
    const navigate = useNavigate();

    const [signInErr, setSignInErr] = React.useState(false);
    const [signInErrMsg, setSignInErrMsg] = React.useState("");
    const [usernameHelperText, setUsernameHelperText] = React.useState("");
    const [passwordHelperText, setPasswordHelperText] = React.useState("");


    const handleSubmit = (event) => {
        event.preventDefault();
        const data = new FormData(event.currentTarget);

        let username = data.get("username")
        let password = data.get("password")

        if (username === ""){
            setUsernameHelperText("Username required")
            return
        }

        if (password === ""){
            setPasswordHelperText("Password required")
            return
        }

        login({
            username: data.get("username"),
            password: data.get("password")
        }).then((res) => {
            if (res.data.code !== 0) {
                setSignInErr(true)
                setSignInErrMsg(res.data.message)
            } else {
                localStorage.setItem("login", "true");
                navigate("/", {replace: true});
            }

        });

    };

    return (
        <ThemeProvider theme={theme}>
            <Container component="main" maxWidth="xs">
                <CssBaseline/>
                <Box
                    sx={{
                        marginTop: 8,
                        display: "flex",
                        flexDirection: "column",
                        alignItems: "center"
                    }}
                >
                    <Avatar sx={{m: 1, bgcolor: "secondary.main"}}>
                        <LockOutlinedIcon/>
                    </Avatar>
                    <Typography component="h1" variant="h5">
                        Sign in
                    </Typography>
                    <Box
                        component="form"
                        onSubmit={handleSubmit}
                        noValidate
                        sx={{mt: 1}}
                    >
                        <TextField
                            error={signInErr || usernameHelperText!==""}
                            margin="normal"
                            required
                            fullWidth
                            id="username"
                            label="Username"
                            name="username"
                            autoComplete="username"
                            autoFocus
                            helperText={usernameHelperText}
                            onInput={() => {
                                setSignInErr(false)
                                setUsernameHelperText("")
                            }}
                        />
                        <TextField
                            error={signInErr || passwordHelperText!==""}
                            margin="normal"
                            required
                            fullWidth
                            name="password"
                            label="Password"
                            type="password"
                            id="password"
                            autoComplete="current-password"
                            helperText={(signInErr && signInErrMsg)|| passwordHelperText}
                            onInput={() => {
                                setSignInErr(false)
                                setPasswordHelperText("")
                            }}
                        />
                        <Button
                            type="submit"
                            fullWidth
                            variant="contained"
                            sx={{mt: 3, mb: 2}}
                        >
                            Sign In
                        </Button>
                    </Box>
                </Box>
                <Copyright sx={{mt: 8, mb: 4}}/>
            </Container>
        </ThemeProvider>
    );
}
