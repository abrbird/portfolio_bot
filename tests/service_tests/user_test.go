package service_tests

import (
	"context"
	"github.com/abrbird/portfolio_bot/internal/domain"
	"github.com/abrbird/portfolio_bot/internal/repository/mock_repository"
	"github.com/abrbird/portfolio_bot/internal/service/service_impl"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserCreate(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	user := domain.User{
		Id:        domain.UserId(1),
		UserName:  "UserName",
		FirstName: "",
		LastName:  "",
	}

	mockRepo := mock_repository.NewUserRepositoryMock(mc)
	mockRepo.CreateMock.Expect(
		context.Background(),
		&user,
	).Return(
		nil,
	)

	userService := service_impl.UserService{}
	err := userService.Create(context.Background(), &user, mockRepo)

	assert.Nil(t, err)
}
func TestUserRetrieve(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	userId := 1
	userName := "UserName"

	user := domain.User{
		Id:        domain.UserId(userId),
		UserName:  userName,
		FirstName: "",
		LastName:  "",
	}

	mockRepo := mock_repository.NewUserRepositoryMock(mc)
	mockRepo.RetrieveMock.Expect(
		context.Background(),
		domain.UserId(userId),
	).Return(
		domain.UserRetrieve{
			User:  &user,
			Error: nil,
		},
	)

	userService := service_impl.UserService{}
	userRetrieve := userService.Retrieve(context.Background(), domain.UserId(userId), mockRepo)

	assert.Nil(t, userRetrieve.Error)
	assert.NotNil(t, userRetrieve.User)
	assert.Equal(t, userRetrieve.User.Id, domain.UserId(userId))
	assert.Equal(t, userRetrieve.User.UserName, userName)
}
func TestUserRetrieveOrCreate(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	userId := 1
	userName := "UserName"

	user := domain.User{
		Id:        domain.UserId(userId),
		UserName:  userName,
		FirstName: "",
		LastName:  "",
	}

	mockRepo := mock_repository.NewUserRepositoryMock(mc)
	mockRepo.RetrieveOrCreateMock.Expect(
		context.Background(),
		&user,
	).Return(
		domain.UserRetrieve{
			User: &domain.User{
				Id:        domain.UserId(userId),
				UserName:  userName,
				FirstName: "",
				LastName:  "",
			},
			Error: nil,
		},
	)

	userService := service_impl.UserService{}
	userRetrieve := userService.RetrieveOrCreate(context.Background(), &user, mockRepo)

	assert.Nil(t, userRetrieve.Error)
	assert.NotNil(t, userRetrieve.User)
	assert.Equal(t, userRetrieve.User.Id, domain.UserId(userId))
	assert.Equal(t, userRetrieve.User.UserName, userName)
}

func TestUserUpdate(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	user := domain.User{
		Id:        domain.UserId(1),
		UserName:  "UserName",
		FirstName: "",
		LastName:  "",
	}

	mockRepo := mock_repository.NewUserRepositoryMock(mc)
	mockRepo.UpdateMock.Expect(
		context.Background(),
		&user,
	).Return(
		nil,
	)

	userService := service_impl.UserService{}
	err := userService.Update(context.Background(), &user, mockRepo)

	assert.Nil(t, err)

}
func TestUserDelete(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	userId := 1

	mockRepo := mock_repository.NewUserRepositoryMock(mc)
	mockRepo.DeleteMock.Expect(
		context.Background(),
		domain.UserId(userId),
	).Return(
		nil,
	)

	userService := service_impl.UserService{}
	err := userService.Delete(context.Background(), domain.UserId(userId), mockRepo)

	assert.Nil(t, err)
}
