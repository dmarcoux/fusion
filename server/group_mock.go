// Code generated by MockGen. DO NOT EDIT.
// Source: group.go
//
// Generated by this command:
//
//	mockgen -destination=group_mock.go -source=group.go -package=server
//

// Package server is a generated GoMock package.
package server

import (
	reflect "reflect"

	model "github.com/0x2e/fusion/model"
	gomock "go.uber.org/mock/gomock"
)

// MockGroupRepo is a mock of GroupRepo interface.
type MockGroupRepo struct {
	ctrl     *gomock.Controller
	recorder *MockGroupRepoMockRecorder
}

// MockGroupRepoMockRecorder is the mock recorder for MockGroupRepo.
type MockGroupRepoMockRecorder struct {
	mock *MockGroupRepo
}

// NewMockGroupRepo creates a new mock instance.
func NewMockGroupRepo(ctrl *gomock.Controller) *MockGroupRepo {
	mock := &MockGroupRepo{ctrl: ctrl}
	mock.recorder = &MockGroupRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGroupRepo) EXPECT() *MockGroupRepoMockRecorder {
	return m.recorder
}

// All mocks base method.
func (m *MockGroupRepo) All() ([]*model.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All")
	ret0, _ := ret[0].([]*model.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All.
func (mr *MockGroupRepoMockRecorder) All() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockGroupRepo)(nil).All))
}

// Create mocks base method.
func (m *MockGroupRepo) Create(group *model.Group) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", group)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockGroupRepoMockRecorder) Create(group any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockGroupRepo)(nil).Create), group)
}

// Delete mocks base method.
func (m *MockGroupRepo) Delete(id uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockGroupRepoMockRecorder) Delete(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockGroupRepo)(nil).Delete), id)
}

// Update mocks base method.
func (m *MockGroupRepo) Update(id uint, group *model.Group) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", id, group)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockGroupRepoMockRecorder) Update(id, group any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockGroupRepo)(nil).Update), id, group)
}
