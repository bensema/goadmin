

layui.config({
    base: 'static/layuiadmin/' //静态资源所在路径
  }).extend({
    index: 'lib/index' //主入口模块
    ,goadmin: '../../js/goadmin'
  }).use(['index', 'user', 'goadmin'], function(){
    var $ = layui.$
    ,setter = layui.setter
    ,admin = layui.admin
    ,form = layui.form
    ,router = layui.router()
    ,search = router.search;

    form.render();

      layui.goadmin.req({
        type: 'get'
        ,url: layui.goadmin.api_rsa_url
        ,data: {}
        ,done: function(res){
            layui.data('crypto', {
              key: 'rsa'
              ,value: res.data
            });
            layui.data('crypto', {
              key: 'aes'
              ,value: generateMixed(16)
            });
        }
      })

    //提交
    form.on('submit(LAY-user-login-submit)', function(obj){

      var params = {}
      var data = {}
      data.Username = obj.field.username
      data.Password = obj.field.password
      params.vercode = obj.field.vercode
      params.username = obj.field.username;
      params.password = obj.field.password;

      //请求登入接口
      layui.goadmin.req({
        type: 'post'
        ,url: layui.goadmin.api_login_url
        ,data: params
        ,done: function(res){

          //请求成功后，写入 access_token
          layui.data(setter.tableName, {
            key: setter.request.tokenName
            ,value: res.data.access_token
          });

          //登入成功的提示与跳转
          layer.msg('登入成功', {
            offset: '15px'
            ,icon: 1
            ,time: 1000
          }, function(){
            location.href = '/'; //后台主页
          });
        }
      });

    });


  });

  function generateMixed(n) {
       var str = ['0','1','2','3','4','5','6','7','8','9','A','B','C','D','E','F','G','H','I','J','K','L','M','N','O','P','Q','R','S','T','U','V','W','X','Y','Z'];
       var res = "";
       for(var i = 0; i < n ; i ++) {
           var id = Math.ceil(Math.random()*35);
           res += str[id];
       }
       return res;
  }