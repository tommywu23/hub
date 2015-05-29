package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

type Image struct {
	ID     string `json:"id"`
	Vesion string `json:"version"`
}

func main() {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("/config", configGet)
	}

	r.Run(":6600")

}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func configGet(c *gin.Context) {
	res := Image{ID: getKey(), Vesion: "v3.0.1"}
	c.JSON(200, res)
}

func getKey() string {
	size := 16

	rb := make([]byte, size)
	_, err := rand.Read(rb)

	if err != nil {
		fmt.Println(err)
	}

	rs := base64.URLEncoding.EncodeToString(rb)

	return rs
}
