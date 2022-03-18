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
import Divider from '@mui/material/Divider';
import Autocomplete from '@mui/material/Autocomplete';
import {useSnackbar} from 'notistack';


import Title from "@/components/Title";
import {roleInfo, roleInfoUpdate, permissionAll} from "@/api/api"


export function RoleInfo() {
    const navigate = useNavigate()
    const {roleId} = useParams();
    const [permissionList, setPermissionList] = useState([])
    const [saveLoading, setSaveLoading] = useState(false);

    const {enqueueSnackbar} = useSnackbar();


    const [role, setRole] = useState({
        id: 0,
        name: "",
        permissions: [],
        remark: "",
    })

    useEffect(() => {
        (async () => {
            await permissionAll().then((res) => {
                setPermissionList(res.data.data||[])
            })
        })();
        (async () => {
            await roleInfo({id: roleId}).then((res) => {
                const data = res.data.data
                if (res.data.code===0){
                    setRole(data)
                }
            })
        })();

        return () => {
        };
    }, [])//eslint-disable-line


    const handleRemarkChange = (event) => {
        setRole((prev) => ({...prev, remark: event.target.value}));
    };

    const handleNameChange = (event) => {
        setRole((prev) => ({...prev, name: event.target.value}));
    };

    const handleRoles = (event, newValue) => {
        setRole((prev) => ({...prev, permissions: newValue}));
    }

    const handleSave = () => {
        (async () => {
            setSaveLoading(true)
            let data = {};
            data.id = role.id
            data.name = role.name
            data.remark = role.remark
            data.permissions = role.permissions.map(item => item.id).join(",")
            await roleInfoUpdate(data).then((res) => {
                setSaveLoading(false)
                if (res.data.code === 0) {
                    enqueueSnackbar('Save success!', {variant: 'success'});
                }
            })
        })();
    }
    const handleBack = () => {
        navigate("/roles", {replace: true})
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
                Roles
            </Button>
            <h1> {role.name} </h1>
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
                                    value={role.name}
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
                                    label="Remark"
                                    value={role.remark}
                                    onChange={handleRemarkChange}
                                />
                            </Grid>
                        </Grid>
                        <Divider/>
                        <Grid
                            sx={{
                                p: 2,
                            }}
                        >
                            <Grid item xs={12}>
                                <Autocomplete
                                    multiple
                                    id="permissions"
                                    filterSelectedOptions
                                    options={permissionList}
                                    isOptionEqualToValue={(option, value) => option.name === value.name}
                                    groupBy={(option) => option.permission_group}
                                    getOptionLabel={(option) => option.name}
                                    value={role.permissions}
                                    onChange={handleRoles}
                                    renderInput={(params) => (
                                        <TextField
                                            {...params}
                                            label="Permissions"
                                            placeholder="Permission"
                                        />
                                    )}
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

export default RoleInfo;
