// Code generated by MockGen. DO NOT EDIT.
// Source: token_balance.go
//
// Generated by this command:
//
//	mockgen -source=token_balance.go -destination=mock/token_balance.go -package=mock -typed
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

// MockITokenBalance is a mock of ITokenBalance interface.
type MockITokenBalance struct {
	ctrl     *gomock.Controller
	recorder *MockITokenBalanceMockRecorder
}

// MockITokenBalanceMockRecorder is the mock recorder for MockITokenBalance.
type MockITokenBalanceMockRecorder struct {
	mock *MockITokenBalance
}

// NewMockITokenBalance creates a new mock instance.
func NewMockITokenBalance(ctrl *gomock.Controller) *MockITokenBalance {
	mock := &MockITokenBalance{ctrl: ctrl}
	mock.recorder = &MockITokenBalanceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockITokenBalance) EXPECT() *MockITokenBalanceMockRecorder {
	return m.recorder
}

// Balances mocks base method.
func (m *MockITokenBalance) Balances(ctx context.Context, contractId uint64, tokenId int64, limit, offset int) ([]storage.TokenBalance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Balances", ctx, contractId, tokenId, limit, offset)
	ret0, _ := ret[0].([]storage.TokenBalance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Balances indicates an expected call of Balances.
func (mr *MockITokenBalanceMockRecorder) Balances(ctx, contractId, tokenId, limit, offset any) *ITokenBalanceBalancesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Balances", reflect.TypeOf((*MockITokenBalance)(nil).Balances), ctx, contractId, tokenId, limit, offset)
	return &ITokenBalanceBalancesCall{Call: call}
}

// ITokenBalanceBalancesCall wrap *gomock.Call
type ITokenBalanceBalancesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *ITokenBalanceBalancesCall) Return(arg0 []storage.TokenBalance, arg1 error) *ITokenBalanceBalancesCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *ITokenBalanceBalancesCall) Do(f func(context.Context, uint64, int64, int, int) ([]storage.TokenBalance, error)) *ITokenBalanceBalancesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *ITokenBalanceBalancesCall) DoAndReturn(f func(context.Context, uint64, int64, int, int) ([]storage.TokenBalance, error)) *ITokenBalanceBalancesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// CursorList mocks base method.
func (m *MockITokenBalance) CursorList(ctx context.Context, id, limit uint64, order storage0.SortOrder, cmp storage0.Comparator) ([]*storage.TokenBalance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CursorList", ctx, id, limit, order, cmp)
	ret0, _ := ret[0].([]*storage.TokenBalance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CursorList indicates an expected call of CursorList.
func (mr *MockITokenBalanceMockRecorder) CursorList(ctx, id, limit, order, cmp any) *ITokenBalanceCursorListCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CursorList", reflect.TypeOf((*MockITokenBalance)(nil).CursorList), ctx, id, limit, order, cmp)
	return &ITokenBalanceCursorListCall{Call: call}
}

// ITokenBalanceCursorListCall wrap *gomock.Call
type ITokenBalanceCursorListCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *ITokenBalanceCursorListCall) Return(arg0 []*storage.TokenBalance, arg1 error) *ITokenBalanceCursorListCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *ITokenBalanceCursorListCall) Do(f func(context.Context, uint64, uint64, storage0.SortOrder, storage0.Comparator) ([]*storage.TokenBalance, error)) *ITokenBalanceCursorListCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *ITokenBalanceCursorListCall) DoAndReturn(f func(context.Context, uint64, uint64, storage0.SortOrder, storage0.Comparator) ([]*storage.TokenBalance, error)) *ITokenBalanceCursorListCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Filter mocks base method.
func (m *MockITokenBalance) Filter(ctx context.Context, flt []storage.TokenBalanceFilter, opts ...storage.FilterOption) ([]storage.TokenBalance, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, flt}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Filter", varargs...)
	ret0, _ := ret[0].([]storage.TokenBalance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Filter indicates an expected call of Filter.
func (mr *MockITokenBalanceMockRecorder) Filter(ctx, flt any, opts ...any) *ITokenBalanceFilterCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, flt}, opts...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Filter", reflect.TypeOf((*MockITokenBalance)(nil).Filter), varargs...)
	return &ITokenBalanceFilterCall{Call: call}
}

// ITokenBalanceFilterCall wrap *gomock.Call
type ITokenBalanceFilterCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *ITokenBalanceFilterCall) Return(arg0 []storage.TokenBalance, arg1 error) *ITokenBalanceFilterCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *ITokenBalanceFilterCall) Do(f func(context.Context, []storage.TokenBalanceFilter, ...storage.FilterOption) ([]storage.TokenBalance, error)) *ITokenBalanceFilterCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *ITokenBalanceFilterCall) DoAndReturn(f func(context.Context, []storage.TokenBalanceFilter, ...storage.FilterOption) ([]storage.TokenBalance, error)) *ITokenBalanceFilterCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetByID mocks base method.
func (m *MockITokenBalance) GetByID(ctx context.Context, id uint64) (*storage.TokenBalance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, id)
	ret0, _ := ret[0].(*storage.TokenBalance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockITokenBalanceMockRecorder) GetByID(ctx, id any) *ITokenBalanceGetByIDCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockITokenBalance)(nil).GetByID), ctx, id)
	return &ITokenBalanceGetByIDCall{Call: call}
}

