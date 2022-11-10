package handler

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type Args struct {
	Query      string   `query:"query"` // query参数
	QuerySlice []string `query:"q"`
	Path       string   `path:"path"`     // path参数
	Header     string   `header:"header"` // header参数
	Form       string   `form:"form"`     // form参数
	Json       string   `json:"json"`     // json body参数
	Vd         int      `query:"vd" vd:"$==0||$==1"`
}

// ?query=hello&q=q1&q=q2&vd=1
func Hello(ctx context.Context, c *app.RequestContext) {
	var arg Args
	if err := c.BindAndValidate(&arg); err != nil {
		panic(err)
	}
	fmt.Println(arg)
	c.JSON(200, utils.H{
		"message": "success",
	})
}
