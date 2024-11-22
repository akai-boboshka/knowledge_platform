package handlers

import (
	"awesomeProject/pkg/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) SendMessageToSupport(c *gin.Context) {
	var message models.Support

	if err := c.BindJSON(&message); err != nil {
		log.Printf("SendMessage - c.BindJSON  error:  %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("SendMessage - data after binding: %v", message)

	//sendMessage, err := h.service.SendMessage(&message)
	//if err != nil {
	//	log.Printf("SendMessage - h.service.SendMessage error:  %v", err)
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//	return
	//}

	//log.Printf("SendMessage - create message after binding: %v", sendMessage)
	c.JSON(http.StatusCreated, gin.H{})
}

func (h *Handler) GetMessageById(c *gin.Context) {
	idStr := c.Param("id")

	if idStr == "" {
		log.Printf("GetMessages - c.Param id not found")
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("GetMessages - c.Param id error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	message, err := h.service.GetMessageById(id)
	if err != nil {
		log.Printf("GetMessages - h.service.GetMessageById error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("GetMessages - message: %v", message)
	c.JSON(http.StatusOK, gin.H{"message": message})
}

func (h *Handler) GetAllMessages(c *gin.Context) {
	messages, err := h.service.GetAllMessages()
	if err != nil {
		log.Printf("GetAllMessages - h.service.GetAllMessages error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("GetAllMessages - messages: %v", messages)
	c.JSON(http.StatusOK, gin.H{"messages": messages})
}

func (h *Handler) DeleteMessage(c *gin.Context) {
	idStr := c.Param("id")

	if idStr == "" {
		log.Printf("DeleteMessage - c.Param id not found")
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("DeleteMessage - c.Param id error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := h.service.DeleteMessage(id); err != nil {
		log.Printf("DeleteMessage - h.service.DeleteMessage error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("DeleteMessage - message: %v", id)
	c.JSON(http.StatusOK, gin.H{})
}
