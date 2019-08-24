package handlers

import (
	"github.com/gin-gonic/gin"
)

type Error struct {
	Message string
	Type string
}

var InvalidTokenError = Error{Message: "Invalid or non existing token", Type: "INVALID_TOKEN"}
var UnbindableError = Error{Message: "Unable to bind request", Type: "UNBINDABLE"}
var WrongCredentialsError = Error{Message: "Wrong credentials", Type: "WRONG_CREDENTIALS"}


func BuildOkApiResponse( context *gin.Context , value interface{} ) {
	context.JSON( 200 , gin.H{
		"isError" : false,
		"data" : value,
	})
}

func BuildErrorApiResponse( context *gin.Context , err Error ) {
	context.JSON( 400 , gin.H{
		"isError" : true,
		"type" : err.Type ,
		"message" : err.Message,
	})
}
