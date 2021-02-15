package handler

import (
	"bytes"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/Yosh11/exemple_gin/model/todo"
	"github.com/Yosh11/exemple_gin/pkg/service"
	mockService "github.com/Yosh11/exemple_gin/pkg/service/mock"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHandler_singUp(t *testing.T) {
	type mockBehavior func(r *mockService.MockAuthorization, user todo.User)

	tests := []struct {
		name                 string
		inputBody            string
		inputUser            todo.User
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "OK",
			inputBody: `{"username": "username", "name": "Test Name", "password": "qwerty"}`,
			inputUser: todo.User{
				Name:     "Test Name",
				Username: "username",
				Password: "qwerty",
			},
			mockBehavior: func(r *mockService.MockAuthorization, user todo.User) {
				r.EXPECT().CreateUser(user).Return(1, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"id":1}`,
		},
		{
			name:      "Wrong Input",
			inputBody: `{"username": "username"}`,
			mockBehavior: func(r *mockService.MockAuthorization, user todo.User) {
			},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"invalid input body"}`,
		},
		{
			name:      "Server Error",
			inputBody: `{"username": "username", "name": "Test Name", "password": "qwerty"}`,
			inputUser: todo.User{
				Name:     "Test Name",
				Username: "username",
				Password: "qwerty",
			},
			mockBehavior: func(r *mockService.MockAuthorization, user todo.User) {
				r.EXPECT().CreateUser(user).Return(0, errors.New("something went wrong"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"something went wrong"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mockService.NewMockAuthorization(c)
			test.mockBehavior(repo, test.inputUser)

			s := &service.Service{Authorization: repo}
			handler := Handler{s}

			r := gin.New()
			r.POST("/sing-up", handler.singUp)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sing-up", bytes.NewBufferString(test.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}
