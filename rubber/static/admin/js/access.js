//config.adminPath已在base.js中已定义，当前页面可以直接使用，引入js的时需要注意先引入base.js，然后再引入当前js
$(function () {
    accessApp.init()
})
var accessApp={
    init: function () {
        this.initSubmitRef()    //同步刷新
        this.accessToggle()    //点击收缩、展示
        this.initAccessToggle() //初始化收缩
    },

    initSubmitRef: function () {
        $("#submitRef").click(function () {
            alert("同步完成，跳至首页！")
            window.parent.location.reload();
        })
    },
    initAccessToggle: function () {
        $(".hide_table").hide()

    },
    accessToggle: function () {
        // 点击图标 收缩
        $(".chShowAccess").click(function () {
            var id = $(this).attr("data-id")

            var el = $(this)
            var accessItemData = ""
            $.get("/" + config.adminPath + "/access/doClickAccess", {
                id: id,
            }, function (response) {
                if (response.success) {
                    accessItemData = response.result;
                    for (var i = 0; i < accessItemData.length; i++) {
                        if (accessItemData[i].checked_list === true) {
                            for (var j = 0; j < accessItemData[i].access_item.length; j++) {
                                var num =document.getElementById(accessItemData[i].access_item[j].id)
                                if(num.style.display===''){
                                    num.style.display='none'

                                }else{
                                    num.style.display=''
                                }
                            }
                        } else {
                            for (var j = 0; j < accessItemData[i].access_item.length; j++) {
                                var hide =document.getElementById(accessItemData[i].access_item[j].id)
                                hide.style.display='none'
                                checked_id = false

                            }
                            console.log(response)
                        }
                    }
                }
            })
        })

    }
}