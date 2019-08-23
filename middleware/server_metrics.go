package middleware

import (
	"github.com/black-engine/base/entities"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
	"time"
)

func ServerMetrics( db *gorm.DB ) gin.HandlerFunc{
	return func(context *gin.Context) {
		start := time.Now()

		context.Next()

		request := entities.Request{}
		request.Path = context.Request.URL.Path
		request.Method = context.Request.Method
		request.Status = context.Writer.Status()

		go func(){
			request.ID = uuid.Must( uuid.NewV4() ).String()
			request.Created = time.Now()
			request.Latency = time.Since( start ).Nanoseconds()
			db.Save( &request )
		}()
	}
}