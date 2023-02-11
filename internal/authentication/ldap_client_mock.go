// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/authelia/authelia/v4/internal/authentication (interfaces: LDAPClient)

// Package authentication is a generated GoMock package.
package authentication

import (
	tls "crypto/tls"
	reflect "reflect"
	time "time"

	ldap "github.com/go-ldap/ldap/v3"
	gomock "github.com/golang/mock/gomock"
)

// MockLDAPClient is a mock of LDAPClient interface.
type MockLDAPClient struct {
	ctrl     *gomock.Controller
	recorder *MockLDAPClientMockRecorder
}

// MockLDAPClientMockRecorder is the mock recorder for MockLDAPClient.
type MockLDAPClientMockRecorder struct {
	mock *MockLDAPClient
}

// NewMockLDAPClient creates a new mock instance.
func NewMockLDAPClient(ctrl *gomock.Controller) *MockLDAPClient {
	mock := &MockLDAPClient{ctrl: ctrl}
	mock.recorder = &MockLDAPClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLDAPClient) EXPECT() *MockLDAPClientMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockLDAPClient) Add(arg0 *ldap.AddRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockLDAPClientMockRecorder) Add(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockLDAPClient)(nil).Add), arg0)
}

// Bind mocks base method.
func (m *MockLDAPClient) Bind(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Bind", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Bind indicates an expected call of Bind.
func (mr *MockLDAPClientMockRecorder) Bind(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bind", reflect.TypeOf((*MockLDAPClient)(nil).Bind), arg0, arg1)
}

// Close mocks base method.
func (m *MockLDAPClient) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockLDAPClientMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockLDAPClient)(nil).Close))
}

