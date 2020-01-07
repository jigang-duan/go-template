package home

import (
	"github.com/kataras/iris/v12"
)

type homeHandler struct {
}

func newHomeHandler() *homeHandler {
	return &homeHandler{}
}

func (h *homeHandler) index(ctx iris.Context) {
	ctx.ViewData("pageTitle", "主页 - Kitchen")
	ctx.View("home/index.pug")
}
