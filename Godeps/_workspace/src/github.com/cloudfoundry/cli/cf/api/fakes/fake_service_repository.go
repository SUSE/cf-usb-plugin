// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry/cli/cf/api"
	"github.com/cloudfoundry/cli/cf/api/resources"
	"github.com/cloudfoundry/cli/cf/models"
)

type FakeServiceRepository struct {
	PurgeServiceOfferingStub        func(offering models.ServiceOffering) error
	purgeServiceOfferingMutex       sync.RWMutex
	purgeServiceOfferingArgsForCall []struct {
		offering models.ServiceOffering
	}
	purgeServiceOfferingReturns struct {
		result1 error
	}
	GetServiceOfferingByGuidStub        func(serviceGuid string) (offering models.ServiceOffering, apiErr error)
	getServiceOfferingByGuidMutex       sync.RWMutex
	getServiceOfferingByGuidArgsForCall []struct {
		serviceGuid string
	}
	getServiceOfferingByGuidReturns struct {
		result1 models.ServiceOffering
		result2 error
	}
	FindServiceOfferingsByLabelStub        func(name string) (offering models.ServiceOfferings, apiErr error)
	findServiceOfferingsByLabelMutex       sync.RWMutex
	findServiceOfferingsByLabelArgsForCall []struct {
		name string
	}
	findServiceOfferingsByLabelReturns struct {
		result1 models.ServiceOfferings
		result2 error
	}
	FindServiceOfferingByLabelAndProviderStub        func(name, provider string) (offering models.ServiceOffering, apiErr error)
	findServiceOfferingByLabelAndProviderMutex       sync.RWMutex
	findServiceOfferingByLabelAndProviderArgsForCall []struct {
		name     string
		provider string
	}
	findServiceOfferingByLabelAndProviderReturns struct {
		result1 models.ServiceOffering
		result2 error
	}
	FindServiceOfferingsForSpaceByLabelStub        func(spaceGuid, name string) (offering models.ServiceOfferings, apiErr error)
	findServiceOfferingsForSpaceByLabelMutex       sync.RWMutex
	findServiceOfferingsForSpaceByLabelArgsForCall []struct {
		spaceGuid string
		name      string
	}
	findServiceOfferingsForSpaceByLabelReturns struct {
		result1 models.ServiceOfferings
		result2 error
	}
	GetAllServiceOfferingsStub        func() (offerings models.ServiceOfferings, apiErr error)
	getAllServiceOfferingsMutex       sync.RWMutex
	getAllServiceOfferingsArgsForCall []struct{}
	getAllServiceOfferingsReturns     struct {
		result1 models.ServiceOfferings
		result2 error
	}
	GetServiceOfferingsForSpaceStub        func(spaceGuid string) (offerings models.ServiceOfferings, apiErr error)
	getServiceOfferingsForSpaceMutex       sync.RWMutex
	getServiceOfferingsForSpaceArgsForCall []struct {
		spaceGuid string
	}
	getServiceOfferingsForSpaceReturns struct {
		result1 models.ServiceOfferings
		result2 error
	}
	FindInstanceByNameStub        func(name string) (instance models.ServiceInstance, apiErr error)
	findInstanceByNameMutex       sync.RWMutex
	findInstanceByNameArgsForCall []struct {
		name string
	}
	findInstanceByNameReturns struct {
		result1 models.ServiceInstance
		result2 error
	}
	PurgeServiceInstanceStub        func(instance models.ServiceInstance) error
	purgeServiceInstanceMutex       sync.RWMutex
	purgeServiceInstanceArgsForCall []struct {
		instance models.ServiceInstance
	}
	purgeServiceInstanceReturns struct {
		result1 error
	}
	CreateServiceInstanceStub        func(name, planGuid string, params map[string]interface{}, tags []string) (apiErr error)
	createServiceInstanceMutex       sync.RWMutex
	createServiceInstanceArgsForCall []struct {
		name     string
		planGuid string
		params   map[string]interface{}
		tags     []string
	}
	createServiceInstanceReturns struct {
		result1 error
	}
	UpdateServiceInstanceStub        func(instanceGuid, planGuid string, params map[string]interface{}, tags []string) (apiErr error)
	updateServiceInstanceMutex       sync.RWMutex
	updateServiceInstanceArgsForCall []struct {
		instanceGuid string
		planGuid     string
		params       map[string]interface{}
		tags         []string
	}
	updateServiceInstanceReturns struct {
		result1 error
	}
	RenameServiceStub        func(instance models.ServiceInstance, newName string) (apiErr error)
	renameServiceMutex       sync.RWMutex
	renameServiceArgsForCall []struct {
		instance models.ServiceInstance
		newName  string
	}
	renameServiceReturns struct {
		result1 error
	}
	DeleteServiceStub        func(instance models.ServiceInstance) (apiErr error)
	deleteServiceMutex       sync.RWMutex
	deleteServiceArgsForCall []struct {
		instance models.ServiceInstance
	}
	deleteServiceReturns struct {
		result1 error
	}
	FindServicePlanByDescriptionStub        func(planDescription resources.ServicePlanDescription) (planGuid string, apiErr error)
	findServicePlanByDescriptionMutex       sync.RWMutex
	findServicePlanByDescriptionArgsForCall []struct {
		planDescription resources.ServicePlanDescription
	}
	findServicePlanByDescriptionReturns struct {
		result1 string
		result2 error
	}
	ListServicesFromBrokerStub        func(brokerGuid string) (services []models.ServiceOffering, err error)
	listServicesFromBrokerMutex       sync.RWMutex
	listServicesFromBrokerArgsForCall []struct {
		brokerGuid string
	}
	listServicesFromBrokerReturns struct {
		result1 []models.ServiceOffering
		result2 error
	}
	ListServicesFromManyBrokersStub        func(brokerGuids []string) (services []models.ServiceOffering, err error)
	listServicesFromManyBrokersMutex       sync.RWMutex
	listServicesFromManyBrokersArgsForCall []struct {
		brokerGuids []string
	}
	listServicesFromManyBrokersReturns struct {
		result1 []models.ServiceOffering
		result2 error
	}
	GetServiceInstanceCountForServicePlanStub        func(v1PlanGuid string) (count int, apiErr error)
	getServiceInstanceCountForServicePlanMutex       sync.RWMutex
	getServiceInstanceCountForServicePlanArgsForCall []struct {
		v1PlanGuid string
	}
	getServiceInstanceCountForServicePlanReturns struct {
		result1 int
		result2 error
	}
	MigrateServicePlanFromV1ToV2Stub        func(v1PlanGuid, v2PlanGuid string) (changedCount int, apiErr error)
	migrateServicePlanFromV1ToV2Mutex       sync.RWMutex
	migrateServicePlanFromV1ToV2ArgsForCall []struct {
		v1PlanGuid string
		v2PlanGuid string
	}
	migrateServicePlanFromV1ToV2Returns struct {
		result1 int
		result2 error
	}
}

