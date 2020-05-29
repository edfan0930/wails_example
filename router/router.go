package router

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/middleware"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

type (
	//CustomValidatora
	CustomValidator struct {
		validator *validator.Validate
	}
)

//Validate registered validator
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

//Router --
func Router(e *echo.Echo) {

	e.Use(middleware.Recover(), middleware.CORSWithConfig(middleware.DefaultCORSConfig))

	e.GET("/first", func(c echo.Context) error {
		fmt.Println("in first")
		return c.JSON(http.StatusOK, "success")
	})
	/* 	tool9527 := lotto.Group("/9527")

	   	tool9527.GET("/version", controllers.Version)

	   	//ping mongo orderlotto
	   	tool9527.GET("/pingorderlotto", controllers.PingMongoOrderLotto)

	   	//ping mongo index
	   	tool9527.GET("/pingindex", controllers.PingMongoIndex)

	   	//switch debug level
	   	tool9527.POST("/debuglevel", controllers.DebugLevel)

	   	//debug := tool9527.Group("/debug/pprof")

	*/ //debug.GET("/", pprof.Index)
}

//Auth --
//middleware
func Auth(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {

		/* 		//recover panic
		   		defer func(c echo.Context) {

		   			if err := recover(); err != nil {
		   				zap.L().Error("recover :", zap.Any("panic", err), zap.String("path", c.Path()))
		   				c.JSON(http.StatusOK, response.Error(response.PanicErr, "has panic"))
		   			}
		   		}(c)

		   		//取token 並判斷是否為空
		   		token := strings.TrimSpace(c.Request().Header.Get("systoken"))
		   		if token != env.SysToken {
		   			return c.JSON(http.StatusOK, response.Error(response.AuthErr, errors.New("Auth Failed").Error()))
		   		}
		*/return next(c)
	}
}

/*
func requestID() echo.MiddlewareFunc {

	return middleware.RequestIDWithConfig(middleware.RequestIDConfig{
		Generator: snowflake.Text,
	})
}
*/
