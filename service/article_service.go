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
	Delete(ctx context.Context, UUID string)
	FindById(ctx context.Context, UUID string) response.ArticleResponse
	FindAll(ctx context.Context) []response.ArticleResponse
	FindByCategory(ctx context.Context, category string) []response.ArticleResponse
}

type ArticleServiceImpl struct {
	ArticleRepository repository.ArticleRepository
}

func NewArticleService(articleRepository repository.ArticleRepository) ArticleService {
	return &ArticleServiceImpl{ArticleRepository: articleRepository}
}

func (service *ArticleServiceImpl) Create(ctx context.Context, request request.ArticleCreateRequest) {
	ERT := util.EstimateReadingTime(request.Body)
	article := model.Article{
		Title:     request.Title,
		Body:      request.Body,
		ERT:       ERT,
		Category:  request.Category,
		Thumbnail: request.Thumbnail,
	}
	service.ArticleRepository.Create(ctx, article)
}

func (service *ArticleServiceImpl) Update(ctx context.Context, request request.ArticleUpdateRequest) {
	article, err := service.ArticleRepository.FindById(ctx, request.UUID)
	util.PanicIfError(err)

	ERT := util.EstimateReadingTime(request.Body)
	article = model.Article{
		UUID:      request.UUID,
		Title:     request.Title,
		Body:      request.Body,
		ERT:       ERT,
		Category:  request.Category,
		Thumbnail: request.Thumbnail,
	}
	service.ArticleRepository.Update(ctx, article)
}

func (service *ArticleServiceImpl) Delete(ctx context.Context, UUID string) {
	_, err := service.ArticleRepository.FindById(ctx, UUID)
	util.PanicIfError(err)
	service.ArticleRepository.Delete(ctx, UUID)
}

func (service *ArticleServiceImpl) FindById(ctx context.Context, UUID string) response.ArticleResponse {
	article, err := service.ArticleRepository.FindById(ctx, UUID)
	util.PanicIfError(err)

	return response.ArticleResponse(article)
}

func (service *ArticleServiceImpl) FindAll(ctx context.Context) []response.ArticleResponse {
	articles := service.ArticleRepository.FindAll(ctx)
	var articleResponse []response.ArticleResponse

	for _, v := range articles {
		article := response.ArticleResponse{UUID: v.UUID, Body: v.Body, Title: v.Title, Date: v.Date, ERT: v.ERT, Category: v.Category}
		if len(article.Body) > 150 {
			article.Body = article.Body[:150] + "..."
		}
		articleResponse = append(articleResponse, article)
	}
	return articleResponse
}

func (service *ArticleServiceImpl) FindByCategory(ctx context.Context, category string) []response.ArticleResponse {
	articles := service.ArticleRepository.FindByCategory(ctx, category)
	var articleResponse []response.ArticleResponse

	for _, v := range articles {
		article := response.ArticleResponse{UUID: v.UUID, Body: v.Body, Title: v.Title, Date: v.Date, ERT: v.ERT, Category: v.Category}
		if len(article.Body) > 150 {
			article.Body = article.Body[:150] + "..."
		}
		articleResponse = append(articleResponse, article)
	}
	return articleResponse
}
