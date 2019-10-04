package models

import (
	"github.com/jinzhu/gorm"
	"simple-REST/pkg/config"
)

var dbArticleClient *gorm.DB
var articlesPerPage = 5

type Article struct {
	gorm.Model
	// Id string `json:"id"`
	AuthorId    int64  `gorm:""json:"author_id"`
	Content     string `gorm:""json:"content"`
	Description string `gorm:""json:"description"`
	Title       string `gorm:""json:"title"`
}

func init() {
	config.Connect()
	dbArticleClient = config.GetDB()
	dbArticleClient.AutoMigrate(&Article{})
}

func (a *Article) NewArticle() *Article {
	dbArticleClient.NewRecord(a)
	dbArticleClient.Create(&a)
	return a
}

func GetAllArticles() []Article {
	var articles []Article
	dbArticleClient.Find(&articles)
	return articles
}

func GetArticleById(id int64) (*Article, *gorm.DB) {
	var article Article
	db := dbArticleClient.Where("id = ?", id).Find(&article)
	return &article, db
}

func GetArticlesByPageNumber(page int64) []Article {
	var wantedArticles []Article
	dbArticleClient.Order("created_at desc").Limit(articlesPerPage).Find(&wantedArticles)
	return wantedArticles
}

func DeleteArticleById(id int64) Article {
	var article Article
	dbArticleClient.Where("id = ?", id).Delete(article)
	return article
}
