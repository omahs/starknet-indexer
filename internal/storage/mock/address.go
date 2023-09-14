// Code generated by MockGen. DO NOT EDIT.
// Source: address.go
//
// Generated by this command:
//
//	mockgen -source=address.go -destination=mock/address.go -package=mock -typed
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

// MockIAddress is a mock of IAddress interface.
type MockIAddress struct {
	ctrl     *gomock.Controller
	recorder *MockIAddressMockRecorder
}

// MockIAddressMockRecorder is the mock recorder for MockIAddress.
type MockIAddressMockRecorder struct {
	mock *MockIAddress
}

// NewMockIAddress creates a new mock instance.
func NewMockIAddress(ctrl *gomock.Controller) *MockIAddress {
	mock := &MockIAddress{ctrl: ctrl}
	mock.recorder = &MockIAddressMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIAddress) EXPECT() *MockIAddressMockRecorder {
	return m.recorder
}

// CursorList mocks base method.
func (m *MockIAddress) CursorList(ctx context.Context, id, limit uint64, order storage0.SortOrder, cmp storage0.Comparator) ([]*storage.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CursorList", ctx, id, limit, order, cmp)
	ret0, _ := ret[0].([]*storage.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CursorList indicates an expected call of CursorList.
func (mr *MockIAddressMockRecorder) CursorList(ctx, id, limit, order, cmp any) *IAddressCursorListCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CursorList", reflect.TypeOf((*MockIAddress)(nil).CursorList), ctx, id, limit, order, cmp)
	return &IAddressCursorListCall{Call: call}
}

// IAddressCursorListCall wrap *gomock.Call
type IAddressCursorListCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *IAddressCursorListCall) Return(arg0 []*storage.Address, arg1 error) *IAddressCursorListCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *IAddressCursorListCall) Do(f func(context.Context, uint64, uint64, storage0.SortOrder, storage0.Comparator) ([]*storage.Address, error)) *IAddressCursorListCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *IAddressCursorListCall) DoAndReturn(f func(context.Context, uint64, uint64, storage0.SortOrder, storage0.Comparator) ([]*storage.Address, error)) *IAddressCursorListCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Filter mocks base method.
func (m *MockIAddress) Filter(ctx context.Context, flt []storage.AddressFilter, opts ...storage.FilterOption) ([]storage.Address, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, flt}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Filter", varargs...)
	ret0, _ := ret[0].([]storage.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Filter indicates an expected call of Filter.
func (mr *MockIAddressMockRecorder) Filter(ctx, flt any, opts ...any) *IAddressFilterCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, flt}, opts...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Filter", reflect.TypeOf((*MockIAddress)(nil).Filter), varargs...)
	return &IAddressFilterCall{Call: call}
}

// IAddressFilterCall wrap *gomock.Call
type IAddressFilterCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *IAddressFilterCall) Return(arg0 []storage.Address, arg1 error) *IAddressFilterCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *IAddressFilterCall) Do(f func(context.Context, []storage.AddressFilter, ...storage.FilterOption) ([]storage.Address, error)) *IAddressFilterCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *IAddressFilterCall) DoAndReturn(f func(context.Context, []storage.AddressFilter, ...storage.FilterOption) ([]storage.Address, error)) *IAddressFilterCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetAddresses mocks base method.
func (m *MockIAddress) GetAddresses(ctx context.Context, ids ...uint64) ([]storage.Address, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx}
	for _, a := range ids {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAddresses", varargs...)
	ret0, _ := ret[0].([]storage.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAddresses indicates an expected call of GetAddresses.
func (mr *MockIAddressMockRecorder) GetAddresses(ctx any, ids ...any) *IAddressGetAddressesCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx}, ids...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAddresses", reflect.TypeOf((*MockIAddress)(nil).GetAddresses), varargs...)
	return &IAddressGetAddressesCall{Call: call}
}

// IAddressGetAddressesCall wrap *gomock.Call
type IAddressGetAddressesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *IAddressGetAddressesCall) Return(arg0 []storage.Address, arg1 error) *IAddressGetAddressesCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *IAddressGetAddressesCall) Do(f func(context.Context, ...uint64) ([]storage.Address, error)) *IAddressGetAddressesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *IAddressGetAddressesCall) DoAndReturn(f func(context.Context, ...uint64) ([]storage.Address, error)) *IAddressGetAddressesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetByHash mocks base method.
func (m *MockIAddress) GetByHash(ctx context.Context, hash []byte) (storage.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByHash", ctx, hash)
	ret0, _ := ret[0].(storage.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByHash indicates an expected call of GetByHash.
func (mr *MockIAddressMockRecorder) GetByHash(ctx, hash any) *IAddressGetByHashCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByHash", reflect.TypeOf((*MockIAddress)(nil).GetByHash), ctx, hash)
	return &IAddressGetByHashCall{Call: call}
}

// IAddressGetByHashCall wrap *gomock.Call
type IAddressGetByHashCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *IAddressGetByHashCall) Return(arg0 storage.Address, arg1 error) *IAddressGetByHashCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *IAddressGetByHashCall) Do(f func(context.Context, []byte) (storage.Address, error)) *IAddressGetByHashCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *IAddressGetByHashCall) DoAndReturn(f func(context.Context, []byte) (storage.Address, error)) *IAddressGetByHashCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetByHashes mocks base method.
func (m *MockIAddress) GetByHashes(ctx context.Context, hash [][]byte) ([]storage.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByHashes", ctx, hash)
	ret0, _ := ret[0].([]storage.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByHashes indicates an expected call of GetByHashes.
func (mr *MockIAddressMockRecorder) GetByHashes(ctx, hash any) *IAddressGetByHashesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByHashes", reflect.TypeOf((*MockIAddress)(nil).GetByHashes), ctx, hash)
	return &IAddressGetByHashesCall{Call: call}
}

// IAddressGetByHashesCall wrap *gomock.Call
type IAddressGetByHashesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *IAddressGetByHashesCall) Return(arg0 []storage.Address, arg1 error) *IAddressGetByHashesCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *IAddressGetByHashesCall) Do(f func(context.Context, [][]byte) ([]storage.Address, error)) *IAddressGetByHashesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *IAddressGetByHashesCall) DoAndReturn(f func(context.Context, [][]byte) ([]storage.Address, error)) *IAddressGetByHashesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetByID mocks base method.
func (m *MockIAddress) GetByID(ctx context.Context, id uint64) (*storage.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, id)
	ret0, _ := ret[0].(*storage.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockIAddressMockRecorder) GetByID(ctx, id any) *IAddressGetByIDCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockIAddress)(nil).GetByID), ctx, id)
	return &IAddressGetByIDCall{Call: call}
}