// Compare mocks base method.
func (m *MockLDAPClient) Compare(arg0, arg1, arg2 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Compare", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Compare indicates an expected call of Compare.
func (mr *MockLDAPClientMockRecorder) Compare(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Compare", reflect.TypeOf((*MockLDAPClient)(nil).Compare), arg0, arg1, arg2)
}

// Del mocks base method.
func (m *MockLDAPClient) Del(arg0 *ldap.DelRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Del", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Del indicates an expected call of Del.
func (mr *MockLDAPClientMockRecorder) Del(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Del", reflect.TypeOf((*MockLDAPClient)(nil).Del), arg0)
}

// DigestMD5Bind mocks base method.
func (m *MockLDAPClient) DigestMD5Bind(arg0 *ldap.DigestMD5BindRequest) (*ldap.DigestMD5BindResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DigestMD5Bind", arg0)
	ret0, _ := ret[0].(*ldap.DigestMD5BindResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DigestMD5Bind indicates an expected call of DigestMD5Bind.
func (mr *MockLDAPClientMockRecorder) DigestMD5Bind(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DigestMD5Bind", reflect.TypeOf((*MockLDAPClient)(nil).DigestMD5Bind), arg0)
}

// ExternalBind mocks base method.
func (m *MockLDAPClient) ExternalBind() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExternalBind")
	ret0, _ := ret[0].(error)
	return ret0
}

// ExternalBind indicates an expected call of ExternalBind.
func (mr *MockLDAPClientMockRecorder) ExternalBind() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExternalBind", reflect.TypeOf((*MockLDAPClient)(nil).ExternalBind))
}

// IsClosing mocks base method.
func (m *MockLDAPClient) IsClosing() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsClosing")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsClosing indicates an expected call of IsClosing.
func (mr *MockLDAPClientMockRecorder) IsClosing() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsClosing", reflect.TypeOf((*MockLDAPClient)(nil).IsClosing))
}

// MD5Bind mocks base method.
func (m *MockLDAPClient) MD5Bind(arg0, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MD5Bind", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// MD5Bind indicates an expected call of MD5Bind.
func (mr *MockLDAPClientMockRecorder) MD5Bind(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MD5Bind", reflect.TypeOf((*MockLDAPClient)(nil).MD5Bind), arg0, arg1, arg2)
}

// Modify mocks base method.
func (m *MockLDAPClient) Modify(arg0 *ldap.ModifyRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Modify", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Modify indicates an expected call of Modify.
func (mr *MockLDAPClientMockRecorder) Modify(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Modify", reflect.TypeOf((*MockLDAPClient)(nil).Modify), arg0)
}

// ModifyDN mocks base method.
func (m *MockLDAPClient) ModifyDN(arg0 *ldap.ModifyDNRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModifyDN", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ModifyDN indicates an expected call of ModifyDN.
func (mr *MockLDAPClientMockRecorder) ModifyDN(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyDN", reflect.TypeOf((*MockLDAPClient)(nil).ModifyDN), arg0)
}

// ModifyWithResult mocks base method.
func (m *MockLDAPClient) ModifyWithResult(arg0 *ldap.ModifyRequest) (*ldap.ModifyResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModifyWithResult", arg0)
	ret0, _ := ret[0].(*ldap.ModifyResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModifyWithResult indicates an expected call of ModifyWithResult.
func (mr *MockLDAPClientMockRecorder) ModifyWithResult(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModifyWithResult", reflect.TypeOf((*MockLDAPClient)(nil).ModifyWithResult), arg0)
}

// PasswordModify mocks base method.
func (m *MockLDAPClient) PasswordModify(arg0 *ldap.PasswordModifyRequest) (*ldap.PasswordModifyResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PasswordModify", arg0)
	ret0, _ := ret[0].(*ldap.PasswordModifyResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PasswordModify indicates an expected call of PasswordModify.
func (mr *MockLDAPClientMockRecorder) PasswordModify(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PasswordModify", reflect.TypeOf((*MockLDAPClient)(nil).PasswordModify), arg0)
}

// Search mocks base method.
func (m *MockLDAPClient) Search(arg0 *ldap.SearchRequest) (*ldap.SearchResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", arg0)
	ret0, _ := ret[0].(*ldap.SearchResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search.
func (mr *MockLDAPClientMockRecorder) Search(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockLDAPClient)(nil).Search), arg0)
}

// SearchWithPaging mocks base method.
func (m *MockLDAPClient) SearchWithPaging(arg0 *ldap.SearchRequest, arg1 uint32) (*ldap.SearchResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchWithPaging", arg0, arg1)
	ret0, _ := ret[0].(*ldap.SearchResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchWithPaging indicates an expected call of SearchWithPaging.
func (mr *MockLDAPClientMockRecorder) SearchWithPaging(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchWithPaging", reflect.TypeOf((*MockLDAPClient)(nil).SearchWithPaging), arg0, arg1)
}

// SetTimeout mocks base method.
func (m *MockLDAPClient) SetTimeout(arg0 time.Duration) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTimeout", arg0)
}

// SetTimeout indicates an expected call of SetTimeout.
func (mr *MockLDAPClientMockRecorder) SetTimeout(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTimeout", reflect.TypeOf((*MockLDAPClient)(nil).SetTimeout), arg0)
}

// SimpleBind mocks base method.
func (m *MockLDAPClient) SimpleBind(arg0 *ldap.SimpleBindRequest) (*ldap.SimpleBindResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SimpleBind", arg0)
	ret0, _ := ret[0].(*ldap.SimpleBindResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SimpleBind indicates an expected call of SimpleBind.
func (mr *MockLDAPClientMockRecorder) SimpleBind(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SimpleBind", reflect.TypeOf((*MockLDAPClient)(nil).SimpleBind), arg0)
}

// Start mocks base method.
func (m *MockLDAPClient) Start() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Start")
}

// Start indicates an expected call of Start.
func (mr *MockLDAPClientMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockLDAPClient)(nil).Start))
}

// StartTLS mocks base method.
func (m *MockLDAPClient) StartTLS(arg0 *tls.Config) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartTLS", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartTLS indicates an expected call of StartTLS.
func (mr *MockLDAPClientMockRecorder) StartTLS(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartTLS", reflect.TypeOf((*MockLDAPClient)(nil).StartTLS), arg0)
}

// TLSConnectionState mocks base method.
func (m *MockLDAPClient) TLSConnectionState() (tls.ConnectionState, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TLSConnectionState")
	ret0, _ := ret[0].(tls.ConnectionState)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// TLSConnectionState indicates an expected call of TLSConnectionState.
func (mr *MockLDAPClientMockRecorder) TLSConnectionState() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TLSConnectionState", reflect.TypeOf((*MockLDAPClient)(nil).TLSConnectionState))
}

// UnauthenticatedBind mocks base method.
func (m *MockLDAPClient) UnauthenticatedBind(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnauthenticatedBind", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnauthenticatedBind indicates an expected call of UnauthenticatedBind.
func (mr *MockLDAPClientMockRecorder) UnauthenticatedBind(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnauthenticatedBind", reflect.TypeOf((*MockLDAPClient)(nil).UnauthenticatedBind), arg0)
}

// Unbind mocks base method.
func (m *MockLDAPClient) Unbind() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unbind")
	ret0, _ := ret[0].(error)
	return ret0
}

// Unbind indicates an expected call of Unbind.
func (mr *MockLDAPClientMockRecorder) Unbind() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unbind", reflect.TypeOf((*MockLDAPClient)(nil).Unbind))
}

// WhoAmI mocks base method.
func (m *MockLDAPClient) WhoAmI(arg0 []ldap.Control) (*ldap.WhoAmIResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WhoAmI", arg0)
	ret0, _ := ret[0].(*ldap.WhoAmIResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WhoAmI indicates an expected call of WhoAmI.
func (mr *MockLDAPClientMockRecorder) WhoAmI(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WhoAmI", reflect.TypeOf((*MockLDAPClient)(nil).WhoAmI), arg0)
}
