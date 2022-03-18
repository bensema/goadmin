import React from 'react';
import Container from '@mui/material/Container';
import {useNavigate} from "react-router-dom"
import Button from '@mui/material/Button';
import Grid from '@mui/material/Grid';
import Paper from '@mui/material/Paper';
import DeleteIcon from '@mui/icons-material/Delete';

import {DataGridPro, GridActionsCellItem} from "@mui/x-data-grid-pro";

import {menuPages} from "@/api/api"

import {MenuAdd} from "./MenuAdd";
import {DeleteDialog} from "./MenuDelete";

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

function findParents(treeData, id) {
    if (treeData.length === 0) return
    for (let i = 0; i < treeData.length; i++) {
        if (treeData[i].id === id) {
            return []
        } else {
            if (treeData[i].children) {
                let res = findParents(treeData[i].children, id)
                if (res !== undefined) {
                    return res.concat(treeData[i].name)
                }
            }
        }
    }
}

function hierarchy(datasource,node){
    const targetData = {};
    function loops(data = [], parent) {
        return data.map(({ children, name: value }) => {
            const node = {
                value,
                parent
            }
            targetData[value] = node;
            node.children = loops(children, node);
            return
        })
    }

    function getNode(value) {
        let node = [];
        let currentNode = targetData[value];
        node.push(currentNode.value);
        if (currentNode.parent) {
            node = [...getNode(currentNode.parent.value), ...node];
        }
        return node
    }


    loops(datasource)
    //获取父节点
    return getNode(node)
}


export default function Menus() {
    const navigate = useNavigate();

    const [openDelete, setOpenDelete] = React.useState(false);
    const [openAdd, setOpenAdd] = React.useState(false);
    const [deleteId, setDeleteId] = React.useState();
    const [deleteName, setDeleteName] = React.useState();

    const [rowsState, setRowsState] = React.useState({
        page: 0,
        pageSize: 1000,
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
            field: "pid",
            headerName: "Parent Id",
            width: 150,
            flex: 0.1,
        },
        {
            field: "index_sort",
            headerName: "Sort",
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
            field: "icon",
            headerName: "Icon",
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
            await menuPages({page: rowsState.page, page_size: rowsState.pageSize}).then((res) => {
                res = res.data
                if (res.code === 0) {
                    let dataList = res.data.rows
                    let dataSource = res.data.rows
                    let dataTree = listToTree(dataSource);
                    dataList.map((item) => item.hierarchy = hierarchy(dataTree, item.name))
                    const newState = {
                        page: res.data.page,
                        pageSize: res.data.page_size,
                        rows: dataList,
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

    const getTreeDataPath = (row) => row.hierarchy;

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
                <h1>Menus </h1>
                <Grid
                    container
                    direction="row"
                    justifyContent="flex-end"
                    alignItems="center"
                >
                    <Button variant="outlined"
                            onClick={handleClickOpenAdd}
                    >Add Menu</Button>
                </Grid>
                <Paper sx={{p: 2, display: 'flex', flexDirection: 'column'}} elevation={2}>
                    <div style={{width: '100%'}}>
                        <DataGridPro
                            autoHeight
                            pagination
                            treeData
                            initialState={{
                                pinnedColumns: {right: ['actions']}
                            }}

                            rowsPerPageOptions={[1000]}

                            onRowClick={(params, event) => {
                                navigate("/menu/" + params.id, {replace: true})
                            }}

                            getTreeDataPath={getTreeDataPath}

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
            <MenuAdd open={openAdd} handleClose={handleCloseAdd}/>
        </Container>
    );
}
