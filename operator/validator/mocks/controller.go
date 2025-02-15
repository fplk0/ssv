// Code generated by MockGen. DO NOT EDIT.
// Source: ./controller.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	phase0 "github.com/attestantio/go-eth2-client/spec/phase0"
	types "github.com/bloxapp/ssv-spec/types"
	network "github.com/bloxapp/ssv/network"
	duties "github.com/bloxapp/ssv/operator/duties"
	beacon "github.com/bloxapp/ssv/protocol/v2/blockchain/beacon"
	protocolp2p "github.com/bloxapp/ssv/protocol/v2/p2p"
	validator "github.com/bloxapp/ssv/protocol/v2/ssv/validator"
	types0 "github.com/bloxapp/ssv/protocol/v2/types"
	storage "github.com/bloxapp/ssv/registry/storage"
	basedb "github.com/bloxapp/ssv/storage/basedb"
	common "github.com/ethereum/go-ethereum/common"
	gomock "github.com/golang/mock/gomock"
	peer "github.com/libp2p/go-libp2p/core/peer"
	zap "go.uber.org/zap"
)

// MockController is a mock of Controller interface.
type MockController struct {
	ctrl     *gomock.Controller
	recorder *MockControllerMockRecorder
}

// MockControllerMockRecorder is the mock recorder for MockController.
type MockControllerMockRecorder struct {
	mock *MockController
}

// NewMockController creates a new mock instance.
func NewMockController(ctrl *gomock.Controller) *MockController {
	mock := &MockController{ctrl: ctrl}
	mock.recorder = &MockControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockController) EXPECT() *MockControllerMockRecorder {
	return m.recorder
}

// AllActiveIndices mocks base method.
func (m *MockController) AllActiveIndices(epoch phase0.Epoch, afterInit bool) []phase0.ValidatorIndex {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllActiveIndices", epoch, afterInit)
	ret0, _ := ret[0].([]phase0.ValidatorIndex)
	return ret0
}

// AllActiveIndices indicates an expected call of AllActiveIndices.
func (mr *MockControllerMockRecorder) AllActiveIndices(epoch, afterInit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllActiveIndices", reflect.TypeOf((*MockController)(nil).AllActiveIndices), epoch, afterInit)
}

// CommitteeActiveIndices mocks base method.
func (m *MockController) CommitteeActiveIndices(epoch phase0.Epoch) []phase0.ValidatorIndex {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CommitteeActiveIndices", epoch)
	ret0, _ := ret[0].([]phase0.ValidatorIndex)
	return ret0
}

// CommitteeActiveIndices indicates an expected call of CommitteeActiveIndices.
func (mr *MockControllerMockRecorder) CommitteeActiveIndices(epoch interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CommitteeActiveIndices", reflect.TypeOf((*MockController)(nil).CommitteeActiveIndices), epoch)
}

// ExecuteDuty mocks base method.
func (m *MockController) ExecuteDuty(logger *zap.Logger, duty *types.Duty) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ExecuteDuty", logger, duty)
}

// ExecuteDuty indicates an expected call of ExecuteDuty.
func (mr *MockControllerMockRecorder) ExecuteDuty(logger, duty interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecuteDuty", reflect.TypeOf((*MockController)(nil).ExecuteDuty), logger, duty)
}

// ExitValidator mocks base method.
func (m *MockController) ExitValidator(pubKey phase0.BLSPubKey, blockNumber uint64, validatorIndex phase0.ValidatorIndex) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExitValidator", pubKey, blockNumber, validatorIndex)
	ret0, _ := ret[0].(error)
	return ret0
}

// ExitValidator indicates an expected call of ExitValidator.
func (mr *MockControllerMockRecorder) ExitValidator(pubKey, blockNumber, validatorIndex interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExitValidator", reflect.TypeOf((*MockController)(nil).ExitValidator), pubKey, blockNumber, validatorIndex)
}

// GetOperatorShares mocks base method.
func (m *MockController) GetOperatorShares() []*types0.SSVShare {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOperatorShares")
	ret0, _ := ret[0].([]*types0.SSVShare)
	return ret0
}

