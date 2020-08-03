// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/state (interfaces: RemoteEntitiesInput,RelationNetworksInput,RemoteApplicationsInput,ApplicationOfferStateDocumentFactory,ApplicationOfferInput,ExternalControllerStateDocumentFactory,ExternalControllersInput,FirewallRulesInput)

// Package state is a generated GoMock package.
package state

import (
	gomock "github.com/golang/mock/gomock"
	description "github.com/juju/description/v2"
	txn "gopkg.in/mgo.v2/txn"
	reflect "reflect"
)

// MockRemoteEntitiesInput is a mock of RemoteEntitiesInput interface
type MockRemoteEntitiesInput struct {
	ctrl     *gomock.Controller
	recorder *MockRemoteEntitiesInputMockRecorder
}

// MockRemoteEntitiesInputMockRecorder is the mock recorder for MockRemoteEntitiesInput
type MockRemoteEntitiesInputMockRecorder struct {
	mock *MockRemoteEntitiesInput
}

// NewMockRemoteEntitiesInput creates a new mock instance
func NewMockRemoteEntitiesInput(ctrl *gomock.Controller) *MockRemoteEntitiesInput {
	mock := &MockRemoteEntitiesInput{ctrl: ctrl}
	mock.recorder = &MockRemoteEntitiesInputMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRemoteEntitiesInput) EXPECT() *MockRemoteEntitiesInputMockRecorder {
	return m.recorder
}

// DocID mocks base method
func (m *MockRemoteEntitiesInput) DocID(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DocID", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// DocID indicates an expected call of DocID
func (mr *MockRemoteEntitiesInputMockRecorder) DocID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DocID", reflect.TypeOf((*MockRemoteEntitiesInput)(nil).DocID), arg0)
}

// RemoteEntities mocks base method
func (m *MockRemoteEntitiesInput) RemoteEntities() []description.RemoteEntity {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoteEntities")
	ret0, _ := ret[0].([]description.RemoteEntity)
	return ret0
}

// RemoteEntities indicates an expected call of RemoteEntities
func (mr *MockRemoteEntitiesInputMockRecorder) RemoteEntities() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoteEntities", reflect.TypeOf((*MockRemoteEntitiesInput)(nil).RemoteEntities))
}

// MockRelationNetworksInput is a mock of RelationNetworksInput interface
type MockRelationNetworksInput struct {
	ctrl     *gomock.Controller
	recorder *MockRelationNetworksInputMockRecorder
}

// MockRelationNetworksInputMockRecorder is the mock recorder for MockRelationNetworksInput
type MockRelationNetworksInputMockRecorder struct {
	mock *MockRelationNetworksInput
}

// NewMockRelationNetworksInput creates a new mock instance
func NewMockRelationNetworksInput(ctrl *gomock.Controller) *MockRelationNetworksInput {
	mock := &MockRelationNetworksInput{ctrl: ctrl}
	mock.recorder = &MockRelationNetworksInputMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRelationNetworksInput) EXPECT() *MockRelationNetworksInputMockRecorder {
	return m.recorder
}

// DocID mocks base method
func (m *MockRelationNetworksInput) DocID(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DocID", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// DocID indicates an expected call of DocID
func (mr *MockRelationNetworksInputMockRecorder) DocID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DocID", reflect.TypeOf((*MockRelationNetworksInput)(nil).DocID), arg0)
}

// RelationNetworks mocks base method
func (m *MockRelationNetworksInput) RelationNetworks() []description.RelationNetwork {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RelationNetworks")
	ret0, _ := ret[0].([]description.RelationNetwork)
	return ret0
}

// RelationNetworks indicates an expected call of RelationNetworks
func (mr *MockRelationNetworksInputMockRecorder) RelationNetworks() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RelationNetworks", reflect.TypeOf((*MockRelationNetworksInput)(nil).RelationNetworks))
}

