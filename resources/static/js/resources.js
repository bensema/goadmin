// 菜单 api 资源
var ClickObj = {};
var MenuClickObj = {};
layui.config({
    base: 'static/layuiadmin/' //静态资源所在路径
  }).extend({
    index: 'lib/index' //主入口模块
    ,goadmin: '../../js/goadmin'
  }).use(['index', 'table', 'form','goadmin', 'layer', 'element'], function(){
      var table = layui.table
      ,form = layui.form
      ,layer = layui.layer;

      treeInit()

      var tree_all = [];
      var tree_setting = {
          view: {
            addDiyDom: addDiyDom,
//            addHoverDom: addHoverDom,
//            removeHoverDom: removeHoverDom,
            fontCss : {width:"90%",padding:"5px"},

          },
          callback:{
              onMouseDown: zTreeOnClick,
          },
          check: {
              enable: false
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

      var menu_tree_setting = {
        view: {
                selectedMulti: false //按住ctrl是否可以多选
            },
            callback:{
                onCheck: menu_zTreeOnClick,
            },
            check: {
                enable: true,
                chkStyle: "radio",
                radioType: "all"
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
            ,content: layui.goadmin.web_resource_add_url
            ,maxmin: true
            ,area: ['650px', '650px']
            ,btn: ['确定', '取消']
            ,yes: function(index, layero){
              //点击确认触发 iframe 内容中的按钮提交
              var submit = layero.find('iframe').contents().find("#layuiadmin-app-form-submit");
              submit.click();
            }
          });
        },
        refresh: function(){
            treeInit()
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
                    layui.layer.msg("操作成功")
                }
            });
        },
        pid_add_api: function(){
              layer.open({
                type: 1
                ,title: '选择菜单'
                ,content: $('.menu_content_wrap')
                ,maxmin: true
                ,area: ['550px', '550px']
                ,btn: ['确定', '取消']
                ,yes: function(index, layero){
                   if (typeof MenuClickObj.id === 'undefined'){
                       layer.alert("请选择")
                       return
                   }
                  layui.$('#add_operation_pid').val(MenuClickObj.id);
                  layui.$('#add_operation_pid_name').val(MenuClickObj.name);

                  layer.close(index);
                }
                ,end: function(){
                    $('.menu_content_wrap').hide()
                }
              });

        },
        pid_add_menu: function(){
              layer.open({
                type: 1
                ,title: '选择菜单'
                ,content: $('.menu_content_wrap')
                ,maxmin: true
                ,area: ['550px', '550px']
                ,btn: ['确定', '取消']
                ,yes: function(index, layero){
                   if (typeof MenuClickObj.id === 'undefined'){
                       layer.alert("请选择")
                       return
                   }
                  layui.$('#add_menu_pid').val(MenuClickObj.id);
                  layui.$('#add_menu_pid_name').val(MenuClickObj.name);

                  layer.close(index);
                }
                ,end: function(){
                    $('.menu_content_wrap').hide()
                }
              });

        },
        edit_menu_set_no_pid: function(){
            layui.$('#edit_menu_pid').val("0");
            layui.$('#edit_menu_pid_name').val("无上级");
        },
        pid_edit_menu: function(){
              layer.open({
                type: 1
                ,title: '选择菜单'
                ,content: $('.menu_content_wrap')
                ,maxmin: true
                ,area: ['550px', '550px']
                ,btn: ['确定', '取消']
                ,yes: function(index, layero){
                   if (typeof MenuClickObj.id === 'undefined'){
                       layer.alert("请选择")
                       return
                   }
                  layui.$('#edit_menu_pid').val(MenuClickObj.id);
                  layui.$('#edit_menu_pid_name').val(MenuClickObj.name);

                  layer.close(index);
                }
                ,end: function(){
                    $('.menu_content_wrap').hide()
                }
              });

        },
        pid_edit_operation: function(){
              layer.open({
                type: 1
                ,title: '选择菜单'
                ,content: $('.menu_content_wrap')
                ,maxmin: true
                ,area: ['550px', '550px']
                ,btn: ['确定', '取消']
                ,yes: function(index, layero){
                   if (typeof MenuClickObj.id === 'undefined'){
                       layer.alert("请选择")
                       return
                   }
                  layui.$('#edit_operation_pid').val(MenuClickObj.id);
                  layui.$('#edit_operation_pid_name').val(MenuClickObj.name);

                  layer.close(index);
                }
                ,end: function(){
                    $('.menu_content_wrap').hide()
                }
              });

        },
        add_operation_reset: function(){
            layui.$('#add_operation_pid').val("")
            layui.$('#add_operation_pid_name').val("")
            layui.$('#add_operation_name').val("")
            layui.$('#add_operation_code').val("")
            layui.$('#add_operation_method').val("")
            layui.$('#add_operation_url').val("")
        },
        add_operation: function(){
            params = {}
            params.name = layui.$('#add_operation_name').val()
            params.code = layui.$('#add_operation_code').val()
            params.method = layui.$('#add_operation_method').val()
            params.url = layui.$('#add_operation_url').val()
            params.pid = layui.$('#add_operation_pid').val().substring(5)

            layui.goadmin.req({
                type: 'post'
                ,url: layui.goadmin.api_operation_add_url
                ,data: params
                ,done: function(res){
                    layui.layer.msg("操作成功")
                    treeInit()
                    layui.$('#add_operation_name').val("")
                    layui.$('#add_operation_code').val("")
                    layui.$('#add_operation_method').val("")
                    layui.$('#add_operation_url').val("")
                }
            });
        },
        add_menu_reset: function(){
            layui.$('#add_menu_pid').val("")
            layui.$('#add_menu_pid_name').val("")
            layui.$('#add_menu_name').val("")
            layui.$('#add_menu_icon').val("")
            layui.$('#add_menu_index_sort').val("")
            layui.$('#add_menu_url').val("")
        },
        add_menu: function(){
            params = {}
            params.name = layui.$('#add_menu_name').val()
            params.pid = layui.$('#add_menu_pid').val().substring(5)
            params.icon = layui.$('#add_menu_icon').val()
            params.url = layui.$('#add_menu_url').val()
            params.sort_index = layui.$('#add_menu_sort_index').val()

            layui.goadmin.req({
                type: 'post'
                ,url: layui.goadmin.api_menu_add_url
                ,data: params
                ,done: function(res){
                    layui.layer.msg("操作成功")
                    treeInit()
                    layui.$('#add_menu_name').val("")
                    layui.$('#add_menu_icon').val("")
                    layui.$('#add_menu_index_sort').val("")
                    layui.$('#add_menu_url').val("")
                }
            });
        },
        update_menu: function(){
            params = {}

            params.id = layui.$('#edit_cur_id').val().substring(5)
            params.pid = layui.$('#edit_menu_pid').val().substring(5)
            params.name = layui.$('#edit_menu_name').val()
            params.icon = layui.$('#edit_menu_icon').val()
            params.url = layui.$('#edit_menu_url').val()
            params.index_sort = layui.$('#edit_menu_index_sort').val()

            layui.goadmin.req({
                type: 'post'
                ,url: layui.goadmin.api_menu_update_url
                ,data: params
                ,done: function(res){
                    layui.layer.msg("操作成功")
                    treeInit()
                    layui.$('#edit_cur_id').val("")
                    layui.$('#edit_menu_pid').val("")
                    layui.$('#edit_menu_pid_name').val("")
                    layui.$('#edit_menu_name').val("")
                    layui.$('#edit_menu_icon').val("")
                    layui.$('#edit_menu_index_sort').val("")
                    layui.$('#edit_menu_url').val("")
                }
            });
        },
        update_operation: function(){
            params = {}

            params.id = layui.$('#edit_cur_id').val().substring(10)
            params.pid = layui.$('#edit_operation_pid').val().substring(5)
            params.name = layui.$('#edit_operation_name').val()
            params.code = layui.$('#edit_operation_code').val()
            params.url = layui.$('#edit_operation_url').val()
            params.method = layui.$('#edit_operation_method').val()

            layui.goadmin.req({
                type: 'post'
                ,url: layui.goadmin.api_operation_update_url
                ,data: params
                ,done: function(res){
                    layui.layer.msg("操作成功")
                    treeInit()
                    layui.$('#edit_cur_id').val("")
                    layui.$('#edit_operation_pid').val("")
                    layui.$('#edit_operation_pid_name').val("")
                    layui.$('#edit_operation_name').val("")
                    layui.$('#edit_operation_code').val("")
                    layui.$('#edit_operation_method').val("")
                    layui.$('#edit_operation_url').val("")
                }
            });
        },

      };

      layui.$('.layui-btn.layuiadmin-btn-list').on('click', function(){
        var type = layui.$(this).data('type');
        active[type] ? active[type].call(this) : '';
      });

      layui.$(document).on('click',"#do_delete",function(){
            resource_type = ClickObj.permission_type;
            params = {};
            url = '';
            tip = '';
            if (resource_type === "menu"){
                tip = '确定删除菜单:'
                params.id = ClickObj.id.substring(5)
                url = layui.goadmin.api_menu_delete_url
            } else if (resource_type === "operation"){
               tip = '确定删除API功能:'
                params.id = ClickObj.id.substring(10)
                url = layui.goadmin.api_operation_delete_url
            }

              layer.confirm( tip+ ClickObj.name, function(index){

              layui.goadmin.req({
                  type: "post",
                  url: url,
                  data: params,
                  done: function(res) {
                      layer.msg("操作成功");
                      treeInit();
                  }
              });
              layer.close(index);
          });
      });

      layui.$(document).on('click',"#do_edit_menu",function(){
            layui.element.tabChange('demo', 'tab_edit_menu');
            layui.$('#edit_cur_id').val(ClickObj.id)
            layui.$('#edit_menu_pid').val(ClickObj.pid)
            layui.$('#edit_menu_pid_name').val('上级id:'+ClickObj.pid)
            layui.$('#edit_menu_name').val(ClickObj.name)
            layui.$('#edit_menu_icon').val(ClickObj._icon)
            layui.$('#edit_menu_index_sort').val(ClickObj.index_sort)
            layui.$('#edit_menu_url').val(ClickObj.url)
        });

        layui.$(document).on('click',"#do_edit_operation",function(){
            layui.element.tabChange('demo', 'tab_edit_operation');
            layui.$('#edit_cur_id').val(ClickObj.id)
            layui.$('#edit_operation_pid').val(ClickObj.pid)
            layui.$('#edit_operation_pid_name').val('上级id:'+ClickObj.pid)
            layui.$('#edit_operation_name').val(ClickObj.name)
            layui.$('#edit_operation_code').val(ClickObj.code)
            layui.$('#edit_operation_method').val(ClickObj.method)
            layui.$('#edit_operation_url').val(ClickObj.url)
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
                    item._icon = item.icon
                    item.icon = ''
                    item.permission_type = 'menu'
               }
               menu_tree_all = menu_data
               zNodes = init_checked(menu_tree_all,[]);
                 $.fn.zTree.init($("#tree_menu"), menu_tree_setting, zNodes);
                 var menu_zTree = $.fn.zTree.getZTreeObj("tree_menu");
                 menu_zTree.expandAll(true);

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

    function addDiyDom(treeId, treeNode) {
        var aObj = $("#" + treeNode.tId + "_a");
        var editStr = "<span class='layui-badge-rim'>" +  treeNode.id + "</span>" +"<span id='diyBtn_space_" +treeNode.id+ "' >" + treeNode.url + "</span>"
            +'<button class=" list_button button3"  id="do_delete" >删除</button>'
        if(treeNode.permission_type==="menu"){
            editStr += '<button class=" list_button button1" id="do_edit_menu">编辑菜单</button>'
        }
        if(treeNode.permission_type==="operation"){
            editStr += '<button class=" list_button button1" id="do_edit_operation">编辑API</button>'
        }
        aObj.append(editStr);
    };

    function zTreeOnClick(event, treeId, treeNode) {
        ClickObj = treeNode;
    };

     function menu_zTreeOnClick(event, treeId, treeNode) {
            MenuClickObj = treeNode;
        };

  });
