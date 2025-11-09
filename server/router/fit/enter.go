package fit

import (
	api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
)

type RouterGroup struct {
	FITUploadRouter
}

var (
	fitFileUploadApi = api.ApiGroupApp.FitApiGroup
)