// MockRemoteApplicationsInput is a mock of RemoteApplicationsInput interface
type MockRemoteApplicationsInput struct {
	ctrl     *gomock.Controller
	recorder *MockRemoteApplicationsInputMockRecorder
}

// MockRemoteApplicationsInputMockRecorder is the mock recorder for MockRemoteApplicationsInput
type MockRemoteApplicationsInputMockRecorder struct {
	mock *MockRemoteApplicationsInput
}

// NewMockRemoteApplicationsInput creates a new mock instance
func NewMockRemoteApplicationsInput(ctrl *gomock.Controller) *MockRemoteApplicationsInput {
	mock := &MockRemoteApplicationsInput{ctrl: ctrl}
	mock.recorder = &MockRemoteApplicationsInputMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRemoteApplicationsInput) EXPECT() *MockRemoteApplicationsInputMockRecorder {
	return m.recorder
}

// DocID mocks base method
func (m *MockRemoteApplicationsInput) DocID(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DocID", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// DocID indicates an expected call of DocID
func (mr *MockRemoteApplicationsInputMockRecorder) DocID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DocID", reflect.TypeOf((*MockRemoteApplicationsInput)(nil).DocID), arg0)
}

// MakeRemoteApplicationDoc mocks base method
func (m *MockRemoteApplicationsInput) MakeRemoteApplicationDoc(arg0 description.RemoteApplication) *remoteApplicationDoc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeRemoteApplicationDoc", arg0)
	ret0, _ := ret[0].(*remoteApplicationDoc)
	return ret0
}

// MakeRemoteApplicationDoc indicates an expected call of MakeRemoteApplicationDoc
func (mr *MockRemoteApplicationsInputMockRecorder) MakeRemoteApplicationDoc(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeRemoteApplicationDoc", reflect.TypeOf((*MockRemoteApplicationsInput)(nil).MakeRemoteApplicationDoc), arg0)
}

// MakeStatusDoc mocks base method
func (m *MockRemoteApplicationsInput) MakeStatusDoc(arg0 description.Status) statusDoc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeStatusDoc", arg0)
	ret0, _ := ret[0].(statusDoc)
	return ret0
}

// MakeStatusDoc indicates an expected call of MakeStatusDoc
func (mr *MockRemoteApplicationsInputMockRecorder) MakeStatusDoc(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeStatusDoc", reflect.TypeOf((*MockRemoteApplicationsInput)(nil).MakeStatusDoc), arg0)
}

// MakeStatusOp mocks base method
func (m *MockRemoteApplicationsInput) MakeStatusOp(arg0 string, arg1 statusDoc) txn.Op {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeStatusOp", arg0, arg1)
	ret0, _ := ret[0].(txn.Op)
	return ret0
}

// MakeStatusOp indicates an expected call of MakeStatusOp
func (mr *MockRemoteApplicationsInputMockRecorder) MakeStatusOp(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeStatusOp", reflect.TypeOf((*MockRemoteApplicationsInput)(nil).MakeStatusOp), arg0, arg1)
}

// NewRemoteApplication mocks base method
func (m *MockRemoteApplicationsInput) NewRemoteApplication(arg0 *remoteApplicationDoc) *RemoteApplication {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewRemoteApplication", arg0)
	ret0, _ := ret[0].(*RemoteApplication)
	return ret0
}

// NewRemoteApplication indicates an expected call of NewRemoteApplication
func (mr *MockRemoteApplicationsInputMockRecorder) NewRemoteApplication(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewRemoteApplication", reflect.TypeOf((*MockRemoteApplicationsInput)(nil).NewRemoteApplication), arg0)
}

// RemoteApplications mocks base method
func (m *MockRemoteApplicationsInput) RemoteApplications() []description.RemoteApplication {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoteApplications")
	ret0, _ := ret[0].([]description.RemoteApplication)
	return ret0
}

