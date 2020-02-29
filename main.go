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
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		// log.Fatal("$PORT must be set")
		port = "3333"
	}
	router := gin.Default()        // r实际上是route的缩写
	router.LoadHTMLGlob("views/*") // glob模式
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html") //单文件模式
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})
	router.GET("/ping", func(ctx *gin.Context) {
		var name string = ctx.Query("name")
		fmt.Println(name)
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// port = "3333"
	router.Run(":" + port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
