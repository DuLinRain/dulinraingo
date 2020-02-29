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
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

// XResponseTime 测试
func XResponseTime() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		log.Println("before request")
		ctx.Next()
		log.Println("after request")
		latency := fmt.Sprintf("%d", time.Since(start).Microseconds())
		// log.Println("after request", ""+latency.Microseconds()+"ms")
		ctx.Header("X-Response-Time", latency+"ms")
	}
}

// ArticlelistController Articlelist控制器
func ArticlelistController() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Println("before")
		ctx.Next()
		category := ctx.DefaultQuery("category", "all")
		log.Println("category", category)
		if category == "" {
			ctx.JSON(200, gin.H{})
		} else {
			ctx.JSON(200, gin.H{
				"title": "Main website",
				"id":    "112212",
			})
		}
	}
}

// MyLogger 测试中间件
func MyLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		t := time.Now()
		// before request
		log.Println("before request")
		ctx.Next()
		// after request
		log.Println("after request")
		latency := time.Since(t)
		log.Print("latency: ", latency)
		var value, ok = ctx.Get("example")
		// access the status we are sending
		status := ctx.Writer.Status()
		log.Println("status", status, value, ok)
		// ctx.JSON(200, gin.H{
		// 	"message1": "pong",
		// })
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		// log.Fatal("$PORT must be set")
		port = "3333"
	}
	router := gin.Default() // gin Default和gin New的区别 Default With the Logger and Recovery middleware already attached
	router.Use(XResponseTime())
	router.LoadHTMLGlob("views/*") // glob模式
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html") //单文件模式
	router.GET("/", func(ctx *gin.Context) {
		ctx.Header("Name", "my name is smallsoup")
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})
	router.GET("/articlelist", ArticlelistController())
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
