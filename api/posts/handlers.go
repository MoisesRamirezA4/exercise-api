package posts

import (
	_ "encoding/json"
	_ "io/ioutil"
	_ "log"
	"net/http"

	"github.com/epa-datos/exercise-api/entity"
	"github.com/epa-datos/exercise-api/repositories/mysql"
	"github.com/epa-datos/exercise-api/services/post"
	"github.com/gin-gonic/gin"
)

type postHandler struct {
}

//Method for http.GET
func (p *postHandler) listPosts(c *gin.Context) {
	var getPost post.Post           //Creating the instance of the service
	c.JSON(200, getPost.AllPosts()) //Printing out all the posts by calling the method AllPosts()
}

// Method for GET --> PATH number

func (p *postHandler) postPath(c *gin.Context) {
	var getPostById post.Post                                           //Creating the instance of the service
	c.JSON(200, getPostById.OnePost(c.Param("number")).ConvertToPost()) //Printing out the Post got by the path number and the method OnePost
}

//Method for http.POST
func (p *postHandler) newPost(c *gin.Context) {
	var newPost entity.Post                  //Creating a variable of the Post struct
	if err := c.Bind(&newPost); err != nil { //Getting the
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var body post.Post
	c.String(http.StatusOK, string(body.CreatePost(newPost)))
}

//Method to post on the MYSQL database
func (p *postHandler) newPostDb(c *gin.Context) {
	var newPost *entity.Post
	if err := c.Bind(&newPost); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	mysql.Db.Create(&newPost)
}

//Method to get a single post from the database
func (p *postHandler) getPostDb(c *gin.Context) {
	post := &entity.Post{}
	id := c.Param("id")
	mysql.Db.First(post, id)
	c.JSON(200, post)
	//return post, mysql.Db.First(post, id).Error
}

//Method to get all posts from the database
func (p *postHandler) getPostsDb(c *gin.Context) {
	posts := []*entity.Post{}
	mysql.Db.Find(&posts)
	c.JSON(200, posts)
	//return posts, mysql.Db.Find(&posts).Error
}

func newPostHandler() *postHandler {
	return &postHandler{}
}
