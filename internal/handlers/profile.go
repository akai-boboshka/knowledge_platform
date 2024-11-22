package handlers

import (
	"awesomeProject/pkg/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) CreateProfile(c *gin.Context) {
	var profile models.Profile

	if err := c.BindJSON(&profile); err != nil {
		log.Printf("CreateProfile - BindJSON error: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "You must enter your details"})
		return
	}

	log.Printf("CreateProfile - data after binding: %v\n", profile)

	createProfile, err := h.service.CreateProfile(&profile)
	if err != nil {
		log.Printf("CreateProfile - CreateProfile error: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	log.Printf("CreateProfile - created profile: %v\n", createProfile)
	c.JSON(http.StatusCreated, gin.H{"profile": createProfile})
}

func (h *Handler) ListProfiles(c *gin.Context) {
	profiles, err := h.service.ListProfiles()
	if err != nil {
		log.Printf("ListProfiles - h.service.ListProfiles error: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("ListProfiles - list profiles: %v\n", profiles)
	c.JSON(http.StatusOK, gin.H{"profiles": profiles})
}

func (h *Handler) GetProfileByID(c *gin.Context) {
	idStr := c.Param("id")

	if idStr == "" {
		log.Printf("GetProfileByID id is empty\n")
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("GetProfileByID id error: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is invalid"})
		return
	}

	profile, err := h.service.GetProfileByID(id)
	if err != nil {
		log.Printf("GetProfileByID id error: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("GetProfileByID id: %v\n", profile)
	c.JSON(http.StatusOK, gin.H{"profile": profile})
}

func (h *Handler) UpdateProfile(c *gin.Context) {
	var profile models.Profile

	if err := c.BindJSON(&profile); err != nil {
		log.Printf("UpdateProfile - BindJSON error: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedProfile, err := h.service.EditProfile(&profile)
	if err != nil {
		log.Printf("UpdateProfile - h.service.EditProfile error: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("UpdateProfile - updated profile: %v\n", updatedProfile)
	c.JSON(http.StatusOK, gin.H{"profile": updatedProfile})
}

func (h *Handler) DeleteProfile(c *gin.Context) {
	idStr := c.Param("id")

	if idStr == "" {
		log.Printf("DeleteProfile id is empty\n")
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("DeleteProfile id error: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	deletedProfile, err := h.service.DeleteProfile(id)
	if err != nil {
		log.Printf("DeleteProfile id error: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("DeleteProfile id: %v\n", deletedProfile)
	c.JSON(http.StatusOK, gin.H{"profile": deletedProfile})
}
