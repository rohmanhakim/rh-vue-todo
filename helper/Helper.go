package helper

import (
	"github.com/rohmanhakim/rh-vue-todo/model"
	"net/http"
	"github.com/unrolled/render"
)

func RenderErrorResponse(status int, res model.CommonResponse, rw http.ResponseWriter, ren *render.Render){
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
