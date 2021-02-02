package handler

import (
	"net/http"
	"strconv"

	"github.com/Yosh11/exemple_gin/model/todo"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createItem(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	listID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	var input todo.TodoItem
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.TodoItem.Create(userID, listID, input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllItems(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	listID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	items, err := h.services.TodoItem.GetAll(userID, listID)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, items)
}

func (h *Handler) getItemByID(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	itemID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	item, err := h.services.TodoItem.GetByID(userID, itemID)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, item)
}

func (h *Handler) updateItem(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	var input todo.UpdateItemInput
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.TodoItem.Update(userID, id, input); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) deleteItem(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	itemID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	err = h.services.TodoItem.DeleteByID(userID, itemID)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}
