package handler

import (
	"github.com/bonNope/makvesTask/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("get-items", h.getEntities)

	return router
}

type GetEntitiesResponse struct {
	Ids []int `json:"ids,omitempty"`
}

func (h *Handler) getEntities(ctx *gin.Context) {
	resp := &GetEntitiesResponse{}
	err := ctx.ShouldBindJSON(&resp)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, `Bad json format: should be {"ids":[1,2,3,4]}`)
	}

	result := h.service.GetEntitiesByIds(resp.Ids)

	ctx.JSON(http.StatusOK, result)
}
