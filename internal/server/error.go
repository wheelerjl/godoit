package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s Server) RouteNotFound(ctx *gin.Context) {
	ErrorResponse(ctx, http.StatusNotFound, "route not found", nil)
}

func ErrorResponse(ctx *gin.Context, code int, msg string, err error) {
	if err != nil {
		msg = fmt.Sprintf("%s: %v", msg, err)
	}
	ctx.JSON(code, DefaultResponse{Status: "error", Message: msg})
}
