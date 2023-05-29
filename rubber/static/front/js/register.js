(function($){
    var app={
        init:function(){
            this.sendCodeStep1(); //发送验证码 step1
            this.returnButton();
            this.submitFormStep3();
            this.verifyUserNameStep3();
        },
        //返回上一页
        returnButton () {
            $("#returnButton").click(function () {
                window.location.href = document.referrer
            })

        },
    //判断全选是否选择
        sendCodeStep1(){
            //发送验证码
            $("#registerButton").click(function () {
                //验证验证码是否正确
                var phone = $('#phone').val();
                var photo_code = $('#photo_code').val();
                var photoCodeId = $("input[name='captcha_id']").val();

                var reg = /^[\d]{11}$/;
                if (!reg.test(phone)) {
                    $(".error").html("Error：手机号输入错误");
                    return false;
                }
                if (photo_code.length < 4) {
                    $(".error").html("Error：图形验证码长度不合法")
                    return false;
                }
                $.get('/pass/RegisterStep1', { phone: phone}, function (response) {
                    console.log(response)
                    if (response.success) {

                    }else{
                        $(".error").html("该手机号已注册!")
                    }
                    })

                $.get('/pass/sendCode', { phone: phone, photo_code: photo_code, photoCodeId: photoCodeId }, function (response) {
                    console.log(response)
                    if (response.success === true) {
                        //跳转到下页面
                        location.href="/pass/registerStep2?sign="+response.sign+"&photo_code="+photo_code;
                    } else {
                        //改变验证码
                        $(".error").html("Error：" + response.msg + ",请重新输入!")

                        //改变验证码
                        var captchaImgSrc = $(".captcha-img").attr("src")
                        $("#photo_code").val("")
                        $(".captcha-img").attr("src", captchaImgSrc + "?reload=" + (new Date()).getTime())
                    }

                });

            })
        },
        submitFormStep3(){
            $("#nextStep").click(function(){
                //在form 中获取，方便传值后端
               document.getElementById("username").value=document.getElementById("rusername").value;
               // 判断
               $(".error").html("")
                var password=$('#password').val();
                var rpassword=$('#rpassword').val();

                if(password.length<6){
                    $(".error").html("密码的长度不能小于6位")
                    return false;
                }
                if(password!==rpassword){
                    $(".error").html("密码和确认密码不一致")
                    return false;
                }
                return true;

            })
        },
        verifyUserNameStep3(){
            $('#verifyUserName').click(function () {
                var username=document.getElementById("rusername").value
                var verifname = /^[\u4E00-\u9FA5\uF900-\uFA2D\w]{2,50}$/;
                if (!verifname.test(username)){
                    $(".success").html("")
                    $(".error").html("请输入正确的用户名格式")
                    return false
                }
                else{
                    $(".error").html("")
                }
                $.get("/pass/validateUsername", {"username": username}, function (response) {
                    if (response.success) {
                        $(".success").html("用户名可以注册")

                    }
                    else {
                        $(".error").html(response.msg)
                    }
                })

            })
        }


    }

    $(function(){
        app.init();
    })
})($)
