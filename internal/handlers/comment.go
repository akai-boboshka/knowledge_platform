package handlers

import (
	"awesomeProject/pkg/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) AddComment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("FindArticle - h.service.FindArticleByID err: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var comment models.Comment

	if err := c.BindJSON(&comment); err != nil {
		log.Printf("AddComment - c.BindJSON err : %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	log.Printf("AddComment - data after binding: %v", comment)

	writeComment, err := h.service.CreateComment(id, &comment)
	if err != nil {
		log.Printf("AddComment - h.service.CreateComment err : %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	log.Printf("AddComment - write comment: %v", writeComment)
	c.JSON(http.StatusCreated, gin.H{"comment": writeComment})
}

//func (h *Handler) EditComment(c *gin.Context) {
//	var comment models.Comment
//
//	if err := c.BindJSON(&comment); err != nil {
//		log.Printf("UpdateArticle - c.BindJSON  error:  %v", err)
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	log.Printf("UpdateArticle - data after binding: %v", comment)
//
//	updatedComment, err := h.service.EditComment(id, &comment)
//	if err != nil {
//		log.Printf("UpdateArticle - h.service.EditArticle error:  %v", err)
//		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//		return
//	}
//
//	log.Printf("UpdateComment - updated comment: %v", updatedComment)
//	c.JSON(http.StatusOK, gin.H{"data": updatedComment})
//}

func (h *Handler) DeleteComment(c *gin.Context) {
	idstr := c.Param("id")

	if idstr == "" {
		log.Printf("DeleteComment - c.Param(\"id\") not found")
		c.JSON(http.StatusNotFound, gin.H{"error": "id not found"})
		return
	}

	id, err := strconv.Atoi(idstr)
	if err != nil {
		log.Printf("DeleteComment - strconv.Atoi error:  %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if _, err := h.service.RemoveArticle(id); err != nil {
		log.Printf("DeleteArticle - h.service.RemoveComment error:  %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("DeleteComment - coment with id: %v removed", id)
	c.JSON(http.StatusNoContent, gin.H{"data": "Comment deleted"})
}

func (h *Handler) GetComments(articleId int, c *gin.Context) {
	comment, err := h.service.AllCommentsOfArticle()
	if err != nil {
		log.Printf("GetComments - h.service.AllCommentsOfArticle err: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("GetComments - h.service.AllCommentsOfArticle result: %v", comment)
	c.JSON(http.StatusOK, gin.H{"data": comment})
}
