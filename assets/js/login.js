let loginForm = document.getElementById("login-form");
loginForm.addEventListener("submit", function (event) {
    event.preventDefault(); // 阻止表单默认提交行为

    const xhr = new XMLHttpRequest();
    const url = "/api/auth";
    const formData = new FormData(loginForm);

    const params = new URLSearchParams(formData).toString(); // 将表单数据序列化为URL参数

    xhr.open("GET", url + "?" + params);

    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4 && xhr.status === 200) {
            const response = JSON.parse(xhr.responseText);
            console.log(response)
            if (response["code"] == 200) {
                window.location.href = "/"
            } else {
                window.alert("login failed, err: " + response["msg"]);
            }
        }
    };

    xhr.send();
});

