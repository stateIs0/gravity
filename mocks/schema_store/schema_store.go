// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/moiot/gravity/schema_store (interfaces: SchemaStore)

// Package mock_schema_store is a generated GoMock package.
package mock_schema_store

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	schema_store "github.com/moiot/gravity/schema_store"
)

// MockSchemaStore is a mock of SchemaStore interface
type MockSchemaStore struct {
	ctrl     *gomock.Controller
	recorder *MockSchemaStoreMockRecorder
}

// MockSchemaStoreMockRecorder is the mock recorder for MockSchemaStore
type MockSchemaStoreMockRecorder struct {
	mock *MockSchemaStore
}

// NewMockSchemaStore creates a new mock instance
func NewMockSchemaStore(ctrl *gomock.Controller) *MockSchemaStore {
	mock := &MockSchemaStore{ctrl: ctrl}
	mock.recorder = &MockSchemaStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSchemaStore) EXPECT() *MockSchemaStoreMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockSchemaStore) Close() {
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockSchemaStoreMockRecorder) Close() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockSchemaStore)(nil).Close))
}

// GetSchema mocks base method
func (m *MockSchemaStore) GetSchema(arg0 string) (schema_store.Schema, error) {
	ret := m.ctrl.Call(m, "GetSchema", arg0)
	ret0, _ := ret[0].(schema_store.Schema)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSchema indicates an expected call of GetSchema
func (mr *MockSchemaStoreMockRecorder) GetSchema(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSchema", reflect.TypeOf((*MockSchemaStore)(nil).GetSchema), arg0)
}

// InvalidateCache mocks base method
func (m *MockSchemaStore) InvalidateCache() error {
	ret := m.ctrl.Call(m, "InvalidateCache")
	ret0, _ := ret[0].(error)
	return ret0
}

// InvalidateCache indicates an expected call of InvalidateCache
func (mr *MockSchemaStoreMockRecorder) InvalidateCache() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InvalidateCache", reflect.TypeOf((*MockSchemaStore)(nil).InvalidateCache))
}

// IsInCache mocks base method
func (m *MockSchemaStore) IsInCache(arg0 string) bool {
	ret := m.ctrl.Call(m, "IsInCache", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsInCache indicates an expected call of IsInCache
func (mr *MockSchemaStoreMockRecorder) IsInCache(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsInCache", reflect.TypeOf((*MockSchemaStore)(nil).IsInCache), arg0)
}