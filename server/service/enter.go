package service

import (
	FIT "github.com/flipped-aurora/gin-vue-admin/server/service/bussiness/fit"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
	FitServiceGroup     FIT.ServiceGroup
}
