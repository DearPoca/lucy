let registerForm = document.getElementById("register-form");
registerForm.addEventListener("submit", function (event) {
    event.preventDefault(); // 阻止表单默认提交行为

    const xhr = new XMLHttpRequest();
    const url = "/api/register";
    const formData = new FormData(registerForm);

    const params = new URLSearchParams(formData).toString(); // 将表单数据序列化为URL参数

    xhr.open("POST", url + "?" + params);

    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4 && xhr.status === 200) {
            const response = JSON.parse(xhr.responseText);
            console.log(response)
            if (response["code"] == 200) {
                window.location.href = "/"
            } else {
                window.alert("register failed, err: " + response["msg"]);
            }
        }
    };

    xhr.send();
});


