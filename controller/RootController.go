package controller

import (
	"net/http"
	"github.com/unrolled/render"
)

func RootHandler(rw http.ResponseWriter, req *http.Request, ren *render.Render){
	ren.HTML(rw,http.StatusOK,"index",nil)
}