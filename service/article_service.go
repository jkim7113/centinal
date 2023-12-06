package service

import (
	"context"

	"github.com/jkim7113/centinal/model"
	"github.com/jkim7113/centinal/model/request"
	"github.com/jkim7113/centinal/model/response"
	"github.com/jkim7113/centinal/repository"
	"github.com/jkim7113/centinal/util"
)

type ArticleService interface {
	Create(ctx context.Context, request request.ArticleCreateRequest)
	Update(ctx context.Context, request request.ArticleUpdateRequest)
	Delete(ctx context.Context, Id int)
	FindById(ctx context.Context, Id int) response.ArticleResponse
	FindAll(ctx context.Context) []response.ArticleResponse
}

type ArticleServiceImpl struct {
	ArticleRepository repository.ArticleRepository
}

func NewArticleService(articleRepository repository.ArticleRepository) ArticleService {
	return &ArticleServiceImpl{ArticleRepository: articleRepository}
}

func (repo *ArticleServiceImpl) Create(ctx context.Context, request request.ArticleCreateRequest) {
	ERT := util.EstimateReadingTime(request.Body)
	article := model.Article{
		Title:    request.Title,
		Body:     request.Body,
		ERT:      ERT,
		Category: request.Category,
	}
	repo.ArticleRepository.Create(ctx, article)
}

func (repo *ArticleServiceImpl) Update(ctx context.Context, request request.ArticleUpdateRequest) {
	article, err := repo.ArticleRepository.FindById(ctx, request.Id)
	util.PanicIfError(err)

	ERT := util.EstimateReadingTime(request.Body)
	article = model.Article{
		Id:       request.Id,
		Title:    request.Title,
		Body:     request.Body,
		ERT:      ERT,
		Category: request.Category,
	}
	repo.ArticleRepository.Update(ctx, article)
}

func (repo *ArticleServiceImpl) Delete(ctx context.Context, Id int) {
	_, err := repo.ArticleRepository.FindById(ctx, Id)
	util.PanicIfError(err)
	repo.ArticleRepository.Delete(ctx, Id)
}

func (repo *ArticleServiceImpl) FindById(ctx context.Context, Id int) response.ArticleResponse {
	article, err := repo.ArticleRepository.FindById(ctx, Id)
	util.PanicIfError(err)

	return response.ArticleResponse(article)
}

func (repo *ArticleServiceImpl) FindAll(ctx context.Context) []response.ArticleResponse {
	articles := repo.ArticleRepository.FindAll(ctx)
	var articleResponse []response.ArticleResponse

	for _, v := range articles {
		article := response.ArticleResponse{Id: v.Id, Body: v.Body, Title: v.Title, Date: v.Date, ERT: v.ERT}
		articleResponse = append(articleResponse, article)
	}
	return articleResponse
}
