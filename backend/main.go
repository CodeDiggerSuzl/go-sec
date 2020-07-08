package main

import (
	"github.com/kataras/iris"
)

func main() {
	// create iris instance
	app := iris.New()
	// set debug mode
	app.Logger().SetLevel("debug")
	// register view
	template := iris.HTML("./backend/web/views/", ".html").Layout("shared/layout.html")
	app.RegisterView(template)
	// set template target
	app.StaticWeb("/assets", "./backend/web/assets")
	// set error page
	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("message", ctx.Values().GetStringDefault("message", "访问的页面出错"))
		ctx.ViewLayout("")
		_ = ctx.View("shared/error.html")
	})

	// register controller

	// run the app
	_ = app.Run(
		iris.Addr(":8080"),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)

}
