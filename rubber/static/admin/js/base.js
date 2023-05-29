$(function () {
    app.init()
    $(window).resize(function () {
        app.resizeIframe()
    })
})
var config = {
    adminPath: "state_grid"
}
var app = {
   init: function () {
    this.prevPageBack()	 //返回上一页
    this.changeModuleId()	 //修改属性则节点类型改变
    this.slideToggle()	 //左侧侧边栏点击收缩
    this.resizeIframe()		// 内容部分高度为浏览器的自适应高度-80
    this.confirmDelete()   	//删除提示
    this.changeStatus() 	//点击修改状态
    this.changeNum()	//点击修改数字

},
prevPageBack: function () {
    $('#prevPage').click(function () {
        window.location.href = document.referrer
    })
},

slideToggle: function () {
    // 调整收缩左侧隐藏模块下子内容
    $('.aside>li:nth-child(1) ul,.aside>li:nth-child(2) ul,.aside>li:nth-child(3) ul,.aside>li:nth-child(4) ul,.aside>li:nth-child(5) ul,.aside>li:nth-child(6) ul').hide()
    $('.aside h4').click(function () {
        $(this).siblings('ul').slideToggle();
    })
},
resizeIframe: function () {
    $("#rightMain").height($(window).height() - 80)

},
confirmDelete: function () {
    $(".delete").click(function () {
        var flag = confirm("您确定要删除吗？")
        return flag

    })
},
// confirSearch :function () {
// 	$(".search").click(function () {
// 		alert("未搜索到相关数据")
// 	})
// },
changeStatus: function () {
    $(".chStatus").click(function () {
        var id = $(this).attr("data-id");
        var table = $(this).attr("data-table");
        var field = $(this).attr("data-field");
        var el = $(this)
        $.get("/" + config.adminPath + "/main/changeStatus", {
            id: id,
            table: table,
            field: field
        }, function (response) {
            if (response.success) {

                if (el.attr("src").indexOf("yes") != -1) {
                    el.attr("src", "/static/admin/images/no.png")
                } else {
                    el.attr("src", "/static/admin/images/yes.png")
                }
            } else {
                console.log(response)
            }
        })
    })
},
changeNum: function () {
    /*
    1、获取el里面的值  var spanNum=$(this).html()
    2、创建一个input的dom节点   var input=$("<input value='' />");
    3、把input放在el里面   $(this).html(input);
    4、让input获取焦点  给input赋值    $(input).trigger('focus').val(val);
    5、点击input的时候阻止冒泡
                $(input).click(function(e){
                    e.stopPropagation();
                })
    6、鼠标离开的时候给span赋值,并触发ajax请求
        $(input).blur(function(){
            var inputNum=$(this).val();
            spanEl.html(inputNum);
            触发ajax请求
        })
    */

    $(".chSpanNum").click(function () {
        var id = $(this).attr("data-id")
        var table = $(this).attr("data-table");
        var field = $(this).attr("data-field");
        var spanNum = $(this).attr("data-num");
        var spanEl = $(this)  //保存span这个dom节点

        var input = $("<input value='' style='width:60px' />");
        $(this).html(input);
        $(input).trigger('focus').val(spanNum);   //让输入框获取焦点并设置值
        $(input).click(function (e) {
            e.stopPropagation();
        })
        $(input).blur(function (e) {
            var inputNum = $(this).val();
            spanEl.html(inputNum);
            //异步请求修改数量
            $.get("/" + config.adminPath + "/main/editNum", {
                id: id,
                table: table,
                field: field,
                num: inputNum
            }, function (response) {
                if (!response.success) {
                    console.log(response)
                }
                // // 重新加载框架
                document.location.reload();
            })
        })
    })
},
// 根据所选的所属模块显示节点类型，近顶级模块显示节点类型：模块
changeModuleId: function () {
    $("#module_id").change(function () {
        var module_id = $(this).val();

        var optionNow = $("#type").find("option").eq(0);
        var optionP = optionNow.parent("span"); //添加span 标签
        var optionLength = $("#type").children("span").length //查看span 长度
        if (module_id === "0") {
            optionP.children().clone().replaceAll(optionP); //删除隐藏的span 使用上面克隆的option 显示
        } else {
            if (optionLength === 0) {
                optionNow.wrap("<span style='display:none'></span>");
            }
        }

    })

}
}

