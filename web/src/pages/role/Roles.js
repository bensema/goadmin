import React from 'react';
import Container from '@mui/material/Container';
import Chip from '@mui/material/Chip';
import Stack from '@mui/material/Stack';
import {useNavigate} from "react-router-dom"
import Button from '@mui/material/Button';
import Grid from '@mui/material/Grid';
import Paper from '@mui/material/Paper';
import DeleteIcon from '@mui/icons-material/Delete';

import {DataGridPro, GridActionsCellItem} from "@mui/x-data-grid-pro";

import {rolePages} from "@/api/api"

import {RoleAdd} from "./RoleAdd";
import {DeleteDialog} from "./RoleDelete";


export default function Roles() {
    const navigate = useNavigate();

    const [openDeleteRole, setOpenDeleteRole] = React.useState(false);
    const [openAddRole, setOpenAddRole] = React.useState(false);
    const [deleteId, setDeleteId] = React.useState();
    const [deleteName, setDeleteName] = React.useState();

    const [rowsState, setRowsState] = React.useState({
        page: 0,
        pageSize: 10,
        rowCount: 0,
        rows: [],
        loading: false,
    });

    const columns = [
        {
            field: "id",
            headerName: "ID",
        },
        {
            field: "name",
            headerName: "Name",
            width: 150,
            flex: 0.1,
        },
        {
            field: "permissions",
            headerName: "Permissions",
            width: 150,
            flex: 1,
            renderCell: (params) => (
                <Stack direction="row" spacing={0.5}>
                    {params.value && params.value.map((item) => <Chip key={item.id} variant="outlined" color="primary"
                                                                      size="small"
                                                                      label={item.name}/>
                    )}
                </Stack>
            ),

        },
        {
            field: 'actions',
            headerName: "Actions",
            type: 'actions',
            width: 100,
            getActions: (params: GridRowParams) => [
                <GridActionsCellItem icon={<DeleteIcon/>} onClick={() => {
                    handleOpenDeleteRole()
                    setDeleteId(params.id)
                    setDeleteName(params.row.name)
                }} label="Delete"/>,
            ],
        },
    ];

    React.useEffect(() => {
        let active = true;
        (async () => {
            setRowsState((prev) => ({...prev, loading: true}));
            await rolePages({page: rowsState.page, page_size: rowsState.pageSize}).then((res) => {
                res = res.data
                if (res.code === 0) {
                    const newState = {
                        page: res.data.page,
                        pageSize: res.data.page_size,
                        rows: res.data.rows,
                        rowCount: res.data.rows_total,
                        loading: false,
                    }

                    if (!active) {
                        return;
                    }
                    setRowsState(newState);

                } else {

                }
            })
        })();
        return () => {
            active = false;
        };

    }, [rowsState.page, rowsState.pageSize, openAddRole, openDeleteRole]);

    const handleClickOpenAddRole = () => {
        setOpenAddRole(true);
    };

    const handleCloseAddRole = () => {
        setOpenAddRole(false);
    };

    const handleOpenDeleteRole = () => {
        setOpenDeleteRole(true);
    };
    const handleCloseDeleteRole = () => {
        setOpenDeleteRole(false);
    };

    return (
        <Container fixed maxWidth="lg" sx={{mt: 4, mb: 4}}>
            <Grid item xs={12}>
                <h1>Roles </h1>
                <Grid
                    container
                    direction="row"
                    justifyContent="flex-end"
                    alignItems="center"
                >
                    <Button variant="outlined"
                            onClick={handleClickOpenAddRole}
                    >Add Role</Button>
                </Grid>
                <Paper sx={{p: 2, display: 'flex', flexDirection: 'column'}} elevation={2}>
                    <div style={{width: '100%'}}>
                        <DataGridPro
                            autoHeight
                            pagination
                            initialState={{
                                pinnedColumns: {right: ['actions']}
                            }}

                            rowsPerPageOptions={[10, 20, 50]}

                            onRowClick={(params, event) => {
                                navigate("/role/" + params.id, {replace: true})
                            }}

                            paginationMode="server"
                            onPageChange={(page) => setRowsState((prev) => ({...prev, page}))}
                            onPageSizeChange={(pageSize) =>
                                setRowsState((prev) => ({...prev, pageSize}))
                            }

                            {...rowsState}
                            columns={columns}/>
                    </div>
                </Paper>
            </Grid>
            {/*delete dialog*/}
            <DeleteDialog deleteId={deleteId} deleteName={deleteName} open={openDeleteRole}
                          handleClose={handleCloseDeleteRole}/>
            {/*add dialog*/}
            <RoleAdd open={openAddRole} handleClose={handleCloseAddRole}/>
        </Container>
    );
}
