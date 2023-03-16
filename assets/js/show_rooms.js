function createTd(father, text) {
    let td = document.createElement("td");
    father.appendChild(td);
    td.innerHTML = text;
}

function createLink(father, url) {
    let link = document.createElement("td");
    father.appendChild(link);
    link.innerHTML = `<a href=${url} target="_blank">Link</a>`
}

function showTable() {

    let ip_addr = document.location.hostname;
    let xhr = new XMLHttpRequest();

    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4 && xhr.status === 200) {
            const resJson = JSON.parse(xhr.response)
            console.log(resJson)
            if (resJson["code"] == 200) {
                let table = document.querySelector("table");
                for (let i = 0; i < resJson["data"].length; i++) {
                    let room = resJson["data"][i]
                    let tr = document.createElement("tr");
                    table.children[1].appendChild(tr);
                    createTd(tr, room["id"])
                    createTd(tr, room["name"])
                    createTd(tr, room["path"])
                    let webrtc_page_url = `/play/webrtc?room_id=${room["id"]}`
                    let flv_page_url = `/play/flv?room_id=${room["id"]}`
                    createLink(tr, webrtc_page_url)
                    createLink(tr, flv_page_url)
                }
            } else {
                window.alert("get rooms failed");
            }
        }
    }

    let url = `http://${ip_addr}/api/v1/get_rooms`
    xhr.open("get", url)
    xhr.send();

}

showTable()

