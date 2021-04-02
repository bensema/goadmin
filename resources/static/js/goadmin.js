/**

    goadmin 全局

 */

layui.define(function(exports){ //提示：模块也可以依赖其它模块，如：layui.define('layer', callback);
    var $ = layui.jquery
      ,setter = layui.setter
      ,v = "0.1.0"
      ,obj = {
        api_rsa_url : '/api/v1/rsa'  // rsa
        ,api_login_url : '/api/v1/login'  // 登陆地址
        ,api_menu_url: '/api/v1/menu'   // 菜单地址
        ,api_admin_page_url: '/api/v1/admin/pages'   // 管理员列表
        ,api_admin_info_url: '/api/v1/admin/info'   // 管理员信息
        ,api_role_all_url: '/api/v1/role/all'   // 获取全部角色
        ,api_admin_update_url: '/api/v1/admin/update'   // 修改账户
        ,api_admin_delete_url: '/api/v1/admin/delete'   // 删除用户
        ,api_admin_add_url: '/api/v1/admin/add'   // 添加用户
        ,api_role_page_url: '/api/v1/role/pages'
        ,api_role_add_url: '/api/v1/role/add'
        ,api_role_delete_url: '/api/v1/role/delete'
        ,api_role_info_url: '/api/v1/role/info'
        ,api_permission_all_url: '/api/v1/permission/all'
        ,api_menu_all_url: '/api/v1/menu/all'
        ,api_operation_all_url: '/api/v1/operation/all'
        ,api_role_update_url: '/api/v1/role/update'
        ,api_permission_page_url: '/api/v1/permission/pages'
        ,api_permission_delete_url: '/api/v1/permission/delete'
        ,api_permission_menu_url: '/api/v1/permission_menu/find'
        ,api_permission_operation_url: '/api/v1/permission_operation/find'
        ,api_permission_add_url: '/api/v1/permission/add'
        ,api_permission_update_url: '/api/v1/permission/update'
        ,api_menu_delete_url: '/api/v1/menu/delete'
        ,api_menu_update_url: '/api/v1/menu/update'
        ,api_operation_update_url: '/api/v1/operation/update'
        ,api_operation_delete_url: '/api/v1/operation/delete'
        ,api_operation_add_url: '/api/v1/operation/add'
        ,api_menu_add_url: '/api/v1/menu/add'
        ,api_log_login_page_url: '/api/v1/log_login/pages'
        ,api_log_operation_page_url: '/api/v1/log_operation/pages'

        ,bb_admin_api_advertise_pages:'/api/v1/advertise/pages'
        ,bb_admin_api_advertise_add:'/api/v1/advertise/add'
        ,bb_admin_api_advertise_del:'/api/v1/advertise/del'
        ,bb_admin_api_advertise_query: '/api/v1/advertise/query'
        ,bb_admin_api_advertise_update: '/api/v1/advertise/update'

        ,bb_admin_api_announcements_pages:'/api/v1/announcements/pages'
        ,bb_admin_api_announcements_add:'/api/v1/announcements/add'
        ,bb_admin_api_announcements_del:'/api/v1/announcements/del'
        ,bb_admin_api_announcements_query: '/api/v1/announcements/query'
        ,bb_admin_api_announcements_update: '/api/v1/announcements/update'

        ,bb_admin_api_game_pages:'/api/v1/game/pages'
        ,bb_admin_api_game_add:'/api/v1/game/add'
        ,bb_admin_api_game_del:'/api/v1/game/del'
        ,bb_admin_api_game_query: '/api/v1/game/query'
        ,bb_admin_api_game_update: '/api/v1/game/update'

        ,bb_admin_api_game_type_pages:'/api/v1/game_type/pages'
        ,bb_admin_api_game_type_add:'/api/v1/game_type/add'
        ,bb_admin_api_game_type_del:'/api/v1/game_type/del'
        ,bb_admin_api_game_type_query: '/api/v1/game_type/query'
        ,bb_admin_api_game_type_update: '/api/v1/game_type/update'

        ,bb_admin_api_game_result_pages:'/api/v1/game_result/pages'
        ,bb_admin_api_game_result_add:'/api/v1/game_result/add'
        ,bb_admin_api_game_result_del:'/api/v1/game_result/del'
        ,bb_admin_api_game_result_query: '/api/v1/game_result/query'
        ,bb_admin_api_game_result_update: '/api/v1/game_result/update'

        ,bb_admin_api_game_group_pages:'/api/v1/game_group/pages'
        ,bb_admin_api_game_group_add:'/api/v1/game_group/add'
        ,bb_admin_api_game_group_del:'/api/v1/game_group/del'
        ,bb_admin_api_game_group_query: '/api/v1/game_group/query'
        ,bb_admin_api_game_group_update: '/api/v1/game_group/update'

        ,bb_admin_api_issue_factory_pages:'/api/v1/issue_factory/pages'
        ,bb_admin_api_issue_factory_add:'/api/v1/issue_factory/add'
        ,bb_admin_api_issue_factory_del:'/api/v1/issue_factory/del'
        ,bb_admin_api_issue_factory_query: '/api/v1/issue_factory/query'
        ,bb_admin_api_issue_factory_update: '/api/v1/issue_factory/update'

        ,web_admin_form_url: '/admin/form'
        ,web_admin_add_url: '/admin/add'
        ,web_role_add_url: '/role/add'
        ,web_role_form_url: '/role/form'
        ,web_permission_add_url: '/permission/add'
        ,web_resource_add_url: '/resource/add'

        ,web_bb_advertise_add_url: '/advertise/add'
        ,web_bb_advertise_form_url: '/advertise/form'

        ,web_bb_announcement_add_url: '/announcement/add'
        ,web_bb_announcement_form_url: '/announcement/form'

        ,web_bb_game_add_url: '/game/add'
        ,web_bb_game_form_url: '/game/form'

        ,web_bb_game_result_detail_url: '/game_result/detail'

        ,timestampToTime: function(timestamp) {  // 1561953956 => yyyy-MM-dd hh:mm:ss
            var d = new Date(timestamp * 1000); //时间戳为10位需*1000，时间戳为13位的话不需乘1000
            year = d.getFullYear();
            month = ((d.getMonth()+1)>9?"":"0")+(d.getMonth()+1);
            day = (d.getDate()>9?"":"0")+d.getDate();
            hh = (d.getHours()>9?"":"0")+d.getHours();
            mm = (d.getMinutes()>9?"":"0")+d.getMinutes();
            ss = (d.getSeconds()>9?"":"0")+d.getSeconds();
            return year + "-" + month + "-" + day + " " + hh + ":" + mm + ":" + ss;
        },
        code: {
        ok: 0
        },
        response: {
          statusName: 'code' //数据状态的字段名称
          ,statusCode: {
            ok: 0 //数据状态一切正常的状态码
            ,nologin: -101 // 未登录
            ,access_denied: -403
            ,logout: 1001 //登录状态失效的状态码
          }
          ,msgName: 'message' //状态信息的字段名称
          ,dataName: 'data' //数据详情的字段名称
        },
        req: function(options){
         var that = this
         ,success = options.success
         ,error = options.error
         ,response = that.response

         options.data = options.data || {};
         options.headers = options.headers || {};


         delete options.success;
         delete options.error;

         return $.ajax($.extend({
           type: 'get'
           ,dataType: 'json'
           ,success: function(res){
             var statusCode = response.statusCode;

             //只有 response 的 code 一切正常才执行 done
             if(res[response.statusName] == statusCode.ok) {
               typeof options.done === 'function' && options.done(res);
             }

             //登录状态失效，清除本地 access_token，并强制跳转到登入页
             else if(res[response.statusName] == statusCode.logout){
               layui.layer.alert(res.message)
             }

             else if(res[response.statusName] == statusCode.nologin){
                layui.layer.alert(res.message)
                location.href = '/login';
              }

             //其它异常
             else {
               layui.layer.alert(res.message)
             }

             //只要 http 状态码正常，无论 response 的 code 是否正常都执行 success
             typeof success === 'function' && success(res);
           }
           ,error: function(e, code){
            if (typeof error !== 'function'){
               layui.layer.alert(e.responseText)
               return
            }
             typeof error === 'function' && error(res);
           }
         }, options));
       },

  };

  exports('goadmin', obj);
});