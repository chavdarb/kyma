// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	"github.com/kyma-project/kyma/components/metadata-service/internal/apperrors"
	"github.com/kyma-project/kyma/components/metadata-service/internal/metadata/model"
)
import "github.com/stretchr/testify/mock"

// ServiceDefinitionService is an autogenerated mock type for the ServiceDefinitionService type
type ServiceDefinitionService struct {
	mock.Mock
}

// Create provides a mock function with given fields: remoteEnvironment, serviceDefinition
func (_m *ServiceDefinitionService) Create(remoteEnvironment string, serviceDefinition *model.ServiceDefinition) (string, apperrors.AppError) {
	ret := _m.Called(remoteEnvironment, serviceDefinition)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, *model.ServiceDefinition) string); ok {
		r0 = rf(remoteEnvironment, serviceDefinition)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 apperrors.AppError
	if rf, ok := ret.Get(1).(func(string, *model.ServiceDefinition) apperrors.AppError); ok {
		r1 = rf(remoteEnvironment, serviceDefinition)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(apperrors.AppError)
		}
	}

	return r0, r1
}

// Delete provides a mock function with given fields: remoteEnvironment, id
func (_m *ServiceDefinitionService) Delete(remoteEnvironment string, id string) apperrors.AppError {
	ret := _m.Called(remoteEnvironment, id)

	var r0 apperrors.AppError
	if rf, ok := ret.Get(0).(func(string, string) apperrors.AppError); ok {
		r0 = rf(remoteEnvironment, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(apperrors.AppError)
		}
	}

	return r0
}

// GetAPI provides a mock function with given fields: remoteEnvironment, serviceId
func (_m *ServiceDefinitionService) GetAPI(remoteEnvironment string, serviceId string) (*model.API, apperrors.AppError) {
	ret := _m.Called(remoteEnvironment, serviceId)

	var r0 *model.API
	if rf, ok := ret.Get(0).(func(string, string) *model.API); ok {
		r0 = rf(remoteEnvironment, serviceId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.API)
		}
	}

	var r1 apperrors.AppError
	if rf, ok := ret.Get(1).(func(string, string) apperrors.AppError); ok {
		r1 = rf(remoteEnvironment, serviceId)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(apperrors.AppError)
		}
	}

	return r0, r1
}

// GetAll provides a mock function with given fields: remoteEnvironment
func (_m *ServiceDefinitionService) GetAll(remoteEnvironment string) ([]model.ServiceDefinition, apperrors.AppError) {
	ret := _m.Called(remoteEnvironment)

	var r0 []model.ServiceDefinition
	if rf, ok := ret.Get(0).(func(string) []model.ServiceDefinition); ok {
		r0 = rf(remoteEnvironment)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.ServiceDefinition)
		}
	}

	var r1 apperrors.AppError
	if rf, ok := ret.Get(1).(func(string) apperrors.AppError); ok {
		r1 = rf(remoteEnvironment)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(apperrors.AppError)
		}
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: remoteEnvironment, id
func (_m *ServiceDefinitionService) GetByID(remoteEnvironment string, id string) (model.ServiceDefinition, apperrors.AppError) {
	ret := _m.Called(remoteEnvironment, id)

	var r0 model.ServiceDefinition
	if rf, ok := ret.Get(0).(func(string, string) model.ServiceDefinition); ok {
		r0 = rf(remoteEnvironment, id)
	} else {
		r0 = ret.Get(0).(model.ServiceDefinition)
	}

	var r1 apperrors.AppError
	if rf, ok := ret.Get(1).(func(string, string) apperrors.AppError); ok {
		r1 = rf(remoteEnvironment, id)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(apperrors.AppError)
		}
	}

	return r0, r1
}

// Update provides a mock function with given fields: remoteEnvironment, id, serviceDef
func (_m *ServiceDefinitionService) Update(remoteEnvironment string, id string, serviceDef *model.ServiceDefinition) (model.ServiceDefinition, apperrors.AppError) {
	ret := _m.Called(remoteEnvironment, id, serviceDef)

	var r0 model.ServiceDefinition
	if rf, ok := ret.Get(0).(func(string, string, *model.ServiceDefinition) model.ServiceDefinition); ok {
		r0 = rf(remoteEnvironment, id, serviceDef)
	} else {
		r0 = ret.Get(0).(model.ServiceDefinition)
	}

	var r1 apperrors.AppError
	if rf, ok := ret.Get(1).(func(string, string, *model.ServiceDefinition) apperrors.AppError); ok {
		r1 = rf(remoteEnvironment, id, serviceDef)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(apperrors.AppError)
		}
	}

	return r0, r1
}
