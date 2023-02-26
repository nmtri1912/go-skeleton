package handler

import (
	"fmt"
	"go-skeleton/internal/request"
	service "go-skeleton/internal/service/example"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterExampleHandler(router *gin.RouterGroup, exampleService service.ExampleService) {
	handler := exampleHandler{
		exampleService: exampleService,
	}

	router.GET("/info", handler.info)
}

type exampleHandler struct {
	exampleService service.ExampleService
}

func (h exampleHandler) info(ctx *gin.Context) {
	var (
		req request.ExampleRequest
	)

	if err := ctx.BindJSON(&req); err != nil {
		fmt.Println("binding request failed", err.Error())
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	example, err := h.exampleService.CreateExample(ctx, &req)
	if err != nil {
		fmt.Println("create example failed", err.Error())
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, example)

}
