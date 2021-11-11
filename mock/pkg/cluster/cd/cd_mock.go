// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/cluster/cd/cd.go

// Package mock_cd is a generated GoMock package.
package mock_cd

import (
	context "context"
	cd "g.hz.netease.com/horizon/pkg/cluster/cd"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockCD is a mock of CD interface
type MockCD struct {
	ctrl     *gomock.Controller
	recorder *MockCDMockRecorder
}

// MockCDMockRecorder is the mock recorder for MockCD
type MockCDMockRecorder struct {
	mock *MockCD
}

// NewMockCD creates a new mock instance
func NewMockCD(ctrl *gomock.Controller) *MockCD {
	mock := &MockCD{ctrl: ctrl}
	mock.recorder = &MockCDMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCD) EXPECT() *MockCDMockRecorder {
	return m.recorder
}

// CreateCluster mocks base method
func (m *MockCD) CreateCluster(ctx context.Context, params *cd.CreateClusterParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCluster", ctx, params)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateCluster indicates an expected call of CreateCluster
func (mr *MockCDMockRecorder) CreateCluster(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCluster", reflect.TypeOf((*MockCD)(nil).CreateCluster), ctx, params)
}

// DeployCluster mocks base method
func (m *MockCD) DeployCluster(ctx context.Context, params *cd.DeployClusterParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeployCluster", ctx, params)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeployCluster indicates an expected call of DeployCluster
func (mr *MockCDMockRecorder) DeployCluster(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeployCluster", reflect.TypeOf((*MockCD)(nil).DeployCluster), ctx, params)
}

// DeleteCluster mocks base method
func (m *MockCD) DeleteCluster(ctx context.Context, params *cd.DeleteClusterParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCluster", ctx, params)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCluster indicates an expected call of DeleteCluster
func (mr *MockCDMockRecorder) DeleteCluster(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCluster", reflect.TypeOf((*MockCD)(nil).DeleteCluster), ctx, params)
}

// Next mocks base method
func (m *MockCD) Next(ctx context.Context, params *cd.ClusterNextParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Next", ctx, params)
	ret0, _ := ret[0].(error)
	return ret0
}

// Next indicates an expected call of Next
func (mr *MockCDMockRecorder) Next(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockCD)(nil).Next), ctx, params)
}

// GetClusterState mocks base method
func (m *MockCD) GetClusterState(ctx context.Context, params *cd.GetClusterStateParams) (*cd.ClusterState, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClusterState", ctx, params)
	ret0, _ := ret[0].(*cd.ClusterState)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClusterState indicates an expected call of GetClusterState
func (mr *MockCDMockRecorder) GetClusterState(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClusterState", reflect.TypeOf((*MockCD)(nil).GetClusterState), ctx, params)
}

// GetContainerLog mocks base method
func (m *MockCD) GetContainerLog(ctx context.Context, params *cd.GetContainerLogParams) (<-chan string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContainerLog", ctx, params)
	ret0, _ := ret[0].(<-chan string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContainerLog indicates an expected call of GetContainerLog
func (mr *MockCDMockRecorder) GetContainerLog(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContainerLog", reflect.TypeOf((*MockCD)(nil).GetContainerLog), ctx, params)
}

// Online mocks base method
func (m *MockCD) Online(ctx context.Context, params *cd.ExecParams) (map[string]cd.ExecResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Online", ctx, params)
	ret0, _ := ret[0].(map[string]cd.ExecResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Online indicates an expected call of Online
func (mr *MockCDMockRecorder) Online(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Online", reflect.TypeOf((*MockCD)(nil).Online), ctx, params)
}

// Offline mocks base method
func (m *MockCD) Offline(ctx context.Context, params *cd.ExecParams) (map[string]cd.ExecResp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Offline", ctx, params)
	ret0, _ := ret[0].(map[string]cd.ExecResp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Offline indicates an expected call of Offline
func (mr *MockCDMockRecorder) Offline(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Offline", reflect.TypeOf((*MockCD)(nil).Offline), ctx, params)
}
