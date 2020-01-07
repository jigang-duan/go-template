package features

import (
	"go-template/bootstrap"
	"go-template/features/heart"
	"go-template/features/home"
)

func Configure(b *bootstrap.Bootstrapper) {
	b.Configure(heart.Configure)
	b.Configure(home.Configure)
}
