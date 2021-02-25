

layui.config({
    base: 'static/layuiadmin/' //静态资源所在路径
  }).extend({
    index: 'lib/index' //主入口模块
    ,goadmin: '../../js/goadmin'
  }).use(['index', 'table', 'form','goadmin', 'layer'], function(){
      var table = layui.table
      ,form = layui.form
      ,layer = layui.layer;

//      treeInit()

      var tree_all = [];
      var tree_setting = {
          view: {
//            addDiyDom: addDiyDom,
            addHoverDom: addHoverDom,
            removeHoverDom: removeHoverDom,
          },
          check: {
              enable: true
          },
          data: {
              simpleData: {
                  enable: true,
                  idKey: "id",
                  pIdKey: "pid",
              },
              key: {
                  name: "name",
                  url: "nourl",

              },
              view: {
                  showIcon: true,
              }
          }
      };

      table.render({
        elem: '#LAY-app-content-list'
        ,url: layui.goadmin.api_permission_page_url
        ,cols: [[
          {type: 'checkbox', fixed: 'left'}
          ,{field: 'id', width: 100, title: '权限ID', sort: true, align: 'center'}
          ,{field: 'name', title: '权限', align: 'center'}
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
              ,title: '敏感操作，请手动输入被删除权限'
            }, function(value, index){
              if (value !== obj.data.name) {    
                layer.alert("输入权限名与被删权限不同");
                return;
              }
              layer.close(index);

              layer.confirm(value + ' 确定删除吗？', function(index) {

                //执行 Ajax 后重载
                layui.goadmin.req({
                    type: "post",
                    url: layui.goadmin.api_permission_delete_url,
                    data: {"id": obj.data.id},
                    done: function(res) {
                        layer.alert("删除成功")
                        layui.table.reload('LAY-app-content-list'); //重载表格

                    },
                });

              });
            });
        } else if(obj.event === 'edit'){
          layui.$('#edit_cur_id').val(obj.data.id);
          layui.$('#edit_cur_name').val(obj.data.name);
          $('#save_btn').removeClass("layui-btn-disabled").attr("disabled",false);


          treeInit(function(){
              params = {}
              params.id = obj.data.id
              layui.goadmin.req({
                  type: 'get'
                  ,url: layui.goadmin.api_permission_menu_url
                  ,data: params
                  ,done: function(res){
                     menu_data = res.data;
                     for (const item of menu_data) {
                          item.pid = 'menu_' + item.pid
                          item.menu_id = 'menu_' + item.menu_id
                          item.icon = ''
                          item.permission_type = 'menu'
                     }
                     menu_select = menu_data.map(function(elem){return elem.menu_id;});

                     layui.goadmin.req({
                         type: 'get'
                         ,url: layui.goadmin.api_permission_operation_url
                         ,data: params
                         ,done: function(res){
                            operation_data = res.data;
                            for (const item of operation_data) {
                                 item.pid = 'menu_' + item.pid
                                 item.operation_id = 'operation_' + item.operation_id
                                 item.permission_type = 'operation'
                            }
                            operation_select = operation_data.map(function(elem){return elem.operation_id;});

                            select_ids  = menu_select.concat(operation_select)
                            zNodes = init_checked(tree_all,select_ids);
                            $.fn.zTree.init($("#edit_tree"), tree_setting, zNodes);
                            var zTree = $.fn.zTree.getZTreeObj("edit_tree");
                            zTree.expandAll(true);
                         }
                     });

                  }
                      });
          })

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

      var active = {
        add: function(){
          layer.open({
            type: 2
            ,title: '添加权限'
            ,content: layui.goadmin.web_permission_add_url
            ,maxmin: true
            ,area: ['550px', '550px']
            ,btn: ['确定', '取消']
            ,yes: function(index, layero){
              //点击确认触发 iframe 内容中的按钮提交
              var submit = layero.find('iframe').contents().find("#layuiadmin-app-form-submit");
              submit.click();
            }
          });
        },
        save: function(){
            var treeObj=$.fn.zTree.getZTreeObj("edit_tree")
            var nodes=treeObj.getCheckedNodes(true);
            var menus = [];
            var operations = [];

            for (const item of nodes){
                if (item.permission_type === "menu"){
                    menus.push(item.id.substring(5))
                }
                if (item.permission_type === "operation"){
                    operations.push(item.id.substring(10))
                }
            }

            var params = {};
            params.id = layui.$('#edit_cur_id').val();
            params.menus = menus.join(',')
            params.operations = operations.join(',')
            layui.goadmin.req({
                type: 'post'
                ,url: layui.goadmin.api_permission_update_url
                ,data: params
                ,done: function(res){
                    layui.layer.alert("操作成功")
                }
            });
        },
      };

      layui.$('.layui-btn.layuiadmin-btn-list').on('click', function(){
        var type = layui.$(this).data('type');
        active[type] ? active[type].call(this) : '';
      });

      function treeInit(callback){
        layui.goadmin.req({
            type: 'get'
            ,url: layui.goadmin.api_menu_all_url
            ,data: {}
            ,done: function(res){
               menu_data = res.data;
               for (const item of menu_data) {
                    item.id = 'menu_' + item.id
                    item.pid = 'menu_' + item.pid
                    item.icon = ''
                    item.permission_type = 'menu'
               }

               layui.goadmin.req({
                   type: 'get'
                   ,url: layui.goadmin.api_operation_all_url
                   ,data: {}
                   ,done: function(res){
                      operation_data = res.data;
                      for (const item of operation_data) {
                           item.id = 'operation_' + item.id
                           item.pid = 'menu_' + item.pid
                           item.permission_type = 'operation'
                      }
                      tree_all = menu_data.concat(operation_data)
                      zNodes = init_checked(tree_all,[]);
                      $.fn.zTree.init($("#edit_tree"), tree_setting, zNodes);
                      var zTree = $.fn.zTree.getZTreeObj("edit_tree");
                      zTree.expandAll(true);
                      typeof callback === 'function' && callback(tree_all);
                   }
               });

            }
        });
      }

    function init_checked(zNodes, checked) {
        $.each(zNodes,function(k,v){
            v.checked = false;

        });
        $.each(zNodes,function(k,v){
            if($.inArray(v.id,checked)>=0){
                v.checked = true;
            }

        });
        return  zNodes;
    }
    function addHoverDom(treeId, treeNode) {
        var aObj = $("#" + treeNode.tId + "_a");
//        if ($("#diyBtn_"+treeNode.id).length>0) return;
        if ($("#diyBtn_space_"+treeNode.id).length>0) return;
        var editStr = "<span id='diyBtn_space_" +treeNode.id+ "' >" + treeNode.url + "</span>"
//            + "<button type='button' class='diyBtn1' id='diyBtn_" + treeNode.id
//            + "' title='"+treeNode.name+"' onfocus='this.blur();'>" + treeNode.permission_type + "</button>";
        aObj.append(editStr);
//        var btn = $("#diyBtn_"+treeNode.id);
//        if (btn) btn.bind("click", function(){alert("diy Button for " + treeNode.name);});
    };
    function removeHoverDom(treeId, treeNode) {
        $("#diyBtn_"+treeNode.id).unbind().remove();
        $("#diyBtn_space_" +treeNode.id).unbind().remove();
    };
    function addDiyDom(treeId, treeNode) {
        var aObj = $("#" + treeNode.tId + "_a");
        var editStr = "<span id='diyBtn_space_" +treeNode.id+ "' >" + treeNode.url + "</span>"
            + "<button type='button' class='diyBtn1' id='diyBtn_" + treeNode.id
            + "' title='"+treeNode.name+"' onfocus='this.blur();'>" + treeNode.permission_type + "</button>";
        aObj.append(editStr);
    };
  });

