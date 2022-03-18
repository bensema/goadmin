import React, {useEffect, useState} from "react"
import Container from '@mui/material/Container';
import Grid from '@mui/material/Grid';
import {useParams} from "react-router-dom"
import Paper from '@mui/material/Paper';
import Button from '@mui/material/Button';
import ArrowBackIcon from '@mui/icons-material/ArrowBack';
import {useNavigate} from "react-router-dom"
import FormControlLabel from '@mui/material/FormControlLabel';
import Switch from '@mui/material/Switch';
import FormGroup from '@mui/material/FormGroup';
import LoadingButton from '@mui/lab/LoadingButton';
import TextField from '@mui/material/TextField';
import Divider from '@mui/material/Divider';
import Autocomplete from '@mui/material/Autocomplete';
import {  useSnackbar } from 'notistack';


import Title from "@/components/Title";
import {adminInfo,adminInfoUpdate, roleAll} from "@/api/api"
import {formatDate} from "@/utils/utils"


export function AdminInfo() {
    const navigate = useNavigate()
    const {adminId} = useParams();
    const [roleList, setRoleList] = useState([])
    const [saveLoading, setSaveLoading] = useState(false);

    const { enqueueSnackbar } = useSnackbar();


    const [admin, setAdmin] = useState({
        id: 0,
        name: "",
        status: "",
        created_at: "",
        updated_at: "",
        roles: [],
        remark: "",

    })

    useEffect(() => {
        (async () => {
            await roleAll().then((res) => {
                setRoleList(res.data.data)
            })
        })();
        (async () => {
            await adminInfo({id: adminId}).then((res) => {
                const data = res.data.data
                setAdmin(data)
            })
        })();

        return () => {};
    }, [])//eslint-disable-line

    const handleStatusChange = (event) => {
        setAdmin((prev) => ({...prev, status: event.target.checked?"Enable":"Disable"}));
    };

    const handleRemarkChange = (event) => {
        setAdmin((prev) => ({...prev, remark: event.target.value}));
    };

    const handleRoles = (event, newValue) => {
        setAdmin((prev) => ({...prev, roles: newValue}));
    }

    const handleSave = () => {
        (async () => {
            setSaveLoading(true)
            let data={};
            data.id = admin.id
            data.status = admin.status
            data.remark = admin.remark
            data.roles = admin.roles.map(item=>item.id).join(",")
            await adminInfoUpdate(data).then((res) => {
                setSaveLoading(false)
                if (res.data.code === 0){
                    enqueueSnackbar('Save success!', {variant:'success'});
                }
            })
        })();
    }
    const handleBack = () => {
        navigate("/admins", {replace: true})
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
                Admins
            </Button>
            <h1> {admin.name} </h1>
            <Grid container spacing={3}>
                <Grid item xs={12} md={8} lg={9}>
                    <Paper
                        sx={{
                            p: 2,
                            display: 'flex',
                            flexDirection: 'column',
                            height: 240,
                        }}
                    >
                        <Title>Information</Title>
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
                            <Grid item >
                                <FormGroup>
                                    <FormControlLabel control={<Switch
                                        checked={admin.status==="Enable"}
                                        onChange={handleStatusChange}
                                        inputProps={{ 'aria-label': 'controlled' }}
                                    />} label={admin.status} />
                                </FormGroup>
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
                                    value={admin.remark}
                                    onChange={handleRemarkChange}
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
                        <Title>History</Title>


                    </Paper>
                </Grid>


                <Grid item xs={12} md={8} lg={9}>
                    <Paper
                        sx={{
                            p: 2,
                            display: 'flex',
                            flexDirection: 'column',
                            height: 240,
                        }}
                    >
                        <Title>Base Information</Title>
                        <Grid
                            container
                            direction="row"
                            justifyContent="center"
                            alignItems="center"
                            spacing={5}
                            sx={{
                                p: 2,
                            }}
                        >
                            <Grid item xs={6}>
                                <TextField
                                    disabled
                                    fullWidth
                                    label="Created at"
                                    defaultValue="Normal"
                                    value={formatDate(admin.created_at * 1000)}/>
                            </Grid>
                            <Grid item xs={6}>
                                <TextField
                                    disabled
                                    fullWidth
                                    label="Updated at"
                                    defaultValue="Normal"
                                    value={formatDate(admin.updated_at * 1000)}/>

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
                                    id="roles"
                                    filterSelectedOptions
                                    options={roleList}
                                    isOptionEqualToValue={(option, value) => option.name === value.name}
                                    getOptionLabel={(option) => option.name}
                                    value={admin.roles}
                                    onChange={handleRoles}
                                    renderInput={(params) => (
                                        <TextField
                                            {...params}
                                            label="Permissions"
                                            placeholder="Role"
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
                {/* Recent Orders */}
                <Grid item xs={12}>
                    <Paper sx={{p: 2, display: 'flex', flexDirection: 'column'}}>
                        {/*<Orders />*/}
                        <Title>Log actions</Title>

                    </Paper>
                </Grid>
                <Grid item xs={12}>
                    <Grid
                        container
                        direction="row"
                        justifyContent="space-between"
                        alignItems="center"
                    >
                        <Button variant="outlined"  onClick={handleBack}> Back </Button>
                        <LoadingButton variant="contained" loading={saveLoading} onClick={handleSave}> Save </LoadingButton>
                    </Grid>

                </Grid>

            </Grid>
        </Container>
    )
}

export default AdminInfo;
