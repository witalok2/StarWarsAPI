// Code generated by MockGen. DO NOT EDIT.
// Source: application/planet.go

// Package mock_application is a generated GoMock package.
package mock_application

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
	application "github.com/witalok2/starwarsplanet/application"
)

// MockPlanetInterface is a mock of PlanetInterface interface.
type MockPlanetInterface struct {
	ctrl     *gomock.Controller
	recorder *MockPlanetInterfaceMockRecorder
}

// MockPlanetInterfaceMockRecorder is the mock recorder for MockPlanetInterface.
type MockPlanetInterfaceMockRecorder struct {
	mock *MockPlanetInterface
}

// NewMockPlanetInterface creates a new mock instance.
func NewMockPlanetInterface(ctrl *gomock.Controller) *MockPlanetInterface {
	mock := &MockPlanetInterface{ctrl: ctrl}
	mock.recorder = &MockPlanetInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPlanetInterface) EXPECT() *MockPlanetInterfaceMockRecorder {
	return m.recorder
}

// GetClimate mocks base method.
func (m *MockPlanetInterface) GetClimate() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClimate")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetClimate indicates an expected call of GetClimate.
func (mr *MockPlanetInterfaceMockRecorder) GetClimate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClimate", reflect.TypeOf((*MockPlanetInterface)(nil).GetClimate))
}

// GetID mocks base method.
func (m *MockPlanetInterface) GetID() uuid.UUID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetID")
	ret0, _ := ret[0].(uuid.UUID)
	return ret0
}

// GetID indicates an expected call of GetID.
func (mr *MockPlanetInterfaceMockRecorder) GetID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetID", reflect.TypeOf((*MockPlanetInterface)(nil).GetID))
}

// GetMovieAppearances mocks base method.
func (m *MockPlanetInterface) GetMovieAppearances() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMovieAppearances")
	ret0, _ := ret[0].(int)
	return ret0
}

// GetMovieAppearances indicates an expected call of GetMovieAppearances.
func (mr *MockPlanetInterfaceMockRecorder) GetMovieAppearances() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMovieAppearances", reflect.TypeOf((*MockPlanetInterface)(nil).GetMovieAppearances))
}

// GetName mocks base method.
func (m *MockPlanetInterface) GetName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetName indicates an expected call of GetName.
func (mr *MockPlanetInterfaceMockRecorder) GetName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetName", reflect.TypeOf((*MockPlanetInterface)(nil).GetName))
}

// GetTerrain mocks base method.
func (m *MockPlanetInterface) GetTerrain() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTerrain")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetTerrain indicates an expected call of GetTerrain.
func (mr *MockPlanetInterfaceMockRecorder) GetTerrain() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTerrain", reflect.TypeOf((*MockPlanetInterface)(nil).GetTerrain))
}

// IsValid mocks base method.
func (m *MockPlanetInterface) IsValid() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsValid")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsValid indicates an expected call of IsValid.
func (mr *MockPlanetInterfaceMockRecorder) IsValid() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsValid", reflect.TypeOf((*MockPlanetInterface)(nil).IsValid))
}

// MockPlanetServiceInterface is a mock of PlanetServiceInterface interface.
type MockPlanetServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockPlanetServiceInterfaceMockRecorder
}

// MockPlanetServiceInterfaceMockRecorder is the mock recorder for MockPlanetServiceInterface.
type MockPlanetServiceInterfaceMockRecorder struct {
	mock *MockPlanetServiceInterface
}

// NewMockPlanetServiceInterface creates a new mock instance.
func NewMockPlanetServiceInterface(ctrl *gomock.Controller) *MockPlanetServiceInterface {
	mock := &MockPlanetServiceInterface{ctrl: ctrl}
	mock.recorder = &MockPlanetServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPlanetServiceInterface) EXPECT() *MockPlanetServiceInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockPlanetServiceInterface) Create(name, climate, terrain string, appearances int) (application.PlanetInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", name, climate, terrain, appearances)
	ret0, _ := ret[0].(application.PlanetInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockPlanetServiceInterfaceMockRecorder) Create(name, climate, terrain, appearances interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockPlanetServiceInterface)(nil).Create), name, climate, terrain, appearances)
}

// Delete mocks base method.
func (m *MockPlanetServiceInterface) Delete(id uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockPlanetServiceInterfaceMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockPlanetServiceInterface)(nil).Delete), id)
}

// Get mocks base method.
func (m *MockPlanetServiceInterface) Get(id uuid.UUID) (application.PlanetInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(application.PlanetInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockPlanetServiceInterfaceMockRecorder) Get(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockPlanetServiceInterface)(nil).Get), id)
}

// GetName mocks base method.
func (m *MockPlanetServiceInterface) GetName(name string) (application.PlanetInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetName", name)
	ret0, _ := ret[0].(application.PlanetInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetName indicates an expected call of GetName.
func (mr *MockPlanetServiceInterfaceMockRecorder) GetName(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetName", reflect.TypeOf((*MockPlanetServiceInterface)(nil).GetName), name)
}

// List mocks base method.
func (m *MockPlanetServiceInterface) List() ([]application.PlanetInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]application.PlanetInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockPlanetServiceInterfaceMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockPlanetServiceInterface)(nil).List))
}

