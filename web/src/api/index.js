/**
 * services模块根据接口文档定义接口的名称和方法
 * 根据模块拆分文件
 * index.js为出口文件，需要引入其他的services并整合
 */
import login from "./api";
import menus from "./api";
// import permission from './permission'
// import menu from './menu'
// import user from './user'

const API = Object.assign(
    {},
    login,
    menus,
);

export default API;
