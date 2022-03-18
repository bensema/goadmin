import React, {useEffect, useState} from "react"
import Container from '@mui/material/Container';
import Grid from '@mui/material/Grid';
import {useParams} from "react-router-dom"
import Paper from '@mui/material/Paper';
import Button from '@mui/material/Button';
import ArrowBackIcon from '@mui/icons-material/ArrowBack';
import {useNavigate} from "react-router-dom"
import LoadingButton from '@mui/lab/LoadingButton';
import TextField from '@mui/material/TextField';
import {useSnackbar} from 'notistack';


import Title from "@/components/Title";
import {menuInfo, menuInfoUpdate} from "@/api/api"


export function MenuInfo() {
    const navigate = useNavigate()
    const {id} = useParams();
    const [saveLoading, setSaveLoading] = useState(false);

    const {enqueueSnackbar} = useSnackbar();


    const [menu, setMenu] = useState({
        id: 0,
        name: "",
        pid: "",
        icon: "",
        url: "",
        index_sort: "",
    })

    useEffect(() => {

        (async () => {
            await menuInfo({id: id}).then((res) => {
                const data = res.data.data
                if (res.data.code === 0) {
                    setMenu(data)
                }
            })
        })();

        return () => {
        };
    }, [])//eslint-disable-line


    const handleNameChange = (event) => {
        setMenu((prev) => ({...prev, name: event.target.value}));
    };
    const handlePidChange = (event) => {
        setMenu((prev) => ({...prev, pid: event.target.value}));
    };
    const handleIconChange = (event) => {
        setMenu((prev) => ({...prev, icon: event.target.value}));
    };
    const handleUrlChange = (event) => {
        setMenu((prev) => ({...prev, url: event.target.value}));
    };
    const handleIndexSortChange = (event) => {
        setMenu((prev) => ({...prev, index_sort: event.target.value}));
    };

    const handleSave = () => {
        (async () => {
            setSaveLoading(true)

            await menuInfoUpdate(menu).then((res) => {
                setSaveLoading(false)
                if (res.data.code === 0) {
                    enqueueSnackbar('Save success!', {variant: 'success'});
                }
            })
        })();
    }
    const handleBack = () => {
        navigate("/menus", {replace: true})
    }

    return (
        <Container maxWidth="lg" sx={{mt: 4, mb: 4}}>
            <Button
                startIcon={<ArrowBackIcon/>}
                size="large"
                variant="text"
                onClick={handleBack}
                sx={{
                    textTransform: "capitalize",
                }}

            >
                Menus
            </Button>
            <h1> {menu.name} </h1>
            <Grid container spacing={3}>
                <Grid item xs={12} md={8} lg={9}>
                    <Paper
                        sx={{
                            p: 2,
                            display: 'flex',
                            flexDirection: 'column',
                            height: 'auto',
                        }}
                    >
                        <Title>Base Information</Title>
                        <Grid
                            container
                            direction="row"
                            justifyContent="flex-start"
                            alignItems="center"
                            spacing={5}
                            sx={{
                                p: 2,
                            }}
                        >

                        </Grid>
                        <Grid
                            sx={{
                                p: 2,
                            }}
                        >
                            <Grid item xs={12}>
                                <TextField
                                    fullWidth
                                    label="Name"
                                    value={menu.name}
                                    onChange={handleNameChange}
                                />
                            </Grid>
                        </Grid>
                        <Grid
                            sx={{
                                p: 2,
                            }}
                        >
                            <Grid item xs={12}>
                                <TextField
                                    fullWidth
                                    label="Pid"
                                    value={menu.pid}
                                    onChange={handlePidChange}
                                />
                            </Grid>
                        </Grid>
                        <Grid
                            sx={{
                                p: 2,
                            }}
                        >
                            <Grid item xs={12}>
                                <TextField
                                    fullWidth
                                    label="Icon"
                                    value={menu.icon}
                                    onChange={handleIconChange}
                                />
                            </Grid>
                        </Grid>
                        <Grid
                            sx={{
                                p: 2,
                            }}
                        >
                            <Grid item xs={12}>
                                <TextField
                                    fullWidth
                                    label="Url"
                                    value={menu.url}
                                    onChange={handleUrlChange}
                                />
                            </Grid>
                        </Grid>
                        <Grid
                            sx={{
                                p: 2,
                            }}
                        >
                            <Grid item xs={12}>
                                <TextField
                                    fullWidth
                                    label="Index sort"
                                    value={menu.index_sort}
                                    onChange={handleIndexSortChange}
                                />
                            </Grid>
                        </Grid>
                    </Paper>
                </Grid>
                <Grid item xs={12} md={4} lg={3}>
                    <Paper
                        sx={{
                            p: 2,
                            display: 'flex',
                            flexDirection: 'column',
                            height: 240,
                            width: '100%'
                        }}
                    >
                        <Title>Others</Title>


                    </Paper>
                </Grid>

                <Grid item xs={12}>
                    <Grid
                        container
                        direction="row"
                        justifyContent="space-between"
                        alignItems="center"
                    >
                        <Button variant="outlined" onClick={handleBack}> Back </Button>
                        <LoadingButton variant="contained" loading={saveLoading}
                                       onClick={handleSave}> Save </LoadingButton>
                    </Grid>

                </Grid>

            </Grid>
        </Container>
    )
}

export default MenuInfo;