// ITokenBalanceGetByIDCall wrap *gomock.Call
type ITokenBalanceGetByIDCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *ITokenBalanceGetByIDCall) Return(arg0 *storage.TokenBalance, arg1 error) *ITokenBalanceGetByIDCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *ITokenBalanceGetByIDCall) Do(f func(context.Context, uint64) (*storage.TokenBalance, error)) *ITokenBalanceGetByIDCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *ITokenBalanceGetByIDCall) DoAndReturn(f func(context.Context, uint64) (*storage.TokenBalance, error)) *ITokenBalanceGetByIDCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// IsNoRows mocks base method.
func (m *MockITokenBalance) IsNoRows(err error) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsNoRows", err)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsNoRows indicates an expected call of IsNoRows.
func (mr *MockITokenBalanceMockRecorder) IsNoRows(err any) *ITokenBalanceIsNoRowsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsNoRows", reflect.TypeOf((*MockITokenBalance)(nil).IsNoRows), err)
	return &ITokenBalanceIsNoRowsCall{Call: call}
}

// ITokenBalanceIsNoRowsCall wrap *gomock.Call
type ITokenBalanceIsNoRowsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *ITokenBalanceIsNoRowsCall) Return(arg0 bool) *ITokenBalanceIsNoRowsCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *ITokenBalanceIsNoRowsCall) Do(f func(error) bool) *ITokenBalanceIsNoRowsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *ITokenBalanceIsNoRowsCall) DoAndReturn(f func(error) bool) *ITokenBalanceIsNoRowsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// LastID mocks base method.
func (m *MockITokenBalance) LastID(ctx context.Context) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LastID", ctx)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LastID indicates an expected call of LastID.
func (mr *MockITokenBalanceMockRecorder) LastID(ctx any) *ITokenBalanceLastIDCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LastID", reflect.TypeOf((*MockITokenBalance)(nil).LastID), ctx)
	return &ITokenBalanceLastIDCall{Call: call}
}

// ITokenBalanceLastIDCall wrap *gomock.Call
type ITokenBalanceLastIDCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *ITokenBalanceLastIDCall) Return(arg0 uint64, arg1 error) *ITokenBalanceLastIDCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *ITokenBalanceLastIDCall) Do(f func(context.Context) (uint64, error)) *ITokenBalanceLastIDCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *ITokenBalanceLastIDCall) DoAndReturn(f func(context.Context) (uint64, error)) *ITokenBalanceLastIDCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// List mocks base method.
func (m *MockITokenBalance) List(ctx context.Context, limit, offset uint64, order storage0.SortOrder) ([]*storage.TokenBalance, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, limit, offset, order)
	ret0, _ := ret[0].([]*storage.TokenBalance)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockITokenBalanceMockRecorder) List(ctx, limit, offset, order any) *ITokenBalanceListCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockITokenBalance)(nil).List), ctx, limit, offset, order)
	return &ITokenBalanceListCall{Call: call}
}

// ITokenBalanceListCall wrap *gomock.Call
type ITokenBalanceListCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *ITokenBalanceListCall) Return(arg0 []*storage.TokenBalance, arg1 error) *ITokenBalanceListCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *ITokenBalanceListCall) Do(f func(context.Context, uint64, uint64, storage0.SortOrder) ([]*storage.TokenBalance, error)) *ITokenBalanceListCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *ITokenBalanceListCall) DoAndReturn(f func(context.Context, uint64, uint64, storage0.SortOrder) ([]*storage.TokenBalance, error)) *ITokenBalanceListCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Save mocks base method.
func (m_2 *MockITokenBalance) Save(ctx context.Context, m *storage.TokenBalance) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Save", ctx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockITokenBalanceMockRecorder) Save(ctx, m any) *ITokenBalanceSaveCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockITokenBalance)(nil).Save), ctx, m)
	return &ITokenBalanceSaveCall{Call: call}
}

// ITokenBalanceSaveCall wrap *gomock.Call
type ITokenBalanceSaveCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *ITokenBalanceSaveCall) Return(arg0 error) *ITokenBalanceSaveCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *ITokenBalanceSaveCall) Do(f func(context.Context, *storage.TokenBalance) error) *ITokenBalanceSaveCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *ITokenBalanceSaveCall) DoAndReturn(f func(context.Context, *storage.TokenBalance) error) *ITokenBalanceSaveCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Update mocks base method.
func (m_2 *MockITokenBalance) Update(ctx context.Context, m *storage.TokenBalance) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Update", ctx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockITokenBalanceMockRecorder) Update(ctx, m any) *ITokenBalanceUpdateCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockITokenBalance)(nil).Update), ctx, m)
	return &ITokenBalanceUpdateCall{Call: call}
}

// ITokenBalanceUpdateCall wrap *gomock.Call
type ITokenBalanceUpdateCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *ITokenBalanceUpdateCall) Return(arg0 error) *ITokenBalanceUpdateCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *ITokenBalanceUpdateCall) Do(f func(context.Context, *storage.TokenBalance) error) *ITokenBalanceUpdateCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *ITokenBalanceUpdateCall) DoAndReturn(f func(context.Context, *storage.TokenBalance) error) *ITokenBalanceUpdateCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
