// Code generated by MockGen. DO NOT EDIT.
// Source: shop/Repo (interfaces: ProdRepository)

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"
	Models "shop/Models"

	gomock "github.com/golang/mock/gomock"
)

// MockProdRepository is a mock of ProdRepository interface.
type MockProdRepository struct {
	ctrl     *gomock.Controller
	recorder *MockProdRepositoryMockRecorder
}

// MockProdRepositoryMockRecorder is the mock recorder for MockProdRepository.
type MockProdRepositoryMockRecorder struct {
	mock *MockProdRepository
}

// NewMockProdRepository creates a new mock instance.
func NewMockProdRepository(ctrl *gomock.Controller) *MockProdRepository {
	mock := &MockProdRepository{ctrl: ctrl}
	mock.recorder = &MockProdRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProdRepository) EXPECT() *MockProdRepositoryMockRecorder {
	return m.recorder
}

// SaveProduct mocks base method.
func (m *MockProdRepository) SaveProduct(arg0 *Models.Product) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveProduct", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveProduct indicates an expected call of SaveProduct.
func (mr *MockProdRepositoryMockRecorder) SaveProduct(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveProduct", reflect.TypeOf((*MockProdRepository)(nil).SaveProduct), arg0)
}