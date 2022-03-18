import React from "react";
import {RequireAuth} from "@/auth/auth";

const lazyLoad = (name) => {
    const Component = React.lazy(() => import("@/pages/" + name));
    return (
        <React.Suspense fallback={<>...</>}>
            <Component/>
        </React.Suspense>
    );
};

const routes = [
    {
        path: "/",
        element: <RequireAuth>{lazyLoad("layout/Layout")}</RequireAuth>,
        children: [
            {
                index: true,
                element: lazyLoad("Dashboard")
            },
            {
                path: "/dashboard",
                element: lazyLoad("Dashboard")
            },
            {
                path: "/admins",
                element: lazyLoad("admin/Admins")
            },
            {
                path: "/admin/:adminId",
                element: lazyLoad("admin/AdminInfo")
            },
            {
                path: "/roles",
                element: lazyLoad("role/Roles")
            },
            {
                path: "/role/:roleId",
                element: lazyLoad("role/RoleInfo")
            },
            {
                path: "/permissions",
                element: lazyLoad("permission/Permissions")
            },
            {
                path: "/permission/:permissionId",
                element: lazyLoad("permission/PermissionInfo")
            },
            {
                path: "/menus",
                element: lazyLoad("menu/Menus")
            },
            {
                path: "/menu/:id",
                element: lazyLoad("menu/MenuInfo")
            },
            {
                path: "/myapis",
                element: lazyLoad("api/Apis")
            },
            {
                path: "/myapi/:id",
                element: lazyLoad("api/ApiInfo")
            },
            {
                path: "/log/login",
                element: lazyLoad("log/LogLogin")
            },
            {
                path: "/log/operation",
                element: lazyLoad("log/LogOperation")
            },
            {
                path: "*",
                element: lazyLoad("Page404")
            }
        ]
    },
    {
        path: "/sign-in",
        element: lazyLoad("SignIn")
    },

];

export default routes;
