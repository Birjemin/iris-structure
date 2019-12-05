package middleware

import (
	"github.com/kataras/iris/v12"
)

func Handler(ctx iris.Context) {
	// do something like auth
	ctx.Next()
}