// GetOperatorShares indicates an expected call of GetOperatorShares.
func (mr *MockControllerMockRecorder) GetOperatorShares() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOperatorShares", reflect.TypeOf((*MockController)(nil).GetOperatorShares))
}

// GetValidator mocks base method.
func (m *MockController) GetValidator(pubKey string) (*validator.Validator, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetValidator", pubKey)
	ret0, _ := ret[0].(*validator.Validator)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetValidator indicates an expected call of GetValidator.
func (mr *MockControllerMockRecorder) GetValidator(pubKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetValidator", reflect.TypeOf((*MockController)(nil).GetValidator), pubKey)
}

// GetValidatorStats mocks base method.
func (m *MockController) GetValidatorStats() (uint64, uint64, uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetValidatorStats")
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(uint64)
	ret2, _ := ret[2].(uint64)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// GetValidatorStats indicates an expected call of GetValidatorStats.
func (mr *MockControllerMockRecorder) GetValidatorStats() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetValidatorStats", reflect.TypeOf((*MockController)(nil).GetValidatorStats))
}

// IndicesChangeChan mocks base method.
func (m *MockController) IndicesChangeChan() chan struct{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IndicesChangeChan")
	ret0, _ := ret[0].(chan struct{})
	return ret0
}

// IndicesChangeChan indicates an expected call of IndicesChangeChan.
func (mr *MockControllerMockRecorder) IndicesChangeChan() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IndicesChangeChan", reflect.TypeOf((*MockController)(nil).IndicesChangeChan))
}

// LiquidateCluster mocks base method.
func (m *MockController) LiquidateCluster(owner common.Address, operatorIDs []uint64, toLiquidate []*types0.SSVShare) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LiquidateCluster", owner, operatorIDs, toLiquidate)
	ret0, _ := ret[0].(error)
	return ret0
}

// LiquidateCluster indicates an expected call of LiquidateCluster.
func (mr *MockControllerMockRecorder) LiquidateCluster(owner, operatorIDs, toLiquidate interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LiquidateCluster", reflect.TypeOf((*MockController)(nil).LiquidateCluster), owner, operatorIDs, toLiquidate)
}

// ReactivateCluster mocks base method.
func (m *MockController) ReactivateCluster(owner common.Address, operatorIDs []uint64, toReactivate []*types0.SSVShare) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReactivateCluster", owner, operatorIDs, toReactivate)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReactivateCluster indicates an expected call of ReactivateCluster.
func (mr *MockControllerMockRecorder) ReactivateCluster(owner, operatorIDs, toReactivate interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReactivateCluster", reflect.TypeOf((*MockController)(nil).ReactivateCluster), owner, operatorIDs, toReactivate)
}

// StartNetworkHandlers mocks base method.
func (m *MockController) StartNetworkHandlers() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "StartNetworkHandlers")
}

// StartNetworkHandlers indicates an expected call of StartNetworkHandlers.
func (mr *MockControllerMockRecorder) StartNetworkHandlers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartNetworkHandlers", reflect.TypeOf((*MockController)(nil).StartNetworkHandlers))
}

// StartValidator mocks base method.
func (m *MockController) StartValidator(share *types0.SSVShare) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartValidator", share)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartValidator indicates an expected call of StartValidator.
func (mr *MockControllerMockRecorder) StartValidator(share interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartValidator", reflect.TypeOf((*MockController)(nil).StartValidator), share)
}

// StartValidators mocks base method.
func (m *MockController) StartValidators() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "StartValidators")
}

// StartValidators indicates an expected call of StartValidators.
func (mr *MockControllerMockRecorder) StartValidators() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartValidators", reflect.TypeOf((*MockController)(nil).StartValidators))
}

// StopValidator mocks base method.
func (m *MockController) StopValidator(pubKey types.ValidatorPK) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StopValidator", pubKey)
	ret0, _ := ret[0].(error)
	return ret0
}

