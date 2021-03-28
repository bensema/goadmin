

  layui.config({
    base: '/static/layuiadmin/' //静态资源所在路径
  }).extend({
    index: 'lib/index' //主入口模块
    ,goadmin: '../../js/goadmin'
  }).use(['index', 'form', 'goadmin', 'transfer', 'laydate'], function(){
    var $ = layui.$
    ,form = layui.form
    ,admin = layui.admin
    ,transfer = layui.transfer
    ,laydate = layui.laydate

    id = getQueryVariable('id')
    params = {}
    params.id = getQueryVariable('id')
    roles = []
    role_data = []

    layui.goadmin.req({
      type: 'get'
      ,url: layui.goadmin.bb_admin_api_game_type_pages
      ,data: {"num":1,"size":100}
      ,done: function(res){
         $.each(res.data.data, function(index, item) {
              $('#edit_game_type').append(new Option(item.game_type_display_name, item.game_type));
         });

         layui.goadmin.req({
           type: 'get'
           ,url: layui.goadmin.bb_admin_api_game_group_pages
           ,data: {"num":1,"size":100}
           ,done: function(res){
              $.each(res.data.data, function(index, item) {
                   $('#edit_game_group').append(new Option(item.game_group_display_name, item.game_group));
              });

              layui.goadmin.req({
                  type: 'get'
                  ,url: layui.goadmin.bb_admin_api_game_query
                  ,data: params
                  ,done: function(res){
                    layui.$('#edit_name').val(res.data.name);
                    layui.$('#edit_display_name').val(res.data.display_name);
                    layui.$('#edit_game_code').val(res.data.game_code);
                    layui.$('#edit_game_type').val(res.data.game_type);
                    layui.$('#edit_game_group').val(res.data.game_group);
                    layui.$('#edit_sort_index').val(res.data.sort_index);
                    layui.$('#edit_status').val(res.data.status);
                    layui.$('#edit_remark').val(res.data.remark);


                      layui.form.render();
                  }
              });
           }
         });
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
        params.id=id;
        params.name = layui.$('#edit_name').val();
        params.display_name = layui.$('#edit_display_name').val();
        params.game_code = layui.$('#edit_game_code').val();
        params.game_type = layui.$('#edit_game_type').val();
        params.game_group = layui.$('#edit_game_group').val();
        params.sort_index = layui.$('#edit_sort_index').val();
        params.status = layui.$('#edit_status').val();
        params.remark = layui.$('#edit_remark').val();

         layui.goadmin.req({
             type: 'post'
             ,url: layui.goadmin.bb_admin_api_game_update
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