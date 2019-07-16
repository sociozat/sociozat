package controllers

import (
	"fmt"
	"math/rand"

	"github.com/revel/revel"
)

type PostC struct {
	AppC
}

type Topic struct {
	ID      int32  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (this PostC) Index() revel.Result {
	text := this.Params.Query.Get("text")
	title := "Post Content"
	id := rand.Intn(100)
	response := Topic{int32(id), title, text}
	fmt.Println(response)

	return this.RenderJSON(response)
	// return this.Render(text, title)

}