// StopValidator indicates an expected call of StopValidator.
func (mr *MockControllerMockRecorder) StopValidator(pubKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopValidator", reflect.TypeOf((*MockController)(nil).StopValidator), pubKey)
}

// UpdateFeeRecipient mocks base method.
func (m *MockController) UpdateFeeRecipient(owner, recipient common.Address) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateFeeRecipient", owner, recipient)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateFeeRecipient indicates an expected call of UpdateFeeRecipient.
func (mr *MockControllerMockRecorder) UpdateFeeRecipient(owner, recipient interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFeeRecipient", reflect.TypeOf((*MockController)(nil).UpdateFeeRecipient), owner, recipient)
}

// UpdateValidatorMetaDataLoop mocks base method.
func (m *MockController) UpdateValidatorMetaDataLoop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateValidatorMetaDataLoop")
}

// UpdateValidatorMetaDataLoop indicates an expected call of UpdateValidatorMetaDataLoop.
func (mr *MockControllerMockRecorder) UpdateValidatorMetaDataLoop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateValidatorMetaDataLoop", reflect.TypeOf((*MockController)(nil).UpdateValidatorMetaDataLoop))
}

// ValidatorExitChan mocks base method.
func (m *MockController) ValidatorExitChan() <-chan duties.ExitDescriptor {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidatorExitChan")
	ret0, _ := ret[0].(<-chan duties.ExitDescriptor)
	return ret0
}

// ValidatorExitChan indicates an expected call of ValidatorExitChan.
func (mr *MockControllerMockRecorder) ValidatorExitChan() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidatorExitChan", reflect.TypeOf((*MockController)(nil).ValidatorExitChan))
}

// MockRecipients is a mock of Recipients interface.
type MockRecipients struct {
	ctrl     *gomock.Controller
	recorder *MockRecipientsMockRecorder
}

// MockRecipientsMockRecorder is the mock recorder for MockRecipients.
type MockRecipientsMockRecorder struct {
	mock *MockRecipients
}

// NewMockRecipients creates a new mock instance.
func NewMockRecipients(ctrl *gomock.Controller) *MockRecipients {
	mock := &MockRecipients{ctrl: ctrl}
	mock.recorder = &MockRecipientsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRecipients) EXPECT() *MockRecipientsMockRecorder {
	return m.recorder
}

// GetRecipientData mocks base method.
func (m *MockRecipients) GetRecipientData(r basedb.Reader, owner common.Address) (*storage.RecipientData, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRecipientData", r, owner)
	ret0, _ := ret[0].(*storage.RecipientData)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetRecipientData indicates an expected call of GetRecipientData.
func (mr *MockRecipientsMockRecorder) GetRecipientData(r, owner interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRecipientData", reflect.TypeOf((*MockRecipients)(nil).GetRecipientData), r, owner)
}

// MockSharesStorage is a mock of SharesStorage interface.
type MockSharesStorage struct {
	ctrl     *gomock.Controller
	recorder *MockSharesStorageMockRecorder
}

// MockSharesStorageMockRecorder is the mock recorder for MockSharesStorage.
type MockSharesStorageMockRecorder struct {
	mock *MockSharesStorage
}

// NewMockSharesStorage creates a new mock instance.
func NewMockSharesStorage(ctrl *gomock.Controller) *MockSharesStorage {
	mock := &MockSharesStorage{ctrl: ctrl}
	mock.recorder = &MockSharesStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSharesStorage) EXPECT() *MockSharesStorageMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockSharesStorage) Get(txn basedb.Reader, pubKey []byte) *types0.SSVShare {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", txn, pubKey)
	ret0, _ := ret[0].(*types0.SSVShare)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockSharesStorageMockRecorder) Get(txn, pubKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockSharesStorage)(nil).Get), txn, pubKey)
}

// List mocks base method.
func (m *MockSharesStorage) List(txn basedb.Reader, filters ...storage.SharesFilter) []*types0.SSVShare {
	m.ctrl.T.Helper()
	varargs := []interface{}{txn}
	for _, a := range filters {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].([]*types0.SSVShare)
	return ret0
}

