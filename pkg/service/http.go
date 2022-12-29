package service

import "github.com/gin-gonic/gin"

func NewHTTP(addr ...string) IHttp {
	return &Http{
		engine: gin.Default(),
		addr:   addr,
	}
}

type IHttp interface {
	Run() error
	HandleGET(path string, hello func(c *gin.Context))
	HandlePOST(path string, hello func(c *gin.Context))
}

type Http struct {
	addr   []string
	engine *gin.Engine
}

func (h *Http) HandleGET(path string, handler func(c *gin.Context)) {
	h.engine.GET(path, handler)
}

func (h *Http) HandlePOST(path string, handler func(c *gin.Context)) {
	h.engine.POST(path, handler)
}

func (h *Http) Run() error {
	return h.engine.Run(h.addr...)
}
