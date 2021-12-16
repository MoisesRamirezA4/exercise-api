package posts

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/epa-datos/exercise/github.com/epa-datos/exercise-api/entity"
	"github.com/epa-datos/exercise/github.com/epa-datos/exercise-api/services/post"
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
	number := c.Param("number")
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/" + number)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var postResponse entity.PostDTO

	json.Unmarshal(body, &postResponse)
	c.JSON(200, postResponse.ConvertToPost())

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
