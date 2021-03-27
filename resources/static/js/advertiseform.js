

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

    // 获取账户信息
    layui.goadmin.req({
        type: 'get'
        ,url: layui.goadmin.bb_admin_api_advertise_query
        ,data: params
        ,done: function(res){
            layui.$('#edit_title').val(res.data.title);
            layui.$('#edit_content').val(res.data.content);
            layui.$('#edit_img_url').val(res.data.img_url);
            layui.$('#edit_link_url').val(res.data.link_url);
            layui.$('#edit_sort_index').val(res.data.sort_index);
            layui.$('#edit_status').val(res.data.status);
            layui.$('#edit_remark').val(res.data.remark);
            layui.$('#edit_popup').val(res.data.popup);


            laydate.render({
                elem: '#edit_start_at'
                ,type: 'datetime'
                ,value: layui.goadmin.timestampToTime(res.data.start_at)
            });
           laydate.render({
                elem: '#edit_end_at'
                ,type: 'datetime'
                ,value: layui.goadmin.timestampToTime(res.data.end_at)
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
        params.id=id;
        params.title = layui.$('#edit_title').val();
        params.content = layui.$('#edit_content').val();
        params.img_url = layui.$('#edit_img_url').val();
        params.link_url = layui.$('#edit_link_url').val();
        params.sort_index = layui.$('#edit_sort_index').val();
        params.status = layui.$('#edit_status').val();
        params.remark = layui.$('#edit_remark').val();
        params.popup = layui.$('#edit_popup').val();

        params.start_at = parseDateTime(layui.$('#edit_start_at').val())
        params.end_at = parseDateTime(layui.$('#edit_end_at').val())

         layui.goadmin.req({
             type: 'post'
             ,url: layui.goadmin.bb_admin_api_advertise_update
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