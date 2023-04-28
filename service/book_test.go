package service

import (
	"challange-10/model"
	"challange-10/repository/mocks"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_BookService_GetBookId(t *testing.T) {
	type testCase struct {
		name           string
		expectedResult model.Book
		expectedError  error
		wantError      bool
		onBookRepo     func(mock *mocks.MockBookRepo)
	}

	var testTable []testCase

	testTable = append(testTable, testCase{
		name:      "success",
		wantError: false,
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().GetBookId(gomock.Any()).Return(model.Book{
				ID:          1,
				Title:       "Buku Golang",
				Author:      "Fadil",
				Description: "Buku tentang golang",
			}, nil).Times(1)
		},
		expectedResult: model.Book{
			ID:          1,
			Title:       "Buku Golang",
			Author:      "Fadil",
			Description: "Buku tentang golang",
		},
	})

	testTable = append(testTable, testCase{
		name:          "record not found",
		wantError:     true,
		expectedError: errors.New("id not found"),
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().GetBookId(gomock.Any()).Return(model.Book{}, errors.New("record not found")).Times(1)
		},
	})

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)

			bookRepo := mocks.NewMockBookRepo(mockCtrl)

			if testCase.onBookRepo != nil {
				testCase.onBookRepo(bookRepo)
			}

			service := Service{
				repo: bookRepo,
			}

			res, err := service.GetBookId(1)

			if testCase.wantError {
				assert.EqualError(t, err, testCase.expectedError.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, testCase.expectedResult, res)
			}
		})
	}
}

func Test_BookService_GetAllBook(t *testing.T) {
	type testCase struct {
		name              string
		expectedResultAll []model.Book
		expectedError     error
		wantError         bool
		onBookRepo        func(mock *mocks.MockBookRepo)
	}

	var testTable []testCase

	testTable = append(testTable, testCase{
		name:      "success",
		wantError: false,
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().GetAllBook().Return([]model.Book{{
				ID:          1,
				Title:       "Buku Golang",
				Author:      "Fadil",
				Description: "Buku tentang golang",
			}, {
				ID:          2,
				Title:       "Buku PHP",
				Author:      "Indra",
				Description: "Buku tentang PHP",
			},
			}, nil).Times(1)
		},
		expectedResultAll: []model.Book{{
			ID:          1,
			Title:       "Buku Golang",
			Author:      "Fadil",
			Description: "Buku tentang golang",
		}, {
			ID:          2,
			Title:       "Buku PHP",
			Author:      "Indra",
			Description: "Buku tentang PHP",
		}},
	})

	testTable = append(testTable, testCase{
		name:          "record not found",
		wantError:     true,
		expectedError: errors.New("record not found"),
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().GetAllBook().Return([]model.Book{}, errors.New("record not found")).Times(1)
		},
	})

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)

			bookRepo := mocks.NewMockBookRepo(mockCtrl)

			if testCase.onBookRepo != nil {
				testCase.onBookRepo(bookRepo)
			}

			service := Service{
				repo: bookRepo,
			}

			res, err := service.GetAllBook()

			if testCase.wantError {
				assert.EqualError(t, err, testCase.expectedError.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, testCase.expectedResultAll, res)
			}
		})
	}

}
