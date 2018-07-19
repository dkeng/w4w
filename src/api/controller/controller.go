package controller

import (
	"github.com/dkeng/w4w/src/store"
)

// AllController 所有控制器
type AllController struct {
	HomeController  *HomeController
	ShortController *ShortController
}

//Init 初始化所有Controller
func (a *AllController) Init(s *store.Store) *AllController {
	// 初始化业务逻辑层
	allStore := new(store.AllStore).Init(s)
	// 初始化Controller
	a.HomeController = new(HomeController).Init(allStore)
	a.ShortController = new(ShortController).Init(allStore)
	return a
}