// RemoteApplications indicates an expected call of RemoteApplications
func (mr *MockRemoteApplicationsInputMockRecorder) RemoteApplications() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoteApplications", reflect.TypeOf((*MockRemoteApplicationsInput)(nil).RemoteApplications))
}

// MockApplicationOfferStateDocumentFactory is a mock of ApplicationOfferStateDocumentFactory interface
type MockApplicationOfferStateDocumentFactory struct {
	ctrl     *gomock.Controller
	recorder *MockApplicationOfferStateDocumentFactoryMockRecorder
}

// MockApplicationOfferStateDocumentFactoryMockRecorder is the mock recorder for MockApplicationOfferStateDocumentFactory
type MockApplicationOfferStateDocumentFactoryMockRecorder struct {
	mock *MockApplicationOfferStateDocumentFactory
}

// NewMockApplicationOfferStateDocumentFactory creates a new mock instance
func NewMockApplicationOfferStateDocumentFactory(ctrl *gomock.Controller) *MockApplicationOfferStateDocumentFactory {
	mock := &MockApplicationOfferStateDocumentFactory{ctrl: ctrl}
	mock.recorder = &MockApplicationOfferStateDocumentFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockApplicationOfferStateDocumentFactory) EXPECT() *MockApplicationOfferStateDocumentFactoryMockRecorder {
	return m.recorder
}

