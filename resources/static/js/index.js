
 layui.config({
    base: 'static/layuiadmin/' //静态资源所在路径
  }).extend({
    index: 'lib/index' //主入口模块
    ,goadmin: '../../js/goadmin'

  }).use(['index','layer', 'element','goadmin'],function(){
    var element = layui.element;
    layui.goadmin.req({
        type: "get",
        url: layui.goadmin.api_menu_url ,
        done: function(res) {
          let m = convert(res.data);
          let mm = sort_menu(m);
          let h = html_menu(convert(mm));
          layui.$("#LAY-system-side-menu").html(h);
          element.render('nav', 'layadmin-system-side-menu')
          return
        }
    });

  });

  function sort_menu(data) {
    data.sort(sortObj("index_sort",0));
    for (const item of data) {
      if (item.children !== undefined && item.children !== null && item.children.length > 0) {
        sort_menu(item.children);
      }
    }
    return data;
  }

  function convert(list) {
	const res = []
	const map = list.reduce((res, v) => (res[v.id] = v, res), {})
	for (const item of list) {
		if (item.pid === 0) {
			res.push(item)
			continue
		}
		if (item.pid in map) {
			const parent = map[item.pid]
			parent.children = parent.children || []
			parent.children.push(item)
		}
	}
	return res
  }

  function html_menu(menu_list) {
    let ulHtml = "";
    for (const item of menu_list){
      if (item.pid === 0){
        ulHtml += '<li data-name="" class="layui-nav-item">';
        if (item.url === ""){
          ulHtml += '<a href="javascript:;" lay-tips="'+item.name + '" lay-direction="2"><i class="layui-icon ' + item.icon + '"></i><cite>'+item.name+'</cite></a>';
         }else{
          ulHtml += '<a href="'+item.url+'" lay-tips="'+item.name + '" lay-direction="2"><i class="layui-icon '+ item.icon +'"></i><cite>'+item.name+'</cite></a>';
         }
        if (item.children !== undefined && item.children !== null && item.children.length > 0) {
            ulHtml += html_menu(item.children);
        }
        ulHtml += '</li>';
        continue
      }

      ulHtml += '<dl class="layui-nav-child">';
      if (item.url === "") {
          ulHtml += '<dd data-name="console"><a href="javascript:;"><i class="layui-icon ' + item.icon + '"></i>' + item.name + '</a>'
      } else {
          ulHtml += '<dd data-name="console"><a lay-href="' + item.url + '"><i class="layui-icon ' + item.icon + '"></i>' + item.name + '</a>'
      }

      if (item.children !== undefined && item.children !== null && item.children.length > 0) {
          ulHtml += html_menu(item.children);
      }

      ulHtml += "</dd></dl>";
    }
    return ulHtml;
  }

  function sortObj(propertyName,cond) {
      return function(object1, object2) {
        var value1 = object1[propertyName];
        var value2 = object2[propertyName];
        if(cond == 1){//降序
          if (value2 < value1) {
            return - 1;
        } else if (value2 > value1) {
            return 1;
        } else {
            return 0;
        }
        }else if(cond == 0){//升序
          if (value2 < value1) {
            return  1;
        } else if (value2 > value1) {
            return - 1;
        } else {
            return 0;
        }
      }
    }
  }