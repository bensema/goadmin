

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



    //监听提交
    form.on('submit(layuiadmin-app-form-submit)', function(data){
        var index = parent.layer.getFrameIndex(window.name); //先得到当前iframe层的索引
        params = {}
        params.name = layui.$('#add_name').val();

         layui.goadmin.req({
             type: 'post'
             ,url: layui.goadmin.api_permission_add_url
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
