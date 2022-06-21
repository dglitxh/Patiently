package auth

import (
	"time"

	"github.com/dglitxh/patiently/controllers/middleware"
	"github.com/dglitxh/patiently/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func SignJWT(username, jwt_secret string) *models.JWTOutput {
	expiry := time.Now().Add(30 * time.Minute)

	claims := &models.Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiry.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString([]byte(jwt_secret))

	if err != nil {
		panic(err)
	}

	jwtOutput := &models.JWTOutput{
		Token:   tokenStr,
		Expires: expiry,
	}

	return jwtOutput
}

func RegAuthRoutes(router *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	r := router.Group("/auth")
	r.POST("/signup", h.SignupHandler)
	r.POST("/login", h.LoginHandler)
	r.GET("/logout", Logout)

	auth := r.Use(middleware.AuthMiddleware())
	auth.GET("/users", h.GetUsers)
	auth.GET("/user/:id", h.GetUser)
	auth.GET("/refresh", RefreshJwt)
}
