package controllers

import (
	"fmt"
	"github.com/revel/revel"
	"math/rand"
)

type Post struct {
	App
}

type Topic struct {
	ID int32 `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
}

func (this Post) Index() revel.Result {
	text := this.Params.Query.Get("text")
	title := "Post Content"
	id := rand.Intn(100)
	response := Topic{int32(id), title, text}
	fmt.Println(response)

	return this.RenderJSON(response)
	// return this.Render(text, title)

}
