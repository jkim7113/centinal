package service

import (
	"context"

	"github.com/jkim7113/centinal/model"
	"github.com/jkim7113/centinal/model/request"
	"github.com/jkim7113/centinal/model/response"
	"github.com/jkim7113/centinal/repository"
	"github.com/jkim7113/centinal/util"
)

type UserService interface {
	Create(ctx context.Context, request request.UserCreateRequest)
	Update(ctx context.Context, request request.UserUpdateRequest)
	UpdatePassword(ctx context.Context, UUID string, password string)
	UpdateRole(ctx context.Context, UUID string, role string)
	VerifyEmail(ctx context.Context, email string)
	Delete(ctx context.Context, UUID string)
	FindById(ctx context.Context, UUID string) response.UserResponse
}

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &UserServiceImpl{UserRepository: userRepository}
}

func (service *UserServiceImpl) Create(ctx context.Context, request request.UserCreateRequest) {
	user := model.User{
		Username: request.Username,
		Email:    request.Email,
		Pw:       request.Pw,
		Bio:      request.Bio,
		PFP:      request.PFP,
	}
	service.UserRepository.Create(ctx, user)
}

func (service *UserServiceImpl) Update(ctx context.Context, request request.UserUpdateRequest) {
	user, err := service.UserRepository.FindById(ctx, request.UUID)
	util.PanicIfError(err)

	user = model.User{
		UUID:     request.UUID,
		Username: request.Username,
		Email:    request.Email,
		Bio:      request.Bio,
		PFP:      request.PFP,
	}
	service.UserRepository.Update(ctx, user)
}

func (service *UserServiceImpl) UpdatePassword(ctx context.Context, UUID string, password string) {
	_, err := service.UserRepository.FindById(ctx, UUID)
	util.PanicIfError(err)

	service.UserRepository.UpdatePassword(ctx, UUID, password)
}

func (service *UserServiceImpl) UpdateRole(ctx context.Context, UUID string, role string) {
	_, err := service.UserRepository.FindById(ctx, UUID)
	util.PanicIfError(err)

	service.UserRepository.UpdateRole(ctx, UUID, role)
}

func (service *UserServiceImpl) VerifyEmail(ctx context.Context, UUID string) {
	_, err := service.UserRepository.FindById(ctx, UUID)
	util.PanicIfError(err)

	service.UserRepository.VerifyEmail(ctx, UUID)
}

func (service *UserServiceImpl) Delete(ctx context.Context, UUID string) {
	_, err := service.UserRepository.FindById(ctx, UUID)
	util.PanicIfError(err)
	service.UserRepository.Delete(ctx, UUID)
}

func (service *UserServiceImpl) FindById(ctx context.Context, UUID string) response.UserResponse {
	user, err := service.UserRepository.FindById(ctx, UUID)
	util.PanicIfError(err)

	return response.UserResponse{
		UUID:     user.UUID,
		Username: user.Username,
		Email:    user.Email,
		Bio:      user.Bio,
		Date:     user.Date,
		PFP:      user.PFP,
		Role:     user.Role,
	}
}
