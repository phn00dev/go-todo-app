package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/phn00dev/go-todo-app"
)

func (h *Handler) signUp(c *gin.Context) {
	var inputs todo.User
	if err := c.BindJSON(&inputs); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Authorization.CreateUser(inputs)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id":      id,
		"message": "user registered successfully",
	})
}

func (h *Handler) signIn(c *gin.Context) {

}
