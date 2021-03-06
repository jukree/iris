package main

import (
	"github.com/jukree/iris/_examples/structuring/bootstrap/bootstrap"
	"github.com/jukree/iris/_examples/structuring/bootstrap/middleware/identity"
	"github.com/jukree/iris/_examples/structuring/bootstrap/routes"
)

func main() {
	app := bootstrap.New("Awesome App", "kataras2006@hotmail.com")
	app.Bootstrap()
	app.Configure(identity.Configure, routes.Configure)
	app.Listen(":8080")
}
