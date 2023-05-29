(function ($) {

    var app = {
        init: function () {
            window.initUserPhone = $('#phone').val();
            window.initUsername = $('#username').val();
            window.initEmail = $('#email').val();

            this.changeUserPassword();
            this.changeUserInfo();
        },
        changeUserInfo: function () {
            // 提交信息
            $("#changeUserInfo").click(function () {
                var id = $('#id').val();
                var phone = $('#phone').val();
                var username = $('#username').val();
                var password = $('#password').val();
                var email = $('#email').val();
                if ( window.initUsername === username && window.initUserPhone === phone &&  window.initEmail===email && (password ==="" || password === undefined))
                {
                    $(".error").html("未修改任何信息!")
                    return false;
                }
                app.changePhone(); //手机号校验
                app.changeUsername(); //用户名校验
                if(password ==="" || password === undefined){
                    password =""
                }else
                {
                    app.changePassword(); //密码校验

                }
                $.post('/user/changeUserInfo', {
                    id: id,
                    phone:  phone,
                    username: username,
                    password: password,
                    email:email,
                }, function (responseUserInfo) {
                    console.log(responseUserInfo)
                    if (responseUserInfo.success) {

                        $(".success").html("修改信息成功!3秒后重新登录")
                        setInterval(function(){
                            $.get('/pass/loginOut', function (loginOut) {
                                location.href = '/pass/login'
                            })
                        }, 3000);
                    } else {
                        $(".error").html("修改信息失败!")
                        return false;
                    }
                })
            });
        },
        changePhone: function () {
            $(".error").html("");
            var phone = $('#phone').val();
            var reg = /^[\d]{11}$/;
            if (!reg.test(phone)) {
                $(".error").html("手机号输入错误");
                return false;
            }
            if (initUserPhone === phone) {
                //    未修改
            } else {
                $.get('/user/changeUserPhone', {phone: phone}, function (response) {
                    console.log(response)
                    if (response.success) {

                    } else {
                        $(".error").html("该手机号已注册!")
                        return false;
                    }
                })
            }
        },

        changeUsername: function () {
            var username = $('#username').val();
            var verifname = /^[\u4E00-\u9FA5\uF900-\uFA2D\w]{2,50}$/;
            if (!verifname.test(username)) {
                $(".error").html("请输入正确的用户名格式")
                return false
            } else if (initUsername === username) {
                //    未修改
            } else {
                $.get("/pass/validateUsername", {"username": username}, function (response) {
                    if (response.success) {
                        //"用户名可以注册"
                    } else {
                        $(".error").html(response.msg)
                    }
                })
            }
        },

        changePassword: function () {
            // var password = $('#password').val();
            // var rpassword = $('#rpassword').val();

            if (password.length < 6) {
                $(".error").html("密码的长度不能小于6位")
                return false;
            }
            if (password !== rpassword) {
                $(".error").html("密码和确认密码不一致")
                return false;
            }
        },

        changeUserPassword:function () {
            $("#changeUserPassword").click(function () {
                var str="";
                str += '<li>修改密码: <input  type="text"  id="password" name="password" placeholder="默认为空，不修改密码"/></li>'
                str += '<li>确认密码: <input  type="text" id="rpassword" name="rpassword" placeholder="默认为空，不修改密码"/></li>'
                $("#changePd").html(str)
            })

        }
    }
    $(function () {
        app.init();
    })


})($)
