package usecase

import (
	"backend/domain"
	"backend/domain/mocks"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestInsertUser(t *testing.T) {
	repo := new(mocks.UserData)
	cost := 10

	mockData := domain.User{Firstname: "Lukman", Lastname: "Hafidz", Username: "NotAPanda",
		Email: "lukman@gmail.com", Password: "polar", Birthdate: "1999-12-05", Photoprofile: "lukman.jpg"}

	emptyMockData := domain.User{ID: 0, Firstname: "", Lastname: "", Username: "",
		Email: "", Password: "", Birthdate: "", Photoprofile: ""}

	returnData := mockData
	returnData.ID = 1
	returnData.Password = "$2a$10$SrMvwwY/QnQ4nZunBvGOuOm2U1w8wcAENOoAMI7l8xH7C1Vmt5mru"

	invalidData := mockData
	invalidData.Email = ""

	noData := mockData
	noData.ID = 0

	t.Run("Success insert data", func(t *testing.T) {
		// useCase := New(&mockUserDataTrue{})
		repo.On("CheckDuplicate", mock.Anything).Return(false).Once()
		repo.On("RegisterData", mock.Anything).Return(returnData).Once()
		useCase := New(repo, validator.New())
		res := useCase.RegisterUser(mockData, cost)

		assert.Equal(t, 200, res)
		repo.AssertExpectations(t)
	})

	t.Run("Validator Error", func(t *testing.T) {
		useCase := New(repo, validator.New())
		res := useCase.RegisterUser(invalidData, cost)

		assert.Equal(t, 400, res)
		repo.AssertExpectations(t)
	})

	t.Run("Generate Hash Error", func(t *testing.T) {
		repo.On("CheckDuplicate", mock.Anything).Return(false).Once()
		useCase := New(repo, validator.New())
		res := useCase.RegisterUser(mockData, 40)

		assert.Equal(t, 500, res)
		repo.AssertExpectations(t)
	})

	t.Run("Data Not Found", func(t *testing.T) {
		repo.On("CheckDuplicate", mock.Anything).Return(false).Once()
		repo.On("RegisterData", mock.Anything).Return(emptyMockData).Once()
		useCase := New(repo, validator.New())
		res := useCase.RegisterUser(noData, cost)

		assert.Equal(t, 404, res)
		repo.AssertExpectations(t)
	})

	t.Run("Duplicate Data", func(t *testing.T) {
		repo.On("CheckDuplicate", mock.Anything).Return(true).Once()
		useCase := New(repo, validator.New())
		res := useCase.RegisterUser(mockData, cost)

		assert.Equal(t, 400, res)
		repo.AssertExpectations(t)
	})
}

func TestUpdateUser(t *testing.T) {
	repo := new(mocks.UserData)
	cost := 10

	mockData := domain.User{Firstname: "Lukman", Lastname: "Hafidz", Username: "NotAPanda",
		Email: "lukman@gmail.com", Password: "polar", Birthdate: "1999-12-05", Photoprofile: "lukman.jpg"}

	returnData := domain.User{ID: 1, Firstname: "Lukman", Lastname: "Hafidz", Username: "NotAPanda",
		Email: "lukman@gmail.com", Password: "polar", Birthdate: "1999-12-05", Photoprofile: "lukman.jpg"}

	invalidData := mockData
	invalidData.Firstname = ""

	t.Run("Success Update", func(t *testing.T) {
		repo.On("CheckDuplicate", mock.Anything).Return(false).Once()
		repo.On("UpdateUserData", mock.Anything).Return(returnData).Once()
		useCase := New(repo, validator.New())
		res := useCase.UpdateUser(mockData, 1, cost)

		assert.Equal(t, 200, res)
		repo.AssertExpectations(t)
	})

	t.Run("Validator Error", func(t *testing.T) {
		useCase := New(repo, validator.New())
		res := useCase.RegisterUser(invalidData, cost)

		assert.Equal(t, 400, res)
		repo.AssertExpectations(t)
	})

	t.Run("Data Not Found", func(t *testing.T) {
		useCase := New(repo, validator.New())
		res := useCase.UpdateUser(mockData, 0, cost)

		assert.Equal(t, 404, res)
		repo.AssertExpectations(t)
	})

	t.Run("Generate Hash Error", func(t *testing.T) {
		repo.On("CheckDuplicate", mock.Anything).Return(false).Once()
		useCase := New(repo, validator.New())
		res := useCase.UpdateUser(mockData, 1, 40)

		assert.Equal(t, 500, res)
		repo.AssertExpectations(t)
	})

	t.Run("Duplicate Data", func(t *testing.T) {
		repo.On("CheckDuplicate", mock.Anything).Return(true).Once()
		useCase := New(repo, validator.New())
		res := useCase.RegisterUser(mockData, cost)

		assert.Equal(t, 400, res)
		repo.AssertExpectations(t)
	})
}

<<<<<<< HEAD
// func TestSearchUser(t *testing.T) {
// 	repo := new(mocks.UserData)
// 	returnData := domain.User{ID: 1, Firstname: "Vanili", Lastname: "Nugroho", Username: "vanili", Email: "vanili@vanili", Password: "d78", Birthdate: "1996-04-25", Photoprofile: "jpg"}
// 	// returnData1 := []domain.UserPosting{{ID: 1, Photo: "jpg", Caption: "a"}}
// 	// returnData2 := []domain.CommentUser{{Id: 1, Firstname: "Vanili", Lastname: "Nugroho", Photoprofile: "apa", Postid: 2, Comment: "apa tuh"}}
// 	t.Run("Success get all Post", func(t *testing.T) {
// 		repo.On("SearchUserData").Return(returnData).Once()
// 		// repo.On("SearchUserPostingData").Return(returnData1).Once()
// 		// repo.On("SearchUserPostingCommentData").Return(returnData2).Once()
// 		usecase := New(repo, validator.New())
// 		res, res2, res3, status := usecase.SearchUser("vanili")
// 		assert.Equal(t, 200, status)
// 		assert.GreaterOrEqual(t, len(res2), 1)
// 		assert.GreaterOrEqual(t, len(res3), 1)
// 		assert.Greater(t, res.ID, 0)
// 		assert.Greater(t, res2[0].ID, 0)
// 		assert.Greater(t, res3[0].Id, 0)
// 		repo.AssertExpectations(t)
// 	})

// 	t.Run("No data found", func(t *testing.T) {
// 		repo.On("SearchUserData").Return(domain.User{}).Once()
// 		// repo.On("SearchUserPostingData").Return([]domain.CommentUser{}).Once()
// 		// repo.On("SearchUserPostingCommentData").Return([]domain.PostComent{}).Once()
// 		usecase := New(repo, validator.New())
// 		res, res2, res3, status := usecase.SearchUser("")
// 		assert.Equal(t, 404, status)
// 		assert.Equal(t, len(res2), 0)
// 		assert.Equal(t, len(res3), 0)
// 		assert.Equal(t, []domain.PostComent([]domain.PostComent(nil)), res)
// 		assert.Equal(t, []domain.PostComent(nil), res)
// 		repo.AssertExpectations(t)
// 	})
// }
=======
func TestSearchUser(t *testing.T) {
	repo := new(mocks.UserData)

	returnDatauser := domain.User{ID: 1, Firstname: "Lukman", Lastname: "Hafidz", Username: "NotAPanda",
		Email: "lukman@gmail.com", Password: "polar", Birthdate: "1999-12-05", Photoprofile: "lukman.jpg"}

	returnDataPostComment := []domain.UserPosting{{ID: 1, Photo: "post.jpg", Caption: "keren bgt"}}

	returnDataCommentUser := []domain.CommentUser{{Id: 1, Firstname: "Lukman", Lastname: "Hafidz", Photoprofile: "lukman.jpg", Postid: 1,
		Comment: "keren bang mamah mu pasti bangga"}}

	t.Run("Success get user", func(t *testing.T) {
		repo.On("SearchUserData", mock.Anything).Return(returnDatauser).Once()
		repo.On("SearchUserPostingData", mock.Anything).Return(returnDataPostComment)
		repo.On("SearchUserPostingCommentData", mock.Anything).Return(returnDataCommentUser)
		useCase := New(repo, validator.New())
		profile, posting, comment, status := useCase.SearchUser("NotAPanda")

		assert.Equal(t, returnDatauser, profile)
		assert.Equal(t, posting, returnDataPostComment)
		assert.Equal(t, comment, returnDataCommentUser)
		assert.Equal(t, 200, status)
		repo.AssertExpectations(t)
	})

	t.Run("wrong input")

	t.Run("Data not found", func(t *testing.T) {
		returnDatauser.ID = 0
		repo.On("SearchUserData", mock.Anything).Return(returnDatauser).Once()
		repo.On("SearchUserPostingData", mock.Anything).Return(returnDataPostComment)
		repo.On("SearchUserPostingCommentData", mock.Anything).Return(returnDataCommentUser)
		useCase := New(repo, validator.New())
		profile, _, _, status := useCase.SearchUser("NotAPanda")

		assert.Equal(t, 0, profile.ID)
		assert.Equal(t, 404, status)
		repo.AssertExpectations(t)
	})
}
>>>>>>> 22265ad (fix update)

func TestDeleteUser(t *testing.T) {
	repo := new(mocks.UserData)

	t.Run("Succes delete", func(t *testing.T) {
		repo.On("DeleteUserData", mock.Anything).Return(true).Once()
		usecase := New(repo, validator.New())
		delete := usecase.DeleteUser(1)

		assert.Equal(t, 200, delete)
		repo.AssertExpectations(t)
	})

	t.Run("No Data Found", func(t *testing.T) {
		repo.On("DeleteUserData", mock.Anything).Return(false).Once()
		usecase := New(repo, validator.New())
		delete := usecase.DeleteUser(100)

		assert.Equal(t, 404, delete)
		repo.AssertExpectations(t)
	})
}
<<<<<<< HEAD
=======

func TestProfileUser(t *testing.T) {
	repo := new(mocks.UserData)

	returnDatauser := domain.User{ID: 1, Firstname: "Lukman", Lastname: "Hafidz", Username: "NotAPanda",
		Email: "lukman@gmail.com", Password: "polar", Birthdate: "1999-12-05", Photoprofile: "lukman.jpg"}

	returnDataPostComment := []domain.UserPosting{{ID: 1, Photo: "post.jpg", Caption: "keren bgt"}}

	returnDataCommentUser := []domain.CommentUser{{Id: 1, Firstname: "Lukman", Lastname: "Hafidz", Photoprofile: "lukman.jpg", Postid: 1,
		Comment: "keren bang mamah mu pasti bangga"}}

	t.Run("Success get user", func(t *testing.T) {
		repo.On("ProfileUserData", mock.Anything).Return(returnDatauser).Once()
		repo.On("GetUserPostingData", mock.Anything).Return(returnDataPostComment)
		repo.On("GetUserCommentData", mock.Anything).Return(returnDataCommentUser)
		useCase := New(repo, validator.New())
		profile, posting, comment, status := useCase.ProfileUser(1)

		assert.Equal(t, returnDatauser, profile)
		assert.Equal(t, posting, returnDataPostComment)
		assert.Equal(t, comment, returnDataCommentUser)
		assert.Equal(t, 200, status)
		repo.AssertExpectations(t)
	})

	t.Run("Data not found", func(t *testing.T) {
		returnDatauser.ID = 0
		repo.On("ProfileUserData", mock.Anything).Return(returnDatauser).Once()
		repo.On("GetUserPostingData", mock.Anything).Return(returnDataPostComment)
		repo.On("GetUserCommentData", mock.Anything).Return(returnDataCommentUser)
		useCase := New(repo, validator.New())
		profile, _, _, status := useCase.ProfileUser(1)

		assert.Equal(t, 0, profile.ID)
		assert.Equal(t, 404, status)
		repo.AssertExpectations(t)
	})
}
>>>>>>> 22265ad (fix update)
