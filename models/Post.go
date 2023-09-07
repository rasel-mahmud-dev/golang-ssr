package models

type Post struct {
	Common
	Posts []map[string]interface{}
}

func InitPost() *Post {
	return &Post{}
}

func (p *Post) GetPosts() []map[string]interface{} {
	return p.Posts
}

func (p *Post) AddPost(post map[string]interface{}) {
	p.Posts = append(p.Posts, post)
}

func (p *Post) DeletePost(postId interface{}) {
	var updatedPosts []map[string]interface{}
	for _, post := range p.Posts {
		if post["id"] != postId {
			updatedPosts = append(updatedPosts, post)
		}
	}
	p.Posts = updatedPosts
}
