package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/phn00dev/go-todo-app"
)

func (h *Handler) createList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input todo.TodoList
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	// call service method
	id, err := h.services.TodoList.Create(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}

type getAllListsResponse struct {
	Data []todo.TodoList `json:"data"`
}

func (h *Handler) getAllLists(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	lists, err := h.services.TodoList.GetAll(userId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, getAllListsResponse{
		Data: lists,
	})
}

func (h *Handler) getListById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	list, err := h.services.TodoList.GetOne(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, list)

}

func (h *Handler) updateList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var inputs todo.UpdateListInput
	if err = c.BindJSON(&inputs); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	// update service
	if err := h.services.TodoList.Update(userId, listId, inputs); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) deleteList(c *gin.Context) {

	userId, err := getUserId(c)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	// delete list service
	err = h.services.TodoList.Delete(userId, listId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
