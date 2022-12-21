package handlers

import (
	"bytes"
	repo "devtool/internal/repository"
	"devtool/internal/service"
	service_mocks "devtool/internal/service/mocks"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestHandler_UpdateMetric(t *testing.T) {
	type mockBehavior func(r *service_mocks.MockUpdateMetric)

	tests := []struct {
		name               string
		url                string
		mockBehavior       mockBehavior
		expectedStatusCode int
	}{
		{
			name: "Ok",
			url:  "http://localhost:8080/update/gauge/Alloc/6.0",
			mockBehavior: func(r *service_mocks.MockUpdateMetric) {
				r.EXPECT().UpdateMetric(&repo.Metrics{Name: "Alloc", Type: "gauge", Value: 6.0}).
					Return(nil).AnyTimes()
			},
			expectedStatusCode: 200,
		},
		{
			name: "Ok, but value doesn't exist yet",
			url:  "http://localhost:8080/update/gauge/AllocNew/36.0",
			mockBehavior: func(r *service_mocks.MockUpdateMetric) {
				r.EXPECT().UpdateMetric(&repo.Metrics{Name: "AllocNew", Type: "gauge", Value: 36.0}).
					Return(nil).AnyTimes()
			},
			expectedStatusCode: 200,
		},
		{
			name: "не определен",
			url:  "http://localhost:8080/update/gauge2/AllocNew/36.0",
			mockBehavior: func(r *service_mocks.MockUpdateMetric) {
				r.EXPECT().UpdateMetric(&repo.Metrics{Name: "AllocNew", Type: "gauge2", Value: 36.0}).
					Return(errors.New("тип не определен")).AnyTimes()
			},
			expectedStatusCode: 501,
		},
		{
			name: "404",
			url:  "http://localhost:8080/update/gauge",
			mockBehavior: func(r *service_mocks.MockUpdateMetric) {
				r.EXPECT().UpdateMetric(&repo.Metrics{}).
					Return(nil).AnyTimes()
			},
			expectedStatusCode: 404,
		},
		{
			name: "no value",
			url:  "http://localhost:8080/update/gauge/Alloc/",
			mockBehavior: func(r *service_mocks.MockUpdateMetric) {
				r.EXPECT().UpdateMetric(&repo.Metrics{}).
					Return(nil).AnyTimes()
			},
			expectedStatusCode: 404,
		},
		{
			name: "wrong value type",
			url:  "http://localhost:8080/update/gauge/Alloc/664q",
			mockBehavior: func(r *service_mocks.MockUpdateMetric) {
				r.EXPECT().UpdateMetric(&repo.Metrics{}).
					Return(nil).AnyTimes()
			},
			expectedStatusCode: 400,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			repo := service_mocks.NewMockUpdateMetric(c)
			test.mockBehavior(repo)

			services := &service.Service{UpdateMetric: repo}
			handler := Handler{services}

			req := httptest.NewRequest("POST", test.url,
				bytes.NewBufferString(""))
			w := httptest.NewRecorder()
			//определяем хендлер
			router := gin.Default()
			router.POST("/update/:type/:name/:value", handler.UpdateMetricHandler)
			router.POST("/update/:type/:name/", handler.CustomNotFound)

			router.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, test.expectedStatusCode, w.Code)
		})
	}
}

func TestHandler_GetMetric(t *testing.T) {
	type mockBehavior func(r *service_mocks.MockGetMetric)

	tests := []struct {
		name               string
		url                string
		mockBehavior       mockBehavior
		expectedStatusCode int
		expectedBody       string
	}{
		{
			name: "Ok",
			url:  "http://localhost:8080/value/gauge/Alloc",
			mockBehavior: func(r *service_mocks.MockGetMetric) {
				r.EXPECT().GetMetric("Alloc").
					Return(6.0, nil).AnyTimes()
			},
			expectedStatusCode: 200,
			expectedBody:       "6",
		},
		{
			name: "404 can't find value by name",
			url:  "http://localhost:8080/value/gauge/Alqweloc",
			mockBehavior: func(r *service_mocks.MockGetMetric) {
				r.EXPECT().GetMetric("Alqweloc").
					Return(0.0, errors.New("значение не установлено")).AnyTimes()
			},
			expectedStatusCode: 404,
			expectedBody:       "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			repo := service_mocks.NewMockGetMetric(c)
			test.mockBehavior(repo)

			services := &service.Service{GetMetric: repo}
			handler := Handler{services}

			req := httptest.NewRequest("GET", test.url,
				bytes.NewBufferString(""))
			w := httptest.NewRecorder()
			//определяем хендлер
			router := gin.Default()
			router.GET("/value/:type/:name", handler.GetMetricHandler)
			//router.Use(handler.GetMetricHandler)

			router.ServeHTTP(w, req)
			// Assert
			assert.Equal(t, test.expectedStatusCode, w.Code)
			assert.Equal(t, test.expectedBody, w.Body.String())
		})
	}
}

// for tests without mocks
//errRedirectBlocked := errors.New("HTTP redirect blocked")
//redirPolicy := resty.RedirectPolicyFunc(func(_ *http.Request, _ []*http.Request) error {
//	return errRedirectBlocked
//})
//
//httpc := resty.NewWithClient(&http.Client{
//	Transport: &http.Transport{
//		DisableCompression: true,
//	},
//}).SetRedirectPolicy(redirPolicy)
//req := httpc.R()
//resp, _ := req.Post(test.url)
