/*
Copyright 2018 the Velero contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import (
	mock "github.com/stretchr/testify/mock"

	"github.com/heptio/velero/pkg/plugin/velero"
)

// Manager is an autogenerated mock type for the Manager type
type Manager struct {
	mock.Mock
}

// CleanupClients provides a mock function with given fields:
func (_m *Manager) CleanupClients() {
	_m.Called()
}

// GetBackupItemAction provides a mock function with given fields: name
func (_m *Manager) GetBackupItemAction(name string) (velero.BackupItemAction, error) {
	ret := _m.Called(name)

	var r0 velero.BackupItemAction
	if rf, ok := ret.Get(0).(func(string) velero.BackupItemAction); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(velero.BackupItemAction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBackupItemActions provides a mock function with given fields:
func (_m *Manager) GetBackupItemActions() ([]velero.BackupItemAction, error) {
	ret := _m.Called()

	var r0 []velero.BackupItemAction
	if rf, ok := ret.Get(0).(func() []velero.BackupItemAction); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]velero.BackupItemAction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBlockStore provides a mock function with given fields: name
func (_m *Manager) GetBlockStore(name string) (velero.BlockStore, error) {
	ret := _m.Called(name)

	var r0 velero.BlockStore
	if rf, ok := ret.Get(0).(func(string) velero.BlockStore); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(velero.BlockStore)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetObjectStore provides a mock function with given fields: name
func (_m *Manager) GetObjectStore(name string) (velero.ObjectStore, error) {
	ret := _m.Called(name)

	var r0 velero.ObjectStore
	if rf, ok := ret.Get(0).(func(string) velero.ObjectStore); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(velero.ObjectStore)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRestoreItemAction provides a mock function with given fields: name
func (_m *Manager) GetRestoreItemAction(name string) (velero.RestoreItemAction, error) {
	ret := _m.Called(name)

	var r0 velero.RestoreItemAction
	if rf, ok := ret.Get(0).(func(string) velero.RestoreItemAction); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(velero.RestoreItemAction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRestoreItemActions provides a mock function with given fields:
func (_m *Manager) GetRestoreItemActions() ([]velero.RestoreItemAction, error) {
	ret := _m.Called()

	var r0 []velero.RestoreItemAction
	if rf, ok := ret.Get(0).(func() []velero.RestoreItemAction); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]velero.RestoreItemAction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
