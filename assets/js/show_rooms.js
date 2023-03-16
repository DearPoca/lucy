function Item_Data(room_id, room_name, room_path) {
    this.room_id = room_id;
    this.room_name = room_name;
    this.room_path = room_path;
}

function createTd(father, text) {
    let td = document.createElement("td");
    father.appendChild(td);
    td.innerHTML = text;
}

function createLink(father, id) {
    let del = document.createElement("td");
    let link = `http://${ip_addr}/play?room_id=${id}`
    del.innerHTML = `<a href=${link}>Link</a>`
    father.appendChild(del);

    return del;
}

let ip_addr = document.location.hostname;
let xhr = new XMLHttpRequest();

xhr.onreadystatechange = function () {
    if (xhr.readyState === 4 && xhr.status === 200) {
        const resJson = JSON.parse(xhr.response)
        console.log(resJson)
        if (resJson["code"] == 200) {
            let items = [];
            let table = document.querySelector("table");
            for (let i = 0; i < resJson["data"].length; i++) {
                let stream = resJson["data"][i]
                items[i] = new Item_Data(stream["id"], stream["name"], stream["url"])
                let tr = document.createElement("tr");
                table.children[1].appendChild(tr);
                for (let j in items[i]) {
                    createTd(tr, items[i][j]);
                }
                let link = createLink(tr, resJson["data"][i]["id"]);
            }
            let as = document.querySelectorAll("a");
            console.log(as);
            as.forEach((a) => {
                a.onclick = function () {
                    console.log(a.parentNode.parentNode);
                    table.children[1].removeChild(a.parentNode.parentNode)
                }
            })
        } else {
            window.alert("get rooms failed");
        }
    }
};

let url = `http://${ip_addr}/api/v1/get_rooms`
xhr.open("get", url)
xhr.send();

