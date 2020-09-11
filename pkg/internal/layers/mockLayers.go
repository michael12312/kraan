// Code generated by MockGen. DO NOT EDIT.
// Source: layers.go

package layers

import (
	context "context"
	v1alpha1 "github.com/fidelity/kraan/pkg/api/v1alpha1"
	logr "github.com/go-logr/logr"
	gomock "github.com/golang/mock/gomock"
	kubernetes "k8s.io/client-go/kubernetes"
	reflect "reflect"
	time "time"
)

// MockLayer is a mock of Layer interface
type MockLayer struct {
	ctrl     *gomock.Controller
	recorder *MockLayerMockRecorder
}

// MockLayerMockRecorder is the mock recorder for MockLayer
type MockLayerMockRecorder struct {
	mock *MockLayer
}

// NewMockLayer creates a new mock instance
func NewMockLayer(ctrl *gomock.Controller) *MockLayer {
	mock := &MockLayer{ctrl: ctrl}
	mock.recorder = &MockLayerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockLayer) EXPECT() *MockLayerMockRecorder {
	return _m.recorder
}

// SetStatusK8sVersion mocks base method
func (_m *MockLayer) SetStatusK8sVersion() {
	_m.ctrl.Call(_m, "SetStatusK8sVersion")
}

// SetStatusK8sVersion indicates an expected call of SetStatusK8sVersion
func (_mr *MockLayerMockRecorder) SetStatusK8sVersion() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetStatusK8sVersion", reflect.TypeOf((*MockLayer)(nil).SetStatusK8sVersion))
}

// SetStatusApplying mocks base method
func (_m *MockLayer) SetStatusApplying() {
	_m.ctrl.Call(_m, "SetStatusApplying")
}

// SetStatusApplying indicates an expected call of SetStatusApplying
func (_mr *MockLayerMockRecorder) SetStatusApplying() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetStatusApplying", reflect.TypeOf((*MockLayer)(nil).SetStatusApplying))
}

// SetStatusPruning mocks base method
func (_m *MockLayer) SetStatusPruning() {
	_m.ctrl.Call(_m, "SetStatusPruning")
}

// SetStatusPruning indicates an expected call of SetStatusPruning
func (_mr *MockLayerMockRecorder) SetStatusPruning() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetStatusPruning", reflect.TypeOf((*MockLayer)(nil).SetStatusPruning))
}

// SetStatusDeployed mocks base method
func (_m *MockLayer) SetStatusDeployed() {
	_m.ctrl.Call(_m, "SetStatusDeployed")
}

// SetStatusDeployed indicates an expected call of SetStatusDeployed
func (_mr *MockLayerMockRecorder) SetStatusDeployed() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetStatusDeployed", reflect.TypeOf((*MockLayer)(nil).SetStatusDeployed))
}

// StatusUpdate mocks base method
func (_m *MockLayer) StatusUpdate(status string, reason string, message string) {
	_m.ctrl.Call(_m, "StatusUpdate", status, reason, message)
}

// StatusUpdate indicates an expected call of StatusUpdate
func (_mr *MockLayerMockRecorder) StatusUpdate(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "StatusUpdate", reflect.TypeOf((*MockLayer)(nil).StatusUpdate), arg0, arg1, arg2)
}

// IsHold mocks base method
func (_m *MockLayer) IsHold() bool {
	ret := _m.ctrl.Call(_m, "IsHold")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsHold indicates an expected call of IsHold
func (_mr *MockLayerMockRecorder) IsHold() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "IsHold", reflect.TypeOf((*MockLayer)(nil).IsHold))
}

// SetHold mocks base method
func (_m *MockLayer) SetHold() {
	_m.ctrl.Call(_m, "SetHold")
}

// SetHold indicates an expected call of SetHold
func (_mr *MockLayerMockRecorder) SetHold() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetHold", reflect.TypeOf((*MockLayer)(nil).SetHold))
}

// DependenciesDeployed mocks base method
func (_m *MockLayer) DependenciesDeployed() bool {
	ret := _m.ctrl.Call(_m, "DependenciesDeployed")
	ret0, _ := ret[0].(bool)
	return ret0
}

// DependenciesDeployed indicates an expected call of DependenciesDeployed
func (_mr *MockLayerMockRecorder) DependenciesDeployed() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "DependenciesDeployed", reflect.TypeOf((*MockLayer)(nil).DependenciesDeployed))
}

// GetStatus mocks base method
func (_m *MockLayer) GetStatus() string {
	ret := _m.ctrl.Call(_m, "GetStatus")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetStatus indicates an expected call of GetStatus
func (_mr *MockLayerMockRecorder) GetStatus() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetStatus", reflect.TypeOf((*MockLayer)(nil).GetStatus))
}

