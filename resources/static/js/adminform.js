

  layui.config({
    base: '/static/layuiadmin/' //静态资源所在路径
  }).extend({
    index: 'lib/index' //主入口模块
    ,goadmin: '../../js/goadmin'
  }).use(['index', 'form', 'goadmin', 'transfer',], function(){
    var $ = layui.$
    ,form = layui.form
    ,admin = layui.admin
    ,transfer = layui.transfer

    admin_id = getQueryVariable('admin_id')
    params = {}
    params.admin_id = getQueryVariable('admin_id')
    roles = []
    role_data = []

    // 获取账户信息
    layui.goadmin.req({
        type: 'get'
        ,url: layui.goadmin.api_admin_info_url
        ,data: params
        ,done: function(res){
            layui.$('#edit_name').val(res.data.name);
            layui.$('#edit_status').val(res.data.status);
            for (const item of res.data.roles) {
              roles.push(item.id);
            }
            // 获取全部角色
            layui.goadmin.req({
              type: 'get'
              ,url: layui.goadmin.api_role_all_url
              ,data: params
              ,done: function(res){
                 role_data = res.data;
                 for (const item of res.data) {
                      item.value = item.id
                      item.title = item.name
                 }
                  transfer.render({
                    elem: '#edit_roles'
                    ,title: ['可选角色', '已有角色']  //自定义标题
                    ,data: role_data
                    ,value: roles
                    ,id: 'edit_roles'
                  })
              }
          });

          layui.form.render();
        }
    });

    //监听提交
    form.on('submit(layuiadmin-app-form-submit)', function(data){
      var field = data.field; //获取提交的字段
      var index = parent.layer.getFrameIndex(window.name); //先得到当前iframe层的索引

      //提交 Ajax 成功后，关闭当前弹层并重载表格

      parent.layui.table.reload('LAY-app-content-list'); //重载表格
      parent.layer.close(index); //再执行关闭
    });
    //监听提交
    form.on('submit(layuiadmin-app-form-edit)', function(data){
        var index = parent.layer.getFrameIndex(window.name); //先得到当前iframe层的索引
        params = {}

        var roles = layui.transfer.getData('edit_roles');

        params.admin_id = getQueryVariable('admin_id');
        params.status = layui.$('#edit_status').val();
        params.roles = roles.map(function(elem){return elem.id;}).join(",");

         layui.goadmin.req({
             type: 'post'
             ,url: layui.goadmin.api_admin_update_url
             ,data: params
             ,done: function(res){
                layer.confirm('操作成功', {
                  btn: ['确定'] //按钮
                }, function(){
                    parent.layui.table.reload('LAY-app-content-list'); //重载表格
                    parent.layer.close(index); //再执行关闭
                });
             },
         });
    });
  })

  function getQueryVariable(variable){
      var query = window.location.search.substring(1);
      var vars = query.split("&");
      for (var i=0;i<vars.length;i++) {
              var pair = vars[i].split("=");
              if(pair[0] == variable){return pair[1];}
      }
      return(false);
   }