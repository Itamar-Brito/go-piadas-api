package main

import (
    "net/http"
    "go-api-piadas/piadas"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    
    router.GET("/hello", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"message": "hello wts item testandoooo"})
    })

    router.GET("/kkk", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"message": "kkk"})
    })

	router.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"message": "bora carai"})
    })

    router.GET("/piadas", func(c *gin.Context) {
        piadas,_ := piadas.FetchPiadas()

        c.JSON(http.StatusOK,gin.H{"piadas": piadas } )
    })

    router.Run(":8000")
}