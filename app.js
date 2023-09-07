class Common {
    findAll(){
        return []
    }
}


class Post  extends  Common {

    posts = []

    constructor() {
        super()
    }

    getPosts(){
        return this.posts;
    }

    addPost(post){
        this.posts.push(post)
    }

    deletePost(postId){
        return this.posts.filter(post=>post.id !== postId)
    }
}

let post = new Post()

post.addPost({title: "Test post", price: 45.10})
post.addPost({title: "Test post", price: 45.10})

post.findAll()

console.log(post.posts)