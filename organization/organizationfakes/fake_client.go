// This file was generated by counterfeiter
package organizationfakes

import (
	"sync"

	"github.com/3dsim/organization-goclient/models"
	"github.com/3dsim/organization-goclient/organization"
)

type FakeClient struct {
	OrganizationsStub        func() ([]*models.Organization, error)
	organizationsMutex       sync.RWMutex
	organizationsArgsForCall []struct{}
	organizationsReturns     struct {
		result1 []*models.Organization
		result2 error
	}
	organizationsReturnsOnCall map[int]struct {
		result1 []*models.Organization
		result2 error
	}
	OrganizationStub        func(organizationID int32) (*models.Organization, error)
	organizationMutex       sync.RWMutex
	organizationArgsForCall []struct {
		organizationID int32
	}
	organizationReturns struct {
		result1 *models.Organization
		result2 error
	}
	organizationReturnsOnCall map[int]struct {
		result1 *models.Organization
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeClient) Organizations() ([]*models.Organization, error) {
	fake.organizationsMutex.Lock()
	ret, specificReturn := fake.organizationsReturnsOnCall[len(fake.organizationsArgsForCall)]
	fake.organizationsArgsForCall = append(fake.organizationsArgsForCall, struct{}{})
	fake.recordInvocation("Organizations", []interface{}{})
	fake.organizationsMutex.Unlock()
	if fake.OrganizationsStub != nil {
		return fake.OrganizationsStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.organizationsReturns.result1, fake.organizationsReturns.result2
}

func (fake *FakeClient) OrganizationsCallCount() int {
	fake.organizationsMutex.RLock()
	defer fake.organizationsMutex.RUnlock()
	return len(fake.organizationsArgsForCall)
}

func (fake *FakeClient) OrganizationsReturns(result1 []*models.Organization, result2 error) {
	fake.OrganizationsStub = nil
	fake.organizationsReturns = struct {
		result1 []*models.Organization
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) OrganizationsReturnsOnCall(i int, result1 []*models.Organization, result2 error) {
	fake.OrganizationsStub = nil
	if fake.organizationsReturnsOnCall == nil {
		fake.organizationsReturnsOnCall = make(map[int]struct {
			result1 []*models.Organization
			result2 error
		})
	}
	fake.organizationsReturnsOnCall[i] = struct {
		result1 []*models.Organization
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) Organization(organizationID int32) (*models.Organization, error) {
	fake.organizationMutex.Lock()
	ret, specificReturn := fake.organizationReturnsOnCall[len(fake.organizationArgsForCall)]
	fake.organizationArgsForCall = append(fake.organizationArgsForCall, struct {
		organizationID int32
	}{organizationID})
	fake.recordInvocation("Organization", []interface{}{organizationID})
	fake.organizationMutex.Unlock()
	if fake.OrganizationStub != nil {
		return fake.OrganizationStub(organizationID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.organizationReturns.result1, fake.organizationReturns.result2
}

func (fake *FakeClient) OrganizationCallCount() int {
	fake.organizationMutex.RLock()
	defer fake.organizationMutex.RUnlock()
	return len(fake.organizationArgsForCall)
}

func (fake *FakeClient) OrganizationArgsForCall(i int) int32 {
	fake.organizationMutex.RLock()
	defer fake.organizationMutex.RUnlock()
	return fake.organizationArgsForCall[i].organizationID
}

func (fake *FakeClient) OrganizationReturns(result1 *models.Organization, result2 error) {
	fake.OrganizationStub = nil
	fake.organizationReturns = struct {
		result1 *models.Organization
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) OrganizationReturnsOnCall(i int, result1 *models.Organization, result2 error) {
	fake.OrganizationStub = nil
	if fake.organizationReturnsOnCall == nil {
		fake.organizationReturnsOnCall = make(map[int]struct {
			result1 *models.Organization
			result2 error
		})
	}
	fake.organizationReturnsOnCall[i] = struct {
		result1 *models.Organization
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.organizationsMutex.RLock()
	defer fake.organizationsMutex.RUnlock()
	fake.organizationMutex.RLock()
	defer fake.organizationMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeClient) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ organization.Client = new(FakeClient)
