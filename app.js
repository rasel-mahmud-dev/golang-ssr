class Common {
    db;
    tableName
    constructor(tableName) {
        this.db = connectDb()
        this.tableName = tableName
    }

}

class Post  extends  Common {
    constructor() {
        super("posts")
    }

    static findAllPost(columns){
        return db.query(`select * from articles`)
    }
}

class User extends  Common {
    constructor() {
        super("users")
    }
    static findAllUsers(columns){
        return db.query(`select * from users`)
    }
}


Post.findAllPost()
User.findAllUsers()