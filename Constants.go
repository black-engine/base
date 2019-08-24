package base

type Error struct {
	Message string
	Type string
}

var InvalidTokenError = Error{Message: "Invalid or non existing token", Type: "INVALID_TOKEN"}
var UnbindableError = Error{Message: "Unable to bind request", Type: "UNBINDABLE"}
var WrongCredentialsError = Error{Message: "Wrong credentials", Type: "WRONG_CREDENTIALS"}
