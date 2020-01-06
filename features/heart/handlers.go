package heart

import (
	"time"

	"github.com/kataras/iris/v12"
)

func ping(ctx iris.Context) {
	ctx.Header("date", time.Now().Local().String())
	ctx.JSON(iris.Map{"message": "pong"})
}
