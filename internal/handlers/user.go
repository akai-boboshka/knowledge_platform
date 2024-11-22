package handlers

import (
	"awesomeProject/pkg/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) SignUp(c *gin.Context) {
	var user models.User

	err := c.BindJSON(&user)
	if err != nil {
		log.Printf("SignUp - c.BindJSON error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("SignUp - data after unmarshalling: %v", user)

	createUser, err := h.service.CreateUser(&user)
	if err != nil {
		log.Printf("SignUp - h.service.CreateUser error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//log.Printf("SignUp - response to client: %v", string(createUser))

	c.JSON(http.StatusOK, gin.H{"data": createUser})
}

func (h *Handler) SignIn(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		log.Printf("SignIn - c.BindJSON error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	token, err := h.service.SignIn(&user)
	if err != nil {
		log.Printf("SignIn - h.service.SignIn error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (h *Handler) GetUsers(c *gin.Context) {
	users, err := h.service.ListOfUsers()
	if err != nil {
		log.Printf("GetUsers - h.service.ListUsers error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": users})
}

func (h *Handler) FindUserById(c *gin.Context) {
	idStr := c.Param("id")

	if idStr == "" {
		log.Printf("GetUserByID - id is required")
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("GetUserByID - strconv.Atoi error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	user, err := h.service.GetUserByID(id)
	if err != nil {
		log.Printf("GetUserByID - h.service.GetUserByID error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func (h *Handler) UpdateUser(c *gin.Context) {
	idstr := c.Param("id")
	log.Println("id from param", idstr)

	if idstr == "" {
		log.Printf("GetUserByID - c.Param(\"id\") not found")
		c.JSON(http.StatusNotFound, gin.H{"error": "id not found"})
		return
	}

	id, err := strconv.Atoi(idstr)
	if err != nil {
		log.Printf("GetUserByID - strconv.Atoi error:  %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	user.ID = id

	log.Println("user_id before binding", id)

	if err = c.ShouldBindJSON(&user); err != nil {
		log.Printf("UpdateUser - c.BindJSON error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Println("user_handler_func", user)

	err = h.service.EditUser(&user)
	if err != nil {
		log.Printf("UpdateUser - h.service.EditUser error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "done"})
}

func (h *Handler) DeleteUser(c *gin.Context) {
	idStr := c.Param("id")

	if idStr == "" {
		log.Printf("DeleteUser - id is required")
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("DeleteUser - strconv.Atoi error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := h.service.DeleteUser(id); err != nil {
		log.Printf("DeleteUser - h.service.DeleteUser error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": true})
}
