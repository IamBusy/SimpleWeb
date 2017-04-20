package app

import (
	"./controller"
)

func init()  {
	Router.GET("/hello",[]string{"auth"},controller.Hello.GET)
}