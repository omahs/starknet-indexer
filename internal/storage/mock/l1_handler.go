// Code generated by MockGen. DO NOT EDIT.
// Source: l1_handler.go
//
// Generated by this command:
//
//	mockgen -source=l1_handler.go -destination=mock/l1_handler.go -package=mock -typed
//
// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	storage "github.com/dipdup-io/starknet-indexer/internal/storage"
	storage0 "github.com/dipdup-net/indexer-sdk/pkg/storage"
	gomock "go.uber.org/mock/gomock"
)

// MockIL1Handler is a mock of IL1Handler interface.
type MockIL1Handler struct {
	ctrl     *gomock.Controller
	recorder *MockIL1HandlerMockRecorder
}

// MockIL1HandlerMockRecorder is the mock recorder for MockIL1Handler.
type MockIL1HandlerMockRecorder struct {
	mock *MockIL1Handler
}

// NewMockIL1Handler creates a new mock instance.
func NewMockIL1Handler(ctrl *gomock.Controller) *MockIL1Handler {
	mock := &MockIL1Handler{ctrl: ctrl}
	mock.recorder = &MockIL1HandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIL1Handler) EXPECT() *MockIL1HandlerMockRecorder {
	return m.recorder
}

// CursorList mocks base method.
func (m *MockIL1Handler) CursorList(ctx context.Context, id, limit uint64, order storage0.SortOrder, cmp storage0.Comparator) ([]*storage.L1Handler, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CursorList", ctx, id, limit, order, cmp)
	ret0, _ := ret[0].([]*storage.L1Handler)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CursorList indicates an expected call of CursorList.
func (mr *MockIL1HandlerMockRecorder) CursorList(ctx, id, limit, order, cmp any) *IL1HandlerCursorListCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CursorList", reflect.TypeOf((*MockIL1Handler)(nil).CursorList), ctx, id, limit, order, cmp)
	return &IL1HandlerCursorListCall{Call: call}
}

// IL1HandlerCursorListCall wrap *gomock.Call
type IL1HandlerCursorListCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *IL1HandlerCursorListCall) Return(arg0 []*storage.L1Handler, arg1 error) *IL1HandlerCursorListCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *IL1HandlerCursorListCall) Do(f func(context.Context, uint64, uint64, storage0.SortOrder, storage0.Comparator) ([]*storage.L1Handler, error)) *IL1HandlerCursorListCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *IL1HandlerCursorListCall) DoAndReturn(f func(context.Context, uint64, uint64, storage0.SortOrder, storage0.Comparator) ([]*storage.L1Handler, error)) *IL1HandlerCursorListCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Filter mocks base method.
func (m *MockIL1Handler) Filter(ctx context.Context, flt []storage.L1HandlerFilter, opts ...storage.FilterOption) ([]storage.L1Handler, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, flt}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Filter", varargs...)
	ret0, _ := ret[0].([]storage.L1Handler)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Filter indicates an expected call of Filter.
func (mr *MockIL1HandlerMockRecorder) Filter(ctx, flt any, opts ...any) *IL1HandlerFilterCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, flt}, opts...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Filter", reflect.TypeOf((*MockIL1Handler)(nil).Filter), varargs...)
	return &IL1HandlerFilterCall{Call: call}
}

// IL1HandlerFilterCall wrap *gomock.Call
type IL1HandlerFilterCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *IL1HandlerFilterCall) Return(arg0 []storage.L1Handler, arg1 error) *IL1HandlerFilterCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *IL1HandlerFilterCall) Do(f func(context.Context, []storage.L1HandlerFilter, ...storage.FilterOption) ([]storage.L1Handler, error)) *IL1HandlerFilterCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *IL1HandlerFilterCall) DoAndReturn(f func(context.Context, []storage.L1HandlerFilter, ...storage.FilterOption) ([]storage.L1Handler, error)) *IL1HandlerFilterCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetByID mocks base method.
func (m *MockIL1Handler) GetByID(ctx context.Context, id uint64) (*storage.L1Handler, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, id)
	ret0, _ := ret[0].(*storage.L1Handler)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockIL1HandlerMockRecorder) GetByID(ctx, id any) *IL1HandlerGetByIDCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockIL1Handler)(nil).GetByID), ctx, id)
	return &IL1HandlerGetByIDCall{Call: call}
}

