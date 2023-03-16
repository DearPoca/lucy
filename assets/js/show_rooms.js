function createTd(father, text) {
    let td = document.createElement("td");
    father.appendChild(td);
    td.innerHTML = text;
}

function createLink(father, id) {
    let ip_addr = document.location.hostname;

    let webrtc_link = document.createElement("td");
    let webrtc_page_url = `/play/webrtc?room_id=${id}`
    father.appendChild(webrtc_link);
    webrtc_link.innerHTML = `<a href=${webrtc_page_url} target="_blank">Link</a>`

    let rtmp_link = document.createElement("td");
    let rtmp_page_url = `/play/rtmp?room_id=${id}`
    father.appendChild(rtmp_link);
    rtmp_link.innerHTML = `<a href=${rtmp_page_url} target="_blank">Link</a>`
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
                    createLink(tr, resJson["data"][i]["id"]);
                }
            } else {
                window.alert("get rooms failed");
            }
        }
    };

    let url = `http://${ip_addr}/api/v1/get_rooms`
    xhr.open("get", url)
    xhr.send();

}

showTable()

