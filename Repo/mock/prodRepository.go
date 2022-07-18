// Code generated by MockGen. DO NOT EDIT.
// Source: Service/productService.go

// Package mock_Service is a generated GoMock package.
package mock_Service

import (
	reflect "reflect"
	Models "shop/Models"
	dto "shop/dto"

	gomock "github.com/golang/mock/gomock"
)

// MockProductService is a mock of ProductService interface.
type MockProductService struct {
	ctrl     *gomock.Controller
	recorder *MockProductServiceMockRecorder
}

// MockProductServiceMockRecorder is the mock recorder for MockProductService.
type MockProductServiceMockRecorder struct {
	mock *MockProductService
}

// NewMockProductService creates a new mock instance.
func NewMockProductService(ctrl *gomock.Controller) *MockProductService {
	mock := &MockProductService{ctrl: ctrl}
	mock.recorder = &MockProductServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductService) EXPECT() *MockProductServiceMockRecorder {
	return m.recorder
}

// CreateProduct mocks base method.
func (m *MockProductService) CreateProduct(data *dto.Product) (*Models.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProduct", data)
	ret0, _ := ret[0].(*Models.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProduct indicates an expected call of CreateProduct.
func (mr *MockProductServiceMockRecorder) CreateProduct(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProduct", reflect.TypeOf((*MockProductService)(nil).CreateProduct), data)
}

// FindProduct mocks base method.
func (m *MockProductService) FindProduct(id int) (*Models.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindProduct", id)
	ret0, _ := ret[0].(*Models.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindProduct indicates an expected call of FindProduct.
func (mr *MockProductServiceMockRecorder) FindProduct(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindProduct", reflect.TypeOf((*MockProductService)(nil).FindProduct), id)
}

// GetAllProducts mocks base method.
func (m *MockProductService) GetAllProducts() (*[]Models.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllProducts")
	ret0, _ := ret[0].(*[]Models.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllProducts indicates an expected call of GetAllProducts.
func (mr *MockProductServiceMockRecorder) GetAllProducts() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllProducts", reflect.TypeOf((*MockProductService)(nil).GetAllProducts))
}

// UpdateProduct mocks base method.
func (m *MockProductService) UpdateProduct(newProduct *Models.Product) (*Models.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProduct", newProduct)
	ret0, _ := ret[0].(*Models.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateProduct indicates an expected call of UpdateProduct.
func (mr *MockProductServiceMockRecorder) UpdateProduct(newProduct interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProduct", reflect.TypeOf((*MockProductService)(nil).UpdateProduct), newProduct)
}