package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jkim7113/centinal/model"
	"github.com/jkim7113/centinal/util"
)

type ArticleRepository interface {
	Create(ctx context.Context, article model.Article)
	Update(ctx context.Context, article model.Article)
	Delete(ctx context.Context, UUID string)
	FindById(ctx context.Context, UUID string) (model.Article, error)
	FindAll(ctx context.Context) []model.Article
	FindByCategory(ctx context.Context, category string) []model.Article
}

type ArticleRepositoryImpl struct {
	Db *sql.DB
}

func NewArticleRepository(Db *sql.DB) ArticleRepository {
	return &ArticleRepositoryImpl{Db: Db}
}

func (repo *ArticleRepositoryImpl) Create(ctx context.Context, article model.Article) {
	tx, err := repo.Db.Begin()
	util.PanicIfError(err)
	defer util.CommitOrRollback(tx)

	SQL := "INSERT INTO articles (Title, Body, ERT, Category, Thumbnail) VALUES (?, ?, ?, ?, ?)"
	_, err = tx.ExecContext(ctx, SQL, article.Title, article.Body, article.ERT, article.Category, article.Thumbnail)
	util.PanicIfError(err)
}

func (repo *ArticleRepositoryImpl) Update(ctx context.Context, article model.Article) {
	tx, err := repo.Db.Begin()
	util.PanicIfError(err)
	defer util.CommitOrRollback(tx)

	SQL := "UPDATE articles SET Title = ?, Body = ?, ERT = ?, Category = ?, Thumbnail = ? WHERE UUID = ?"
	_, err = tx.ExecContext(ctx, SQL, article.Title, article.Body, article.ERT, article.Category, article.Thumbnail, article.UUID)
	util.PanicIfError(err)
}

func (repo *ArticleRepositoryImpl) Delete(ctx context.Context, UUID string) {
	tx, err := repo.Db.Begin()
	util.PanicIfError(err)
	defer util.CommitOrRollback(tx)

	SQL := "DELETE FROM articles WHERE UUID = UNHEX(?)"
	_, err = tx.ExecContext(ctx, SQL, UUID)
	util.PanicIfError(err)
}

func (repo *ArticleRepositoryImpl) FindById(ctx context.Context, UUID string) (model.Article, error) {
	tx, err := repo.Db.Begin()
	util.PanicIfError(err)
	defer util.CommitOrRollback(tx)

	SQL := "SELECT HEX(UUID), Title, Body, Date, ERT, Thumbnail, Category FROM articles WHERE UUID = UNHEX(?)"
	result, errQuery := tx.QueryContext(ctx, SQL, UUID)
	util.PanicIfError(errQuery)
	defer result.Close()

	article := model.Article{}

	if result.Next() {
		err := result.Scan(&article.UUID, &article.Title, &article.Body, &article.Date, &article.ERT, &article.Thumbnail, &article.Category)
		util.PanicIfError(err)
		return article, nil
	} else {
		return article, errors.New("Couldn't find an article with such UUID")
	}
}

func (repo *ArticleRepositoryImpl) FindAll(ctx context.Context) []model.Article {
	tx, err := repo.Db.Begin()
	util.PanicIfError(err)
	defer util.CommitOrRollback(tx)

	SQL := "SELECT HEX(UUID), Title, Body, Date, ERT, Thumbnail, Category FROM articles"
	result, errQuery := tx.QueryContext(ctx, SQL)
	util.PanicIfError(errQuery)
	defer result.Close()

	var articles []model.Article

	for result.Next() {
		article := model.Article{}
		err := result.Scan(&article.UUID, &article.Title, &article.Body, &article.Date, &article.ERT, &article.Thumbnail, &article.Category)
		util.PanicIfError(err)

		articles = append(articles, article)
	}

	return articles
}

func (repo *ArticleRepositoryImpl) FindByCategory(ctx context.Context, category string) []model.Article {
	tx, err := repo.Db.Begin()
	util.PanicIfError(err)
	defer util.CommitOrRollback(tx)

	SQL := "SELECT HEX(UUID), Title, Body, Date, ERT, Thumbnail, Category FROM articles WHERE Category = ?"
	result, errQuery := tx.QueryContext(ctx, SQL, category)
	util.PanicIfError(errQuery)
	defer result.Close()

	var articles []model.Article

	for result.Next() {
		article := model.Article{}
		err := result.Scan(&article.UUID, &article.Title, &article.Body, &article.Date, &article.ERT, &article.Thumbnail, &article.Category)
		util.PanicIfError(err)

		articles = append(articles, article)
	}

	return articles
}