func (fake *FakeServiceRepository) PurgeServiceOffering(offering models.ServiceOffering) error {
	fake.purgeServiceOfferingMutex.Lock()
	fake.purgeServiceOfferingArgsForCall = append(fake.purgeServiceOfferingArgsForCall, struct {
		offering models.ServiceOffering
	}{offering})
	fake.purgeServiceOfferingMutex.Unlock()
	if fake.PurgeServiceOfferingStub != nil {
		return fake.PurgeServiceOfferingStub(offering)
	} else {
		return fake.purgeServiceOfferingReturns.result1
	}
}

func (fake *FakeServiceRepository) PurgeServiceOfferingCallCount() int {
	fake.purgeServiceOfferingMutex.RLock()
	defer fake.purgeServiceOfferingMutex.RUnlock()
	return len(fake.purgeServiceOfferingArgsForCall)
}

func (fake *FakeServiceRepository) PurgeServiceOfferingArgsForCall(i int) models.ServiceOffering {
	fake.purgeServiceOfferingMutex.RLock()
	defer fake.purgeServiceOfferingMutex.RUnlock()
	return fake.purgeServiceOfferingArgsForCall[i].offering
}

func (fake *FakeServiceRepository) PurgeServiceOfferingReturns(result1 error) {
	fake.PurgeServiceOfferingStub = nil
	fake.purgeServiceOfferingReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeServiceRepository) GetServiceOfferingByGuid(serviceGuid string) (offering models.ServiceOffering, apiErr error) {
	fake.getServiceOfferingByGuidMutex.Lock()
	fake.getServiceOfferingByGuidArgsForCall = append(fake.getServiceOfferingByGuidArgsForCall, struct {
		serviceGuid string
	}{serviceGuid})
	fake.getServiceOfferingByGuidMutex.Unlock()
	if fake.GetServiceOfferingByGuidStub != nil {
		return fake.GetServiceOfferingByGuidStub(serviceGuid)
	} else {
		return fake.getServiceOfferingByGuidReturns.result1, fake.getServiceOfferingByGuidReturns.result2
	}
}

