package middleware

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
)

func MyMiddleware2() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// pre-handle
		fmt.Println("pre-handle")

		c.Next(ctx) // call the next middleware(handler)
		// post-handle
		fmt.Println("post-handle")
	}
}
