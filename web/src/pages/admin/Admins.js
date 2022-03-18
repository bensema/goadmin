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

import {adminPages} from "@/api/api"

import {formatDate} from "@/utils/utils"
import {AdminAdd} from "./AdminAdd";
import {DeleteDialog} from "./AdminDelete";


export default function Admins() {
    const navigate = useNavigate();

    const [openDeleteAdmin, setOpenDeleteAdmin] = React.useState(false);
    const [openAddAdmin, setOpenAddAdmin] = React.useState(false);
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
            flex: 1,
        },
        {
            field: "roles",
            headerName: "Roles",
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
            field: "status",
            headerName: "Status",
            width: 150,
        },
        {
            field: "created_at",
            headerName: "Created at ",
            flex: 0.1,
            minWidth: 200,
            valueFormatter: (params) => {
                return formatDate(params.value * 1000);
            },
        },
        {
            // headerAlign: 'center',
            field: "updated_at",
            headerName: "Updated at",
            flex: 0.1,
            minWidth: 200,
            valueFormatter: (params) => {
                return formatDate(params.value * 1000);
            },
        },
        {
            field: 'actions',
            headerName: "Actions",
            type: 'actions',
            width: 100,
            getActions: (params: GridRowParams) => [
                <GridActionsCellItem icon={<DeleteIcon/>} onClick={() => {
                    handleOpenDeleteAdmin()
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
            await adminPages({page: rowsState.page, page_size: rowsState.pageSize}).then((res) => {
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

    }, [rowsState.page, rowsState.pageSize, openAddAdmin, openDeleteAdmin]);

    const handleClickOpenAddAdmin = () => {
        setOpenAddAdmin(true);
    };

    const handleCloseAddAdmin = () => {
        setOpenAddAdmin(false);
    };

    const handleOpenDeleteAdmin = () => {
        setOpenDeleteAdmin(true);
    };
    const handleCloseDeleteAdmin = () => {
        setOpenDeleteAdmin(false);
    };

    return (
        <Container fixed maxWidth="lg" sx={{mt: 4, mb: 4}}>
            <Grid item xs={12}>
                <h1>Admins </h1>
                <Grid
                    container
                    direction="row"
                    justifyContent="flex-end"
                    alignItems="center"
                >
                    <Button variant="outlined"
                            onClick={handleClickOpenAddAdmin}
                    >Add Admin</Button>
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
                                navigate("/admin/" + params.id, {replace: true})
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
            <DeleteDialog deleteId={deleteId} deleteName={deleteName} open={openDeleteAdmin}
                          handleClose={handleCloseDeleteAdmin}/>
            {/*add dialog*/}
            <AdminAdd open={openAddAdmin} handleClose={handleCloseAddAdmin}/>
        </Container>
    );
}
