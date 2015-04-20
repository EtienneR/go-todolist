package main

import (
	"github.com/gin-gonic/gin"
	"todolist/db"
)

var dbmap = db.InitDb()

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("views/*")

	r.GET("/", indexHandler)
	r.POST("/submit", postHandler)
	r.GET("/delete/:id", deleteHandler)

	r.Run(":3000")
}

func indexHandler(c *gin.Context) {
	messages := []db.Messages{}

	_, err := dbmap.Select(&messages, "SELECT * FROM messages")

	if err != nil || len(messages) == 0 {
		obj := gin.H{"error": "No message in the database"}
		c.HTML(200, "index.html", obj)
	} else {
		obj := gin.H{"title": messages}
		c.HTML(200, "index.html", obj)
	}
}

func postHandler(c *gin.Context) {
	c.Request.ParseForm()

	title := c.Request.Form.Get("title")

	if title != "" {
		data := &db.Messages{0, title}
		dbmap.Insert(data)
	}

	c.Redirect(301, "/")
}

func deleteHandler(c *gin.Context) {
	id := c.Params.ByName("id")
	dbmap.Exec("DELETE FROM messages WHERE id=?", id)

	c.Redirect(301, "/")
}
