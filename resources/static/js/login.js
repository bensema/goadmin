

(function($){

    $(document).ready(function(){
      $("form").submit(function(e){
        e.preventDefault();
         var username = $( "#username" ).val();
         var password = $( "#password" ).val();
         $.ajax({
        		type : 'post',
                url : API_LOGIN_URL,
                dataType : 'json',
                data : {
                    "username" : username,
                    "password" : password,
                },
                success : function(res) {
                    if (res.code != Code.Success){
                        Swal.fire('Oops...', res.message, 'error')
                    } else {
                        console.log(res)
                        window.location.href = '/';
                    }
                },
                error : function(err ){
                    Swal.fire('Oops...', err.responseText, 'error')
                    console.log(err)
                }
           })

      });
    });

})(jQuery)