// GetName mocks base method
func (_m *MockLayer) GetName() string {
	ret := _m.ctrl.Call(_m, "GetName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetName indicates an expected call of GetName
func (_mr *MockLayerMockRecorder) GetName() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetName", reflect.TypeOf((*MockLayer)(nil).GetName))
}

// GetLogger mocks base method
func (_m *MockLayer) GetLogger() logr.Logger {
	ret := _m.ctrl.Call(_m, "GetLogger")
	ret0, _ := ret[0].(logr.Logger)
	return ret0
}

// GetLogger indicates an expected call of GetLogger
func (_mr *MockLayerMockRecorder) GetLogger() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetLogger", reflect.TypeOf((*MockLayer)(nil).GetLogger))
}

// GetContext mocks base method
func (_m *MockLayer) GetContext() context.Context {
	ret := _m.ctrl.Call(_m, "GetContext")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// GetContext indicates an expected call of GetContext
func (_mr *MockLayerMockRecorder) GetContext() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetContext", reflect.TypeOf((*MockLayer)(nil).GetContext))
}

// GetSourcePath mocks base method
func (_m *MockLayer) GetSourcePath() string {
	ret := _m.ctrl.Call(_m, "GetSourcePath")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetSourcePath indicates an expected call of GetSourcePath
func (_mr *MockLayerMockRecorder) GetSourcePath() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetSourcePath", reflect.TypeOf((*MockLayer)(nil).GetSourcePath))
}

// GetTimeout mocks base method
func (_m *MockLayer) GetTimeout() time.Duration {
	ret := _m.ctrl.Call(_m, "GetTimeout")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// GetTimeout indicates an expected call of GetTimeout
func (_mr *MockLayerMockRecorder) GetTimeout() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetTimeout", reflect.TypeOf((*MockLayer)(nil).GetTimeout))
}

// IsUpdated mocks base method
func (_m *MockLayer) IsUpdated() bool {
	ret := _m.ctrl.Call(_m, "IsUpdated")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsUpdated indicates an expected call of IsUpdated
func (_mr *MockLayerMockRecorder) IsUpdated() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "IsUpdated", reflect.TypeOf((*MockLayer)(nil).IsUpdated))
}

// NeedsRequeue mocks base method
func (_m *MockLayer) NeedsRequeue() bool {
	ret := _m.ctrl.Call(_m, "NeedsRequeue")
	ret0, _ := ret[0].(bool)
	return ret0
}

// NeedsRequeue indicates an expected call of NeedsRequeue
func (_mr *MockLayerMockRecorder) NeedsRequeue() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "NeedsRequeue", reflect.TypeOf((*MockLayer)(nil).NeedsRequeue))
}

// IsDelayed mocks base method
func (_m *MockLayer) IsDelayed() bool {
	ret := _m.ctrl.Call(_m, "IsDelayed")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsDelayed indicates an expected call of IsDelayed
func (_mr *MockLayerMockRecorder) IsDelayed() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "IsDelayed", reflect.TypeOf((*MockLayer)(nil).IsDelayed))
}

// GetDelay mocks base method
func (_m *MockLayer) GetDelay() time.Duration {
	ret := _m.ctrl.Call(_m, "GetDelay")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// GetDelay indicates an expected call of GetDelay
func (_mr *MockLayerMockRecorder) GetDelay() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetDelay", reflect.TypeOf((*MockLayer)(nil).GetDelay))
}

// SetRequeue mocks base method
func (_m *MockLayer) SetRequeue() {
	_m.ctrl.Call(_m, "SetRequeue")
}

// SetRequeue indicates an expected call of SetRequeue
func (_mr *MockLayerMockRecorder) SetRequeue() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetRequeue", reflect.TypeOf((*MockLayer)(nil).SetRequeue))
}

// SetDelayedRequeue mocks base method
func (_m *MockLayer) SetDelayedRequeue() {
	_m.ctrl.Call(_m, "SetDelayedRequeue")
}

// SetDelayedRequeue indicates an expected call of SetDelayedRequeue
func (_mr *MockLayerMockRecorder) SetDelayedRequeue() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetDelayedRequeue", reflect.TypeOf((*MockLayer)(nil).SetDelayedRequeue))
}

// SetUpdated mocks base method
func (_m *MockLayer) SetUpdated() {
	_m.ctrl.Call(_m, "SetUpdated")
}

// SetUpdated indicates an expected call of SetUpdated
func (_mr *MockLayerMockRecorder) SetUpdated() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SetUpdated", reflect.TypeOf((*MockLayer)(nil).SetUpdated))
}

