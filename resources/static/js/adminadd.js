

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

    roles = []
    role_data = []


    layui.goadmin.req({
      type: 'get'
      ,url: layui.goadmin.api_role_all_url
      ,data: {}
      ,done: function(res){
         role_data = res.data;
         for (const item of res.data) {
              item.value = item.id
              item.title = item.name
         }
          transfer.render({
            elem: '#add_roles'
            ,title: ['可选角色', '已有角色']  //自定义标题
            ,data: role_data
            ,value: roles
            ,id: 'add_roles'
          })
      }
  });

  layui.form.render();


    //监听提交
    form.on('submit(layuiadmin-app-form-submit)', function(data){
        var index = parent.layer.getFrameIndex(window.name); //先得到当前iframe层的索引
        params = {}

        var roles = layui.transfer.getData('add_roles');

        params.name = layui.$('#add_name').val();
        params.password = layui.$('#add_password').val();
        params.status = layui.$('#add_status').val();
        params.roles = roles.map(function(elem){return elem.id;}).join(",");

         layui.goadmin.req({
             type: 'post'
             ,url: layui.goadmin.api_admin_add_url
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
