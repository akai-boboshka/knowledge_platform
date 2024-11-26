package service

import (
	"awesomeProject/pkg/models"
	"bufio"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"log"
	"os"
	"strings"
	"unicode"
)

var (
	ErrRecordNotFound  = errors.New("record not found")
	ErrNoArticlesFound = errors.New("no articles found")
	ErrUsersIDNotFound = errors.New("ID not found")
)

func (s *Service) CreateArticle(a *models.Article) (*models.Article, error) {
	//_, err := s.Repository.GetProfileByID(a.ProfileID)
	//if err == nil {
	//	if errors.As(err, &ErrRecordNotFound) {
	//		return nil, fmt.Errorf("Profile with ID %d doesn't exist", a.ProfileID)
	//	}
	//	return nil, err
	//}

	//articles, err := s.Repository.GetArticleByProfile(a.ID)
	//if err != nil {
	//	return nil, err
	//}
	//
	//if len(articles) > 0 {
	//	for _, article := range articles {
	//		if article.Title == a.Title {
	//			return nil, fmt.Errorf("The article with title %s is already exists", a.Title)
	//		}
	//	}
	//}

	// Create the file
	file, err := os.Create("article.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Write the content of article to file
	data := []byte(a.ArticlesText)
	_, err = file.Write(data)
	if err != nil {
		return nil, err
	}
	file.Close()

	//Слова, которые не должны присутсвовать в содержании статьи
	badWords := []string{"Убийство", "Терроризм", "Насилие", "Сигареты"}

	file, err = os.Open("article.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	isSeparator := func(r rune) bool {
		return unicode.IsSpace(r) || unicode.IsPunct(r)
	}

	scanner := bufio.NewScanner(file)
	var words []string
	for scanner.Scan() {
		// Разбиваем строку на слова по пробелам
		//lineWords := strings.Fields(scanner.Text())
		lineWords := strings.FieldsFunc(scanner.Text(), isSeparator)
		words = append(words, lineWords...)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка при сканировании:", err)
		return nil, err
	}

	for _, badWord := range badWords {
		badWordLower := strings.ToLower(badWord)
		for _, word := range words {
			wordLower := strings.ToLower(word)
			if wordLower == badWordLower {
				log.Printf("Your article contains inappropriate words like: %s, so it will be removed", word)
			}
		}
	}

	return s.Repository.AddArticle(a)
}

func (s *Service) ListArticles() ([]models.Article, error) {
	articles, err := s.Repository.GetArticles()
	if err != nil {
		return nil, err
	}
	if len(articles) == 0 {
		return nil, ErrNoArticlesFound
	}

	return s.Repository.GetArticles()
}

func (s *Service) FindArticleByID(id int) (*models.Article, error) {
	article, err := s.Repository.GetArticleByID(id)
	if err != nil {
		s.Log.WithFields(logrus.Fields{
			"id":    id,
			"error": err.Error(),
		}).Error("Failed to fetch article")
		return nil, err
	}
	if article == nil {
		s.Log.WithFields(logrus.Fields{
			"id": id,
		}).Warn("Article not found")
		return nil, fmt.Errorf("Article with ID: %d didn't found", id)
	}

	return article, nil
}

func (s *Service) FindArticleByProfile(profileID int) ([]models.Article, error) {
	article, err := s.Repository.GetArticleByProfile(profileID)
	if err != nil {
		return nil, err
	}
	if len(article) == 0 {
		return nil, fmt.Errorf("Article with profileID %d didn't found", profileID)
	}

	return article, nil
}

func (s *Service) EditArticle(a *models.Article) (*models.Article, error) {
	_, err := s.Repository.GetArticleByID(a.ID)
	if err != nil {
		if errors.Is(err, ErrRecordNotFound) {
			return nil, fmt.Errorf("Article with id %d not found", a.ID)
		}
		return nil, err
	}

	return s.Repository.UpdateArticle(a)
}

func (s *Service) RemoveArticle(id int) (int, error) {
	_, err := s.Repository.GetArticleByID(id)
	if err != nil {
		if errors.Is(err, ErrRecordNotFound) {
			return 0, fmt.Errorf("Article with id: %d not found", id)
		}
		return 0, err
	}

	return s.Repository.DeleteArticle(id)
}

//func (s *Service) AddArticleToReadLater(ID int) (int, error) {
//	_, err := s.Repository.GetArticleFromFavoriteById(ID) // ошибка возвращается в случае если такой записи нету или другие случаи
//	if err != nil {
//		if errors.As(err, &gorm.ErrRecordNotFound) { // проверяем на существование такой записи
//			err := s.Repository.AddToReadLater(&ID) // добавляем запись
//			if err != nil {
//				return 0, err
//			}
//			return ID, nil
//		}
//		return 0, err
//	}
//
//	return 0, nil
//}

func (s *Service) ListOfFavoritesArticles() ([]models.ReadLater, error) {
	articles, err := s.Repository.GetListOfFavoritesArticles()
	if err != nil {
		return nil, err
	}
	if len(articles) == 0 {
		return nil, ErrNoArticlesFound
	}

	return s.Repository.GetListOfFavoritesArticles()
}

func (s *Service) GetArticleFromFavoritesById(id int) (*models.ReadLater, error) {
	article, err := s.Repository.GetArticleFromFavoriteById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrRecordNotFound
		}
		return nil, err
	}

	return article, nil
}

func (s *Service) DeleteArticleFromFavorites(articleId int) (int, error) {
	_, err := s.Repository.GetArticleFromFavoriteById(articleId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, ErrRecordNotFound
		}
		return 0, err
	}

	return s.Repository.RemoveArticleFromReadLater(articleId)
}