// GetRequiredK8sVersion mocks base method
func (_m *MockLayer) GetRequiredK8sVersion() string {
	ret := _m.ctrl.Call(_m, "GetRequiredK8sVersion")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetRequiredK8sVersion indicates an expected call of GetRequiredK8sVersion
func (_mr *MockLayerMockRecorder) GetRequiredK8sVersion() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetRequiredK8sVersion", reflect.TypeOf((*MockLayer)(nil).GetRequiredK8sVersion))
}

// CheckK8sVersion mocks base method
func (_m *MockLayer) CheckK8sVersion() bool {
	ret := _m.ctrl.Call(_m, "CheckK8sVersion")
	ret0, _ := ret[0].(bool)
	return ret0
}

// CheckK8sVersion indicates an expected call of CheckK8sVersion
func (_mr *MockLayerMockRecorder) CheckK8sVersion() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "CheckK8sVersion", reflect.TypeOf((*MockLayer)(nil).CheckK8sVersion))
}

// GetFullStatus mocks base method
func (_m *MockLayer) GetFullStatus() *v1alpha1.AddonsLayerStatus {
	ret := _m.ctrl.Call(_m, "GetFullStatus")
	ret0, _ := ret[0].(*v1alpha1.AddonsLayerStatus)
	return ret0
}

// GetFullStatus indicates an expected call of GetFullStatus
func (_mr *MockLayerMockRecorder) GetFullStatus() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetFullStatus", reflect.TypeOf((*MockLayer)(nil).GetFullStatus))
}

// GetSpec mocks base method
func (_m *MockLayer) GetSpec() *v1alpha1.AddonsLayerSpec {
	ret := _m.ctrl.Call(_m, "GetSpec")
	ret0, _ := ret[0].(*v1alpha1.AddonsLayerSpec)
	return ret0
}

// GetSpec indicates an expected call of GetSpec
func (_mr *MockLayerMockRecorder) GetSpec() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetSpec", reflect.TypeOf((*MockLayer)(nil).GetSpec))
}

// GetAddonsLayer mocks base method
func (_m *MockLayer) GetAddonsLayer() *v1alpha1.AddonsLayer {
	ret := _m.ctrl.Call(_m, "GetAddonsLayer")
	ret0, _ := ret[0].(*v1alpha1.AddonsLayer)
	return ret0
}

// GetAddonsLayer indicates an expected call of GetAddonsLayer
func (_mr *MockLayerMockRecorder) GetAddonsLayer() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetAddonsLayer", reflect.TypeOf((*MockLayer)(nil).GetAddonsLayer))
}

// getOtherAddonsLayer mocks base method
func (_m *MockLayer) getOtherAddonsLayer(name string) (*v1alpha1.AddonsLayer, error) {
	ret := _m.ctrl.Call(_m, "getOtherAddonsLayer", name)
	ret0, _ := ret[0].(*v1alpha1.AddonsLayer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// getOtherAddonsLayer indicates an expected call of getOtherAddonsLayer
func (_mr *MockLayerMockRecorder) getOtherAddonsLayer(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "getOtherAddonsLayer", reflect.TypeOf((*MockLayer)(nil).getOtherAddonsLayer), arg0)
}

// getK8sClient mocks base method
func (_m *MockLayer) getK8sClient() kubernetes.Interface {
	ret := _m.ctrl.Call(_m, "getK8sClient")
	ret0, _ := ret[0].(kubernetes.Interface)
	return ret0
}

// getK8sClient indicates an expected call of getK8sClient
func (_mr *MockLayerMockRecorder) getK8sClient() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "getK8sClient", reflect.TypeOf((*MockLayer)(nil).getK8sClient))
}

// setStatus mocks base method
func (_m *MockLayer) setStatus(status string, reason string, message string) {
	_m.ctrl.Call(_m, "setStatus", status, reason, message)
}

// setStatus indicates an expected call of setStatus
func (_mr *MockLayerMockRecorder) setStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "setStatus", reflect.TypeOf((*MockLayer)(nil).setStatus), arg0, arg1, arg2)
}

// isOtherDeployed mocks base method
func (_m *MockLayer) isOtherDeployed(otherVersion string, otherLayer *v1alpha1.AddonsLayer) bool {
	ret := _m.ctrl.Call(_m, "isOtherDeployed", otherVersion, otherLayer)
	ret0, _ := ret[0].(bool)
	return ret0
}

// isOtherDeployed indicates an expected call of isOtherDeployed
func (_mr *MockLayerMockRecorder) isOtherDeployed(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "isOtherDeployed", reflect.TypeOf((*MockLayer)(nil).isOtherDeployed), arg0, arg1)
}
