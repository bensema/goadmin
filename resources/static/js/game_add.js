

  layui.config({
    base: '/static/layuiadmin/' //静态资源所在路径
  }).extend({
    index: 'lib/index' //主入口模块
    ,goadmin: '../../js/goadmin'
  }).use(['index', 'form', 'goadmin', 'transfer','laydate','upload'], function(){
    var $ = layui.$
    ,form = layui.form
    ,admin = layui.admin
    ,transfer = layui.transfer
    ,upload = layui.upload
    ,laydate = layui.laydate;

    laydate.render({
        elem: '#add_start_at'
        ,type: 'datetime'
        ,value: get_current_day_begin_time()
    });
   laydate.render({
        elem: '#add_end_at'
        ,type: 'datetime'
        ,value: get_current_day_end_time()
    });

    roles = []
    role_data = []

    layui.goadmin.req({
      type: 'get'
      ,url: layui.goadmin.bb_admin_api_game_type_pages
      ,data: {"num":1,"size":100}
      ,done: function(res){
         $.each(res.data.data, function(index, item) {
              $('#add_game_type').append(new Option(item.game_type_display_name, item.game_type));
         });

         layui.goadmin.req({
           type: 'get'
           ,url: layui.goadmin.bb_admin_api_game_group_pages
           ,data: {"num":1,"size":100}
           ,done: function(res){
              $.each(res.data.data, function(index, item) {
                   $('#add_game_group').append(new Option(item.game_group_display_name, item.game_group));
              });
              layui.form.render();

           }
         });
      }
    });


    //监听提交
    form.on('submit(layuiadmin-app-form-submit)', function(data){
        var index = parent.layer.getFrameIndex(window.name); //先得到当前iframe层的索引
        params = {}

        params.name = layui.$('#add_name').val();
        params.display_name = layui.$('#add_display_name').val();
        params.game_code = layui.$('#add_game_code').val();
        params.game_type = layui.$('#add_game_type').val();
        params.game_group = layui.$('#add_game_group').val();
        params.sort_index = layui.$('#add_sort_index').val();
        params.status = layui.$('#add_status').val();
        params.remark = layui.$('#add_remark').val();


         layui.goadmin.req({
             type: 'post'
             ,url: layui.goadmin.bb_admin_api_game_add
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
