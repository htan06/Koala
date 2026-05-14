package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"koala.com/internal/config"
	"koala.com/internal/repository"
)

func main() {
	Db := config.ConnectionDb()
	defer Db.Close()

	userRepository := repository.NewUserRepository(Db)

	r := gin.Default()
	
	r.GET("/users", func(c *gin.Context) {
		users := userRepository.FindAll()
    
		var usersString string
		for _, u := range users {
			usersString += u.ToString()
		}

    	c.JSON(http.StatusOK, gin.H{
      	"message": usersString,
    	})
  	})

	r.Run()
}