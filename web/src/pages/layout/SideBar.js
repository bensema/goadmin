import React, {Fragment, useEffect} from "react";
import {styled} from "@mui/material/styles";
import MuiDrawer from "@mui/material/Drawer";
import Toolbar from "@mui/material/Toolbar";
import List from "@mui/material/List";

import ListItem from "@mui/material/ListItem";
import ListItemIcon from "@mui/material/ListItemIcon";
import ListItemText from "@mui/material/ListItemText";

import ExpandLess from '@mui/icons-material/ExpandLess';
import ExpandMore from '@mui/icons-material/ExpandMore';

import Collapse from '@mui/material/Collapse';

import FindInPageIcon from '@mui/icons-material/FindInPage';
import LabelIcon from '@mui/icons-material/Label';
import Dashboard from '@mui/icons-material/Dashboard';
import ManageAccountsIcon from '@mui/icons-material/ManageAccounts';
import {Link} from "react-router-dom"
import {menus} from "@/api/api"


const drawerWidth = 240;

const Drawer = styled(MuiDrawer, {shouldForwardProp: (prop) => prop !== 'open'})(
    ({theme, open}) => ({
        '& .MuiDrawer-paper': {
            position: 'relative',
            whiteSpace: 'nowrap',
            width: drawerWidth,
            transition: theme.transitions.create('width', {
                easing: theme.transitions.easing.sharp,
                duration: theme.transitions.duration.enteringScreen,
            }),
            boxSizing: 'border-box',
            ...(!open && {
                overflowX: 'hidden',
                transition: theme.transitions.create('width', {
                    easing: theme.transitions.easing.sharp,
                    duration: theme.transitions.duration.leavingScreen,
                }),
                width: theme.spacing(7),
                [theme.breakpoints.up('sm')]: {
                    width: theme.spacing(9),
                },
            }),
        },
    }),
);


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

    oldArr.sort(function(a,b){
        return a.index_sort - b.index_sort
    })

    oldArr.forEach(element=>{
        element.children && element.children.sort(function(a,b){
            return a.index_sort - b.index_sort
        })
    })

    oldArr = oldArr.filter(ele => ele.pid === 0); //这一步是过滤，按树展开，将多余的数组剔除；
    return oldArr;
}

function switchIcon(name) {
    switch (name) {
        case "Dashboard":
            return <Dashboard/>
        case "ManageAccountsIcon":
            return <ManageAccountsIcon/>
        case "FindInPageIcon":
            return <FindInPageIcon/>
        default:
            return <LabelIcon/>
    }
}


function SubList(props) {
    const {icon, text, handleListItemClick, myKey} = props
    const [open, setOpen] = React.useState(false);

    const handleClick = () => {
        setOpen(!open);
        handleListItemClick(myKey)
    };

    return (<Fragment>
        <ListItem
            button
            onClick={handleClick}>
            <ListItemIcon>
                {icon}
            </ListItemIcon>
            <ListItemText primary={text}/>
            {open ? <ExpandLess/> : <ExpandMore/>}
        </ListItem>
        <Collapse in={open} timeout="auto" unmountOnExit>
            <List component="div" disablePadding
                  // sx={{ pl: 4 }}
            >
                {props.children}
            </List>
        </Collapse>

    </Fragment>)
}

function ListTree(props) {
    const {menuTree, selectedIndex, handleListItemClick} = props
    return (<>
        {menuTree.map(function (item) {

            if (item.pid === 0) {
                if (item.children === undefined || item.children.length === 0) {
                    return <ListItem
                        button
                        key={item.id}
                        selected={selectedIndex === item.id}
                        onClick={(event) => {
                            handleListItemClick(event, item.id)
                        }}
                        component={Link}
                        to={item.url}
                    >
                        <ListItemIcon>
                            {switchIcon(item.icon)}
                        </ListItemIcon>
                        <ListItemText primary={item.name}/>
                    </ListItem>
                } else {
                    return <SubList
                        key={"subList" + item.id}
                        myKey={item.id}
                        icon={switchIcon(item.icon)}
                        text={item.name}
                        selectedIndex={selectedIndex}
                        handleListItemClick={handleListItemClick}
                    >
                        <ListTree
                            menuTree={item.children}
                            selectedIndex={selectedIndex}
                            handleListItemClick={handleListItemClick}
                        />
                    </SubList>
                }
            } else {
                if (item.children === undefined || item.children.length === 0) {
                    return <ListItem
                        button
                        key={item.id}
                        selected={selectedIndex === item.id}
                        onClick={(event) => {
                            handleListItemClick(event, item.id)
                        }}
                        component={Link}
                        to={item.url}
                    >
                        <ListItemIcon>
                            {/*{switchIcon(item.icon)}*/}
                        </ListItemIcon>
                        <ListItemText primary={item.name}/>
                    </ListItem>
                } else {
                    return <SubList
                        key={"subList" + item.id}
                        myKey={item.id}
                        icon={switchIcon(item.icon)}
                        text={item.name}
                        selectedIndex={selectedIndex}
                        handleListItemClick={handleListItemClick}
                    >
                        <ListTree
                            menuTree={item.children}
                            selectedIndex={selectedIndex}
                            handleListItemClick={handleListItemClick}
                        />
                    </SubList>
                }
            }


        })}
    </>)

}

function SideBar(props) {
    const [menusList, setMenusList] = React.useState([]);
    const [selectedIndex, setSelectedIndex] = React.useState();
    useEffect(() => {
        async function conectar() {
            await menus().then((res) => {
                if (res.data.code === 0){
                    setMenusList(listToTree(res.data.data))
                }
            })
        }

        conectar();


    }, []);


    const handleListItemClick = (event, index) => {
        setSelectedIndex(index)
        props.setSideBarOpen(true)
    };

    return (
        <Drawer variant="permanent" open={props.sideBarOpen}>
            <Toolbar
                sx={{
                    display: 'flex',
                    alignItems: 'center',
                    justifyContent: 'flex-end',
                    px: [1],
                }}
            />
            <List>
                <ListTree
                    selectedIndex={selectedIndex}
                    menuTree={menusList}
                    handleListItemClick={handleListItemClick}
                />
            </List>
        </Drawer>)

}

export default SideBar
