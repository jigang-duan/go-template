package heart

import "github.com/kataras/iris/v12"

func middleware(ctx iris.Context) {
	ctx.Application().Logger().Infof("Runs before %s", ctx.Path())
	ctx.Next()
}
