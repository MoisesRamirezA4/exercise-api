package posts

import "github.com/gin-gonic/gin"

func RegisterRoutes(e *gin.Engine) {
	handler := newPostHandler()
	e.GET("/posts", handler.listPosts)
	e.GET("/posts/:number", handler.postPath)
	e.POST("", handler.newPost)
	e.GET("/postsdb", handler.getPostsDb)
	e.GET("postsdb/:id", handler.getPostDb)
	e.POST("postdb", handler.newPostDb)
}
