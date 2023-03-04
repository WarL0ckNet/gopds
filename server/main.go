package server

import (
	"fmt"

	"github.com/kataras/iris/v12"
)

func Start(host string, port int) {
	app := iris.New()
	app.Favicon("./server/assets/img/favicon.png")
	app.HandleDir("/img", "./server/assets/img")
	app.HandleDir("/css", "./server/assets/css")
	app.HandleDir("/js", "./server/assets/js")
	app.RegisterView(iris.Amber("./server/views", ".amber"))
	app.I18n.Load("./server/locales/*/*")
	app.I18n.SetDefault("en-US")
	app.Get("/", index)
	app.Listen(fmt.Sprintf("%s:%d", host, port))
}

func index(ctx iris.Context) {
	data := iris.Map{
		"Title":      "Page Title",
		"FooterText": "Footer contents",
		"Message":    ctx.Tr("hi"),
	}

	if err := ctx.View("index.amber", data); err != nil {
		ctx.HTML("<h3>%s</h3>", err.Error())
		return
	}
}
