function Register() {
    let ip_addr = document.location.hostname;
    let xhr = new XMLHttpRequest();

    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4 && xhr.status === 200) {
            const resJson = JSON.parse(xhr.response)
            console.log(resJson)
            if (resJson["code"] == 200) {
                window.location.replace(`http://${ip_addr}`)
            } else {
                window.alert("register failed, err: " + resJson["msg"]);
            }
        }
    };

    let username = document.getElementById("Username").value
    let password = document.getElementById("Password").value
    let email = document.getElementById("Email").value
    let telephone = document.getElementById("Telephone").value
    let url = `http://${ip_addr}/api/register?username=${username}&password=${password}&email=${email}&telephone=${telephone}`
    xhr.open("post", url)
    xhr.send();
}
