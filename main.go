package main
import (
  "net/http"
  "os"
)
func main() {
  port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}
  http.ListenAndServe(":" + port, http.FileServer(http.Dir(".")))
}
