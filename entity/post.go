package entity

type Post struct {
	UserID        int    `json:"user_id"`
	Identificador int    `json:"id" binding:"required"` //tags are used to show the response as the object
	Title         string `json:"title"`
	Body          string `json:"body"`
}

type PostDTO struct {
	UserID        int    `json:"userId"`
	Identificador int    `json:"id" binding:"required"` //tags are used to show the response as the object
	Title         string `json:"title"`
	Body          string `json:"body"`
}

func (p *PostDTO) ConvertToPost() *Post {
	return &Post{
		UserID:        p.UserID,
		Identificador: p.Identificador,
		Title:         p.Title,
		Body:          p.Body,
	}
}
