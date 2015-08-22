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

// Display all messages
func indexHandler(c *gin.Context) {
	messages := []db.Messages{}
	_, err := dbmap.Select(&messages, "SELECT * FROM messages")

	var obj interface{}

	if err != nil || len(messages) == 0 {
		obj = gin.H{"error": "No message in the database"}
	} else {
		obj = gin.H{"title": messages}
	}

	c.HTML(200, "index.html", obj)
}

// Insert message from the form
func postHandler(c *gin.Context) {
	c.Request.ParseForm()

	// Get value from field title ("name=title")
	title := c.Request.Form.Get("title")

	if title != "" {
		data := &db.Messages{0, title}
		dbmap.Insert(data)
	}

	c.Redirect(301, "/")
}

// Delete a message (by id)
func deleteHandler(c *gin.Context) {
	// Get id value from URL
	id := c.Params.ByName("id")
	dbmap.Exec("DELETE FROM messages WHERE id=?", id)

	c.Redirect(301, "/")
}