// IL1HandlerGetByIDCall wrap *gomock.Call
type IL1HandlerGetByIDCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *IL1HandlerGetByIDCall) Return(arg0 *storage.L1Handler, arg1 error) *IL1HandlerGetByIDCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *IL1HandlerGetByIDCall) Do(f func(context.Context, uint64) (*storage.L1Handler, error)) *IL1HandlerGetByIDCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *IL1HandlerGetByIDCall) DoAndReturn(f func(context.Context, uint64) (*storage.L1Handler, error)) *IL1HandlerGetByIDCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// IsNoRows mocks base method.
func (m *MockIL1Handler) IsNoRows(err error) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsNoRows", err)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsNoRows indicates an expected call of IsNoRows.
func (mr *MockIL1HandlerMockRecorder) IsNoRows(err any) *IL1HandlerIsNoRowsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsNoRows", reflect.TypeOf((*MockIL1Handler)(nil).IsNoRows), err)
	return &IL1HandlerIsNoRowsCall{Call: call}
}

// IL1HandlerIsNoRowsCall wrap *gomock.Call
type IL1HandlerIsNoRowsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *IL1HandlerIsNoRowsCall) Return(arg0 bool) *IL1HandlerIsNoRowsCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *IL1HandlerIsNoRowsCall) Do(f func(error) bool) *IL1HandlerIsNoRowsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *IL1HandlerIsNoRowsCall) DoAndReturn(f func(error) bool) *IL1HandlerIsNoRowsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// LastID mocks base method.
func (m *MockIL1Handler) LastID(ctx context.Context) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LastID", ctx)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LastID indicates an expected call of LastID.
func (mr *MockIL1HandlerMockRecorder) LastID(ctx any) *IL1HandlerLastIDCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LastID", reflect.TypeOf((*MockIL1Handler)(nil).LastID), ctx)
	return &IL1HandlerLastIDCall{Call: call}
}

// IL1HandlerLastIDCall wrap *gomock.Call
type IL1HandlerLastIDCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *IL1HandlerLastIDCall) Return(arg0 uint64, arg1 error) *IL1HandlerLastIDCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *IL1HandlerLastIDCall) Do(f func(context.Context) (uint64, error)) *IL1HandlerLastIDCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *IL1HandlerLastIDCall) DoAndReturn(f func(context.Context) (uint64, error)) *IL1HandlerLastIDCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// List mocks base method.
func (m *MockIL1Handler) List(ctx context.Context, limit, offset uint64, order storage0.SortOrder) ([]*storage.L1Handler, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, limit, offset, order)
	ret0, _ := ret[0].([]*storage.L1Handler)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockIL1HandlerMockRecorder) List(ctx, limit, offset, order any) *IL1HandlerListCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockIL1Handler)(nil).List), ctx, limit, offset, order)
	return &IL1HandlerListCall{Call: call}
}

// IL1HandlerListCall wrap *gomock.Call
type IL1HandlerListCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *IL1HandlerListCall) Return(arg0 []*storage.L1Handler, arg1 error) *IL1HandlerListCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *IL1HandlerListCall) Do(f func(context.Context, uint64, uint64, storage0.SortOrder) ([]*storage.L1Handler, error)) *IL1HandlerListCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *IL1HandlerListCall) DoAndReturn(f func(context.Context, uint64, uint64, storage0.SortOrder) ([]*storage.L1Handler, error)) *IL1HandlerListCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Save mocks base method.
func (m_2 *MockIL1Handler) Save(ctx context.Context, m *storage.L1Handler) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Save", ctx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockIL1HandlerMockRecorder) Save(ctx, m any) *IL1HandlerSaveCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockIL1Handler)(nil).Save), ctx, m)
	return &IL1HandlerSaveCall{Call: call}
}

// IL1HandlerSaveCall wrap *gomock.Call
type IL1HandlerSaveCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *IL1HandlerSaveCall) Return(arg0 error) *IL1HandlerSaveCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *IL1HandlerSaveCall) Do(f func(context.Context, *storage.L1Handler) error) *IL1HandlerSaveCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *IL1HandlerSaveCall) DoAndReturn(f func(context.Context, *storage.L1Handler) error) *IL1HandlerSaveCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Update mocks base method.
func (m_2 *MockIL1Handler) Update(ctx context.Context, m *storage.L1Handler) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Update", ctx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockIL1HandlerMockRecorder) Update(ctx, m any) *IL1HandlerUpdateCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIL1Handler)(nil).Update), ctx, m)
	return &IL1HandlerUpdateCall{Call: call}
}

// IL1HandlerUpdateCall wrap *gomock.Call
type IL1HandlerUpdateCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *IL1HandlerUpdateCall) Return(arg0 error) *IL1HandlerUpdateCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *IL1HandlerUpdateCall) Do(f func(context.Context, *storage.L1Handler) error) *IL1HandlerUpdateCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *IL1HandlerUpdateCall) DoAndReturn(f func(context.Context, *storage.L1Handler) error) *IL1HandlerUpdateCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
