package controller

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"fmt"
	"../"
	"../service"
)

var (
	Hello *helloController
)

func init() {
	Hello = &helloController{}
	app.Router.GET("/hello",[]string{"auth"},Hello.GET)
}

type helloController struct {

}

func (controller helloController) GET(r *http.Request, ps httprouter.Params) string  {
	result := service.UserService.All()
	var (
		id int
		name string
	)

	for result.Next() {

		var attr1 string
		var attr2 string
		var attr3 string
		var attr4 string
		var attr5 string
		var attr6 string
		var attr7 string
		var attr8 string
		var attr9 string
		var attr10 string
		var attr11 string
		var attr12 string
		var attr13 string
		var attr14 string
		err := result.Scan(&attr1, &attr2, &attr3, &attr4, &attr5, &attr6, &attr7, &attr8, &attr9, &attr10, &attr11, &attr12, &attr13, &attr14)
		fmt.Print(err)
		fmt.Println(attr1)
		fmt.Println(attr2)
		fmt.Println(attr3)
		fmt.Println(attr4)
	}
	result.Scan(&id,&name)
	fmt.Print(id,name)
	return "hello from helloController"
}