package handlers

import (
	"bytes"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
	"devtool/internal/service"
	service_mocks "url-shortener/internal/service/mocks"
)

func TestHandler_CustomNotFound(t *testing.T) {
	type mockBehavior func(r *service_mocks.MockUpdateMetric)

	tests := []struct {
		name                 string
		inputBody            string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Ok",
			inputBody: "http://localhost:8080/update/gauge/Alloc/123",
			mockBehavior: func(r *service_mocks.MockUpdateMetric) {
				r.EXPECT().UpdateMetric("http://localhost:8080/update/gauge/Alloc/123").
				Return(nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: "",
		},
		{
			name:      "already exists",
			inputBody: "http://zrnzqddy.ru/hlc65i",
			mockBehavior: func(r *service_mocks.MockCreateLink) {
				r.EXPECT().CreateLink("http://zrnzqddy.ru/hlc65i").Return(
					"", gin.Error{Err: errors.New("URL уже существует")})
			},
			expectedStatusCode:   500,
			expectedResponseBody: "",
		},
		{
			name:      "server error",
			inputBody: "q",
			mockBehavior: func(r *service_mocks.MockCreateLink) {
				r.EXPECT().CreateLink("q").Return(
					"", gin.Error{Err: errors.New("недопустимый URL")}).
					AnyTimes()
			},
			expectedStatusCode:   500,
			expectedResponseBody: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := service_mocks.NewMockCreateLink(c)
			test.mockBehavior(repo)

			services := &service.Service{CreateLink: repo}
			handler := Handler{services}

			req := httptest.NewRequest("POST", "/",
				bytes.NewBufferString(test.inputBody))
			w := httptest.NewRecorder()
			// определяем хендлер
			router := gin.Default()
			router.Use(handler.CreateLinkHandler)

			router.ServeHTTP(w, req)
			// Assert
			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}