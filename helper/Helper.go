package helper

import (
	"github.com/rohmanhakim/rh-vue-todo/model"
	"net/http"
	"github.com/unrolled/render"
)

func RenderErrorResponse(status int, rw http.ResponseWriter, ren *render.Render){
	var res model.CommonResponse
	switch status {
	case 500:
		res.Status = 500
		res.Success = false
		ren.JSON(rw,http.StatusInternalServerError,res)
	case 502:
		res.Status = 502
		res.Success = false
		ren.JSON(rw,http.StatusBadGateway,res)
	}

}

func RenderOKResponse(rw http.ResponseWriter, ren *render.Render){
	var res model.CommonResponse
	res.Status = 200
	res.Success = true
	ren.JSON(rw,http.StatusOK,res)
}