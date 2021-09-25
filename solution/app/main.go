package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Invoice struct {
	Currency   map[string]string `json:"currency"`
	Value      float64           `json:"value"`
	CustomerId int               `json:"customer_id"`
}

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func ChargeInvoice(c *gin.Context) {
	// Read request body and try to unmarshal (just verification example)
	body := c.Request.Body
	bs, err := ioutil.ReadAll(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status_code": http.StatusInternalServerError,
			"message":     "could not read request body",
		})
		panic(err)
	}

	fmt.Println(string(bs))

	var i Invoice
	err = json.Unmarshal(bs, &i)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": http.StatusBadRequest,
			"message":     "could not unmarshal request body",
		})
		return
	}

	rand.Seed(time.Now().UnixNano())
	r := rand.Float32() < 0.5
	c.JSON(200, gin.H{
		"result": r,
	})
}

const listenAddr = "0.0.0.0:8080"

func main() {
	r := gin.Default()

	r.GET("/ping", Ping)
	r.POST("/", ChargeInvoice)

	fmt.Println("Starting listening on", listenAddr)
	err := r.Run(listenAddr)
	if err != nil {
		panic(err)
	}
}
