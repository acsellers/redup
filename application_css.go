package main

const (
	APPLICATIONCSS = `
html {
  background-color: #e2e2e2;
  margin: 0 auto;
}
.row-shrunk {
  padding: 12px; 
  width: 97%;
}
.page-title {
  padding-left: 24px
}
.navbar-key-items {
  font-family: 'Signika Negative', sans-serif;
}
.flat-bottom {
  margin-bottom: 0px;
}
.active.key-link {
  color: #454545;
}
  `

	APPLICATIONJS = `
  $(document).ready(function(){
    $(".key-link").click(function(ev){
      console.log(ev)
      key_val = ev.currentTarget.dataset["key"];
      console.log(key_val);
      $(".seperator").hide();
      $(".keycontent").hide();
      $(".key-link").removeClass("active");
      $('a[data-key="'+key_val+'"]').addClass("active");
      $("#key_"+key_val).show();
    });
    $(".all_keys").click(function(){
      $(".seperator").show();
      $(".keycontent").show();
      $(".key-link").removeClass("active");
    });
    $(".redis-item").each(function(i,v){
      try {
        value = JSON.parse($(v).html());
        $(v).html(JSON.stringify(value, null, "    "));
      }
      catch(e) {
      }
    });
  })
  `
)
