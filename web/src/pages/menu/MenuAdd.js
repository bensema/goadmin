import React, {useState} from "react"
import TextField from '@mui/material/TextField';
import Dialog from '@mui/material/Dialog';
import Paper from '@mui/material/Paper';
import AppBar from '@mui/material/AppBar';
import Toolbar from '@mui/material/Toolbar';
import IconButton from '@mui/material/IconButton';
import Typography from '@mui/material/Typography';
import CloseIcon from '@mui/icons-material/Close';
import Box from "@mui/material/Box";
import Container from '@mui/material/Container';
import Button from '@mui/material/Button';
import Grid from '@mui/material/Grid';
import LoadingButton from '@mui/lab/LoadingButton';
import Slide from '@mui/material/Slide';
import {useSnackbar} from 'notistack';

import {menuAdd} from "@/api/api"


const Transition = React.forwardRef(function Transition(props, ref) {
    return <Slide direction="up" ref={ref} {...props} />;
});

export const MenuAdd = (props) => {
    const {open, handleClose} = props
    const [loading, setLoading] = useState(false)

    const [nameHelperText, setNameHelperText] = React.useState("");

    const [menu, setMenu] = useState({
        name: "",
        pid: "",
        icon: "",
        url: "",
        index_sort: "",
    })

    const {enqueueSnackbar} = useSnackbar();

    const handleKV = (k, e) => {
        setMenu((prev) => ({...prev, [k]: e.target.value}));
    }

    const handleSubmit = () => {

        if (menu.name === "") {
            setNameHelperText("Name required")
            return
        }

        (async () => {
            setLoading(true)
            await menuAdd(menu).then((res) => {
                setLoading(false)
                if (res.data.code === 0) {
                    enqueueSnackbar('Add Menu success!', {variant: 'success'});
                    handleClose()
                } else {
                    enqueueSnackbar('Add Menu fail!', {variant: 'error'});
                    handleClose()
                }
            })
        })();
    }

    return (
        <Dialog
            fullScreen
            open={open}
            onClose={handleClose}
            TransitionComponent={Transition}
        >
            <AppBar sx={{position: 'relative'}}>
                <Toolbar>
                    <IconButton
                        edge="start"
                        color="inherit"
                        onClick={handleClose}
                        aria-label="close"
                    >
                        <CloseIcon/>
                    </IconButton>
                    <Typography sx={{ml: 2, flex: 1}} variant="h6" component="div">
                        Add Menu
                    </Typography>

                </Toolbar>
            </AppBar>
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
                <Container fixed maxWidth="lg" sx={{mt: 4, mb: 4}}>
                    <Paper sx={{p: 2, display: 'flex', flexDirection: 'column', width: '100%'}} elevation={2}>
                        <Grid
                            container
                            direction="column"
                            justifyContent="center"
                            alignItems="stretch"
                            spacing={2}
                        >

                            <Grid item xs={12}>
                                <TextField
                                    fullWidth
                                    error={nameHelperText !== ""}
                                    helperText={nameHelperText}
                                    label="Name"
                                    defaultValue=""
                                    onChange={(e)=>handleKV("name", e)}
                                    onInput={() => {
                                        setNameHelperText("")
                                    }}
                                />
                            </Grid>

                            <Grid item xs={12}>
                                <TextField
                                    fullWidth
                                    label="Pid"
                                    defaultValue=""
                                    onChange={(e)=>handleKV("pid", e)}
                                />
                            </Grid>

                            <Grid item xs={12}>
                                <TextField
                                    fullWidth
                                    label="Icon"
                                    defaultValue=""
                                    onChange={(e)=>handleKV("icon", e)}
                                />
                            </Grid>

                            <Grid item xs={12}>
                                <TextField
                                    fullWidth
                                    label="Url"
                                    defaultValue=""
                                    onChange={(e)=>handleKV("url", e)}
                                />
                            </Grid>
                            <Grid item xs={12}>
                                <TextField
                                    fullWidth
                                    label="Index Sort"
                                    defaultValue=""
                                    onChange={(e)=>handleKV("index_sort", e)}
                                />
                            </Grid>


                            <Grid item xs={12}>
                                <Grid
                                    container
                                    direction="row"
                                    justifyContent="space-between"
                                    alignItems="center"
                                >
                                    <Button variant="outlined" onClick={handleClose}>Cancel</Button>
                                    <LoadingButton variant="contained" loading={loading}
                                                   onClick={handleSubmit}> Save </LoadingButton>
                                </Grid>

                            </Grid>
                        </Grid>
                    </Paper>

                </Container>
            </Box>
        </Dialog>
    )
}
