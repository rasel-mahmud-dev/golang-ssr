const http = require("http")

const fs = require("fs")
const express = require("express")
const {resolve} = require("path");

const queryString = "https://rr4---sn-ax8xaj5ggpxg-q5j6.googlevideo.com/videoplayback?expire=1698186618&ei=GfE3ZYLSO6bMz7sPsLe18As&ip=103.204.209.121&id=o-AM4CqM8iZK9ZY9KVgby_cc1IaPF8T9qwpN867LIPP7wI&itag=251&source=youtube&requiressl=yes&mh=6s&mm=31%2C29&mn=sn-ax8xaj5ggpxg-q5j6%2Csn-npoeeney&ms=au%2Crdu&mv=m&mvi=4&pl=24&initcwndbps=592500&siu=1&spc=UWF9fx3S6y9KgX3YJJ9gdtZOF22klevGJghYXR1wV0FdUrnQUg9k5-8&vprv=1&svpuc=1&mime=audio%2Fwebm&ns=iHeSetPu3lpdnNLYb24cJwgP&gir=yes&clen=25611817&dur=1563.701&lmt=1696042184947379&mt=1698164745&fvip=1&keepalive=yes&fexp=24007246&beids=24350018&c=WEB&txp=4432434&n=4RHnCSih_yoThw&sparams=expire%2Cei%2Cip%2Cid%2Citag%2Csource%2Crequiressl%2Csiu%2Cspc%2Cvprv%2Csvpuc%2Cmime%2Cns%2Cgir%2Cclen%2Cdur%2Clmt&sig=AGM4YrMwRgIhAOdSmstFp1uCQiyj2W7kKIpWPVJs99KRY04WXAPcFdZeAiEA3j2wGiCV9UZ6qzmPU8eP1xEJu9zWTcD8Yr8NuB3foN0%3D&lsparams=mh%2Cmm%2Cmn%2Cms%2Cmv%2Cmvi%2Cpl%2Cinitcwndbps&lsig=AK1ks_kwRAIgT221ZEa2_FVbh2UhAU2iO1_WDHM_ytnjmKbbDlEsfgkCICOUWzOR5EOLrx5L4RoB72XZn4dKzRUYuTajUA6VkmI1&alr=yes&cpn=eZp0EgnwUeJqaHgd&cver=2.20231020.00.01&range=7691481-7757016&rn=22&rbuf=0&pot=MltKDvZpwQy_mko20Wc0cc9xDyiiCye9_D3sX2WA_xLclB5XOauNMCQLyHIEnZYqm9f0ZsGog4dbJgimdMu6Tor7WHuRZCcHdi1Eh9BKYA2I4RbZEFSJD0D2PrMn&ump=1&srfvp=1";

function getStats(filePath) {
    return new Promise((resolve) => {
        fs.stat(filePath, (err, stats) => {
            if (err) return resolve(null)
            resolve(stats)
        });
    })
}


// getChunk(0, 4085)

const app = express()

app.get("/", (req, res) => {
    res.setHeader("Content-Type", "text/html")
    let p = resolve("view/index.html")
    res.sendFile(p)

})

app.get("/video.html", (req, res) => {
    res.setHeader("Content-Type", "text/html")
    let p = resolve("src/views/watch.html")
    res.sendFile(p)

})

app.get("/videos", (req, res) => {
    let files = fs.readdirSync("public")
    res.json(files)
})

app.get("/watch", async (req, res) => {
    const name = req.query.title
    const filePath = "public/" + name

    let stats = await getStats(filePath)
    console.log(stats)
    res.json({
        size: stats.size,
        url: "/" + filePath
    })
})



function getChunk(req, res) {
    const fileStream = fs.createReadStream("public/Sandra N - Lion Heart (Official Video) ( 360 X 640 ).mp4", {
        start: 0,
    })
    fileStream.on('open', () => {
        fileStream.pipe(res);
    });
    fileStream.on('end', () => {
        res.end(null)
    });
}

app.get("/get-chunk", async (req, res) => {
    const {name, start, end} = req.query
    const path = 'public/Sandra N - Lion Heart (Official Video) ( 360 X 640 ).mp4';
    const stat = fs.statSync(path);
    const fileSize = stat.size;
    const range = req.headers.range;

    if (range) {
        const parts = range.replace(/bytes=/, '').split('-');
        const start = parseInt(parts[0], 10);
        const end = parts[1] ? parseInt(parts[1], 10) : fileSize - 1;

        const chunkSize = (end - start) + 1;
        const file = fs.createReadStream(path, { start, end });
        const headers = {
            'Content-Range': `bytes ${start}-${end}/${fileSize}`,
            'Accept-Ranges': 'bytes',
            'Content-Length': chunkSize,
            'Content-Type': 'video/mp4',
        };

        res.writeHead(206, headers);
        file.pipe(res);
    } else {
        const headers = {
            'Content-Length': fileSize,
            'Content-Type': 'video/mp4',
        };

        res.writeHead(200, headers);
        fs.createReadStream(path).pipe(res);
    }

    // const filePath = "public/" + "Sandra N - Lion Heart (Official Video) ( 360 X 640 ).mp4"
    //
    // const fileStream = fs.createReadStream(filePath, {
    //     start: Number(start),
    //     end: Number(end)
    // })
    // fileStream.on('open', () => {
    //     const stat = fs.statSync(filePath);
    //     res.set({
    //         'Content-Type': 'video/mp4', // This is the Content-Type header for video files
    //         'Content-Length': end - start + 1,
    //         'Accept-Ranges': 'bytes',
    //         'Content-Range': `bytes ${start}-${end}/${stat.size}`,
    //     });
    //     fileStream.pipe(res);
    // });
    // fileStream.on('end', () => {
    //     res.end(null)
    // });
})



const server = http.createServer(app)
server.listen(1000, () => console.log("server is running on port 1000"))

// id: for expire: 1698188064
// o-AM4CqM8iZK9ZY9KVgby_cc1IaPF8T9qwpN867LIPP7wI
// o-ALG1m6cYhLq5xuOhAvP4L0AUzur8ec43dxCMSkNZ7aOu
// o-AEFerKf4naU4IGhC5X5iOSyXtMdYIdU3u1iSvuq7i9dz

