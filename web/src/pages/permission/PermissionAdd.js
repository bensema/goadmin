import React, {useState, useEffect} from "react"
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
import Autocomplete from '@mui/material/Autocomplete';
import {useSnackbar} from 'notistack';

import {menuAll, apiAll, permissionAdd} from "@/api/api"


const Transition = React.forwardRef(function Transition(props, ref) {
    return <Slide direction="up" ref={ref} {...props} />;
});

export const PermissionAdd = (props) => {
    const {open, handleClose} = props
    const [menuList, setMenuList] = useState([])
    const [apiList, setApiList] = useState([])
    const [loading, setLoading] = useState(false)

    const [nameHelperText, setNameHelperText] = React.useState("");

    const [role, setRole] = useState({
        name: "",
        permission_group: "",
        remark: "",
        menus: "",
        apis: "",
    })

    const {enqueueSnackbar} = useSnackbar();

    useEffect(() => {
        (async () => {
            await menuAll().then((res) => {
                setMenuList(res.data.data)
            })
        })();
        (async () => {
            await apiAll().then((res) => {
                setApiList(res.data.data)
            })
        })();

    }, [])

    const handleName = (e) => {
        setRole((prev) => ({...prev, name: e.target.value}));
    }

    const handlePermissionGroup = (e) => {
        setRole((prev) => ({...prev, permission_group: e.target.value}));
    }

    const handleRemark = (e) => {
        setRole((prev) => ({...prev, remark: e.target.value}));
    }

    const handleMenus = (event: object, value: T | T[], reason: string, details?: string) => {

        setRole((prev) => ({...prev, menus: value.map(item => item.id).join(",")}));
    }

    const handleApis = (event: object, value: T | T[], reason: string, details?: string) => {

        setRole((prev) => ({...prev, apis: value.map(item => item.id).join(",")}));
    }

    const handleSubmit = () => {

        if (role.name === "") {
            setNameHelperText("Name required")
            return
        }

        (async () => {
            setLoading(true)
            await permissionAdd(role).then((res) => {
                setLoading(false)
                if (res.data.code === 0) {
                    enqueueSnackbar('Add Role success!', {variant: 'success'});
                    handleClose()
                }else {
                    enqueueSnackbar('Add Role fail!', {variant: 'error'});
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
                        Add Permission
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
                                    onChange={handleName}
                                    onInput={() => {
                                        setNameHelperText("")
                                    }}
                                />
                            </Grid>

                            <Grid item xs={12}>
                                <TextField
                                    fullWidth
                                    label="Permission Group"
                                    defaultValue=""
                                    onChange={handlePermissionGroup}

                                />
                            </Grid>

                            <Grid item xs={12}>
                                <TextField
                                    fullWidth
                                    label="Remark"
                                    defaultValue=""
                                    onChange={handleRemark}
                                />
                            </Grid>

                            <Grid item xs={12}>
                                <Autocomplete
                                    multiple
                                    id="menus"
                                    filterSelectedOptions
                                    options={menuList}
                                    isOptionEqualToValue={(option, value) => option.name === value.name}
                                    getOptionLabel={(option) => option.name}
                                    onChange={handleMenus}
                                    renderInput={(params) => (
                                        <TextField
                                            {...params}
                                            label="Menus"
                                            placeholder="Menu"
                                        />
                                    )}
                                />
                            </Grid>
                            <Grid item xs={12}>
                                <Autocomplete
                                    multiple
                                    id="apis"
                                    filterSelectedOptions
                                    options={apiList}
                                    groupBy={(option) => option.api_group}
                                    isOptionEqualToValue={(option, value) => option.name === value.name}
                                    getOptionLabel={(option) => option.name}
                                    onChange={handleApis}
                                    renderInput={(params) => (
                                        <TextField
                                            {...params}
                                            label="Apis"
                                            placeholder="Api"
                                        />
                                    )}
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
