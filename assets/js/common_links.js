let common_links = [
    {title: "live rooms", url: "/"},
    {title: "my live", url: "/my_live"},
];

function addLinks() {
    let container = document.createElement('div');
    container.style.position = 'fixed';
    container.style.top = '10px';
    container.style.right = '10px';
    container.style.display = 'flex';
    container.style.flexDirection = 'row';
    container.style.gap = '10px';
    container.style.zIndex = '9999';
    container.style.background = '#f5f5f5';

    for (let i = 0; i < common_links.length; i++) {
        let link = document.createElement('a');
        link.href = common_links[i].url;
        link.textContent = common_links[i].title;
        container.appendChild(link);
    }
    // Add userinfo
    let userinfo = document.createElement('a');
    userinfo.href = "/userinfo";
    userinfo.textContent = document.getElementById("username").dataset.value;

    let separator = document.createElement('span');
    separator.textContent = '|';
    separator.style.margin = '0px -6px';

    // Add logout
    let logout = document.createElement('a');
    logout.href = '/login';
    logout.textContent = 'logout';

    container.appendChild(userinfo);
    container.appendChild(separator);
    container.appendChild(logout);

    document.body.appendChild(container);
}

function addReturn() {
    let container = document.createElement('div');
    container.style.position = 'fixed';
    container.style.top = '10px';
    container.style.left = '10px';
    container.style.display = 'flex';
    container.style.flexDirection = 'row';
    container.style.gap = '10px';
    container.style.zIndex = '9999';
    container.style.background = '#f5f5f5';

    let backLink = document.createElement('a');
    backLink.href = 'javascript:history.back()';
    backLink.textContent = 'return';
    container.appendChild(backLink);

    document.body.appendChild(container);
}

window.addEventListener('load', addLinks);
window.addEventListener('load', addReturn);