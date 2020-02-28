// package main
// import (
//   "net/http"
//   "os"
//   "log"
// )
// func main() {
//   port := os.Getenv("PORT")

// 	if port == "" {
// 		log.Fatal("$PORT must be set")
// 	}
//   http.ListenAndServe(":" + port, http.FileServer(http.Dir(".")))
// }

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	r := gin.Default() // r实际上是route的缩写
	r.GET("/ping", func(c *gin.Context) {
		var name string = c.Query("name")
		fmt.Println(name)
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":" + port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
