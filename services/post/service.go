package post

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/epa-datos/exercise-api/entity"
)

type PostService interface {
}

type Post struct {
}

//Creating a new POST
func (p *Post) CreatePost(newPost entity.Post) []byte {
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

	return body
}

//Returning all the posts
func (p *Post) AllPosts() []*entity.Post {
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
	return posts
}

// Returning one post by the number ID within the Path
func (p *Post) OnePost(number string) *entity.PostDTO {
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
	return &postResponse
}
