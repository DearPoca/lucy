function Login() {
    let ip_addr = document.location.hostname;
    let xhr = new XMLHttpRequest();

    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4 && xhr.status === 200) {
            const resJson = JSON.parse(xhr.response)
            console.log(resJson)
            if (resJson["code"] == 200) {
                window.location.replace(`http://${ip_addr}`)
            } else {
                window.alert("username or password error");
            }
        }
    };

    let username = document.getElementById("Username").value
    let password = document.getElementById("Password").value
    let url = `http://${ip_addr}/api/auth?username=${username}&password=${password}`
    xhr.open("get", url)
    xhr.send();
}
