import React from 'react';
import Container from '@mui/material/Container';
import PropTypes from 'prop-types';
import Stack from '@mui/material/Stack';
import Button from '@mui/material/Button';
import Grid from '@mui/material/Grid';
import Paper from '@mui/material/Paper';
import Typography from '@mui/material/Typography';
import TextField from '@mui/material/TextField';
import DateTimePicker from '@mui/lab/DateTimePicker';
import LocalizationProvider from '@mui/lab/LocalizationProvider';
import AdapterDateFns from '@mui/lab/AdapterDateFns';
import Box from "@mui/material/Box";
import InputLabel from '@mui/material/InputLabel';
import MenuItem from '@mui/material/MenuItem';
import FormControl from '@mui/material/FormControl';
import Select from '@mui/material/Select';

import {DataGridPro} from "@mui/x-data-grid-pro";
import {logLoginPages} from "@/api/api"
import {formatDate, DateTimeStart, DateTimeEnd, parseDateTime} from "@/utils/utils"

import {renderCellExpand} from "./renderCell";

const MyTextField = (props) => {
    return (
        <TextField
            size="small"
            required
            fullWidth
            readOnly
            {...props}
        />
    )
}
const MyGrid = (props) => {
    return (
        <Grid item sx={{p: 1,}} xs={12}>
            {props.children}
        </Grid>
    )
}

function DetailPanelContent({row}) {

    return (
        <Stack sx={{py: 2, height: 1, boxSizing: 'border-box'}} direction="column">
            <Paper sx={{flex: 1, mx: 'auto', width: '90%', p: 1}}>
                <Stack
                    component="form"
                    justifyContent="space-between"
                    sx={{height: 1}}
                >
                    <Typography variant="h6">{`Login ID   #${row.id}`}</Typography>
                    <Grid
                        container
                        direction="row"
                        justifyContent="center"
                        alignItems="stretch"
                    >
                        <MyGrid>
                            <MyTextField
                                label="Admin ID"
                                value={row.admin_id}
                            />
                        </MyGrid>
                        <MyGrid>
                            <MyTextField
                                label="Name"
                                value={row.name}
                            />
                        </MyGrid>
                        <MyGrid>
                            <MyTextField
                                label="Location"
                                value={row.location}
                            />
                        </MyGrid>
                        <MyGrid>
                            <MyTextField
                                label="Os"
                                value={row.os}
                            />
                        </MyGrid>
                        <MyGrid>
                            <MyTextField
                                label="Browser"
                                value={row.browser}
                            />
                        </MyGrid>
                        <MyGrid>
                            <MyTextField
                                label="User_Agent"
                                value={row.user_agent}
                            />
                        </MyGrid>
                        <MyGrid>
                            <MyTextField
                                label="Url"
                                value={row.url}
                            />
                        </MyGrid>
                        <MyGrid>
                            <MyTextField
                                label="IP"
                                value={row.ip}
                            />
                        </MyGrid>
                        <MyGrid>
                            <MyTextField
                                label="Result"
                                value={row.result}
                            />
                        </MyGrid>
                        <MyGrid>
                            <MyTextField
                                label="Record at"
                                value={formatDate(row.record_at * 1000)}
                            />
                        </MyGrid>
                        <MyGrid>
                            <MyTextField
                                label="Remark"
                                value={row.remark}
                            />
                        </MyGrid>

                    </Grid>


                </Stack>
            </Paper>
        </Stack>
    );
}

DetailPanelContent.propTypes = {
    row: PropTypes.any.isRequired,
};


