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
import {apiInfo, apiInfoUpdate} from "@/api/api"


export function ApiInfo() {
    const navigate = useNavigate()
    const {id} = useParams();
    const [saveLoading, setSaveLoading] = useState(false);

    const {enqueueSnackbar} = useSnackbar();


    const [aInfo, setAInfo] = useState({
        id: 0,
        name: "",
        api_group: "",
        method: "",
        url: "",
    })

    useEffect(() => {
        (async () => {
            await apiInfo({id: id}).then((res) => {
                const data = res.data.data
                if (res.data.code === 0) {
                    setAInfo(data)
                }
            })
        })();

        return () => {
        };
    }, [])//eslint-disable-line

    const handleChangeKV = (k, event) => {
        setAInfo((prev) => ({...prev, [k]: event.target.value}));
    };


    const handleSave = () => {
        (async () => {
            setSaveLoading(true)

            await apiInfoUpdate(aInfo).then((res) => {
                setSaveLoading(false)
                if (res.data.code === 0) {
                    enqueueSnackbar('Save success!', {variant: 'success'});
                }
            })
        })();
    }
    const handleBack = () => {
        navigate("/myapis", {replace: true})
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
                APIS
            </Button>
            <h1> {aInfo.name} </h1>
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
                                    value={aInfo.name}
                                    onChange={(e)=>handleChangeKV("name", e)}
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
                                    label="API Group"
                                    value={aInfo.api_group}
                                    onChange={(e)=>handleChangeKV("api_group", e)}
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
                                    label="Method"
                                    value={aInfo.method}
                                    onChange={(e)=>handleChangeKV("method", e)}
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
                                    value={aInfo.url}
                                    onChange={(e)=>handleChangeKV("url", e)}
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

export default ApiInfo;
