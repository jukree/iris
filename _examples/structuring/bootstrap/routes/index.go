package routes

import (
	"github.com/jukree/iris"
)

// GetIndexHandler handles the GET: /
func GetIndexHandler(ctx iris.Context) {
	ctx.ViewData("Title", "Index Page")
	ctx.View("index.html")
}
