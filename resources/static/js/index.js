

(function($){
    $.ajax({
        type : 'get',
        url : API_MENU_URL,
        dataType : 'json',
        data : {
        },
        success : function(res) {
            if (res.code != Code.Success){
                Swal.fire('Oops...', res.message, 'error')
            } else {
                let m = convert(res.data);
                let mm = sort_menu(m);
                let h = html_menu(convert(mm));
                $("#menu").html(h);
            }
        },
        error : function(err ){
            Swal.fire('Oops...', err.responseText, 'error')
            console.log(err)
        }
   })

})(jQuery)

function sort_menu(data) {
    data.sort(sortObj("sort",0));
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
        ulHtml += '<li class="nav-item">';
        if (item.url === "" && item.children !== undefined && item.children !== null && item.children.length > 0){
          ulHtml += '<a href="#" class="nav-link"> <i class="nav-icon fas '+item.icon + '"></i><p>' + item.name + '<i class="right fas fa-angle-left"></i></p></a>';
         } else if (item.url === ""){
          ulHtml += '<a href="#" class="nav-link"> <i class="nav-icon fas '+item.icon + '"></i><p>' + item.name + '</p></a>';
         } else{
          ulHtml += '<a href="'+ item.url + '" class="nav-link"> <i class="nav-icon fas '+item.icon + '"></i><p>' + item.name + '<i class="right fas fa-angle-left"></i></p></a>';
         }
        if (item.children !== undefined && item.children !== null && item.children.length > 0) {
            ulHtml += '<ul class="nav nav-treeview">';
            ulHtml += html_menu(item.children);
            ulHtml += "</ul>";
        }
        ulHtml += '</li>';
        continue
      }


      if (item.url === ""  && item.children !== undefined && item.children !== null && item.children.length > 0) {
          ulHtml += '<li class="nav-item"><a href="#" class="nav-link"><i class="far '+item.icon+' nav-icon"></i><p>' +  item.name + '<i class="right fas fa-angle-left"></i></p></a>'
      } else if (item.url === "" ) {
            ulHtml += '<li class="nav-item"><a href="#" class="nav-link"><i class="far '+item.icon+' nav-icon"></i><p>' +  item.name + '</p></a>'
      } else {
          ulHtml += '<li class="nav-item"><a href="' + item.url + '" class="nav-link"><i class="far '+item.icon+' nav-icon"></i><p>' +  item.name + '</p></a>'
      }

      if (item.children !== undefined && item.children !== null && item.children.length > 0) {
          ulHtml += '<ul class="nav nav-treeview">';
          ulHtml += html_menu(item.children);
          ulHtml += "</ul>";
      }
      ulHtml += "</li>";

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