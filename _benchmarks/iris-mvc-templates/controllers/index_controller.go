package controllers

import "github.com/jukree/iris/mvc"

type IndexController struct{ mvc.Controller }

func (c *IndexController) Get() {
	c.Data["Title"] = "Home Page"
	c.Tmpl = "index.html"
}
