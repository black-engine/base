package middleware

import (
	"github.com/black-engine/base/entities"
	"github.com/black-engine/base/helpers"
	"github.com/black-engine/uasurfer"
	"github.com/gin-gonic/gin"
	"github.com/go-http-utils/headers"
	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
	"net/url"
	"time"
)

func Prospects( db *gorm.DB ) gin.HandlerFunc{
	return func(context *gin.Context) {
		if p , err := context.Cookie( "p" ); err != nil && len( p ) == 36 {
			context.Set( "p" , p )
			return //it has prospect
		}

		prospect := entities.Prospect{}
		prospect.ID = uuid.Must( uuid.NewV4() ).String()
		prospect.Created = time.Now()
		prospect.Updated = prospect.Created

		if c , err := context.Cookie( "c" ); err != nil && len( c ) == 36 {
			prospect.CampaignID = &c
		}

		ua := uasurfer.Parse( context.GetHeader( headers.UserAgent ) )

		prospect.Ip = context.GetHeader( "CF-Connecting-IP" )
		prospect.Country = context.GetHeader( "CF-IPCountry" )
		prospect.IsBot = ua.IsBot()
		prospect.BrowserName = ua.Browser.Name.StringTrimPrefix()
		prospect.BrowserVersion = ua.Browser.Version.Major
		prospect.Platform = ua.OS.Platform.StringTrimPrefix()
		prospect.OsName = ua.OS.Name.StringTrimPrefix()
		prospect.OsVersion = ua.OS.Version.Major
		prospect.DeviceType = ua.DeviceType.StringTrimPrefix()
		prospect.Visits = 1
		prospect.Timestamp = time.Now()

		prospect.Referrer = context.GetHeader( headers.Referer )

		if u , e2 := url.Parse( prospect.Referrer ) ; e2 == nil {
			prospect.ReferrerHost = u.Host
		}
		prospect.Language = context.GetHeader( headers.AcceptLanguage )
		if len( prospect.Language ) > 5 {
			prospect.Language = prospect.Language[0:5]
		}

		context.SetCookie( "p" , prospect.ID , 2419200 , "/" , helpers.GetCookieDomainFromContext( context ) , false , false )
		context.Set( "p" , prospect.ID )

		go func(){
			db.Save( &prospect )
		}()
	}
}