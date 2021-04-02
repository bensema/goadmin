

  layui.config({
    base: '/static/layuiadmin/' //静态资源所在路径
  }).extend({
    index: 'lib/index' //主入口模块
    ,goadmin: '../../js/goadmin'
    ,hashes: '../../js/hashes'
  }).use(['index', 'form', 'goadmin', 'transfer', 'laydate','upload','hashes'], function(){
    var $ = layui.$
    ,form = layui.form
    ,admin = layui.admin
    ,transfer = layui.transfer
    ,laydate = layui.laydate
    ,upload = layui.upload
    ,hashes = layui.hashes

    id = getQueryVariable('id')
    game_code = getQueryVariable('game_code')
    params = {}
    params.id = getQueryVariable('id')
    params.game_code = getQueryVariable('game_code')
//                console.log('sha512:'+hashes.sha512('123456'));


    layui.goadmin.req({
        type: 'get'
        ,url: layui.goadmin.bb_admin_api_game_result_query
        ,data: params
        ,done: function(res){
          layui.$('#seed').val(res.data.seed);
          layui.$('#seed_hash').val(res.data.seed_hash);
          layui.$('#game_code').val(res.data.game_code);
          layui.$('#issue').val(res.data.issue);

          layui.form.render();
        }
    });

    function Check512(){
        var seed = layui.$('#seed').val();
        layui.$('#check_seed_hash').val(hashes.sha512(seed));
        layui.form.render();
    }
    form.on('submit(LAY-app-contlist-search)', function(data){
        var seed = layui.$('#seed').val();
        var hash = hashes.sha512(seed);
         layui.$('#check_seed_hash').css("color", "black");
         layui.$('#seed_hash').css("color", "black");
        layui.$('#check_seed_hash').val(hashes.sha512(seed));
        if (hash == layui.$('#seed_hash').val()){
            layui.$('#check_seed_hash').css("color", "green");
            layui.$('#seed_hash').css("color", "green");
            layer.msg("验证成功")
        }else{
           layui.$('#check_seed_hash').css("color", "red");
           layer.alert("验证失败")
        }
        layui.form.render();
    })

    //监听提交
    form.on('submit(layuiadmin-app-form-submit)', function(data){
      var field = data.field; //获取提交的字段
      var index = parent.layer.getFrameIndex(window.name); //先得到当前iframe层的索引

      //提交 Ajax 成功后，关闭当前弹层并重载表格

      parent.layui.table.reload('LAY-app-content-list'); //重载表格
      parent.layer.close(index); //再执行关闭
    });
    //监听提交

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