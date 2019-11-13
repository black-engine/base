package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/black-engine/base/handlers"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strings"
)

func ExecuteJsonPost(url string, headers map[string]string, body interface{}) (*http.Response, error) {
	data , err := json.Marshal( body )
	if err != nil {
		return nil, err
	}

	request , err := http.NewRequest( "POST" , url , bytes.NewBuffer( data ) )
	if err != nil {
		return nil, err
	}

	for key , value := range headers {
		request.Header.Set( key , value )
	}
	request.Header.Set( "Content-Type" , "application/json" )

	return http.DefaultClient.Do( request )
}

func ExecuteGet(url string, headers map[string]string) (*http.Response, error) {
	request , err := http.NewRequest( "GET" , url , nil )
	if err != nil {
		return nil, err
	}

	for key , value := range headers {
		request.Header.Set( key , value )
	}
	request.Header.Set( "Content-Type" , "application/json" )

	return http.DefaultClient.Do( request )
}

func GetBodyFromResponse( response *http.Response ) []byte {
	defer response.Body.Close()

	responseBody , err := ioutil.ReadAll( response.Body )

	if err != nil {
		return nil
	}

	return responseBody
}

func GetCookieDomainFromContext( context *gin.Context ) string {
	host := context.Request.Host
	if strings.LastIndex( host , "127.0.0.1" ) > -1 {
		return host
	}
	if strings.LastIndex( host , "localhost" ) > -1 {
		return host
	}
	elements := strings.Split( host , "." )
	if len( elements ) < 2 {
		return fmt.Sprintf(".%s" , host )
	}
	for i:= 2; i <= len( elements ); i++{
		if len( elements[ len( elements ) - i ] ) > 3 { //if block is longer than 3 chars is not considered a part of tld
			return fmt.Sprintf(".%s" , strings.Join( elements[len(elements)-i:] , ".") )
		}
	}
	return ""
}

func GetBaseDomainFromContext( context *gin.Context ) string {
	host := context.Request.Host
	if strings.LastIndex( host , "127.0.0.1" ) > -1 {
		return host
	}
	if strings.LastIndex( host , "localhost" ) > -1 {
		return host
	}
	elements := strings.Split( host , "." )
	if len( elements ) < 2 {
		return host
	}
	for i:= 2; i <= len( elements ); i++{
		if len( elements[ len( elements ) - i ] ) > 3 { //if block is longer than 3 chars is not considered a part of tld
			return strings.Join( elements[len(elements)-i:] , ".")
		}
	}
	return ""
}

func MustBoolean(u bool, err error) bool {
	if err != nil {
		panic(err)
	}
	return u
}

func MustByteArray(u []byte, err error) []byte {
	if err != nil {
		panic(err)
	}
	return u
}

func MustExist( v interface{}, exists bool ) interface{} {
	if !exists {
		panic( handlers.Error{ Message: "Value is not present" } )
	}
	return v
}