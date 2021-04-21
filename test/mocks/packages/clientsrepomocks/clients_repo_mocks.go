// Code generated by MockGen. DO NOT EDIT.
// Source: ../../internal/repositories/clients/repository.go

// Package clientsrepomocks is a generated GoMock package.
package clientsrepomocks

import (
	context "context"
	entities "rebrainme/gotest/internal/entities"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// FindClientByEmail mocks base method.
func (m *MockRepository) FindClientByEmail(ctx context.Context, email string) (*entities.Client, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindClientByEmail", ctx, email)
	ret0, _ := ret[0].(*entities.Client)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindClientByEmail indicates an expected call of FindClientByEmail.
func (mr *MockRepositoryMockRecorder) FindClientByEmail(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindClientByEmail", reflect.TypeOf((*MockRepository)(nil).FindClientByEmail), ctx, email)
}

// InsertOrUpdateClient mocks base method.
func (m *MockRepository) InsertOrUpdateClient(ctx context.Context, client entities.Client) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertOrUpdateClient", ctx, client)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertOrUpdateClient indicates an expected call of InsertOrUpdateClient.
func (mr *MockRepositoryMockRecorder) InsertOrUpdateClient(ctx, client interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertOrUpdateClient", reflect.TypeOf((*MockRepository)(nil).InsertOrUpdateClient), ctx, client)
}