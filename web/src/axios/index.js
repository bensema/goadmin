// request.js
import axios from "axios";
import Qs from 'qs'
import {Tip} from "@/components/Tip"
import React from 'react';
import ReactDOM from 'react-dom';


// 显示loading
function showTip(message) {
    var dom = document.createElement('div')
    dom.setAttribute('id', 'tip')
    document.body.appendChild(dom)
    ReactDOM.render(<Tip message={message} onClose={function () {
        document.body.removeChild(document.getElementById('tip'))
        // localStorage.setItem("login", "false");
    }
    }/>, dom)

}

const service = axios.create({
    headers: {
        post: {
            "Content-Type": "application/x-www-form-urlencoded; charset=UTF-8"
        },
        get: {
            "Content-Type": "application/json"
        }
    },
    // 允许携带cookie
    withCredentials: true,
    timeout: 30000,
    //请求数据转换
    transformRequest: [
        (data) => {
            return Qs.stringify(data);
        }
    ],
    //相应数据转换
    transformResponse: [
        (data) => {
            if (typeof data === "string" && data.startsWith("{")) {
                data = JSON.parse(data);
            }
            return data;
        }
    ]
});

// 请求拦截器
service.interceptors.request.use(
    (config) => {
        // 让每个请求携带自定义 token
        config.headers["Authorization"] = "token";
        return config;
    },
    (error) => {
        // 错误抛到业务代码
        error.data = {
            code: -1,
            message: "服务器异常，请联系管理员！",
            data: {},
            t: -1
        };
        return Promise.resolve(error);
    }
);

// 响应拦截器
service.interceptors.response.use(
    (response) => {
        const status = response.status;
        if (status !== 200) {
            // 处理http错误，抛到业务代码
            let message = showStatus(status);
            response.data = {
                code: -1,
                message: message,
                data: {},
                t: -1
            };
        }
        if (response.data.code !== 0) {
            showTip(response.data.message)
        }
        return response;
    },
    (error) => {
        // 错误抛到业务代码
        error.data = {
            code: -1,
            message: "请求超时或服务器异常，请检查网络或联系管理员！",
            data: {},
            t: -1
        };
        showTip(error.data.message)
        return Promise.resolve(error);
    }
);

const codeMessage = {
    200: "服务器成功返回请求的数据。",
    201: "新建或修改数据成功。",
    202: "一个请求已经进入后台排队（异步任务）。",
    204: "删除数据成功。",
    400: "发出的请求有错误，服务器没有进行新建或修改数据的操作。",
    401: "用户没有权限（令牌、用户名、密码错误）。",
    403: "用户得到授权，但是访问是被禁止的。",
    404: "发出的请求针对的是不存在的记录，服务器没有进行操作。",
    406: "请求的格式不可得。",
    410: "请求的资源被永久删除，且不会再得到的。",
    422: "当创建一个对象时，发生一个验证错误。",
    500: "服务器发生错误，请检查服务器。",
    502: "网关错误。",
    503: "服务不可用，服务器暂时过载或维护。",
    504: "网关超时。"
};

const showStatus = (status) => {
    let message = codeMessage[status] || `连接出错(${status})!`;
    return `${message}，请检查网络或联系管理员！`;
};

export {service as axios};
