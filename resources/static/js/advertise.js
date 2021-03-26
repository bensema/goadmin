

layui.config({
    base: 'static/layuiadmin/' //静态资源所在路径
  }).extend({
    index: 'lib/index' //主入口模块
    ,goadmin: '../../js/goadmin'
  }).use(['index', 'table', 'form','goadmin', 'layer'], function(){
      var $ = layui.$
      ,table = layui.table
      ,form = layui.form
      ,layer = layui.layer;

      function start_at(d) {
         return layui.goadmin.timestampToTime(d.start_at);
      }
       function end_at(d) {
           return layui.goadmin.timestampToTime(d.end_at);
       }

      function admin_status(d) {
        if (d.status == "1"){
            return '<button class="layui-btn layui-btn-xs">开启</button>'
        }
        if (d.status == "2"){
           return '<button class="layui-btn layui-btn-primary layui-btn-xs">关闭</button>'
        }
      }

      function imgUrl(d) {
        return '<img style="display: inline-block; width: 100%; height: 100%;" src="'+d.img_url+'">'
      }

      //广告管理
      table.render({
        elem: '#LAY-app-content-list'
        ,url: layui.goadmin.bb_admin_api_advertise_pages
        ,cols: [[
          {type: 'checkbox', fixed: 'left'}
          ,{field: 'id', width: 100, title: 'ID', sort: true, align: 'center'}
          ,{field: 'title', title: '标题', minWidth: 100, align: 'center'}
          ,{field: 'img_url', title: '图片', event: 'img_event', templet: imgUrl, align: 'center'}
          ,{field: 'sort_index', title: '排序', align: 'center'}
          ,{field: 'status', title: '状态', templet: admin_status, align: 'center'}
          ,{field: 'start_at', title: '起始时间', templet: start_at, align: 'center'}
          ,{field: 'end_at', title: '结束时间', templet: end_at, align: 'center'}
          ,{title: '操作', minWidth: 150, align: 'center', fixed: 'right', toolbar: '#table-content-list'}
        ]]
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
          layer.confirm(' 确定删除？', function(index) {
              //执行 Ajax 后重载
              layui.goadmin.req({
                  type: "post",
                  url: layui.goadmin.bb_admin_api_advertise_del,
                  data: {"id": obj.data.id},
                  done: function(res) {
                      layer.alert("删除成功")
                      layui.table.reload('LAY-app-content-list'); //重载表格
                  },
              });

            })
        } else if(obj.event === 'edit'){
          layer.open({
            type: 2
            ,title: '编辑管理员'
            ,content: layui.goadmin.web_admin_form_url + '?admin_id='+ data.admin_id
            ,maxmin: true
            ,area: ['650px', '650px']
            ,btn: ['确定', '取消']
            ,yes: function(index, layero){
              var iframeWindow = window['layui-layer-iframe'+ index]
              ,submit = layero.find('iframe').contents().find("#layuiadmin-app-form-edit");
               submit.click();
            }
          });
        } else if(obj.event === 'img_event'){
            layer.photos({
                photos: { "data": [{"src": data.img_url}] },
                anim:5
            });
        }
      });


       //监听搜索
      form.on('submit(LAY-app-contlist-search)', function(data){
        var field = data.field;

        //执行重载
        table.reload('LAY-app-content-list', {
          where: field
        });
      });

      var $ = layui.$, active = {
        add: function(){
          layer.open({
            type: 2
            ,title: '添加广告'
            ,content: layui.goadmin.web_bb_advertise_add_url
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

