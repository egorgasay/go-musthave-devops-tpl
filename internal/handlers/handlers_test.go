package handlers

import (
	"bytes"
	"devtool/internal/service"
	service_mocks "devtool/internal/service/mocks"
	"devtool/internal/storage"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestHandler_UpdateMetric(t *testing.T) {
	type mockBehavior func(r *service_mocks.MockIService)
	fval := float64(6)

	tests := []struct {
		name               string
		url                string
		mockBehavior       mockBehavior
		expectedStatusCode int
	}{
		{
			name: "Ok",
			url:  "http://localhost:8080/update/gauge/Alloc/6.0",
			mockBehavior: func(r *service_mocks.MockIService) {
				r.EXPECT().UpdateMetric(&storage.Metrics{ID: "Alloc", MType: "gauge", Value: &fval}).
					Return(fval, nil).AnyTimes()
			},
			expectedStatusCode: 200,
		},
		{
			name: "Ok, but value doesn't exist yet",
			url:  "http://localhost:8080/update/gauge/AllocNew/6.0",
			mockBehavior: func(r *service_mocks.MockIService) {
				r.EXPECT().UpdateMetric(&storage.Metrics{ID: "AllocNew", MType: "gauge", Value: &fval}).
					Return(fval, nil).AnyTimes()
			},
			expectedStatusCode: 200,
		},
		{
			name: "не определен",
			url:  "http://localhost:8080/update/gauge2/AllocNew/6.0",
			mockBehavior: func(r *service_mocks.MockIService) {
				r.EXPECT().UpdateMetric(&storage.Metrics{ID: "AllocNew", MType: "gauge2", Value: &fval}).
					Return(fval, errors.New("тип не определен")).AnyTimes()
			},
			expectedStatusCode: 501,
		},
		{
			name: "404",
			url:  "http://localhost:8080/update/gauge",
			mockBehavior: func(r *service_mocks.MockIService) {
				r.EXPECT().UpdateMetric(&storage.Metrics{}).
					Return(0.0, nil).AnyTimes()
			},
			expectedStatusCode: 404,
		},
		{
			name: "no value",
			url:  "http://localhost:8080/update/gauge/Alloc/",
			mockBehavior: func(r *service_mocks.MockIService) {
				r.EXPECT().UpdateMetric(&storage.Metrics{}).
					Return(0.0, nil).AnyTimes()
			},
			expectedStatusCode: 404,
		},
		{
			name: "wrong value type",
			url:  "http://localhost:8080/update/gauge/Alloc/664q",
			mockBehavior: func(r *service_mocks.MockIService) {
				r.EXPECT().UpdateMetric(&storage.Metrics{}).
					Return(0.0, nil).AnyTimes()
			},
			expectedStatusCode: 400,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			repos := service_mocks.NewMockIService(c)
			test.mockBehavior(repos)

			services := &service.Service{DB: repos}
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
	type mockBehavior func(r *service_mocks.MockIService)

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
			mockBehavior: func(r *service_mocks.MockIService) {
				r.EXPECT().GetMetric("Alloc").
					Return(6.0, nil).AnyTimes()
			},
			expectedStatusCode: 200,
			expectedBody:       "6",
		},
		{
			name: "404 can't find value by name",
			url:  "http://localhost:8080/value/gauge/Alqweloc",
			mockBehavior: func(r *service_mocks.MockIService) {
				r.EXPECT().GetMetric("Alqweloc").
					Return(0.0, errors.New("значение не установлено")).
					AnyTimes()
			},
			expectedStatusCode: 404,
			expectedBody:       "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			repos := service_mocks.NewMockIService(c)
			test.mockBehavior(repos)

			services := &service.Service{DB: repos}
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

//func TestHandler_GetAllMetric(t *testing.T) {
//	type mockBehavior func(r *service_mocks.MockIService)
//	testOk := []*repo.Metrics{{Type: "Test", Name: "Alloc", Value: 555},
//		{Type: "Test", Name: "BuckHashSys", Value: 123}}
//	testOkBody, err := ioutil.ReadFile("index_test.html")
//	if err != nil {
//		fmt.Println(err)
//	}
//	tests := []struct {
//		name               string
//		url                string
//		mockBehavior       mockBehavior
//		expectedStatusCode int
//		expectedBody       string
//	}{
//		{
//			name: "Get all",
//			url:  "http://127.0.0.1:8080/",
//			mockBehavior: func(r *service_mocks.MockIService) {
//				r.EXPECT().GetAllMetrics().
//					Return(testOk, nil).AnyTimes()
//			},
//			expectedStatusCode: 200,
//			expectedBody:       string(testOkBody),
//		},
//	}
//	for _, test := range tests {
//		t.Run(test.name, func(t *testing.T) {
//			c := gomock.NewController(t)
//			defer c.Finish()
//			repo := service_mocks.NewMockIService(c)
//			test.mockBehavior(repo)
//
//			services := &service.Service{DB: repo}
//			handler := Handler{services}
//
//			req := httptest.NewRequest("GET", test.url,
//				bytes.NewBufferString(""))
//			w := httptest.NewRecorder()
//			//определяем хендлер
//			router := gin.Default()
//			router.GET("/", handler.GetAllMetricsHandler)
//			//router.Use(handler.GetAllMetricsHandler)
//
//			router.ServeHTTP(w, req)
//			// Assert
//			assert.Equal(t, test.expectedStatusCode, w.Code)
//			assert.Equal(t, test.expectedBody, w.Body.String())
//		})
//	}
//}

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
