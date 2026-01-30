package handler

import (
	"log"
	"net/http"
	"time"

	ginJwt "github.com/appleboy/gin-jwt/v3"
	"github.com/gin-gonic/gin"
	goJwt "github.com/golang-jwt/jwt/v5"
)

type loginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type User struct {
	UserName  string
	FirstName string
	LastName  string
	password  string
}

// in-memory cache or temporary store for authorised users
var authorisedUserStore = map[string]User{
	"admin": User{UserName: "admin", FirstName: "ray", LastName: "snow", password: "admin"},
	"test":  User{UserName: "test", FirstName: "john", LastName: "biz", password: "test"},
}

var identityKey = "id"

func EnableJwtSecurity(engine *gin.Engine, publicGroupRouter *gin.RouterGroup, privateGroupRouter *gin.RouterGroup) {
	//create middleware instance
	authMiddleware, errIns := ginJwt.New(&ginJwt.GinJWTMiddleware{
		Realm:           "album zone",
		Key:             signingSecretKey(),
		Timeout:         (10 * time.Minute),
		MaxRefresh:      (15 * time.Minute),
		IdentityKey:     identityKey,
		PayloadFunc:     payloadFunc(),
		IdentityHandler: identityHandler(),
		Authenticator:   loginAuthenticator(),
		Authorizer:      authorizer(),
		Unauthorized:    unauthorized(),
		LogoutResponse:  logoutResponse(),
		TokenHeadName:   "Bearer",
		TimeFunc:        time.Now,
	})
	if errIns != nil {
		log.Fatalln("middleware instance creation failed", errIns)
		panic("middleware instance creation failed !!")
	}
	// initialize middleware
	errInit := authMiddleware.MiddlewareInit()
	if errInit != nil {
		log.Fatalln("middleware initialization failed", errInit)
		panic("middleware initialization failed !!")
	}

	// public endpoints
	publicGroupRouter.POST("/login", authMiddleware.LoginHandler)
	engine.NoRoute(authMiddleware.MiddlewareFunc(), noRoute())

	//private or protected endpoints
	privateGroupRouter.Use(authMiddleware.MiddlewareFunc())
	privateGroupRouter.POST("/refresh", authMiddleware.RefreshHandler)
	privateGroupRouter.POST("/logout", authMiddleware.LogoutHandler)

}

func signingSecretKey() []byte {
	return []byte("secret key value")
}

func payloadFunc() func(data any) goJwt.MapClaims {
	return func(data any) goJwt.MapClaims {
		log.Println("processing in payloadFunc..")
		if v, ok := data.(*User); ok {
			return goJwt.MapClaims{
				identityKey: v.UserName,
			}
		}
		return goJwt.MapClaims{}
	}
}

func identityHandler() func(c *gin.Context) any {
	return func(c *gin.Context) any {
		log.Println("processing in identityHandler..")
		claims := ginJwt.ExtractClaims(c)
		return &User{
			UserName: claims[identityKey].(string),
		}
	}
}

func noRoute() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(404, gin.H{"code": http.StatusNotFound, "message": "Requested page not found"})
	}
}

func loginAuthenticator() func(c *gin.Context) (any, error) {
	return func(c *gin.Context) (any, error) {
		log.Println("processing in loginAuthenticator..")
		var loginVals loginRequest
		if err := c.ShouldBind(&loginVals); err != nil {
			return "", ginJwt.ErrMissingLoginValues
		}
		userID := loginVals.Username
		password := loginVals.Password

		if user, exist := authorisedUserStore[userID]; exist && (user.password == password) {
			return &User{
				UserName:  userID,
				LastName:  user.LastName,
				FirstName: user.FirstName,
			}, nil
		}
		return nil, ginJwt.ErrFailedAuthentication
	}
}

func authorizer() func(c *gin.Context, data any) bool {
	return func(c *gin.Context, data any) bool {
		log.Println("processing in authorizer..")
		if u, ok := data.(*User); ok {
			if _, exist := authorisedUserStore[u.UserName]; exist {
				return true
			}
		}
		return false
	}
}

func unauthorized() func(c *gin.Context, code int, message string) {
	return func(c *gin.Context, code int, message string) {
		log.Println("processing in unauthorized..")
		c.JSON(code, gin.H{
			"code":    code,
			"message": message,
		})
	}
}

func logoutResponse() func(c *gin.Context) {
	return func(c *gin.Context) {
		claims := ginJwt.ExtractClaims(c)
		user, exists := c.Get(identityKey)
		log.Println("user logout:", claims, user, exists)

		response := gin.H{
			"code":    http.StatusOK,
			"message": "Successfully logged out",
		}

		if len(claims) > 0 {
			response["logged_out_user"] = claims[identityKey]
		}
		if exists {
			response["user_info"] = user.(*User).UserName
		}

		c.JSON(http.StatusOK, response)
	}
}