func (fake *FakeServiceRepository) GetServiceOfferingByGuidCallCount() int {
	fake.getServiceOfferingByGuidMutex.RLock()
	defer fake.getServiceOfferingByGuidMutex.RUnlock()
	return len(fake.getServiceOfferingByGuidArgsForCall)
}

func (fake *FakeServiceRepository) GetServiceOfferingByGuidArgsForCall(i int) string {
	fake.getServiceOfferingByGuidMutex.RLock()
	defer fake.getServiceOfferingByGuidMutex.RUnlock()
	return fake.getServiceOfferingByGuidArgsForCall[i].serviceGuid
}

func (fake *FakeServiceRepository) GetServiceOfferingByGuidReturns(result1 models.ServiceOffering, result2 error) {
	fake.GetServiceOfferingByGuidStub = nil
	fake.getServiceOfferingByGuidReturns = struct {
		result1 models.ServiceOffering
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceRepository) FindServiceOfferingsByLabel(name string) (offering models.ServiceOfferings, apiErr error) {
	fake.findServiceOfferingsByLabelMutex.Lock()
	fake.findServiceOfferingsByLabelArgsForCall = append(fake.findServiceOfferingsByLabelArgsForCall, struct {
		name string
	}{name})
	fake.findServiceOfferingsByLabelMutex.Unlock()
	if fake.FindServiceOfferingsByLabelStub != nil {
		return fake.FindServiceOfferingsByLabelStub(name)
	} else {
		return fake.findServiceOfferingsByLabelReturns.result1, fake.findServiceOfferingsByLabelReturns.result2
	}
}

func (fake *FakeServiceRepository) FindServiceOfferingsByLabelCallCount() int {
	fake.findServiceOfferingsByLabelMutex.RLock()
	defer fake.findServiceOfferingsByLabelMutex.RUnlock()
	return len(fake.findServiceOfferingsByLabelArgsForCall)
}

func (fake *FakeServiceRepository) FindServiceOfferingsByLabelArgsForCall(i int) string {
	fake.findServiceOfferingsByLabelMutex.RLock()
	defer fake.findServiceOfferingsByLabelMutex.RUnlock()
	return fake.findServiceOfferingsByLabelArgsForCall[i].name
}

func (fake *FakeServiceRepository) FindServiceOfferingsByLabelReturns(result1 models.ServiceOfferings, result2 error) {
	fake.FindServiceOfferingsByLabelStub = nil
	fake.findServiceOfferingsByLabelReturns = struct {
		result1 models.ServiceOfferings
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceRepository) FindServiceOfferingByLabelAndProvider(name string, provider string) (offering models.ServiceOffering, apiErr error) {
	fake.findServiceOfferingByLabelAndProviderMutex.Lock()
	fake.findServiceOfferingByLabelAndProviderArgsForCall = append(fake.findServiceOfferingByLabelAndProviderArgsForCall, struct {
		name     string
		provider string
	}{name, provider})
	fake.findServiceOfferingByLabelAndProviderMutex.Unlock()
	if fake.FindServiceOfferingByLabelAndProviderStub != nil {
		return fake.FindServiceOfferingByLabelAndProviderStub(name, provider)
	} else {
		return fake.findServiceOfferingByLabelAndProviderReturns.result1, fake.findServiceOfferingByLabelAndProviderReturns.result2
	}
}

func (fake *FakeServiceRepository) FindServiceOfferingByLabelAndProviderCallCount() int {
	fake.findServiceOfferingByLabelAndProviderMutex.RLock()
	defer fake.findServiceOfferingByLabelAndProviderMutex.RUnlock()
	return len(fake.findServiceOfferingByLabelAndProviderArgsForCall)
}

func (fake *FakeServiceRepository) FindServiceOfferingByLabelAndProviderArgsForCall(i int) (string, string) {
	fake.findServiceOfferingByLabelAndProviderMutex.RLock()
	defer fake.findServiceOfferingByLabelAndProviderMutex.RUnlock()
	return fake.findServiceOfferingByLabelAndProviderArgsForCall[i].name, fake.findServiceOfferingByLabelAndProviderArgsForCall[i].provider
}

func (fake *FakeServiceRepository) FindServiceOfferingByLabelAndProviderReturns(result1 models.ServiceOffering, result2 error) {
	fake.FindServiceOfferingByLabelAndProviderStub = nil
	fake.findServiceOfferingByLabelAndProviderReturns = struct {
		result1 models.ServiceOffering
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceRepository) FindServiceOfferingsForSpaceByLabel(spaceGuid string, name string) (offering models.ServiceOfferings, apiErr error) {
	fake.findServiceOfferingsForSpaceByLabelMutex.Lock()
	fake.findServiceOfferingsForSpaceByLabelArgsForCall = append(fake.findServiceOfferingsForSpaceByLabelArgsForCall, struct {
		spaceGuid string
		name      string
	}{spaceGuid, name})
	fake.findServiceOfferingsForSpaceByLabelMutex.Unlock()
	if fake.FindServiceOfferingsForSpaceByLabelStub != nil {
		return fake.FindServiceOfferingsForSpaceByLabelStub(spaceGuid, name)
	} else {
		return fake.findServiceOfferingsForSpaceByLabelReturns.result1, fake.findServiceOfferingsForSpaceByLabelReturns.result2
	}
}

func (fake *FakeServiceRepository) FindServiceOfferingsForSpaceByLabelCallCount() int {
	fake.findServiceOfferingsForSpaceByLabelMutex.RLock()
	defer fake.findServiceOfferingsForSpaceByLabelMutex.RUnlock()
	return len(fake.findServiceOfferingsForSpaceByLabelArgsForCall)
}

func (fake *FakeServiceRepository) FindServiceOfferingsForSpaceByLabelArgsForCall(i int) (string, string) {
	fake.findServiceOfferingsForSpaceByLabelMutex.RLock()
	defer fake.findServiceOfferingsForSpaceByLabelMutex.RUnlock()
	return fake.findServiceOfferingsForSpaceByLabelArgsForCall[i].spaceGuid, fake.findServiceOfferingsForSpaceByLabelArgsForCall[i].name
}

func (fake *FakeServiceRepository) FindServiceOfferingsForSpaceByLabelReturns(result1 models.ServiceOfferings, result2 error) {
	fake.FindServiceOfferingsForSpaceByLabelStub = nil
	fake.findServiceOfferingsForSpaceByLabelReturns = struct {
		result1 models.ServiceOfferings
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceRepository) GetAllServiceOfferings() (offerings models.ServiceOfferings, apiErr error) {
	fake.getAllServiceOfferingsMutex.Lock()
	fake.getAllServiceOfferingsArgsForCall = append(fake.getAllServiceOfferingsArgsForCall, struct{}{})
	fake.getAllServiceOfferingsMutex.Unlock()
	if fake.GetAllServiceOfferingsStub != nil {
		return fake.GetAllServiceOfferingsStub()
	} else {
		return fake.getAllServiceOfferingsReturns.result1, fake.getAllServiceOfferingsReturns.result2
	}
}

func (fake *FakeServiceRepository) GetAllServiceOfferingsCallCount() int {
	fake.getAllServiceOfferingsMutex.RLock()
	defer fake.getAllServiceOfferingsMutex.RUnlock()
	return len(fake.getAllServiceOfferingsArgsForCall)
}

func (fake *FakeServiceRepository) GetAllServiceOfferingsReturns(result1 models.ServiceOfferings, result2 error) {
	fake.GetAllServiceOfferingsStub = nil
	fake.getAllServiceOfferingsReturns = struct {
		result1 models.ServiceOfferings
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceRepository) GetServiceOfferingsForSpace(spaceGuid string) (offerings models.ServiceOfferings, apiErr error) {
	fake.getServiceOfferingsForSpaceMutex.Lock()
	fake.getServiceOfferingsForSpaceArgsForCall = append(fake.getServiceOfferingsForSpaceArgsForCall, struct {
		spaceGuid string
	}{spaceGuid})
	fake.getServiceOfferingsForSpaceMutex.Unlock()
	if fake.GetServiceOfferingsForSpaceStub != nil {
		return fake.GetServiceOfferingsForSpaceStub(spaceGuid)
	} else {
		return fake.getServiceOfferingsForSpaceReturns.result1, fake.getServiceOfferingsForSpaceReturns.result2
	}
}

func (fake *FakeServiceRepository) GetServiceOfferingsForSpaceCallCount() int {
	fake.getServiceOfferingsForSpaceMutex.RLock()
	defer fake.getServiceOfferingsForSpaceMutex.RUnlock()
	return len(fake.getServiceOfferingsForSpaceArgsForCall)
}

func (fake *FakeServiceRepository) GetServiceOfferingsForSpaceArgsForCall(i int) string {
	fake.getServiceOfferingsForSpaceMutex.RLock()
	defer fake.getServiceOfferingsForSpaceMutex.RUnlock()
	return fake.getServiceOfferingsForSpaceArgsForCall[i].spaceGuid
}

func (fake *FakeServiceRepository) GetServiceOfferingsForSpaceReturns(result1 models.ServiceOfferings, result2 error) {
	fake.GetServiceOfferingsForSpaceStub = nil
	fake.getServiceOfferingsForSpaceReturns = struct {
		result1 models.ServiceOfferings
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceRepository) FindInstanceByName(name string) (instance models.ServiceInstance, apiErr error) {
	fake.findInstanceByNameMutex.Lock()
	fake.findInstanceByNameArgsForCall = append(fake.findInstanceByNameArgsForCall, struct {
		name string
	}{name})
	fake.findInstanceByNameMutex.Unlock()
	if fake.FindInstanceByNameStub != nil {
		return fake.FindInstanceByNameStub(name)
	} else {
		return fake.findInstanceByNameReturns.result1, fake.findInstanceByNameReturns.result2
	}
}

func (fake *FakeServiceRepository) FindInstanceByNameCallCount() int {
	fake.findInstanceByNameMutex.RLock()
	defer fake.findInstanceByNameMutex.RUnlock()
	return len(fake.findInstanceByNameArgsForCall)
}

func (fake *FakeServiceRepository) FindInstanceByNameArgsForCall(i int) string {
	fake.findInstanceByNameMutex.RLock()
	defer fake.findInstanceByNameMutex.RUnlock()
	return fake.findInstanceByNameArgsForCall[i].name
}

func (fake *FakeServiceRepository) FindInstanceByNameReturns(result1 models.ServiceInstance, result2 error) {
	fake.FindInstanceByNameStub = nil
	fake.findInstanceByNameReturns = struct {
		result1 models.ServiceInstance
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceRepository) PurgeServiceInstance(instance models.ServiceInstance) error {
	fake.purgeServiceInstanceMutex.Lock()
	fake.purgeServiceInstanceArgsForCall = append(fake.purgeServiceInstanceArgsForCall, struct {
		instance models.ServiceInstance
	}{instance})
	fake.purgeServiceInstanceMutex.Unlock()
	if fake.PurgeServiceInstanceStub != nil {
		return fake.PurgeServiceInstanceStub(instance)
	} else {
		return fake.purgeServiceInstanceReturns.result1
	}
}

func (fake *FakeServiceRepository) PurgeServiceInstanceCallCount() int {
	fake.purgeServiceInstanceMutex.RLock()
	defer fake.purgeServiceInstanceMutex.RUnlock()
	return len(fake.purgeServiceInstanceArgsForCall)
}

func (fake *FakeServiceRepository) PurgeServiceInstanceArgsForCall(i int) models.ServiceInstance {
	fake.purgeServiceInstanceMutex.RLock()
	defer fake.purgeServiceInstanceMutex.RUnlock()
	return fake.purgeServiceInstanceArgsForCall[i].instance
}

func (fake *FakeServiceRepository) PurgeServiceInstanceReturns(result1 error) {
	fake.PurgeServiceInstanceStub = nil
	fake.purgeServiceInstanceReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeServiceRepository) CreateServiceInstance(name string, planGuid string, params map[string]interface{}, tags []string) (apiErr error) {
	fake.createServiceInstanceMutex.Lock()
	fake.createServiceInstanceArgsForCall = append(fake.createServiceInstanceArgsForCall, struct {
		name     string
		planGuid string
		params   map[string]interface{}
		tags     []string
	}{name, planGuid, params, tags})
	fake.createServiceInstanceMutex.Unlock()
	if fake.CreateServiceInstanceStub != nil {
		return fake.CreateServiceInstanceStub(name, planGuid, params, tags)
	} else {
		return fake.createServiceInstanceReturns.result1
	}
}

func (fake *FakeServiceRepository) CreateServiceInstanceCallCount() int {
	fake.createServiceInstanceMutex.RLock()
	defer fake.createServiceInstanceMutex.RUnlock()
	return len(fake.createServiceInstanceArgsForCall)
}

func (fake *FakeServiceRepository) CreateServiceInstanceArgsForCall(i int) (string, string, map[string]interface{}, []string) {
	fake.createServiceInstanceMutex.RLock()
	defer fake.createServiceInstanceMutex.RUnlock()
	return fake.createServiceInstanceArgsForCall[i].name, fake.createServiceInstanceArgsForCall[i].planGuid, fake.createServiceInstanceArgsForCall[i].params, fake.createServiceInstanceArgsForCall[i].tags
}

func (fake *FakeServiceRepository) CreateServiceInstanceReturns(result1 error) {
	fake.CreateServiceInstanceStub = nil
	fake.createServiceInstanceReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeServiceRepository) UpdateServiceInstance(instanceGuid string, planGuid string, params map[string]interface{}, tags []string) (apiErr error) {
	fake.updateServiceInstanceMutex.Lock()
	fake.updateServiceInstanceArgsForCall = append(fake.updateServiceInstanceArgsForCall, struct {
		instanceGuid string
		planGuid     string
		params       map[string]interface{}
		tags         []string
	}{instanceGuid, planGuid, params, tags})
	fake.updateServiceInstanceMutex.Unlock()
	if fake.UpdateServiceInstanceStub != nil {
		return fake.UpdateServiceInstanceStub(instanceGuid, planGuid, params, tags)
	} else {
		return fake.updateServiceInstanceReturns.result1
	}
}

func (fake *FakeServiceRepository) UpdateServiceInstanceCallCount() int {
	fake.updateServiceInstanceMutex.RLock()
	defer fake.updateServiceInstanceMutex.RUnlock()
	return len(fake.updateServiceInstanceArgsForCall)
}

func (fake *FakeServiceRepository) UpdateServiceInstanceArgsForCall(i int) (string, string, map[string]interface{}, []string) {
	fake.updateServiceInstanceMutex.RLock()
	defer fake.updateServiceInstanceMutex.RUnlock()
	return fake.updateServiceInstanceArgsForCall[i].instanceGuid, fake.updateServiceInstanceArgsForCall[i].planGuid, fake.updateServiceInstanceArgsForCall[i].params, fake.updateServiceInstanceArgsForCall[i].tags
}

func (fake *FakeServiceRepository) UpdateServiceInstanceReturns(result1 error) {
	fake.UpdateServiceInstanceStub = nil
	fake.updateServiceInstanceReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeServiceRepository) RenameService(instance models.ServiceInstance, newName string) (apiErr error) {
	fake.renameServiceMutex.Lock()
	fake.renameServiceArgsForCall = append(fake.renameServiceArgsForCall, struct {
		instance models.ServiceInstance
		newName  string
	}{instance, newName})
	fake.renameServiceMutex.Unlock()
	if fake.RenameServiceStub != nil {
		return fake.RenameServiceStub(instance, newName)
	} else {
		return fake.renameServiceReturns.result1
	}
}

func (fake *FakeServiceRepository) RenameServiceCallCount() int {
	fake.renameServiceMutex.RLock()
	defer fake.renameServiceMutex.RUnlock()
	return len(fake.renameServiceArgsForCall)
}

func (fake *FakeServiceRepository) RenameServiceArgsForCall(i int) (models.ServiceInstance, string) {
	fake.renameServiceMutex.RLock()
	defer fake.renameServiceMutex.RUnlock()
	return fake.renameServiceArgsForCall[i].instance, fake.renameServiceArgsForCall[i].newName
}

func (fake *FakeServiceRepository) RenameServiceReturns(result1 error) {
	fake.RenameServiceStub = nil
	fake.renameServiceReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeServiceRepository) DeleteService(instance models.ServiceInstance) (apiErr error) {
	fake.deleteServiceMutex.Lock()
	fake.deleteServiceArgsForCall = append(fake.deleteServiceArgsForCall, struct {
		instance models.ServiceInstance
	}{instance})
	fake.deleteServiceMutex.Unlock()
	if fake.DeleteServiceStub != nil {
		return fake.DeleteServiceStub(instance)
	} else {
		return fake.deleteServiceReturns.result1
	}
}

func (fake *FakeServiceRepository) DeleteServiceCallCount() int {
	fake.deleteServiceMutex.RLock()
	defer fake.deleteServiceMutex.RUnlock()
	return len(fake.deleteServiceArgsForCall)
}

func (fake *FakeServiceRepository) DeleteServiceArgsForCall(i int) models.ServiceInstance {
	fake.deleteServiceMutex.RLock()
	defer fake.deleteServiceMutex.RUnlock()
	return fake.deleteServiceArgsForCall[i].instance
}

func (fake *FakeServiceRepository) DeleteServiceReturns(result1 error) {
	fake.DeleteServiceStub = nil
	fake.deleteServiceReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeServiceRepository) FindServicePlanByDescription(planDescription resources.ServicePlanDescription) (planGuid string, apiErr error) {
	fake.findServicePlanByDescriptionMutex.Lock()
	fake.findServicePlanByDescriptionArgsForCall = append(fake.findServicePlanByDescriptionArgsForCall, struct {
		planDescription resources.ServicePlanDescription
	}{planDescription})
	fake.findServicePlanByDescriptionMutex.Unlock()
	if fake.FindServicePlanByDescriptionStub != nil {
		return fake.FindServicePlanByDescriptionStub(planDescription)
	} else {
		return fake.findServicePlanByDescriptionReturns.result1, fake.findServicePlanByDescriptionReturns.result2
	}
}

func (fake *FakeServiceRepository) FindServicePlanByDescriptionCallCount() int {
	fake.findServicePlanByDescriptionMutex.RLock()
	defer fake.findServicePlanByDescriptionMutex.RUnlock()
	return len(fake.findServicePlanByDescriptionArgsForCall)
}

func (fake *FakeServiceRepository) FindServicePlanByDescriptionArgsForCall(i int) resources.ServicePlanDescription {
	fake.findServicePlanByDescriptionMutex.RLock()
	defer fake.findServicePlanByDescriptionMutex.RUnlock()
	return fake.findServicePlanByDescriptionArgsForCall[i].planDescription
}

func (fake *FakeServiceRepository) FindServicePlanByDescriptionReturns(result1 string, result2 error) {
	fake.FindServicePlanByDescriptionStub = nil
	fake.findServicePlanByDescriptionReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceRepository) ListServicesFromBroker(brokerGuid string) (services []models.ServiceOffering, err error) {
	fake.listServicesFromBrokerMutex.Lock()
	fake.listServicesFromBrokerArgsForCall = append(fake.listServicesFromBrokerArgsForCall, struct {
		brokerGuid string
	}{brokerGuid})
	fake.listServicesFromBrokerMutex.Unlock()
	if fake.ListServicesFromBrokerStub != nil {
		return fake.ListServicesFromBrokerStub(brokerGuid)
	} else {
		return fake.listServicesFromBrokerReturns.result1, fake.listServicesFromBrokerReturns.result2
	}
}

func (fake *FakeServiceRepository) ListServicesFromBrokerCallCount() int {
	fake.listServicesFromBrokerMutex.RLock()
	defer fake.listServicesFromBrokerMutex.RUnlock()
	return len(fake.listServicesFromBrokerArgsForCall)
}

func (fake *FakeServiceRepository) ListServicesFromBrokerArgsForCall(i int) string {
	fake.listServicesFromBrokerMutex.RLock()
	defer fake.listServicesFromBrokerMutex.RUnlock()
	return fake.listServicesFromBrokerArgsForCall[i].brokerGuid
}

func (fake *FakeServiceRepository) ListServicesFromBrokerReturns(result1 []models.ServiceOffering, result2 error) {
	fake.ListServicesFromBrokerStub = nil
	fake.listServicesFromBrokerReturns = struct {
		result1 []models.ServiceOffering
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceRepository) ListServicesFromManyBrokers(brokerGuids []string) (services []models.ServiceOffering, err error) {
	fake.listServicesFromManyBrokersMutex.Lock()
	fake.listServicesFromManyBrokersArgsForCall = append(fake.listServicesFromManyBrokersArgsForCall, struct {
		brokerGuids []string
	}{brokerGuids})
	fake.listServicesFromManyBrokersMutex.Unlock()
	if fake.ListServicesFromManyBrokersStub != nil {
		return fake.ListServicesFromManyBrokersStub(brokerGuids)
	} else {
		return fake.listServicesFromManyBrokersReturns.result1, fake.listServicesFromManyBrokersReturns.result2
	}
}

func (fake *FakeServiceRepository) ListServicesFromManyBrokersCallCount() int {
	fake.listServicesFromManyBrokersMutex.RLock()
	defer fake.listServicesFromManyBrokersMutex.RUnlock()
	return len(fake.listServicesFromManyBrokersArgsForCall)
}

func (fake *FakeServiceRepository) ListServicesFromManyBrokersArgsForCall(i int) []string {
	fake.listServicesFromManyBrokersMutex.RLock()
	defer fake.listServicesFromManyBrokersMutex.RUnlock()
	return fake.listServicesFromManyBrokersArgsForCall[i].brokerGuids
}

func (fake *FakeServiceRepository) ListServicesFromManyBrokersReturns(result1 []models.ServiceOffering, result2 error) {
	fake.ListServicesFromManyBrokersStub = nil
	fake.listServicesFromManyBrokersReturns = struct {
		result1 []models.ServiceOffering
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceRepository) GetServiceInstanceCountForServicePlan(v1PlanGuid string) (count int, apiErr error) {
	fake.getServiceInstanceCountForServicePlanMutex.Lock()
	fake.getServiceInstanceCountForServicePlanArgsForCall = append(fake.getServiceInstanceCountForServicePlanArgsForCall, struct {
		v1PlanGuid string
	}{v1PlanGuid})
	fake.getServiceInstanceCountForServicePlanMutex.Unlock()
	if fake.GetServiceInstanceCountForServicePlanStub != nil {
		return fake.GetServiceInstanceCountForServicePlanStub(v1PlanGuid)
	} else {
		return fake.getServiceInstanceCountForServicePlanReturns.result1, fake.getServiceInstanceCountForServicePlanReturns.result2
	}
}

func (fake *FakeServiceRepository) GetServiceInstanceCountForServicePlanCallCount() int {
	fake.getServiceInstanceCountForServicePlanMutex.RLock()
	defer fake.getServiceInstanceCountForServicePlanMutex.RUnlock()
	return len(fake.getServiceInstanceCountForServicePlanArgsForCall)
}

func (fake *FakeServiceRepository) GetServiceInstanceCountForServicePlanArgsForCall(i int) string {
	fake.getServiceInstanceCountForServicePlanMutex.RLock()
	defer fake.getServiceInstanceCountForServicePlanMutex.RUnlock()
	return fake.getServiceInstanceCountForServicePlanArgsForCall[i].v1PlanGuid
}

func (fake *FakeServiceRepository) GetServiceInstanceCountForServicePlanReturns(result1 int, result2 error) {
	fake.GetServiceInstanceCountForServicePlanStub = nil
	fake.getServiceInstanceCountForServicePlanReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceRepository) MigrateServicePlanFromV1ToV2(v1PlanGuid string, v2PlanGuid string) (changedCount int, apiErr error) {
	fake.migrateServicePlanFromV1ToV2Mutex.Lock()
	fake.migrateServicePlanFromV1ToV2ArgsForCall = append(fake.migrateServicePlanFromV1ToV2ArgsForCall, struct {
		v1PlanGuid string
		v2PlanGuid string
	}{v1PlanGuid, v2PlanGuid})
	fake.migrateServicePlanFromV1ToV2Mutex.Unlock()
	if fake.MigrateServicePlanFromV1ToV2Stub != nil {
		return fake.MigrateServicePlanFromV1ToV2Stub(v1PlanGuid, v2PlanGuid)
	} else {
		return fake.migrateServicePlanFromV1ToV2Returns.result1, fake.migrateServicePlanFromV1ToV2Returns.result2
	}
}

func (fake *FakeServiceRepository) MigrateServicePlanFromV1ToV2CallCount() int {
	fake.migrateServicePlanFromV1ToV2Mutex.RLock()
	defer fake.migrateServicePlanFromV1ToV2Mutex.RUnlock()
	return len(fake.migrateServicePlanFromV1ToV2ArgsForCall)
}

func (fake *FakeServiceRepository) MigrateServicePlanFromV1ToV2ArgsForCall(i int) (string, string) {
	fake.migrateServicePlanFromV1ToV2Mutex.RLock()
	defer fake.migrateServicePlanFromV1ToV2Mutex.RUnlock()
	return fake.migrateServicePlanFromV1ToV2ArgsForCall[i].v1PlanGuid, fake.migrateServicePlanFromV1ToV2ArgsForCall[i].v2PlanGuid
}

func (fake *FakeServiceRepository) MigrateServicePlanFromV1ToV2Returns(result1 int, result2 error) {
	fake.MigrateServicePlanFromV1ToV2Stub = nil
	fake.migrateServicePlanFromV1ToV2Returns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

var _ api.ServiceRepository = new(FakeServiceRepository)