// MakeApplicationOfferDoc mocks base method
func (m *MockApplicationOfferStateDocumentFactory) MakeApplicationOfferDoc(arg0 description.ApplicationOffer) (applicationOfferDoc, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeApplicationOfferDoc", arg0)
	ret0, _ := ret[0].(applicationOfferDoc)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MakeApplicationOfferDoc indicates an expected call of MakeApplicationOfferDoc
func (mr *MockApplicationOfferStateDocumentFactoryMockRecorder) MakeApplicationOfferDoc(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeApplicationOfferDoc", reflect.TypeOf((*MockApplicationOfferStateDocumentFactory)(nil).MakeApplicationOfferDoc), arg0)
}

// MakeIncApplicationOffersRefOp mocks base method
func (m *MockApplicationOfferStateDocumentFactory) MakeIncApplicationOffersRefOp(arg0 string) (txn.Op, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeIncApplicationOffersRefOp", arg0)
	ret0, _ := ret[0].(txn.Op)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MakeIncApplicationOffersRefOp indicates an expected call of MakeIncApplicationOffersRefOp
func (mr *MockApplicationOfferStateDocumentFactoryMockRecorder) MakeIncApplicationOffersRefOp(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeIncApplicationOffersRefOp", reflect.TypeOf((*MockApplicationOfferStateDocumentFactory)(nil).MakeIncApplicationOffersRefOp), arg0)
}

// MockApplicationOfferInput is a mock of ApplicationOfferInput interface
type MockApplicationOfferInput struct {
	ctrl     *gomock.Controller
	recorder *MockApplicationOfferInputMockRecorder
}

// MockApplicationOfferInputMockRecorder is the mock recorder for MockApplicationOfferInput
type MockApplicationOfferInputMockRecorder struct {
	mock *MockApplicationOfferInput
}

// NewMockApplicationOfferInput creates a new mock instance
func NewMockApplicationOfferInput(ctrl *gomock.Controller) *MockApplicationOfferInput {
	mock := &MockApplicationOfferInput{ctrl: ctrl}
	mock.recorder = &MockApplicationOfferInputMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockApplicationOfferInput) EXPECT() *MockApplicationOfferInputMockRecorder {
	return m.recorder
}

// DocID mocks base method
func (m *MockApplicationOfferInput) DocID(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DocID", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// DocID indicates an expected call of DocID
func (mr *MockApplicationOfferInputMockRecorder) DocID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DocID", reflect.TypeOf((*MockApplicationOfferInput)(nil).DocID), arg0)
}

// MakeApplicationOfferDoc mocks base method
func (m *MockApplicationOfferInput) MakeApplicationOfferDoc(arg0 description.ApplicationOffer) (applicationOfferDoc, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeApplicationOfferDoc", arg0)
	ret0, _ := ret[0].(applicationOfferDoc)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MakeApplicationOfferDoc indicates an expected call of MakeApplicationOfferDoc
func (mr *MockApplicationOfferInputMockRecorder) MakeApplicationOfferDoc(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeApplicationOfferDoc", reflect.TypeOf((*MockApplicationOfferInput)(nil).MakeApplicationOfferDoc), arg0)
}

// MakeIncApplicationOffersRefOp mocks base method
func (m *MockApplicationOfferInput) MakeIncApplicationOffersRefOp(arg0 string) (txn.Op, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeIncApplicationOffersRefOp", arg0)
	ret0, _ := ret[0].(txn.Op)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MakeIncApplicationOffersRefOp indicates an expected call of MakeIncApplicationOffersRefOp
func (mr *MockApplicationOfferInputMockRecorder) MakeIncApplicationOffersRefOp(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeIncApplicationOffersRefOp", reflect.TypeOf((*MockApplicationOfferInput)(nil).MakeIncApplicationOffersRefOp), arg0)
}

// Offers mocks base method
func (m *MockApplicationOfferInput) Offers() []description.ApplicationOffer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Offers")
	ret0, _ := ret[0].([]description.ApplicationOffer)
	return ret0
}

// Offers indicates an expected call of Offers
func (mr *MockApplicationOfferInputMockRecorder) Offers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Offers", reflect.TypeOf((*MockApplicationOfferInput)(nil).Offers))
}

// MockExternalControllerStateDocumentFactory is a mock of ExternalControllerStateDocumentFactory interface
type MockExternalControllerStateDocumentFactory struct {
	ctrl     *gomock.Controller
	recorder *MockExternalControllerStateDocumentFactoryMockRecorder
}

// MockExternalControllerStateDocumentFactoryMockRecorder is the mock recorder for MockExternalControllerStateDocumentFactory
type MockExternalControllerStateDocumentFactoryMockRecorder struct {
	mock *MockExternalControllerStateDocumentFactory
}

// NewMockExternalControllerStateDocumentFactory creates a new mock instance
func NewMockExternalControllerStateDocumentFactory(ctrl *gomock.Controller) *MockExternalControllerStateDocumentFactory {
	mock := &MockExternalControllerStateDocumentFactory{ctrl: ctrl}
	mock.recorder = &MockExternalControllerStateDocumentFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockExternalControllerStateDocumentFactory) EXPECT() *MockExternalControllerStateDocumentFactoryMockRecorder {
	return m.recorder
}

// ExternalControllerDoc mocks base method
func (m *MockExternalControllerStateDocumentFactory) ExternalControllerDoc(arg0 string) (*externalControllerDoc, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExternalControllerDoc", arg0)
	ret0, _ := ret[0].(*externalControllerDoc)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExternalControllerDoc indicates an expected call of ExternalControllerDoc
func (mr *MockExternalControllerStateDocumentFactoryMockRecorder) ExternalControllerDoc(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExternalControllerDoc", reflect.TypeOf((*MockExternalControllerStateDocumentFactory)(nil).ExternalControllerDoc), arg0)
}

// MakeExternalControllerOp mocks base method
func (m *MockExternalControllerStateDocumentFactory) MakeExternalControllerOp(arg0 externalControllerDoc, arg1 *externalControllerDoc) txn.Op {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeExternalControllerOp", arg0, arg1)
	ret0, _ := ret[0].(txn.Op)
	return ret0
}

// MakeExternalControllerOp indicates an expected call of MakeExternalControllerOp
func (mr *MockExternalControllerStateDocumentFactoryMockRecorder) MakeExternalControllerOp(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeExternalControllerOp", reflect.TypeOf((*MockExternalControllerStateDocumentFactory)(nil).MakeExternalControllerOp), arg0, arg1)
}

// MockExternalControllersInput is a mock of ExternalControllersInput interface
type MockExternalControllersInput struct {
	ctrl     *gomock.Controller
	recorder *MockExternalControllersInputMockRecorder
}

// MockExternalControllersInputMockRecorder is the mock recorder for MockExternalControllersInput
type MockExternalControllersInputMockRecorder struct {
	mock *MockExternalControllersInput
}

// NewMockExternalControllersInput creates a new mock instance
func NewMockExternalControllersInput(ctrl *gomock.Controller) *MockExternalControllersInput {
	mock := &MockExternalControllersInput{ctrl: ctrl}
	mock.recorder = &MockExternalControllersInputMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockExternalControllersInput) EXPECT() *MockExternalControllersInputMockRecorder {
	return m.recorder
}

// ExternalControllerDoc mocks base method
func (m *MockExternalControllersInput) ExternalControllerDoc(arg0 string) (*externalControllerDoc, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExternalControllerDoc", arg0)
	ret0, _ := ret[0].(*externalControllerDoc)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExternalControllerDoc indicates an expected call of ExternalControllerDoc
func (mr *MockExternalControllersInputMockRecorder) ExternalControllerDoc(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExternalControllerDoc", reflect.TypeOf((*MockExternalControllersInput)(nil).ExternalControllerDoc), arg0)
}

// ExternalControllers mocks base method
func (m *MockExternalControllersInput) ExternalControllers() []description.ExternalController {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExternalControllers")
	ret0, _ := ret[0].([]description.ExternalController)
	return ret0
}

// ExternalControllers indicates an expected call of ExternalControllers
func (mr *MockExternalControllersInputMockRecorder) ExternalControllers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExternalControllers", reflect.TypeOf((*MockExternalControllersInput)(nil).ExternalControllers))
}

// MakeExternalControllerOp mocks base method
func (m *MockExternalControllersInput) MakeExternalControllerOp(arg0 externalControllerDoc, arg1 *externalControllerDoc) txn.Op {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MakeExternalControllerOp", arg0, arg1)
	ret0, _ := ret[0].(txn.Op)
	return ret0
}

// MakeExternalControllerOp indicates an expected call of MakeExternalControllerOp
func (mr *MockExternalControllersInputMockRecorder) MakeExternalControllerOp(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeExternalControllerOp", reflect.TypeOf((*MockExternalControllersInput)(nil).MakeExternalControllerOp), arg0, arg1)
}

// MockFirewallRulesInput is a mock of FirewallRulesInput interface
type MockFirewallRulesInput struct {
	ctrl     *gomock.Controller
	recorder *MockFirewallRulesInputMockRecorder
}

// MockFirewallRulesInputMockRecorder is the mock recorder for MockFirewallRulesInput
type MockFirewallRulesInputMockRecorder struct {
	mock *MockFirewallRulesInput
}

// NewMockFirewallRulesInput creates a new mock instance
func NewMockFirewallRulesInput(ctrl *gomock.Controller) *MockFirewallRulesInput {
	mock := &MockFirewallRulesInput{ctrl: ctrl}
	mock.recorder = &MockFirewallRulesInputMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFirewallRulesInput) EXPECT() *MockFirewallRulesInputMockRecorder {
	return m.recorder
}

// FirewallRules mocks base method
func (m *MockFirewallRulesInput) FirewallRules() []description.FirewallRule {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FirewallRules")
	ret0, _ := ret[0].([]description.FirewallRule)
	return ret0
}

// FirewallRules indicates an expected call of FirewallRules
func (mr *MockFirewallRulesInputMockRecorder) FirewallRules() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FirewallRules", reflect.TypeOf((*MockFirewallRulesInput)(nil).FirewallRules))
}
