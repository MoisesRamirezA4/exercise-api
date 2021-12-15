package post

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/epa-datos/exercise/github.com/epa-datos/exercise-api/entity"
)

type postService interface {
}

type post struct {
}

func (p post) allPosts(newPost entity.Post) []byte {
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
