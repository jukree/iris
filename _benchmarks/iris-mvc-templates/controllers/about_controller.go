package controllers

import "github.com/jukree/iris/mvc"

type AboutController struct{ mvc.Controller }

func (c *AboutController) Get() {
	c.Data["Title"] = "About"
	c.Data["Message"] = "Your application description page."
	c.Tmpl = "about.html"
}
