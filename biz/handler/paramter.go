package handler

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func QueryParam(ctx context.Context, c *app.RequestContext) {
	firstname := c.DefaultQuery("firstname", "Guest")
	// shortcut for c.Request.URL.Query().Get("lastname")
	lastname := c.Query("lastname")
	// Iterate all queries and store the one with meeting the conditions in favoriteFood
	// 遍历所有query参数
	var favoriteFood []string
	c.QueryArgs().VisitAll(func(key, value []byte) {
		if string(key) == "food" {
			favoriteFood = append(favoriteFood, string(value))
		}
	})

	c.String(consts.StatusOK, "Hello %s %s, favorite food: %s", firstname, lastname, favoriteFood)
}

func FormParam(ctx context.Context, c *app.RequestContext) {
	// content-type : application/x-www-form-urlencoded
	name := c.PostForm("name")
	message := c.PostForm("message")

	c.PostArgs().VisitAll(func(key, value []byte) {
		if string(key) == "name" {
			fmt.Printf("This is %s!", string(value))
		}
	})
	c.String(consts.StatusOK, "name: %s; message: %s", name, message)

	// content-type : multipart/form-data
	id := c.FormValue("id")
	name_ := c.FormValue("name")
	message_ := c.FormValue("message")

	c.String(consts.StatusOK, "id: %s; name: %s; message: %s\n", id, name_, message_)
}

func CookieParam(ctx context.Context, c *app.RequestContext) {
}
