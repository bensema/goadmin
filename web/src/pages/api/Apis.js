import React from 'react';
import Container from '@mui/material/Container';
import {useNavigate} from "react-router-dom"
import Button from '@mui/material/Button';
import Grid from '@mui/material/Grid';
import Paper from '@mui/material/Paper';
import DeleteIcon from '@mui/icons-material/Delete';

import {DataGridPro, GridActionsCellItem} from "@mui/x-data-grid-pro";

import {apiPages} from "@/api/api"

import {ApiAdd} from "./ApiAdd";
import {DeleteDialog} from "./ApiDelete";


export default function Apis() {
    const navigate = useNavigate();

    const [openDelete, setOpenDelete] = React.useState(false);
    const [openAdd, setOpenAdd] = React.useState(false);
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
            field: "api_group",
            headerName: "API Group",
            width: 150,
            flex: 0.1,
        },
        {
            field: "method",
            headerName: "Method",
            width: 150,
            flex: 0.1,
        },
        {
            field: "url",
            headerName: "Url",
            width: 150,
            flex: 0.1,
        },
        {
            field: 'actions',
            headerName: "Actions",
            type: 'actions',
            width: 100,
            getActions: (params: GridRowParams) => [
                <GridActionsCellItem icon={<DeleteIcon/>} onClick={() => {
                    handleOpenDelete()
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
            await apiPages({page: rowsState.page, page_size: rowsState.pageSize}).then((res) => {
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

    }, [rowsState.page, rowsState.pageSize, openAdd, openDelete]);

    const handleClickOpenAdd = () => {
        setOpenAdd(true);
    };

    const handleCloseAdd = () => {
        setOpenAdd(false);
    };

    const handleOpenDelete = () => {
        setOpenDelete(true);
    };
    const handleCloseDelete = () => {
        setOpenDelete(false);
    };

    return (
        <Container fixed maxWidth="lg" sx={{mt: 4, mb: 4}}>
            <Grid item xs={12}>
                <h1>APIs </h1>
                <Grid
                    container
                    direction="row"
                    justifyContent="flex-end"
                    alignItems="center"
                >
                    <Button variant="outlined"
                            onClick={handleClickOpenAdd}
                    >Add API</Button>
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
                                navigate("/myapi/" + params.id, {replace: true})
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
            <DeleteDialog deleteId={deleteId} deleteName={deleteName} open={openDelete}
                          handleClose={handleCloseDelete}/>
            {/*add dialog*/}
            <ApiAdd open={openAdd} handleClose={handleCloseAdd}/>
        </Container>
    );
}