export default function LogLogin() {

    const getDetailPanelContent = React.useCallback(
        ({row}) => <DetailPanelContent row={row}/>,
        [],
    );

    const getDetailPanelHeight = React.useCallback(() => 700, []);
    const [search, setSearch] = React.useState(true)
    const [searchData, setSearchData] = React.useState({
        name: "",
        ip: "",
        result: "",
        record_at_from: DateTimeStart(),
        record_at_to: DateTimeEnd(),
    })
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
        },
        {
            field: "location",
            headerName: "Location",
            width: 150,
            flex: 1,
        },
        {
            field: "os",
            headerName: "Os",
            width: 150,
            flex: 1,
            renderCell: renderCellExpand,
        },
        {
            field: "browser",
            headerName: "Browser",
            width: 150,
            flex: 1,
            renderCell: renderCellExpand,
        },
        {
            field: "ip",
            headerName: "IP",
            width: 150,
        },
        {
            // headerAlign: 'center',
            field: "record_at",
            headerName: "Record at",
            flex: 0.1,
            minWidth: 200,
            valueFormatter: (params) => {
                return formatDate(params.value * 1000);
            },
        },
    ];

    React.useEffect(() => {
        let active = true;
        (async () => {
            setRowsState((prev) => ({...prev, loading: true}));
            let params = {
                page: rowsState.page,
                page_size: rowsState.pageSize,
                order_by_direction: 'desc',
                order_by_filed: 'id',
            }
            params.record_at_from = parseDateTime(searchData.record_at_from)
            params.record_at_to = parseDateTime(searchData.record_at_to)

            if (searchData.name !== "") {
                params.name = searchData.name
            }
            if (searchData.ip !== "") {
                params.ip = searchData.ip
            }
            if (searchData.result !== "") {
                params.result = searchData.result
            }
            await logLoginPages(params).then((res) => {
                res = res.data
                if (res.code === 0) {
                    const newState = {
                        page: res.data.page,
                        pageSize: res.data.page_size,
                        rows: res.data.rows || [],
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

    }, [rowsState.page, rowsState.pageSize, search]);

    const handleSearchKV = (k, e) => {
        setSearchData((prev) => ({...prev, [k]: e.target.value}));
    }

    const handleChangeRecordAtFrom = (newValue) => {
        setSearchData((prev) => ({...prev, record_at_from: newValue}));
    }

    const handleChangeRecordAtTo = (newValue) => {
        setSearchData((prev) => ({...prev, record_at_to: newValue}));
    }

    return (
        <LocalizationProvider dateAdapter={AdapterDateFns}>
            <Container fixed maxWidth="lg" sx={{mt: 4, mb: 4}}>
                <Grid item xs={12}>
                    <h1>Logins </h1>

                    <Paper sx={{p: 2, display: 'flex', flexDirection: 'column'}} elevation={2}>
                        <div style={{width: '100%'}}>
                            <Box sx={{mt: 2, mb: 2}}>
                                <Grid
                                    container
                                    direction="row"
                                    justifyContent="flex-end"
                                    alignItems="stretch"
                                    spacing={{xs: 2, md: 3}}
                                    columns={{xs: 4, sm: 8, md: 12}}
                                >
                                    <Grid item xs={4}>
                                        <TextField
                                            fullWidth
                                            label="Name"
                                            defaultValue=""
                                            onChange={(e) => handleSearchKV("name", e)}
                                        />
                                    </Grid>
                                    <Grid item xs={4}>
                                        <TextField
                                            fullWidth
                                            label="IP"
                                            defaultValue=""
                                            onChange={(e) => handleSearchKV("ip", e)}
                                        />
                                    </Grid>
                                    <Grid item xs={4}>
                                        <FormControl fullWidth>
                                            <InputLabel id="demo-simple-select-label">Result</InputLabel>
                                            <Select
                                                label="Result"
                                                defaultValue={""}
                                                onChange={(e) => handleSearchKV("result", e)}
                                            >
                                                <MenuItem value="">
                                                    <em>None</em>
                                                </MenuItem>
                                                <MenuItem value={"SUCCESS"}>SUCCESS</MenuItem>
                                                <MenuItem value={"FAIL"}>FAIL</MenuItem>
                                            </Select>
                                        </FormControl>

                                    </Grid>
                                    <Grid item xs={4}>
                                        <DateTimePicker
                                            label="Record at from"
                                            value={searchData.record_at_from}
                                            views={['year', 'month', 'day', 'hours', 'minutes', 'seconds']}
                                            inputFormat="yyyy-MM-dd HH:mm:ss"
                                            mask="__-__-__ __:__:__"
                                            onChange={handleChangeRecordAtFrom}
                                            renderInput={(params) => <TextField fullWidth {...params} />}
                                        />
                                    </Grid>
                                    <Grid item xs={4}>
                                        <DateTimePicker
                                            label="Record at to"
                                            value={searchData.record_at_to}
                                            views={['year', 'month', 'day', 'hours', 'minutes', 'seconds']}
                                            inputFormat="yyyy-MM-dd HH:mm:ss"
                                            mask="__-__-__ __:__:__"
                                            onChange={handleChangeRecordAtTo}
                                            renderInput={(params) => <TextField fullWidth {...params} />}
                                        />
                                    </Grid>
                                    <Grid item xs={4} container
                                          direction="row"
                                          justifyContent="center"
                                          alignItems="center">
                                        <Button
                                            size="large"
                                            variant="contained"
                                            onClick={() => {
                                                setSearch(!search)
                                            }}
                                        >Search</Button>
                                    </Grid>
                                </Grid>

                            </Box>

                            <DataGridPro
                                autoHeight
                                pagination

                                rowsPerPageOptions={[10, 20, 50]}

                                getDetailPanelHeight={getDetailPanelHeight}
                                getDetailPanelContent={getDetailPanelContent}

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

            </Container>
        </LocalizationProvider>
    );
}
