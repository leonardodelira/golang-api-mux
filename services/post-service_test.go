package services

import (
	"goapimux/entities"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (mock *MockRepository) Save(post *entities.Post) (*entities.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*entities.Post), args.Error(1)
}

func (mock *MockRepository) FindAll() (posts []entities.Post, err error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entities.Post), args.Error(1)
}
 
func TestFindAll(t *testing.T) {
	mockRepo := new(MockRepository)
	post := entities.Post{ID: 1, Title: "Title", Text: "Text"}
	mockRepo.On("FindAll").Return([]entities.Post{post}, nil)
	testService := NewPostService(mockRepo)
	result, _ := testService.FindAll()
	mockRepo.AssertExpectations(t)
	assert.Equal(t, 1, result[0].ID)
	assert.Equal(t, "Title", result[0].Title)
	assert.Equal(t, "Text", result[0].Text)
}

func TestCreate(t *testing.T) {
	mockRepo := new(MockRepository)
	post := entities.Post{ID: 1, Title: "Title", Text: "Text"}
	mockRepo.On("Save").Return(&post, nil)
	testService := NewPostService(mockRepo)
	result, _ := testService.Create(&post)
	mockRepo.AssertExpectations(t)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "Title", result.Title)
	assert.Equal(t, "Text", result.Text)
}

func TestValidateEmptyPost(t *testing.T) {
	mockRepo := new(MockRepository)
	testService := NewPostService(mockRepo)
	err := testService.Validate(nil)
	assert.NotNil(t, err)
	assert.Equal(t, "The post is empty", err.Error())
}

func TestValidateEmptyTitle(t *testing.T) {
	mockRepo := new(MockRepository)
	post := entities.Post{}
	testService := NewPostService(mockRepo)
	err := testService.Validate(&post)
	assert.NotNil(t, err)
	assert.Equal(t, "The post Title is empty", err.Error())
}

func TestValidateEmptyText(t *testing.T) {
	mockRepo := new(MockRepository)
	post := entities.Post{
		Title: "Any Title",
	}
	testService := NewPostService(mockRepo)
	err := testService.Validate(&post)
	assert.NotNil(t, err)
	assert.Equal(t, "The post Text is empty", err.Error())
}