// List indicates an expected call of List.
func (mr *MockSharesStorageMockRecorder) List(txn interface{}, filters ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{txn}, filters...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockSharesStorage)(nil).List), varargs...)
}

// UpdateValidatorMetadata mocks base method.
func (m *MockSharesStorage) UpdateValidatorMetadata(pk string, metadata *beacon.ValidatorMetadata) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateValidatorMetadata", pk, metadata)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateValidatorMetadata indicates an expected call of UpdateValidatorMetadata.
func (mr *MockSharesStorageMockRecorder) UpdateValidatorMetadata(pk, metadata interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateValidatorMetadata", reflect.TypeOf((*MockSharesStorage)(nil).UpdateValidatorMetadata), pk, metadata)
}

// MockP2PNetwork is a mock of P2PNetwork interface.
type MockP2PNetwork struct {
	ctrl     *gomock.Controller
	recorder *MockP2PNetworkMockRecorder
}

// MockP2PNetworkMockRecorder is the mock recorder for MockP2PNetwork.
type MockP2PNetworkMockRecorder struct {
	mock *MockP2PNetwork
}

// NewMockP2PNetwork creates a new mock instance.
func NewMockP2PNetwork(ctrl *gomock.Controller) *MockP2PNetwork {
	mock := &MockP2PNetwork{ctrl: ctrl}
	mock.recorder = &MockP2PNetworkMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockP2PNetwork) EXPECT() *MockP2PNetworkMockRecorder {
	return m.recorder
}

// Broadcast mocks base method.
func (m *MockP2PNetwork) Broadcast(message *types.SSVMessage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Broadcast", message)
	ret0, _ := ret[0].(error)
	return ret0
}

// Broadcast indicates an expected call of Broadcast.
func (mr *MockP2PNetworkMockRecorder) Broadcast(message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Broadcast", reflect.TypeOf((*MockP2PNetwork)(nil).Broadcast), message)
}

// Peers mocks base method.
func (m *MockP2PNetwork) Peers(pk types.ValidatorPK) ([]peer.ID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Peers", pk)
	ret0, _ := ret[0].([]peer.ID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Peers indicates an expected call of Peers.
func (mr *MockP2PNetworkMockRecorder) Peers(pk interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Peers", reflect.TypeOf((*MockP2PNetwork)(nil).Peers), pk)
}

// RegisterHandlers mocks base method.
func (m *MockP2PNetwork) RegisterHandlers(logger *zap.Logger, handlers ...*protocolp2p.SyncHandler) {
	m.ctrl.T.Helper()
	varargs := []interface{}{logger}
	for _, a := range handlers {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "RegisterHandlers", varargs...)
}

// RegisterHandlers indicates an expected call of RegisterHandlers.
func (mr *MockP2PNetworkMockRecorder) RegisterHandlers(logger interface{}, handlers ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{logger}, handlers...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterHandlers", reflect.TypeOf((*MockP2PNetwork)(nil).RegisterHandlers), varargs...)
}

// SubscribeRandoms mocks base method.
func (m *MockP2PNetwork) SubscribeRandoms(logger *zap.Logger, numSubnets int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscribeRandoms", logger, numSubnets)
	ret0, _ := ret[0].(error)
	return ret0
}

// SubscribeRandoms indicates an expected call of SubscribeRandoms.
func (mr *MockP2PNetworkMockRecorder) SubscribeRandoms(logger, numSubnets interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeRandoms", reflect.TypeOf((*MockP2PNetwork)(nil).SubscribeRandoms), logger, numSubnets)
}

// UseMessageRouter mocks base method.
func (m *MockP2PNetwork) UseMessageRouter(router network.MessageRouter) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UseMessageRouter", router)
}

// UseMessageRouter indicates an expected call of UseMessageRouter.
func (mr *MockP2PNetworkMockRecorder) UseMessageRouter(router interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UseMessageRouter", reflect.TypeOf((*MockP2PNetwork)(nil).UseMessageRouter), router)
}