// MockPlanetPersistenceInterface is a mock of PlanetPersistenceInterface interface.
type MockPlanetPersistenceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockPlanetPersistenceInterfaceMockRecorder
}

// MockPlanetPersistenceInterfaceMockRecorder is the mock recorder for MockPlanetPersistenceInterface.
type MockPlanetPersistenceInterfaceMockRecorder struct {
	mock *MockPlanetPersistenceInterface
}

// NewMockPlanetPersistenceInterface creates a new mock instance.
func NewMockPlanetPersistenceInterface(ctrl *gomock.Controller) *MockPlanetPersistenceInterface {
	mock := &MockPlanetPersistenceInterface{ctrl: ctrl}
	mock.recorder = &MockPlanetPersistenceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPlanetPersistenceInterface) EXPECT() *MockPlanetPersistenceInterfaceMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockPlanetPersistenceInterface) Delete(id uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockPlanetPersistenceInterfaceMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockPlanetPersistenceInterface)(nil).Delete), id)
}

// Get mocks base method.
func (m *MockPlanetPersistenceInterface) Get(id uuid.UUID) (application.PlanetInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(application.PlanetInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockPlanetPersistenceInterfaceMockRecorder) Get(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockPlanetPersistenceInterface)(nil).Get), id)
}

// GetName mocks base method.
func (m *MockPlanetPersistenceInterface) GetName(name string) (application.PlanetInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetName", name)
	ret0, _ := ret[0].(application.PlanetInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetName indicates an expected call of GetName.
func (mr *MockPlanetPersistenceInterfaceMockRecorder) GetName(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetName", reflect.TypeOf((*MockPlanetPersistenceInterface)(nil).GetName), name)
}

// List mocks base method.
func (m *MockPlanetPersistenceInterface) List() ([]application.PlanetInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]application.PlanetInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockPlanetPersistenceInterfaceMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockPlanetPersistenceInterface)(nil).List))
}

// Save mocks base method.
func (m *MockPlanetPersistenceInterface) Save(planet application.PlanetInterface) (application.PlanetInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", planet)
	ret0, _ := ret[0].(application.PlanetInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Save indicates an expected call of Save.
func (mr *MockPlanetPersistenceInterfaceMockRecorder) Save(planet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockPlanetPersistenceInterface)(nil).Save), planet)
}

// MockPlanetReader is a mock of PlanetReader interface.
type MockPlanetReader struct {
	ctrl     *gomock.Controller
	recorder *MockPlanetReaderMockRecorder
}

// MockPlanetReaderMockRecorder is the mock recorder for MockPlanetReader.
type MockPlanetReaderMockRecorder struct {
	mock *MockPlanetReader
}

// NewMockPlanetReader creates a new mock instance.
func NewMockPlanetReader(ctrl *gomock.Controller) *MockPlanetReader {
	mock := &MockPlanetReader{ctrl: ctrl}
	mock.recorder = &MockPlanetReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPlanetReader) EXPECT() *MockPlanetReaderMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockPlanetReader) Get(id uuid.UUID) (application.PlanetInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(application.PlanetInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockPlanetReaderMockRecorder) Get(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockPlanetReader)(nil).Get), id)
}

// GetName mocks base method.
func (m *MockPlanetReader) GetName(name string) (application.PlanetInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetName", name)
	ret0, _ := ret[0].(application.PlanetInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetName indicates an expected call of GetName.
func (mr *MockPlanetReaderMockRecorder) GetName(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetName", reflect.TypeOf((*MockPlanetReader)(nil).GetName), name)
}

// List mocks base method.
func (m *MockPlanetReader) List() ([]application.PlanetInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]application.PlanetInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockPlanetReaderMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockPlanetReader)(nil).List))
}

// MockPlanetWriter is a mock of PlanetWriter interface.
type MockPlanetWriter struct {
	ctrl     *gomock.Controller
	recorder *MockPlanetWriterMockRecorder
}

// MockPlanetWriterMockRecorder is the mock recorder for MockPlanetWriter.
type MockPlanetWriterMockRecorder struct {
	mock *MockPlanetWriter
}

// NewMockPlanetWriter creates a new mock instance.
func NewMockPlanetWriter(ctrl *gomock.Controller) *MockPlanetWriter {
	mock := &MockPlanetWriter{ctrl: ctrl}
	mock.recorder = &MockPlanetWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPlanetWriter) EXPECT() *MockPlanetWriterMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockPlanetWriter) Delete(id uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockPlanetWriterMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockPlanetWriter)(nil).Delete), id)
}

// Save mocks base method.
func (m *MockPlanetWriter) Save(planet application.PlanetInterface) (application.PlanetInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", planet)
	ret0, _ := ret[0].(application.PlanetInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Save indicates an expected call of Save.
func (mr *MockPlanetWriterMockRecorder) Save(planet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockPlanetWriter)(nil).Save), planet)
}
