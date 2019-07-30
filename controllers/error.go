package controllers

type ErrorController struct {
	BaseController
}

func (c *ErrorController) Error404() {
	c.Data["content"] = "Page not found"
	c.TplName = "error/404.html"
}
