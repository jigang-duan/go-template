package heart

import (
	"go-template/bootstrap"

	"github.com/kataras/iris/v12/core/router"
)

func Configure(b *bootstrap.Bootstrapper) {
	b.PartyFunc("/heart", func(p router.Party) {
		p.Use(middleware)

		p.Get("/ping", ping)
	})
}
