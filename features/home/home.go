package home

import (
	"go-template/bootstrap"

	"github.com/kataras/iris/v12/core/router"
)

var handler = newHomeHandler()

func Configure(b *bootstrap.Bootstrapper) {
	b.PartyFunc("/", func(p router.Party) {

		p.Get("/", handler.index)
	})
}
