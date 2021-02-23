

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


      function permissions(d) {
        text = '<div class="layui-btn-group">'
        if (d.permissions !== undefined && d.permissions !== null && d.permissions.length > 0) {
            for (const item of d.permissions) {
                text += '<button class="layui-btn layui-btn-xs">'+item.name+'</button>'
            }
        }
        text += '</div>'
        return text
      }

      //文章管理
      table.render({
        elem: '#LAY-app-content-list'
        ,url: layui.goadmin.api_role_page_url
        ,cols: [[
          {type: 'checkbox', fixed: 'left'}
          ,{field: 'id', width: 100, title: '角色ID', sort: true, align: 'center'}
          ,{field: 'name', title: '角色', align: 'center'}
          ,{field: 'permissions', title: '权限', templet: permissions, align: 'center'}
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
      });

      //监听工具条
      table.on('tool(LAY-app-content-list)', function(obj){
        var data = obj.data;
        if(obj.event === 'del'){
          layer.prompt({
              formType: 0
              ,title: '敏感操作，请手动输入被删除角色'
            }, function(value, index){
              if (value !== obj.data.name) {
                layer.alert("输入角色名与被删角色不同");
                return;
              }
              layer.close(index);

              layer.confirm(value + ' 确定删除吗？', function(index) {

                //执行 Ajax 后重载
                layui.goadmin.req({
                    type: "post",
                    url: layui.goadmin.api_role_delete_url,
                    data: {"id": obj.data.id},
                    done: function(res) {
                        layer.alert("删除成功")
                        layui.table.reload('LAY-app-content-list'); //重载表格

                    },
                });

              });
            });
        } else if(obj.event === 'edit'){
          layer.open({
            type: 2
            ,title: '编辑角色'
            ,content: layui.goadmin.web_role_form_url + '?id='+ data.id
            ,maxmin: true
            ,area: ['650px', '650px']
            ,btn: ['确定', '取消']
            ,yes: function(index, layero){
              var iframeWindow = window['layui-layer-iframe'+ index]
              ,submit = layero.find('iframe').contents().find("#layuiadmin-app-form-edit");
               submit.click();
            }
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
            ,title: '添加角色'
            ,content: layui.goadmin.web_role_add_url
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