// IAddressGetByIDCall wrap *gomock.Call
type IAddressGetByIDCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *IAddressGetByIDCall) Return(arg0 *storage.Address, arg1 error) *IAddressGetByIDCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *IAddressGetByIDCall) Do(f func(context.Context, uint64) (*storage.Address, error)) *IAddressGetByIDCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *IAddressGetByIDCall) DoAndReturn(f func(context.Context, uint64) (*storage.Address, error)) *IAddressGetByIDCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// IsNoRows mocks base method.
func (m *MockIAddress) IsNoRows(err error) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsNoRows", err)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsNoRows indicates an expected call of IsNoRows.
func (mr *MockIAddressMockRecorder) IsNoRows(err any) *IAddressIsNoRowsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsNoRows", reflect.TypeOf((*MockIAddress)(nil).IsNoRows), err)
	return &IAddressIsNoRowsCall{Call: call}
}

// IAddressIsNoRowsCall wrap *gomock.Call
type IAddressIsNoRowsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *IAddressIsNoRowsCall) Return(arg0 bool) *IAddressIsNoRowsCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *IAddressIsNoRowsCall) Do(f func(error) bool) *IAddressIsNoRowsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *IAddressIsNoRowsCall) DoAndReturn(f func(error) bool) *IAddressIsNoRowsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// LastID mocks base method.
func (m *MockIAddress) LastID(ctx context.Context) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LastID", ctx)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LastID indicates an expected call of LastID.
func (mr *MockIAddressMockRecorder) LastID(ctx any) *IAddressLastIDCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LastID", reflect.TypeOf((*MockIAddress)(nil).LastID), ctx)
	return &IAddressLastIDCall{Call: call}
}

// IAddressLastIDCall wrap *gomock.Call
type IAddressLastIDCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *IAddressLastIDCall) Return(arg0 uint64, arg1 error) *IAddressLastIDCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *IAddressLastIDCall) Do(f func(context.Context) (uint64, error)) *IAddressLastIDCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *IAddressLastIDCall) DoAndReturn(f func(context.Context) (uint64, error)) *IAddressLastIDCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// List mocks base method.
func (m *MockIAddress) List(ctx context.Context, limit, offset uint64, order storage0.SortOrder) ([]*storage.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, limit, offset, order)
	ret0, _ := ret[0].([]*storage.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockIAddressMockRecorder) List(ctx, limit, offset, order any) *IAddressListCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockIAddress)(nil).List), ctx, limit, offset, order)
	return &IAddressListCall{Call: call}
}

// IAddressListCall wrap *gomock.Call
type IAddressListCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *IAddressListCall) Return(arg0 []*storage.Address, arg1 error) *IAddressListCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *IAddressListCall) Do(f func(context.Context, uint64, uint64, storage0.SortOrder) ([]*storage.Address, error)) *IAddressListCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *IAddressListCall) DoAndReturn(f func(context.Context, uint64, uint64, storage0.SortOrder) ([]*storage.Address, error)) *IAddressListCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Save mocks base method.
func (m_2 *MockIAddress) Save(ctx context.Context, m *storage.Address) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Save", ctx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockIAddressMockRecorder) Save(ctx, m any) *IAddressSaveCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockIAddress)(nil).Save), ctx, m)
	return &IAddressSaveCall{Call: call}
}

// IAddressSaveCall wrap *gomock.Call
type IAddressSaveCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *IAddressSaveCall) Return(arg0 error) *IAddressSaveCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *IAddressSaveCall) Do(f func(context.Context, *storage.Address) error) *IAddressSaveCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *IAddressSaveCall) DoAndReturn(f func(context.Context, *storage.Address) error) *IAddressSaveCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Update mocks base method.
func (m_2 *MockIAddress) Update(ctx context.Context, m *storage.Address) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Update", ctx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockIAddressMockRecorder) Update(ctx, m any) *IAddressUpdateCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIAddress)(nil).Update), ctx, m)
	return &IAddressUpdateCall{Call: call}
}

// IAddressUpdateCall wrap *gomock.Call
type IAddressUpdateCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *IAddressUpdateCall) Return(arg0 error) *IAddressUpdateCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *IAddressUpdateCall) Do(f func(context.Context, *storage.Address) error) *IAddressUpdateCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *IAddressUpdateCall) DoAndReturn(f func(context.Context, *storage.Address) error) *IAddressUpdateCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
