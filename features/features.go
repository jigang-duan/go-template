package features

import (
	"go-template/bootstrap"
	"go-template/features/heart"
)

func Configure(b *bootstrap.Bootstrapper) {
	b.Configure(heart.Configure)
}
