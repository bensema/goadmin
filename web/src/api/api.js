// 登录接口
import {axios} from "../axios";

const userApi = {
    login: "/api/sign-in",
    logout: "/api/logout",
    menus: "/api/menus",

    roleAll: "/api/role/all",
    permissionAll: "/api/permission/all",
    menuAll: "/api/menu/all",
    apiAll: "/api/api/all",

    logLoginPages: "/api/log_login/pages",
    logOperationPages: "/api/log_operation/pages",

    adminPages: "/api/admin/pages",
    adminInfo: "/api/admin/info",
    adminInfoUpdate: "/api/admin/update",
    adminDelete: "/api/admin/delete",
    adminAdd: "/api/admin/add",

    rolePages: "/api/role/pages",
    roleInfo: "/api/role/info",
    roleInfoUpdate: "/api/role/update",
    roleDelete: "/api/role/delete",
    roleAdd: "/api/role/add",

    permissionPages: "/api/permission/pages",
    permissionInfo: "/api/permission/info",
    permissionInfoUpdate: "/api/permission/update",
    permissionDelete: "/api/permission/delete",
    permissionAdd: "/api/permission/add",

    menuPages: "/api/menu/pages",
    menuInfo: "/api/menu/info",
    menuInfoUpdate: "/api/menu/update",
    menuDelete: "/api/menu/delete",
    menuAdd: "/api/menu/add",

    apiPages: "/api/api/pages",
    apiInfo: "/api/api/info",
    apiInfoUpdate: "/api/api/update",
    apiDelete: "/api/api/delete",
    apiAdd: "/api/api/add",

};

export function login(parameter) {
    return axios({
        url: userApi.login,
        method: "post",
        data: parameter
    });
}

export function logout(parameter) {
    return axios({
        url: userApi.logout,
        method: "get",
        params: parameter
    });
}

export function menus(parameter) {
    return axios({
        url: userApi.menus,
        method: "get",
        params: parameter
    });
}

export function roleAll(parameter) {
    return axios({
        url: userApi.roleAll,
        method: "get",
        params: parameter
    });
}

export function permissionAll(parameter) {
    return axios({
        url: userApi.permissionAll,
        method: "get",
        params: parameter
    });
}

export function menuAll(parameter) {
    return axios({
        url: userApi.menuAll,
        method: "get",
        params: parameter
    });
}

export function apiAll(parameter) {
    return axios({
        url: userApi.apiAll,
        method: "get",
        params: parameter
    });
}


// admin
export function adminPages(parameter) {
    return axios({
        url: userApi.adminPages,
        method: "get",
        params: parameter
    });
}

export function adminInfo(parameter) {
    return axios({
        url: userApi.adminInfo,
        method: "get",
        params: parameter
    });
}

export function adminInfoUpdate(parameter) {
    return axios({
        url: userApi.adminInfoUpdate,
        method: "post",
        data: parameter
    });
}

export function adminDelete(parameter) {
    return axios({
        url: userApi.adminDelete,
        method: "post",
        data: parameter
    });
}

export function adminAdd(parameter) {
    return axios({
        url: userApi.adminAdd,
        method: "post",
        data: parameter
    });
}

// role
export function rolePages(parameter) {
    return axios({
        url: userApi.rolePages,
        method: "get",
        params: parameter
    });
}

export function roleInfo(parameter) {
    return axios({
        url: userApi.roleInfo,
        method: "get",
        params: parameter
    });
}

export function roleInfoUpdate(parameter) {
    return axios({
        url: userApi.roleInfoUpdate,
        method: "post",
        data: parameter
    });
}

export function roleDelete(parameter) {
    return axios({
        url: userApi.roleDelete,
        method: "post",
        data: parameter
    });
}

export function roleAdd(parameter) {
    return axios({
        url: userApi.roleAdd,
        method: "post",
        data: parameter
    });
}


// permission
export function permissionPages(parameter) {
    return axios({
        url: userApi.permissionPages,
        method: "get",
        params: parameter
    });
}

export function permissionInfo(parameter) {
    return axios({
        url: userApi.permissionInfo,
        method: "get",
        params: parameter
    });
}

export function permissionInfoUpdate(parameter) {
    return axios({
        url: userApi.permissionInfoUpdate,
        method: "post",
        data: parameter
    });
}

export function permissionDelete(parameter) {
    return axios({
        url: userApi.permissionDelete,
        method: "post",
        data: parameter
    });
}

export function permissionAdd(parameter) {
    return axios({
        url: userApi.permissionAdd,
        method: "post",
        data: parameter
    });
}

// menu
export function menuPages(parameter) {
    return axios({
        url: userApi.menuPages,
        method: "get",
        params: parameter
    });
}

export function menuInfo(parameter) {
    return axios({
        url: userApi.menuInfo,
        method: "get",
        params: parameter
    });
}

export function menuInfoUpdate(parameter) {
    return axios({
        url: userApi.menuInfoUpdate,
        method: "post",
        data: parameter
    });
}

export function menuDelete(parameter) {
    return axios({
        url: userApi.menuDelete,
        method: "post",
        data: parameter
    });
}

export function menuAdd(parameter) {
    return axios({
        url: userApi.menuAdd,
        method: "post",
        data: parameter
    });
}

// api
export function apiPages(parameter) {
    return axios({
        url: userApi.apiPages,
        method: "get",
        params: parameter
    });
}

export function apiInfo(parameter) {
    return axios({
        url: userApi.apiInfo,
        method: "get",
        params: parameter
    });
}

export function apiInfoUpdate(parameter) {
    return axios({
        url: userApi.apiInfoUpdate,
        method: "post",
        data: parameter
    });
}

export function apiDelete(parameter) {
    return axios({
        url: userApi.apiDelete,
        method: "post",
        data: parameter
    });
}

export function apiAdd(parameter) {
    return axios({
        url: userApi.apiAdd,
        method: "post",
        data: parameter
    });
}


// logs
export function logLoginPages(parameter) {
    return axios({
        url: userApi.logLoginPages,
        method: "get",
        params: parameter
    });
}

export function logOperationPages(parameter) {
    return axios({
        url: userApi.logOperationPages,
        method: "get",
        params: parameter
    });
}
