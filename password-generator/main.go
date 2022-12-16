package main

import (
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset)-1)]
	}
	return string(b)
}

func NewPassword() string{
	rangeStart := 5
	rangeEnd := 10
	offset := rangeEnd - rangeStart
	randLength := seededRand.Intn(offset) + rangeStart

	charSet := "aAbBcCdDeEfFgGhHiIjJkKlLmMnNoOpPqQrRsStTuUvVwWxXyYzZ"
	randString := StringWithCharset(randLength, charSet)
  return randString

}

func main() {

  //Start Server
	r := gin.Default()
  r.GET("/new", func(ctx *gin.Context) {
    ctx.JSON(200, gin.H{
      "password": NewPassword(),
      "teste": "123",
    })
  })
  r.Run()
}
