window.onpopstate = function(e){
    if(history.state){
        var state = e.state;
        console.log("state: " + state);
        ListFile(state);
    }
}
function clickFileorDir(type, path, name){
    var e = window.event || arguments[0]
    if(type == "file"){
        return true;
    }
    loadFileList(path, name);
    e.preventDefault();
}

function loadFileList(path, name){
    if(name != ""){
        path = path + "/" + name;
    }
    window.history.pushState(path, "", path);
    //console.log("path: " + path);
    ListFile(path);
}

function ListFile(path) {
    $(".myrow").remove();
    $.ajax({
        type: "GET",//使用get方法访问后台
        dataType: "json",//返回json格式的数据
        url: path,//要访问的后台地址
        cache: false,
        beforeSend: function(request){
            request.setRequestHeader("isAjax", "true")
        },
        success: function (msg) {//msg为返回的json数据
            // console.log(msg);
            $("#mytemplate").show();
            $.each(msg, function (i, n) {
                var row = $("#mytemplate").clone();
                row.addClass("myrow");
                row.find("#NO").text(i);
                row.find("#Name").html("<a onclick=clickFileorDir('" + n.type + "','" + path + "','" + n.name + "')" + " href='/" + n.path + "/" + n.name + "' target='_blank'>" + n.name + "</a>")
                row.find("#Size").text(n.size);
                row.find("#Operation").html("<input class='mycheckbox'type='checkbox' value='" + i + "'>");
                row.appendTo("#datas");//添加到模板的容器中
            });
            $("#mytemplate").hide();   //隐藏模板
        },
        error: function () {
            alert("连接服务器错误");
        }
    });
}