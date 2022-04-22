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

func IsHeaderValidate(e *gin.Context) bool {
	header := &model.XHeader{}
	err := e.ShouldBindHeader(header)
	var state bool = true

	if err != nil {
		e.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"state": false,
			"message": []string{
				fmt.Sprintf("token error: %v", err.Error()),
			},
		})
		state = false
		return state
	}

	token := e.GetHeader("token")
	if token != os.Getenv("TOKEN") {
		e.AbortWithStatusJSON(http.StatusForbidden, gin.H{
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
func (h *InstanceHandler) GetInstanceList(e *gin.Context) {
	headerValidate := IsHeaderValidate(e)
	if !headerValidate {
		return
	}

	var r model.GetInstanceForm
	e.Bind(&r)

	validate := validator.New()
	err := validate.Struct(&r)
	if err != nil {
		e.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	data := h.Svc.GetInstanceList(r)
	e.JSON(http.StatusOK, gin.H{
		"state": true,
		"data":  data,
	})
}

func (h *InstanceHandler) Actions(e *gin.Context) {
	headerValidate := IsHeaderValidate(e)
	if !headerValidate {
		return
	}

	var r model.ActionRequestBody
	e.BindJSON(&r)

	validate := validator.New()
	err := validate.Struct(&r)
	if err != nil {
		e.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	if err != nil {
		e.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"state":   false,
			"message": err.Error(),
		})
		return
	}
	e.JSON(http.StatusOK, r)
}
