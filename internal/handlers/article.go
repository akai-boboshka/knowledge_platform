package handlers

import (
	"awesomeProject/pkg/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) AddArticle(c *gin.Context) {
	var article models.Article
	log.Printf("data before binding")
	if err := c.BindJSON(&article); err != nil {
		log.Printf("AddArtcile - c.BindJSON  error:  %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//article =

	log.Printf("AddArticle - data after binding: %v", article)

	createArticle, err := h.service.CreateArticle(&article)
	if err != nil {
		log.Printf("AddArticle - h.service.CreateArticle error:  %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("AddArticle - create article after binding: %v", createArticle)
	c.JSON(http.StatusCreated, gin.H{})
}

func (h *Handler) GetArticles(c *gin.Context) {
	articles, err := h.service.ListArticles()
	if err != nil {
		log.Printf("ListArticles - h.service.ListArticles error:  %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("ListArticles - list articles: %v", articles)
	c.JSON(http.StatusOK, gin.H{"data": articles})
}

func (h *Handler) GetArticleByID(c *gin.Context) {
	idstr := c.Param("id")

	if idstr == "" {
		log.Printf("GetArticleByID - c.Param(\"id\") not found")
		c.JSON(http.StatusNotFound, gin.H{"error": "id not found"})
		return
	}

	id, err := strconv.Atoi(idstr)
	if err != nil {
		log.Printf("GetArticleByID - strconv.Atoi error:  %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	article, err := h.service.FindArticleByID(id)
	if err != nil {
		log.Printf("GetArticleByID - h.service.FindArticleByID error:  %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("GetArticleByID - article: %v", article)
	c.JSON(http.StatusOK, gin.H{"data": article})
}

func (h *Handler) UpdateArticle(c *gin.Context) {
	idstr := c.Param("id")

	if idstr == "" {
		log.Printf("GetArticleByID - c.Param(\"id\") not found")
		c.JSON(http.StatusNotFound, gin.H{"error": "id not found"})
		return
	}

	id, err := strconv.Atoi(idstr)
	if err != nil {
		log.Printf("GetArticleByID - strconv.Atoi error:  %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var article models.Article
	article.ID = id

	log.Printf("data before binding")
	if err := c.BindJSON(&article); err != nil {
		log.Printf("UpdateArticle - c.BindJSON  error:  %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("UpdateArticle - data after binding: %v", article)

	updatedArticle, err := h.service.EditArticle(&article)
	if err != nil {
		log.Printf("UpdateArticle - h.service.EditArticle error:  %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("UpdateArticle - updated article: %v", updatedArticle)
	c.JSON(http.StatusOK, gin.H{"data": updatedArticle})
}

func (h *Handler) DeleteArticle(c *gin.Context) {
	idstr := c.Param("id")

	if idstr == "" {
		log.Printf("DeleteArticle - c.Param(\"id\") not found")
		c.JSON(http.StatusNotFound, gin.H{"error": "id not found"})
		return
	}

	id, err := strconv.Atoi(idstr)
	if err != nil {
		log.Printf("DeleteArticle - strconv.Atoi error:  %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if _, err := h.service.RemoveArticle(id); err != nil {
		log.Printf("DeleteArticle - h.service.RemoveArticle error:  %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("DeleteArticle - article with id: %v removed", id)
	c.JSON(http.StatusNoContent, gin.H{"data": "Article deleted"})
}

func (h *Handler) AddToReadLater(c *gin.Context) {
	idstr := c.Param("id")

	if idstr == "" {
		log.Printf("GetArticleByID - c.Param(\"id\") not found")
		c.JSON(http.StatusNotFound, gin.H{"error": "id not found"})
		return
	}

	id, err := strconv.Atoi(idstr)
	if err != nil {
		log.Printf("GetArticleByID - strconv.Atoi error:  %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	readLater, err := h.service.AddArticleToReadLater(id)
	if err != nil {
		log.Printf("AddArticleToReadLater - h.service.AddArticleToReadLater error:  %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("AddArticleToReadLater - readLater: %v", readLater)
	c.JSON(http.StatusCreated, gin.H{"data": readLater})
}

func (h *Handler) GetArticlesFromFavorites(c *gin.Context) {
	articlesFromFavorites, err := h.service.ListOfFavoritesArticles()
	if err != nil {
		log.Printf("ListOfFavoritesArticles error:  %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("ListOfFavoritesArticles - list articles: %v", articlesFromFavorites)
	c.JSON(http.StatusOK, gin.H{"data": articlesFromFavorites})
}

func (h *Handler) DeleteFromFavorites(c *gin.Context) {
	idstr := c.Param("id")

	if idstr == "" {
		log.Printf("GetArticleByID - c.Param(\"id\") not found")
		c.JSON(http.StatusNotFound, gin.H{"error": "id not found"})
		return
	}

	id, err := strconv.Atoi(idstr)
	if err != nil {
		log.Printf("GetArticleByID - strconv.Atoi error:  %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	deleteArticle, err := h.service.DeleteArticleFromFavorites(id)
	if err != nil {
		log.Printf("DeleteArticle - h.service.DeleteArticleFromFavorites error:  %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("DeleteArticle - deleteArticle: %v", deleteArticle)
	c.JSON(http.StatusNoContent, gin.H{"data": "Article deleted from favorites"})
}
