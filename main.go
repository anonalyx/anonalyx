package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type post struct {
	ID    string            `json:"id"`
	Title string            `json:"title"`
	Body  map[string]string `json:"body"`
	Date  string            `json:"price"`
}

func getHome(c *gin.Context) []post {
	return posts
}
func getPosts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, posts)
}

func getPost(c *gin.Context) *post {
	postID, _ := strconv.Atoi(c.Param("id"))
	fmt.Printf("%v", postID)
	return &posts[postID]
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
	{ID: "2", Title: "Week 2", Body: map[string]string{
		"What I accomplished last week":       "Over the last week I talked to my engineering manager about projects that might coincide with the requirements of this class and be a good fit for the semester long project. I am currently thinking about a relevant exercise with Docker, Kubernetes, and AWS, as these are three technologies I am very interested in.\n\n",
		"What I plan to accomplish this week": "Over the next week, I will finalize a proposal for that plan and submit it.\n\n",
		"What went well":                      "I think that it was encouraging that my Engineering manager was supportive of my project, and he had great advice on things I should consider\n\n",
		"How can I improve the process":       "I am obviously just getting started, so a regular cadence of accountability and deliverables will be critical to establishing a productive work process",
	}, Date: "08/28/2023"},
	{ID: "3", Title: "Week 3", Body: map[string]string{
		"What I accomplished last week":       "",
		"What I plan to accomplish this week": "",
		"What went well":                      "",
		"How can I improve the process":       "",
	}, Date: "08/28/2023"},
	{ID: "4", Title: "Week 4", Body: map[string]string{
		"What I accomplished last week":       "This week, I decided to pivot from my original docker, kubernetes, AWS project idea to one that more closely aligns with my interest in blockchain and cryptocurrencies. I happened to work through the short CryptoZombies tutorial that is an introduction to working with web3 libraries, and at the end they give you 10 ideas on how to round out the tutorial into a full blown application however they leave the actual research and implementation decisions up to you.",
		"What I plan to accomplish this week": "Complete project proposal",
		"What blockers did I face":            "I don't yet know whether I can sketch out and complete a full project proposal ",
		"How can I improve the process":       "Put into practice concrete project management processes, document milestones, and record artifacts",
	}, Date: "08/28/2023"},
	{ID: "5", Title: "Week 5", Body: map[string]string{
		"What I accomplished last week":       "",
		"What I plan to accomplish this week": "",
		"What went well":                      "",
		"How can I improve the process":       "",
	}, Date: "08/28/2023"},
	{ID: "6", Title: "Week 6", Body: map[string]string{
		"What I accomplished last week":       "Last week and this week, I've been heavily focused on learning the Go programming language, as it is one of the languages that is compatible with many of the Web3 libraries. I completed several tutorials, walked through much of the documentation, and began to work with the Gin web application framework.\n",
		"What I plan to accomplish this week": "I plan to refactor my personal website in Go, and host it on a cloud service.\n",
		"What went well":                      "The documentation on Gin specifically is very light, and it's been a bit of a process tracking down all the relevant tutorials and references.\n",
		"How can I improve the process":       "This past week, I was focused on really just learning the language, however over the coming weeks, I will be more focused on implementing and building things in the language to solidify my knowledge.\n",
	}, Date: "08/28/2023"},
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Alex Cullen",
			"posts": posts,
		})
	})
	router.GET("/posts", getPosts)
	router.GET("/post/:id", func(c *gin.Context) {
		c.HTML(http.StatusOK, "post.html", gin.H{
			"title": getPost(c).Title,
			"body":  getPost(c).Body,
		})
	})
	router.POST("/posts", addPosts)

	// Listen and Server in 0.0.0.0:8080
	router.Run(":8080")
}
