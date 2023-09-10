// Code generated by MockGen. DO NOT EDIT.
// Source: validator.go
//
// Generated by this command:
//
//	mockgen -source=validator.go -destination=mock/validator.go -package=mock -typed
//
// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	storage "github.com/dipdup-io/celestia-indexer/internal/storage"
	storage0 "github.com/dipdup-net/indexer-sdk/pkg/storage"
	gomock "go.uber.org/mock/gomock"
)

// MockIValidator is a mock of IValidator interface.
type MockIValidator struct {
	ctrl     *gomock.Controller
	recorder *MockIValidatorMockRecorder
}

// MockIValidatorMockRecorder is the mock recorder for MockIValidator.
type MockIValidatorMockRecorder struct {
	mock *MockIValidator
}

// NewMockIValidator creates a new mock instance.
func NewMockIValidator(ctrl *gomock.Controller) *MockIValidator {
	mock := &MockIValidator{ctrl: ctrl}
	mock.recorder = &MockIValidatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIValidator) EXPECT() *MockIValidatorMockRecorder {
	return m.recorder
}

// ByAddress mocks base method.
func (m *MockIValidator) ByAddress(ctx context.Context, address string) (storage.Validator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ByAddress", ctx, address)
	ret0, _ := ret[0].(storage.Validator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ByAddress indicates an expected call of ByAddress.
func (mr *MockIValidatorMockRecorder) ByAddress(ctx, address any) *IValidatorByAddressCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ByAddress", reflect.TypeOf((*MockIValidator)(nil).ByAddress), ctx, address)
	return &IValidatorByAddressCall{Call: call}
}

// IValidatorByAddressCall wrap *gomock.Call
type IValidatorByAddressCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *IValidatorByAddressCall) Return(arg0 storage.Validator, arg1 error) *IValidatorByAddressCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *IValidatorByAddressCall) Do(f func(context.Context, string) (storage.Validator, error)) *IValidatorByAddressCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *IValidatorByAddressCall) DoAndReturn(f func(context.Context, string) (storage.Validator, error)) *IValidatorByAddressCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// CursorList mocks base method.
func (m *MockIValidator) CursorList(ctx context.Context, id, limit uint64, order storage0.SortOrder, cmp storage0.Comparator) ([]*storage.Validator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CursorList", ctx, id, limit, order, cmp)
	ret0, _ := ret[0].([]*storage.Validator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CursorList indicates an expected call of CursorList.
func (mr *MockIValidatorMockRecorder) CursorList(ctx, id, limit, order, cmp any) *IValidatorCursorListCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CursorList", reflect.TypeOf((*MockIValidator)(nil).CursorList), ctx, id, limit, order, cmp)
	return &IValidatorCursorListCall{Call: call}
}

// IValidatorCursorListCall wrap *gomock.Call
type IValidatorCursorListCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *IValidatorCursorListCall) Return(arg0 []*storage.Validator, arg1 error) *IValidatorCursorListCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *IValidatorCursorListCall) Do(f func(context.Context, uint64, uint64, storage0.SortOrder, storage0.Comparator) ([]*storage.Validator, error)) *IValidatorCursorListCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *IValidatorCursorListCall) DoAndReturn(f func(context.Context, uint64, uint64, storage0.SortOrder, storage0.Comparator) ([]*storage.Validator, error)) *IValidatorCursorListCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetByID mocks base method.
func (m *MockIValidator) GetByID(ctx context.Context, id uint64) (*storage.Validator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, id)
	ret0, _ := ret[0].(*storage.Validator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockIValidatorMockRecorder) GetByID(ctx, id any) *IValidatorGetByIDCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockIValidator)(nil).GetByID), ctx, id)
	return &IValidatorGetByIDCall{Call: call}
}

