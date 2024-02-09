First Run go mod init
it will create a go.mod file in the project folder

#install gin 
go get -u github.com/gin-gonic/gin

#install orom 
go get -u gorm.io/gorm

#install mysql driver
go get -u gorm.io/driver/mysql

# hot reload
go install github.com/cosmtrek/air@latest



now create a main.go file and paste this code 
############################
package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping",
		})
	})
	r.Run()
}

##############

#not run
go run main.go 

this will start the project by default this project run on localhost:8080
