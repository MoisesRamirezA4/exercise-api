package posts

import (
	_ "encoding/json"
	_ "io/ioutil"
	_ "log"
	"net/http"

	"github.com/epa-datos/exercise-api/entity"
	"github.com/epa-datos/exercise-api/services/post"
	"github.com/gin-gonic/gin"
)

type postHandler struct {
}

//Method for http.GET
func (p *postHandler) listPosts(c *gin.Context) {
	var getPost post.Post
	c.JSON(200, getPost.AllPosts())
}

// Method for GET --> PATH number

func (p *postHandler) postPath(c *gin.Context) {
	var getPostById post.Post
	c.JSON(200, getPostById.OnePost(c.Param("number")).ConvertToPost())
}

//Method for http.POST
func (p *postHandler) pruebaPosts(c *gin.Context) {
	var newPost entity.Post
	if err := c.Bind(&newPost); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var body post.Post
	c.String(http.StatusOK, string(body.CreatePost(newPost)))
}

func newPostHandler() *postHandler {
	return &postHandler{}
}
