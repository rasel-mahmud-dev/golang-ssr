// class Common {
//     db;
//     tableName
//     constructor(tableName) {
//         this.db = connectDb()
//         this.tableName = tableName
//     }
//
// }
//
// class Post  extends  Common {
//     constructor() {
//         super("posts")
//     }
//
//     static findAllPost(columns){
//         return db.query(`select * from articles`)
//     }
// }
//
// class User extends  Common {
//     constructor() {
//         super("users")
//     }
//     static findAllUsers(columns){
//         return db.query(`select * from users`)
//     }
// }
//
//
// Post.findAllPost()
// User.findAllUsers()


const queryString = "https://rr4---sn-ax8xaj5ggpxg-q5j6.googlevideo.com/videoplayback?expire=1698186618&ei=GfE3ZYLSO6bMz7sPsLe18As&ip=103.204.209.121&id=o-AM4CqM8iZK9ZY9KVgby_cc1IaPF8T9qwpN867LIPP7wI&itag=251&source=youtube&requiressl=yes&mh=6s&mm=31%2C29&mn=sn-ax8xaj5ggpxg-q5j6%2Csn-npoeeney&ms=au%2Crdu&mv=m&mvi=4&pl=24&initcwndbps=592500&siu=1&spc=UWF9fx3S6y9KgX3YJJ9gdtZOF22klevGJghYXR1wV0FdUrnQUg9k5-8&vprv=1&svpuc=1&mime=audio%2Fwebm&ns=iHeSetPu3lpdnNLYb24cJwgP&gir=yes&clen=25611817&dur=1563.701&lmt=1696042184947379&mt=1698164745&fvip=1&keepalive=yes&fexp=24007246&beids=24350018&c=WEB&txp=4432434&n=4RHnCSih_yoThw&sparams=expire%2Cei%2Cip%2Cid%2Citag%2Csource%2Crequiressl%2Csiu%2Cspc%2Cvprv%2Csvpuc%2Cmime%2Cns%2Cgir%2Cclen%2Cdur%2Clmt&sig=AGM4YrMwRgIhAOdSmstFp1uCQiyj2W7kKIpWPVJs99KRY04WXAPcFdZeAiEA3j2wGiCV9UZ6qzmPU8eP1xEJu9zWTcD8Yr8NuB3foN0%3D&lsparams=mh%2Cmm%2Cmn%2Cms%2Cmv%2Cmvi%2Cpl%2Cinitcwndbps&lsig=AK1ks_kwRAIgT221ZEa2_FVbh2UhAU2iO1_WDHM_ytnjmKbbDlEsfgkCICOUWzOR5EOLrx5L4RoB72XZn4dKzRUYuTajUA6VkmI1&alr=yes&cpn=eZp0EgnwUeJqaHgd&cver=2.20231020.00.01&range=7691481-7757016&rn=22&rbuf=0&pot=MltKDvZpwQy_mko20Wc0cc9xDyiiCye9_D3sX2WA_xLclB5XOauNMCQLyHIEnZYqm9f0ZsGog4dbJgimdMu6Tor7WHuRZCcHdi1Eh9BKYA2I4RbZEFSJD0D2PrMn&ump=1&srfvp=1";

const urlParams = new URLSearchParams(queryString);
const params = {};

for (const [key, value] of urlParams.entries()) {
    params[key] = value;
}

let a  = {
    'https://rr4---sn-ax8xaj5ggpxg-q5j6.googlevideo.com/videoplayback?expire': '1698186618',
    ei: 'GfE3ZYLSO6bMz7sPsLe18As',
    ip: '103.204.209.121',
    id: 'o-AM4CqM8iZK9ZY9KVgby_cc1IaPF8T9qwpN867LIPP7wI',
    itag: '251',
    source: 'youtube',
    requiressl: 'yes',
    mh: '6s',
    mm: '31,29',
    mn: 'sn-ax8xaj5ggpxg-q5j6,sn-npoeeney',
    ms: 'au,rdu',
    mv: 'm',
    mvi: '4',
    pl: '24',
    initcwndbps: '592500',
    siu: '1',
    spc: 'UWF9fx3S6y9KgX3YJJ9gdtZOF22klevGJghYXR1wV0FdUrnQUg9k5-8',
    vprv: '1',
    svpuc: '1',
    mime: 'audio/webm',
    ns: 'iHeSetPu3lpdnNLYb24cJwgP',
    lsparams: 'mh,mm,mn,ms,mv,mvi,pl,initcwndbps',
    lsig: 'AK1ks_kwRAIgT221ZEa2_FVbh2UhAU2iO1_WDHM_ytnjmKbbDlEsfgkCICOUWzOR5EOLrx5L4RoB72XZn4dKzRUYuTajUA6VkmI1',
    alr: 'yes',
    cpn: 'eZp0EgnwUeJqaHgd',
    cver: '2.20231020.00.01',
    range: '7691481-7757016',
    rn: '22',
    rbuf: '0',
    pot: 'MltKDvZpwQy_mko20Wc0cc9xDyiiCye9_D3sX2WA_xLclB5XOauNMCQLyHIEnZYqm9f0ZsGog4dbJgimdMu6Tor7WHuRZCcHdi1Eh9BKYA2I4RbZEFSJD0D2PrMn',
    ump: '1',
    srfvp: '1'
}
// 4086-69621
// 0-4085