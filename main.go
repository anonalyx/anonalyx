package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type post struct {
	ID    string            `json:"id"`
	Title string            `json:"title"`
	Body  map[string]string `json:"body"`
	Date  string            `json:"price"`
}

func getPosts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, posts)
}

func getPost(c *gin.Context) {
	var postID int
	if err := c.BindJSON(&postID); err != nil {
		return
	}
	c.IndentedJSON(http.StatusOK, posts[postID])
}

func addPosts(c *gin.Context) {
	id := c.Param("id")

	for _, p := range posts {
		if p.ID == id {
			c.IndentedJSON(http.StatusOK, p)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "post not found"})
}

var posts = []post{
	{ID: "1", Title: "Week 1", Body: map[string]string{
		"What I accomplished last week":       "",
		"What I plan to accomplish this week": "",
		"What went well":                      "",
		"How can I improve the process":       "",
	}, Date: "08/28/2023"},
	{ID: "2", Title: "Week 1", Body: map[string]string{
		"What I accomplished last week":       "",
		"What I plan to accomplish this week": "",
		"What went well":                      "",
		"How can I improve the process":       "",
	}, Date: "08/28/2023"},
	{ID: "3", Title: "Week 1", Body: map[string]string{
		"What I accomplished last week":       "",
		"What I plan to accomplish this week": "",
		"What went well":                      "",
		"How can I improve the process":       "",
	}, Date: "08/28/2023"},
	{ID: "4", Title: "Week 1", Body: map[string]string{
		"What I accomplished last week":       "",
		"What I plan to accomplish this week": "",
		"What went well":                      "",
		"How can I improve the process":       "",
	}, Date: "08/28/2023"},
	{ID: "5", Title: "Week 1", Body: map[string]string{
		"What I accomplished last week":       "",
		"What I plan to accomplish this week": "",
		"What went well":                      "",
		"How can I improve the process":       "",
	}, Date: "08/28/2023"},
	{ID: "6", Title: "Week 1", Body: map[string]string{
		"What I accomplished last week":       "",
		"What I plan to accomplish this week": "",
		"What went well":                      "",
		"How can I improve the process":       "",
	}, Date: "08/28/2023"},
}

func main() {
	router := gin.Default()
	router.GET("/posts", getPosts)
	router.GET("/posts/:id", getPost)
	router.POST("/posts", addPosts)

	// Listen and Server in 0.0.0.0:8080
	router.Run(":8080")
}
