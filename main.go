package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

func main() {
	server := gin.Default()
	server.GET("/", getTasks)

	server.Run("localhost:8080")
}

func getTasks(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "I am going to be a go developer"})
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
