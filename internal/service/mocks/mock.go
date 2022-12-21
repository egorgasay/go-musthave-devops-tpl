// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	repository "devtool/internal/repository"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUpdateMetric is a mock of UpdateMetric interface.
type MockUpdateMetric struct {
	ctrl     *gomock.Controller
	recorder *MockUpdateMetricMockRecorder
}

// MockUpdateMetricMockRecorder is the mock recorder for MockUpdateMetric.
type MockUpdateMetricMockRecorder struct {
	mock *MockUpdateMetric
}

// NewMockUpdateMetric creates a new mock instance.
func NewMockUpdateMetric(ctrl *gomock.Controller) *MockUpdateMetric {
	mock := &MockUpdateMetric{ctrl: ctrl}
	mock.recorder = &MockUpdateMetricMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUpdateMetric) EXPECT() *MockUpdateMetricMockRecorder {
	return m.recorder
}

// UpdateMetric mocks base method.
func (m *MockUpdateMetric) UpdateMetric(arg0 *repository.Metrics) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMetric", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMetric indicates an expected call of UpdateMetric.
func (mr *MockUpdateMetricMockRecorder) UpdateMetric(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMetric", reflect.TypeOf((*MockUpdateMetric)(nil).UpdateMetric), arg0)
}

// MockGetMetric is a mock of GetMetric interface.
type MockGetMetric struct {
	ctrl     *gomock.Controller
	recorder *MockGetMetricMockRecorder
}

// MockGetMetricMockRecorder is the mock recorder for MockGetMetric.
type MockGetMetricMockRecorder struct {
	mock *MockGetMetric
}

// NewMockGetMetric creates a new mock instance.
func NewMockGetMetric(ctrl *gomock.Controller) *MockGetMetric {
	mock := &MockGetMetric{ctrl: ctrl}
	mock.recorder = &MockGetMetricMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGetMetric) EXPECT() *MockGetMetricMockRecorder {
	return m.recorder
}

// GetMetric mocks base method.
func (m *MockGetMetric) GetMetric(arg0 string) (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMetric", arg0)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMetric indicates an expected call of GetMetric.
func (mr *MockGetMetricMockRecorder) GetMetric(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMetric", reflect.TypeOf((*MockGetMetric)(nil).GetMetric), arg0)
}

// MockGetAllMetrics is a mock of GetAllMetrics interface.
type MockGetAllMetrics struct {
	ctrl     *gomock.Controller
	recorder *MockGetAllMetricsMockRecorder
}

// MockGetAllMetricsMockRecorder is the mock recorder for MockGetAllMetrics.
type MockGetAllMetricsMockRecorder struct {
	mock *MockGetAllMetrics
}

// NewMockGetAllMetrics creates a new mock instance.
func NewMockGetAllMetrics(ctrl *gomock.Controller) *MockGetAllMetrics {
	mock := &MockGetAllMetrics{ctrl: ctrl}
	mock.recorder = &MockGetAllMetricsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGetAllMetrics) EXPECT() *MockGetAllMetricsMockRecorder {
	return m.recorder
}

// GetAllMetrics mocks base method.
func (m *MockGetAllMetrics) GetAllMetrics(arg0 []repository.Metrics) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllMetrics", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetAllMetrics indicates an expected call of GetAllMetrics.
func (mr *MockGetAllMetricsMockRecorder) GetAllMetrics(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllMetrics", reflect.TypeOf((*MockGetAllMetrics)(nil).GetAllMetrics), arg0)
}
