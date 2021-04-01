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

      function opened_at(d) {
         return layui.goadmin.timestampToTime(d.opened_at)+'<br>'+layui.goadmin.timestampToTime(d.closed_at);
      }

      function closed_at(d) {
           return layui.goadmin.timestampToTime(d.closed_at);
        }
      function draw_at(d) {
           return layui.goadmin.timestampToTime(d.draw_at)+'<br>'+layui.goadmin.timestampToTime(d.update_at);
        }

      function status(d) {
        if (d.status == "1"){
            return '<button class="layui-btn layui-btn-xs">已开奖</button>'
        }
        if (d.status == "2"){
           return '<button class="layui-btn layui-btn-primary layui-btn-xs">未开奖</button>'
        }
      }
      var cur_game_code="";

      layui.goadmin.req({
        type: 'get'
        ,url: layui.goadmin.bb_admin_api_game_pages
        ,data: {"num":1,"size":100}
        ,done: function(res){
           $.each(res.data.data, function(index, item) {
                $('#game_code').append(new Option(item.display_name, item.game_code));
           });
           layui.form.render();

           table.render({
               elem: '#LAY-app-content-list'
               ,url: layui.goadmin.bb_admin_api_game_result_pages
               ,where: {game_code: layui.$('#game_code').val()}
               ,cols: [[
                 {field: 'id', width: 100,  title: 'ID', sort: true, align: 'center'}
                 ,{field: 'issue', width: 150, title: '局号',  align: 'center'}
                 ,{field: 'result', title: '结果',  align: 'center'}
                 ,{field: 'opened_at',width: 170, title: '开局时间<br>封单时间',  templet: opened_at, align: 'center'}
                 ,{field: 'draw_at', width: 170,title: '搅珠时间<br>更新时间', templet: draw_at, align: 'center'}
                 ,{field: 'status', width: 100, title: '状态', templet: status, align: 'center'}
                 ,{title: '操作', width: 200,align: 'center', fixed: 'right', toolbar: '#table-content-list'}
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
                  url: layui.goadmin.bb_admin_api_game_del,
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
            ,title: '编辑管游戏'
            ,content: layui.goadmin.web_bb_game_form_url + '?id='+ data.id
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
        var page={};

        if (cur_game_code !== field.game_code){
            page.curr = 1
        }else{
        }
        cur_game_code = field.game_code

        //执行重载
        table.reload('LAY-app-content-list', {
          where: field,
          page:page,
        });
      });

      var $ = layui.$, active = {
        add: function(){
          layer.open({
            type: 2
            ,title: '添加游戏'
            ,content: layui.goadmin.web_bb_game_add_url
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

