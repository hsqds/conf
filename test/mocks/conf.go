// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/hsqds/conf (interfaces: Source,SourcesStorage,Config,ConfigsStorage,Loader)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	conf "github.com/hsqds/conf"
)

// MockSource is a mock of Source interface.
type MockSource struct {
	ctrl     *gomock.Controller
	recorder *MockSourceMockRecorder
}

// MockSourceMockRecorder is the mock recorder for MockSource.
type MockSourceMockRecorder struct {
	mock *MockSource
}

// NewMockSource creates a new mock instance.
func NewMockSource(ctrl *gomock.Controller) *MockSource {
	mock := &MockSource{ctrl: ctrl}
	mock.recorder = &MockSourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSource) EXPECT() *MockSourceMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockSource) Close(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockSourceMockRecorder) Close(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockSource)(nil).Close), arg0)
}

// ID mocks base method.
func (m *MockSource) ID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ID indicates an expected call of ID.
func (mr *MockSourceMockRecorder) ID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockSource)(nil).ID))
}

// Load mocks base method.
func (m *MockSource) Load(arg0 context.Context, arg1 []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Load", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Load indicates an expected call of Load.
func (mr *MockSourceMockRecorder) Load(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Load", reflect.TypeOf((*MockSource)(nil).Load), arg0, arg1)
}

// Priority mocks base method.
func (m *MockSource) Priority() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Priority")
	ret0, _ := ret[0].(int)
	return ret0
}

// Priority indicates an expected call of Priority.
func (mr *MockSourceMockRecorder) Priority() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Priority", reflect.TypeOf((*MockSource)(nil).Priority))
}

// ServiceConfig mocks base method.
func (m *MockSource) ServiceConfig(arg0 string) (conf.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServiceConfig", arg0)
	ret0, _ := ret[0].(conf.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ServiceConfig indicates an expected call of ServiceConfig.
func (mr *MockSourceMockRecorder) ServiceConfig(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServiceConfig", reflect.TypeOf((*MockSource)(nil).ServiceConfig), arg0)
}

// MockSourcesStorage is a mock of SourcesStorage interface.
type MockSourcesStorage struct {
	ctrl     *gomock.Controller
	recorder *MockSourcesStorageMockRecorder
}

// MockSourcesStorageMockRecorder is the mock recorder for MockSourcesStorage.
type MockSourcesStorageMockRecorder struct {
	mock *MockSourcesStorage
}

// NewMockSourcesStorage creates a new mock instance.
func NewMockSourcesStorage(ctrl *gomock.Controller) *MockSourcesStorage {
	mock := &MockSourcesStorage{ctrl: ctrl}
	mock.recorder = &MockSourcesStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSourcesStorage) EXPECT() *MockSourcesStorageMockRecorder {
	return m.recorder
}

// Append mocks base method.
func (m *MockSourcesStorage) Append(arg0 conf.Source) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Append", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Append indicates an expected call of Append.
func (mr *MockSourcesStorageMockRecorder) Append(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Append", reflect.TypeOf((*MockSourcesStorage)(nil).Append), arg0)
}

// Get mocks base method.
func (m *MockSourcesStorage) Get(arg0 string) (conf.Source, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(conf.Source)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockSourcesStorageMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockSourcesStorage)(nil).Get), arg0)
}

// List mocks base method.
func (m *MockSourcesStorage) List() []conf.Source {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]conf.Source)
	return ret0
}

// List indicates an expected call of List.
func (mr *MockSourcesStorageMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockSourcesStorage)(nil).List))
}

// MockConfig is a mock of Config interface.
type MockConfig struct {
	ctrl     *gomock.Controller
	recorder *MockConfigMockRecorder
}

// MockConfigMockRecorder is the mock recorder for MockConfig.
type MockConfigMockRecorder struct {
	mock *MockConfig
}

// NewMockConfig creates a new mock instance.
func NewMockConfig(ctrl *gomock.Controller) *MockConfig {
	mock := &MockConfig{ctrl: ctrl}
	mock.recorder = &MockConfigMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConfig) EXPECT() *MockConfigMockRecorder {
	return m.recorder
}

// Fmt mocks base method.
func (m *MockConfig) Fmt(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fmt", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Fmt indicates an expected call of Fmt.
func (mr *MockConfigMockRecorder) Fmt(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fmt", reflect.TypeOf((*MockConfig)(nil).Fmt), arg0)
}

// Get mocks base method.
func (m *MockConfig) Get(arg0, arg1 string) (string, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockConfigMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockConfig)(nil).Get), arg0, arg1)
}

// MockConfigStorage is a mock of ConfigsStorage interface.
type MockConfigStorage struct {
	ctrl     *gomock.Controller
	recorder *MockConfigStorageMockRecorder
}

// MockConfigStorageMockRecorder is the mock recorder for MockConfigStorage.
type MockConfigStorageMockRecorder struct {
	mock *MockConfigStorage
}

// NewMockConfigStorage creates a new mock instance.
func NewMockConfigStorage(ctrl *gomock.Controller) *MockConfigStorage {
	mock := &MockConfigStorage{ctrl: ctrl}
	mock.recorder = &MockConfigStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConfigStorage) EXPECT() *MockConfigStorageMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockConfigStorage) Get(arg0 string) (conf.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(conf.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockConfigStorageMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockConfigStorage)(nil).Get), arg0)
}

// Set mocks base method.
func (m *MockConfigStorage) Set(arg0 string, arg1 conf.Config) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockConfigStorageMockRecorder) Set(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockConfigStorage)(nil).Set), arg0, arg1)
}

// MockLoader is a mock of Loader interface.
type MockLoader struct {
	ctrl     *gomock.Controller
	recorder *MockLoaderMockRecorder
}

// MockLoaderMockRecorder is the mock recorder for MockLoader.
type MockLoaderMockRecorder struct {
	mock *MockLoader
}

// NewMockLoader creates a new mock instance.
func NewMockLoader(ctrl *gomock.Controller) *MockLoader {
	mock := &MockLoader{ctrl: ctrl}
	mock.recorder = &MockLoaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLoader) EXPECT() *MockLoaderMockRecorder {
	return m.recorder
}

// Load mocks base method.
func (m *MockLoader) Load(arg0 context.Context, arg1 []conf.Source, arg2 []string) []conf.LoadResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Load", arg0, arg1, arg2)
	ret0, _ := ret[0].([]conf.LoadResult)
	return ret0
}

// Load indicates an expected call of Load.
func (mr *MockLoaderMockRecorder) Load(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Load", reflect.TypeOf((*MockLoader)(nil).Load), arg0, arg1, arg2)
}
