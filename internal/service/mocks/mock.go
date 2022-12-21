// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	repository "devtool/internal/repository"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIService is a mock of IService interface.
type MockIService struct {
	ctrl     *gomock.Controller
	recorder *MockIServiceMockRecorder
}

// MockIServiceMockRecorder is the mock recorder for MockIService.
type MockIServiceMockRecorder struct {
	mock *MockIService
}

// NewMockIService creates a new mock instance.
func NewMockIService(ctrl *gomock.Controller) *MockIService {
	mock := &MockIService{ctrl: ctrl}
	mock.recorder = &MockIServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIService) EXPECT() *MockIServiceMockRecorder {
	return m.recorder
}

// GetAllMetrics mocks base method.
func (m *MockIService) GetAllMetrics() ([]*repository.Metrics, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllMetrics")
	ret0, _ := ret[0].([]*repository.Metrics)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllMetrics indicates an expected call of GetAllMetrics.
func (mr *MockIServiceMockRecorder) GetAllMetrics() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllMetrics", reflect.TypeOf((*MockIService)(nil).GetAllMetrics))
}

// GetMetric mocks base method.
func (m *MockIService) GetMetric(arg0 string) (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMetric", arg0)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMetric indicates an expected call of GetMetric.
func (mr *MockIServiceMockRecorder) GetMetric(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMetric", reflect.TypeOf((*MockIService)(nil).GetMetric), arg0)
}

// UpdateMetric mocks base method.
func (m *MockIService) UpdateMetric(arg0 *repository.Metrics) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMetric", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMetric indicates an expected call of UpdateMetric.
func (mr *MockIServiceMockRecorder) UpdateMetric(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMetric", reflect.TypeOf((*MockIService)(nil).UpdateMetric), arg0)
}
