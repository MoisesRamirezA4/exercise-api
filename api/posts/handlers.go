package posts

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/epa-datos/exercise/github.com/epa-datos/exercise-api/entity"
	"github.com/gin-gonic/gin"
)

type postHandler struct {
}

//Method for http.GET
func (p *postHandler) listPosts(c *gin.Context) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var postResponse []entity.PostDTO
	json.Unmarshal(body, &postResponse)

	var posts []*entity.Post
	for _, p := range postResponse {
		posts = append(posts, p.ConvertToPost())
	}

	c.JSON(200, posts)
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
	// requestBody, err := json.Marshal(map[string]string{
	// 	"name":  "Abu Ashraf Masnun",
	// 	"email": "masnun@gmail.com",
	// })

	var newPost entity.Post

	if err := c.Bind(&newPost); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	// This must be moved

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(newPost)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post("https://jsonplaceholder.typicode.com/posts", "application/json", &buf)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	c.String(http.StatusOK, string(body))
}

//Context stores information
func newPostHandler() *postHandler {
	return &postHandler{}
}
