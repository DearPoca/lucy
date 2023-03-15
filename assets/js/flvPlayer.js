if (flvjs.isSupported()) {
    var ip_addr = document.location.hostname
    const xhr = new XMLHttpRequest()
    xhr.addEventListener("readystatechange", function () {
        // 状态不为4的话直接return
        if (xhr.readyState !== XMLHttpRequest.DONE) return

        // 拿到的结果是一个字符串, 我们可以转成js对象
        const resJSON = JSON.parse(xhr.response)
        console.log(resJSON)
        var videoElement = document.getElementById('player')
        var flvPlayer = flvjs.createPlayer({
            type: 'flv',
            url: resJSON['video_url']
        })
        flvPlayer.attachMediaElement(videoElement)
        flvPlayer.load()
        flvPlayer.play()
    })

    xhr.open("get", "http://"+ip_addr+"/video_url")
    xhr.send()
}