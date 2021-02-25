

layui.config({
    base: '/static/layuiadmin/' //静态资源所在路径
  }).extend({
    index: 'lib/index' //主入口模块
    ,goadmin: '../../js/goadmin'
  }).use(['index', 'table', 'form','goadmin', 'layer', 'laydate'], function(){
      var $ = layui.$
      ,table = layui.table
      ,form = layui.form
      ,layer = layui.layer
      ,laydate = layui.laydate;

        laydate.render({
            elem: '#query_begin_time'
            ,type: 'datetime'
            ,value: get_current_day_begin_time()
        });
       laydate.render({
            elem: '#query_end_time'
            ,type: 'datetime'
            ,value: get_current_day_end_time()
        });

      function record_at(d) {
         return layui.goadmin.timestampToTime(d.record_at);
      }

      function result(d) {
        if (d.result === 1){
            return '<button class="layui-btn layui-btn-xs">成功</button>'
        } else {
            return '<button class="layui-btn layui-btn-primary layui-btn-xs">失败</button>'
        }
      }

      function roles(d) {
        text = '<div class="layui-btn-group">'
        if (d.roles !== undefined && d.roles !== null && d.roles.length > 0) {
            for (const item of d.roles) {
                text += '<button class="layui-btn layui-btn-xs">'+item.name+'</button>'
            }
        }
        text += '</div>'
        return text
      }

      var field = {};
      var start_at = layui.$('#query_begin_time').val();
        if (start_at !== ""){
            field.record_at_from = parseDateTime(start_at)
        }else{
            field.record_at_from = parseDateTime(get_current_day_begin_time())
        }

        var end_at = layui.$('#query_end_time').val();
        if (end_at !== ""){
            field.record_at_to = parseDateTime(end_at)
        }else{
            field.record_at_to = parseDateTime(get_current_day_end_time())
        }
        field.sort = "desc"
        field.order_by = "id"

      table.render({
        elem: '#LAY-app-content-list'
        ,url: layui.goadmin.api_log_login_page_url
        ,cols: [[
          {type: 'checkbox', fixed: 'left'}
          ,{field: 'id', width: 100, title: '编号', sort: true, align: 'center'}
          ,{field: 'name', title: '账户', minWidth: 100, align: 'center'}
          ,{field: 'location', title: '位置', align: 'center'}
          ,{field: 'os', title: '系统', align: 'center'}
          ,{field: 'browser', title: '浏览器', align: 'center'}
          ,{field: 'ip', title: 'IP', align: 'center'}
          ,{field: 'result', title: '状态', templet: result, align: 'center'}
          ,{field: 'record_at', title: '记录时间', templet: record_at, align: 'center'}
        ]]
        ,where: field
        ,page: true
        ,limit: 10
        ,limits: [10, 15, 20, 25, 30]
        ,text: '对不起，加载出现异常！'
        ,request: {
              pageName: 'num' //页码的参数名称，默认：page
              ,limitName: 'size' //每页数据量的参数名，默认：limit
            }
          ,parseData: function(res){ //res 即为原始返回的数据
              return {
                  "code": res.code, //解析接口状态
                  "msg": res.message, //解析提示文本
                  "count": res.data.total, //解析数据长度
                  "data": res.data.data //解析数据列表
              };
          }
          ,done : function(res, curr, count){
            if (res.count == 0)
            {
                $(".layui-table-main").html('<div class="layui-none">暂无数据</div>');
            }
         }
      });

      //监听工具条
      table.on('tool(LAY-app-content-list)', function(obj){
        var data = obj.data;
        if(obj.event === 'del'){

        }
      });


       //监听搜索
      form.on('submit(LAY-app-contlist-search)', function(data){
        var field = data.field;
        var start_at = layui.$('#query_begin_time').val();
          if (start_at !== ""){
              field.record_at_from = parseDateTime(start_at)
          }

          var end_at = layui.$('#query_end_time').val();
          if (end_at !== ""){
              field.record_at_to = parseDateTime(end_at)
          }
          field.sort = "desc"
          field.order_by = "id"

        //执行重载
        table.reload('LAY-app-content-list', {
          where: field
        });
      });

      var $ = layui.$, active = {
        add: function(){
          layer.open({
            type: 2
            ,title: '添加账户'
            ,content: layui.goadmin.web_admin_add_url
            ,maxmin: true
            ,area: ['650px', '650px']
            ,btn: ['确定', '取消']
            ,yes: function(index, layero){
              //点击确认触发 iframe 内容中的按钮提交
              var submit = layero.find('iframe').contents().find("#layuiadmin-app-form-submit");
              submit.click();
            }
          });
        }
      };

      $('.layui-btn.layuiadmin-btn-list').on('click', function(){
        var type = $(this).data('type');
        active[type] ? active[type].call(this) : '';
      });

  });