// IValidatorGetByIDCall wrap *gomock.Call
type IValidatorGetByIDCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *IValidatorGetByIDCall) Return(arg0 *storage.Validator, arg1 error) *IValidatorGetByIDCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *IValidatorGetByIDCall) Do(f func(context.Context, uint64) (*storage.Validator, error)) *IValidatorGetByIDCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *IValidatorGetByIDCall) DoAndReturn(f func(context.Context, uint64) (*storage.Validator, error)) *IValidatorGetByIDCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// IsNoRows mocks base method.
func (m *MockIValidator) IsNoRows(err error) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsNoRows", err)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsNoRows indicates an expected call of IsNoRows.
func (mr *MockIValidatorMockRecorder) IsNoRows(err any) *IValidatorIsNoRowsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsNoRows", reflect.TypeOf((*MockIValidator)(nil).IsNoRows), err)
	return &IValidatorIsNoRowsCall{Call: call}
}

// IValidatorIsNoRowsCall wrap *gomock.Call
type IValidatorIsNoRowsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *IValidatorIsNoRowsCall) Return(arg0 bool) *IValidatorIsNoRowsCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *IValidatorIsNoRowsCall) Do(f func(error) bool) *IValidatorIsNoRowsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *IValidatorIsNoRowsCall) DoAndReturn(f func(error) bool) *IValidatorIsNoRowsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// LastID mocks base method.
func (m *MockIValidator) LastID(ctx context.Context) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LastID", ctx)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LastID indicates an expected call of LastID.
func (mr *MockIValidatorMockRecorder) LastID(ctx any) *IValidatorLastIDCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LastID", reflect.TypeOf((*MockIValidator)(nil).LastID), ctx)
	return &IValidatorLastIDCall{Call: call}
}

// IValidatorLastIDCall wrap *gomock.Call
type IValidatorLastIDCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *IValidatorLastIDCall) Return(arg0 uint64, arg1 error) *IValidatorLastIDCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *IValidatorLastIDCall) Do(f func(context.Context) (uint64, error)) *IValidatorLastIDCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *IValidatorLastIDCall) DoAndReturn(f func(context.Context) (uint64, error)) *IValidatorLastIDCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// List mocks base method.
func (m *MockIValidator) List(ctx context.Context, limit, offset uint64, order storage0.SortOrder) ([]*storage.Validator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, limit, offset, order)
	ret0, _ := ret[0].([]*storage.Validator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockIValidatorMockRecorder) List(ctx, limit, offset, order any) *IValidatorListCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockIValidator)(nil).List), ctx, limit, offset, order)
	return &IValidatorListCall{Call: call}
}

// IValidatorListCall wrap *gomock.Call
type IValidatorListCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *IValidatorListCall) Return(arg0 []*storage.Validator, arg1 error) *IValidatorListCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *IValidatorListCall) Do(f func(context.Context, uint64, uint64, storage0.SortOrder) ([]*storage.Validator, error)) *IValidatorListCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *IValidatorListCall) DoAndReturn(f func(context.Context, uint64, uint64, storage0.SortOrder) ([]*storage.Validator, error)) *IValidatorListCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Save mocks base method.
func (m_2 *MockIValidator) Save(ctx context.Context, m *storage.Validator) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Save", ctx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockIValidatorMockRecorder) Save(ctx, m any) *IValidatorSaveCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockIValidator)(nil).Save), ctx, m)
	return &IValidatorSaveCall{Call: call}
}

// IValidatorSaveCall wrap *gomock.Call
type IValidatorSaveCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *IValidatorSaveCall) Return(arg0 error) *IValidatorSaveCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *IValidatorSaveCall) Do(f func(context.Context, *storage.Validator) error) *IValidatorSaveCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *IValidatorSaveCall) DoAndReturn(f func(context.Context, *storage.Validator) error) *IValidatorSaveCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Update mocks base method.
func (m_2 *MockIValidator) Update(ctx context.Context, m *storage.Validator) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Update", ctx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockIValidatorMockRecorder) Update(ctx, m any) *IValidatorUpdateCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIValidator)(nil).Update), ctx, m)
	return &IValidatorUpdateCall{Call: call}
}

// IValidatorUpdateCall wrap *gomock.Call
type IValidatorUpdateCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *IValidatorUpdateCall) Return(arg0 error) *IValidatorUpdateCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *IValidatorUpdateCall) Do(f func(context.Context, *storage.Validator) error) *IValidatorUpdateCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *IValidatorUpdateCall) DoAndReturn(f func(context.Context, *storage.Validator) error) *IValidatorUpdateCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}