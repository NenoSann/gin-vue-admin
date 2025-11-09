package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	fit "github.com/flipped-aurora/gin-vue-admin/server/router/fit"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
	FIT     fit.RouterGroup
}
