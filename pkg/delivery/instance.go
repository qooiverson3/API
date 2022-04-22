package delivery

import (
	"ces-api/pkg/model"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type InstanceHandler struct {
	Svc model.Service
}

func NewInstanceHandler(s model.Service) *InstanceHandler {
	return &InstanceHandler{
		Svc: s,
	}
}

// Router path
func (h *InstanceHandler) Router(e *gin.Engine) {
	e.GET("/api/v1/instance", h.GetInstanceList)
	e.PATCH("/api/v1/instance/action", h.Actions)
}

func IsHeaderValidate(ctx *gin.Context) bool {
	header := &model.XHeader{}
	err := ctx.ShouldBindHeader(header)
	var state bool = true

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"state": false,
			"message": []string{
				fmt.Sprintf("token error: %v", err.Error()),
			},
		})
		state = false
		return state
	}

	token := ctx.GetHeader("token")
	if token != os.Getenv("TOKEN") {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"state": false,
			"message": []string{
				"illegal token.",
			},
		})
		state = false
		return state
	}

	return state
}

// GetInstance _
func (h *InstanceHandler) GetInstanceList(ctx *gin.Context) {
	headerValidate := IsHeaderValidate(ctx)
	if !headerValidate {
		return
	}

	var r model.GetInstanceForm
	ctx.Bind(&r)

	validate := validator.New()
	err := validate.Struct(&r)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	data := h.Svc.GetInstanceList(r)
	ctx.JSON(http.StatusOK, gin.H{
		"state": true,
		"data":  data,
	})
}

func (h *InstanceHandler) Actions(ctx *gin.Context) {
	headerValidate := IsHeaderValidate(ctx)
	if !headerValidate {
		return
	}

	var body model.ActionRequestBody
	ctx.BindJSON(&body)

	validate := validator.New()
	err := validate.Struct(&body)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"state":   false,
			"message": err.Error(),
		})
		return
	}

	result := h.Svc.Actions(body)
	if result == 0 {
		ctx.AbortWithStatusJSON(http.StatusNoContent, gin.H{
			"state": false,
			"message": []string{
				"illegal UUID.",
			},
		})
		return
	}
	ctx.JSON(http.StatusAccepted, gin.H{})
}
