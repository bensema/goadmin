import React, {useEffect, useState} from "react"
import Container from '@mui/material/Container';
import Grid from '@mui/material/Grid';
import {useParams} from "react-router-dom"
import Paper from '@mui/material/Paper';
import Button from '@mui/material/Button';
import ArrowBackIcon from '@mui/icons-material/ArrowBack';
import {useNavigate} from "react-router-dom"
import Autocomplete from '@mui/material/Autocomplete';

import {Tree} from 'antd';
import 'antd/dist/antd.css';

import LoadingButton from '@mui/lab/LoadingButton';
import TextField from '@mui/material/TextField';
import {useSnackbar} from 'notistack';


import Title from "@/components/Title";
import {permissionInfo, permissionInfoUpdate, menuAll, apiAll} from "@/api/api"

function listToTree(oldArr) {
    oldArr.forEach(element => {
        let pid = element.pid;
        if (pid !== 0) {
            oldArr.forEach(ele => {
                if (ele.id === pid) { //当内层循环的ID== 外层循环的parendId时，（说明有children），需要往该内层id里建个children并push对应的数组；
                    if (!ele.children) {
                        ele.children = [];
                    }
                    ele.children.push(element);
                }
            });
        }
    });

    oldArr.sort(function (a, b) {
        return a.index_sort - b.index_sort
    })

    oldArr.forEach(element => {
        element.children && element.children.sort(function (a, b) {
            return a.index_sort - b.index_sort
        })
    })

    oldArr = oldArr.filter(ele => ele.pid === 0); //这一步是过滤，按树展开，将多余的数组剔除；
    return oldArr;
}

export function PermissionInfo() {
    const navigate = useNavigate()
    const {permissionId} = useParams();
    const [treeMenu, setTreeMenu] = useState([])
    const [apiList, setApiList] = useState([])
    const [saveLoading, setSaveLoading] = useState(false);

    const {enqueueSnackbar} = useSnackbar();


    const [permission, setPermission] = useState({
        id: 0,
        name: "",
        permission_group: "",
        remark: "",
        menus: [],
        apis: [],
    })

    useEffect(() => {
        (async () => {
            await menuAll().then((res) => {
                setTreeMenu(listToTree(res.data.data || []))
            })
        })();
        (async () => {
            await apiAll().then((res) => {
                setApiList(res.data.data || [])
            })
        })();
        (async () => {
            await permissionInfo({id: permissionId}).then((res) => {
                const data = res.data.data
                if (res.data.code === 0) {
                    setPermission(data)
                }
            })
        })();

        return () => {
        };
    }, [])//eslint-disable-line

    const onMenuCheck = (checkedKeysValue) => {
        console.log('onCheck', checkedKeysValue);
        setPermission((prev) => ({...prev, menus: checkedKeysValue}));
    };

    const onMenuSelect = (selectedKeysValue, info) => {
        console.log('onSelect', info);
    };


    const handleRemarkChange = (event) => {
        setPermission((prev) => ({...prev, remark: event.target.value}));
    };

    const handleNameChange = (event) => {
        setPermission((prev) => ({...prev, name: event.target.value}));
    };

    const handleApis = (event, newValue) => {
        setPermission((prev) => ({...prev, apis: newValue}));
    }

    const handleSave = () => {
        (async () => {
            setSaveLoading(true)
            let data = {};
            data.id = permission.id
            data.name = permission.name
            data.permission_group = permission.permission_group
            data.remark = permission.remark
            data.menus = permission.menus.map(item => item.id || item).join(",")
            data.apis = permission.apis.map(item => item.id).join(",")
            await permissionInfoUpdate(data).then((res) => {
                setSaveLoading(false)
                if (res.data.code === 0) {
                    enqueueSnackbar('Save success!', {variant: 'success'});
                }
            })
        })();
    }
    const handleBack = () => {
        navigate("/permissions", {replace: true})
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
                Permissions
            </Button>
            <h1> {permission.name} </h1>
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
                                    value={permission.name}
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
                                    value={permission.remark}
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
                        <Title>Others</Title>


                    </Paper>
                </Grid>

                <Grid item xs={12} md={6} lg={6}>
                    <Paper
                        sx={{
                            p: 2,
                            display: 'flex',
                            flexDirection: 'column',
                            height: 'auto',
                        }}
                    >
                        <Title>Api Information</Title>
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
                                <Autocomplete
                                    multiple
                                    id="apis"
                                    filterSelectedOptions
                                    options={apiList}
                                    isOptionEqualToValue={(option, value) => option.name === value.name}
                                    groupBy={(option) => option.api_group}
                                    getOptionLabel={(option) => option.name}
                                    value={permission.apis}
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
                        </Grid>

                    </Paper>
                </Grid>
                <Grid item xs={12} md={6} lg={6}>
                    <Paper
                        sx={{
                            p: 2,
                            display: 'flex',
                            flexDirection: 'column',
                            height: 'auto',
                        }}
                    >
                        <Title>Menu Information</Title>
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
                                <Tree
                                    checkable
                                    defaultExpandAll
                                    fieldNames={{
                                        title: 'name',
                                        key: 'id',
                                    }}
                                    onSelect={onMenuSelect}
                                    onCheck={onMenuCheck}
                                    checkedKeys={permission.menus.map(item => item.id || item)}
                                    treeData={treeMenu}
                                />
                            </Grid>
                        </Grid>

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

export default PermissionInfo;
