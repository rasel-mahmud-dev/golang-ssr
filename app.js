class Common {
    db;
    tableName
    constructor(tableName) {
        this.db = connectDb()
        this.tableName = tableName
    }
    findAll(columns){
        let s = columns.join(",").replace(/,$/, "");
        return db.query(`select ${s} from ${this.tableName}`)
    }
}

class Post  extends  Common {
    constructor() {
        super("posts")
    }
}

class User extends  Common {
    constructor() {
        super("users")
    }
}

let post = new Post()
post.findAll(["id", "slug", "title", "author_id", "created_at"])

let user = new User()
user.findAll()
