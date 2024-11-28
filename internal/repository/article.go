package repository

import (
	"awesomeProject/pkg/models"
	"fmt"
	"log"
)

func (r *Repository) AddArticle(a *models.Article) (*models.Article, error) {
	r.Log.Info("Start adding the article to the database")
	result := r.db.Model(&a).Select("title", "articles_text").Create(&a)
	if result.Error != nil {
		return nil, result.Error
	}
	r.Log.Info("End adding the article to the database")
	return a, nil
}

func (r *Repository) GetArticles() ([]models.Article, error) {
	var a []models.Article

	// select * from articles;
	err := r.db.Find(&a).Error
	if err != nil {
		log.Printf("GetArticles: Failed to get articles: %v\n", err)
		return nil, fmt.Errorf("Cannot find article with error: : %v", err)
	}

	return a, nil

}

func (r *Repository) GetArticleByID(id int) (*models.Article, error) {
	var article *models.Article
	log.Printf("get article by id: %v\n", id)

	result := r.db.First(&article, id)
	if result.Error != nil {
		fmt.Println("Ошибка при выполнении запроса:", result.Error)
		return nil, result.Error
	}

	return article, nil
}

func (r *Repository) GetArticleByProfile(profileID int) ([]models.Article, error) {
	var a []models.Article
	err := r.db.Where("profile_id = ?", profileID).Find(&a).Error
	if err != nil {
		log.Printf("GetArticleByProfile: Failed to get articles: %v\n", err)
		return nil, fmt.Errorf("cannot find article by profileID with error: %v", err)
	}

	return a, nil
}

func (r *Repository) UpdateArticle(a *models.Article) (*models.Article, error) {
	result := r.db.Model(&a).Updates(&a)
	if result.Error != nil {
		log.Printf("UpdateArticle: Failed to update article: %v\n", result.Error)
		return nil, fmt.Errorf("Failed to update article: %v\n", result.Error)
	}

	return a, nil
}

func (r *Repository) DeleteArticle(id int) (int, error) {
	// delete from articles where article_id = id returning id
	result := r.db.Model(&models.Article{}).
		Where("id = ?", id).
		Update("active", false)
	if result.Error != nil {
		log.Printf("DeleteArticle: Failed to delete article: %d\n", result.Error)
		return 0, fmt.Errorf("Failed to delete article: %d\n", result.Error)
	}

	return id, nil
}

func (r *Repository) AddToReadLater(a *models.Article) error {
	//readLater := models.ReadLater{ArticleID: ID}
	//err := r.db.Create(&)
	//if err != nil {
	//	log.Printf("AddToReadLater: Failed to add article to readLater: %v\n", err)
	//	return fmt.Errorf("Failed to add article to readLater: %v\n", err)
	//}

	query := `insert into readLater (article_id) values(?) returning *`
	err := r.db.Raw(query, a.ID)
	if err != nil {
		log.Printf("AddToReadLater: Failed to insert article: %v\n", err)
		return fmt.Errorf("Failed to insert article: %v\n", err)
	}
	return nil
}

func (r *Repository) RemoveArticleFromReadLater(ID int) (int, error) {
	result := r.db.Model(&models.ReadLater{}).
		Where("id = ?", ID).
		Update("active", false)
	if result.Error != nil {
		log.Printf("DeleteArticle: Failed to delete article: %v\n", result.Error)
		return 0, fmt.Errorf("Failed to delete article from favorites: %v\n", result.Error)
	}

	return ID, nil
}

func (r *Repository) GetListOfFavoritesArticles() ([]models.ReadLater, error) {
	var rl []models.ReadLater

	err := r.db.Find(&rl).Error
	if err != nil {
		log.Printf("GetListOfFavoritesArticles: Failed to get the list of readLater: %v\n", err)
		return nil, fmt.Errorf("Failed to get the list of readLater: %v\n", err)
	}

	return rl, nil
}

func (r *Repository) GetArticleFromFavoriteById(ID int) (*models.ReadLater, error) {
	var a *models.ReadLater
	err := r.db.Where("id = ?", ID).First(&a).Error
	if err != nil {
		log.Printf("GetArticleFromFavourite: Failed to get the article from readLater: %v\n", err)
		fmt.Errorf("Failed to get the article from readLater: %v\n", err)
		return nil, err
	}

	return a, nil
}
