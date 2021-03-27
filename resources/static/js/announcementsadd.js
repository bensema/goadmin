

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

      upload.render({
        elem: '#test8'
//        ,url: 'https://httpbin.org/post' //改成您自己的上传接口
        ,url: 'https://www.niupic.com/api/upload' //改成您自己的上传接口
        ,auto: false
        //,multiple: true
        ,bindAction: '#test9'
        ,done: function(res){
          layer.msg('上传成功');
          layui.$('#add_img_url').val(res.data);
          console.log(res)
        }
      });

    roles = []
    role_data = []


    //监听提交
    form.on('submit(layuiadmin-app-form-submit)', function(data){
        var index = parent.layer.getFrameIndex(window.name); //先得到当前iframe层的索引
        params = {}

        params.title = layui.$('#add_title').val();
        params.content = layui.$('#add_content').val();
        params.img_url = layui.$('#add_img_url').val();
        params.sort_index = layui.$('#add_sort_index').val();
        params.status = layui.$('#add_status').val();
        params.hot = layui.$('#add_hot').val();
        params.popup = layui.$('#add_popup').val();

        params.start_at = parseDateTime(layui.$('#add_start_at').val())
        params.end_at = parseDateTime(layui.$('#add_end_at').val())

         layui.goadmin.req({
             type: 'post'
             ,url: layui.goadmin.bb_admin_api_announcement_add
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
