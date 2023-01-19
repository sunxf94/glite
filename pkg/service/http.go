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
	Handle(method, path string, handler func(c *gin.Context))
	Use(handler gin.HandlerFunc)
}

type Http struct {
	addr   []string
	engine *gin.Engine
}

func (h *Http) Handle(method, path string, handler func(c *gin.Context)) {
	h.engine.Handle(method, path, handler)
}

func (h *Http) Use(handler gin.HandlerFunc) {
	h.engine.Use(handler)
}

func (h *Http) Run() error {
	h.engine.Use()
	return h.engine.Run(h.addr...)
}
