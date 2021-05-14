package controllers

import (
	"github.com/gomodule/redigo/redis"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

type Controller struct {
	Db      *gorm.DB
	Redis   redis.Conn
	RedisDb string
}

type Meta struct {
	Code    uint              `json:"code,omitempty"`
	Message string            `json:"message,omitempty"`
	Error   string            `json:"error,omitempty"`
	Errors  map[string]string `json:"errors,omitempty"`
}

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data,omitempty"`
}

func (c *Controller) Data(ctx echo.Context, data interface{}) error {
	response := Response{
		Meta: Meta{
			Code:    http.StatusOK,
			Message: "Operation successfully executed.",
			Error:   "",
			Errors:  nil,
		},
		Data: data,
	}
	return ctx.JSON(http.StatusOK, response)
}

func (c *Controller) Ok(ctx echo.Context) error {
	response := Response{
		Meta: Meta{
			Code:    http.StatusOK,
			Message: "Operation successfully executed.",
			Error:   "",
			Errors:  nil,
		},
		Data: nil,
	}
	return ctx.JSON(http.StatusOK, response)
}

func (c *Controller) BadRequest(ctx echo.Context, err string) error {
	response := Response{
		Meta: Meta{
			Code:    http.StatusBadRequest,
			Message: "Oops... Something went wrong.",
			Error:   err,
			Errors:  nil,
		},
		Data: nil,
	}
	return ctx.JSON(http.StatusBadRequest, response)
}

func (c *Controller) InternalServerError(ctx echo.Context, err error) error {
	response := Response{
		Meta: Meta{
			Code:    http.StatusInternalServerError,
			Message: "Oops... Something went wrong.",
			Error:   err.Error(),
			Errors:  nil,
		},
		Data: nil,
	}
	return ctx.JSON(http.StatusInternalServerError, response)
}
