package storage

import (
	"fmt"

	"github.com/jyangorch/hello-go/exercise-licensing-singlerepo/internal/domain/licensing"
)

type PackageRepoInMem struct {
	storage map[string]*licensing.Package
}

func NewPackageRepoInMem() *PackageRepoInMem {
	r := PackageRepoInMem{}
	r.storage = make(map[string]*licensing.Package)

	pkg1 := newAccelerateVersion2022Package()
	pkg2 := newOptimizeVersion2022Package()
	pkg3 := newOchestrateVersion2022Package()
	r.storage[pkg1.Id] = pkg1
	r.storage[pkg2.Id] = pkg2
	r.storage[pkg3.Id] = pkg3
	return &r
}

func (r *PackageRepoInMem) GetPackageById(pkgId string) (*licensing.Package, error) {
	if result, ok := r.storage[pkgId]; ok {
		return result, nil
	}
	return nil, fmt.Errorf("package not found for pkgId=%s", pkgId)
}

var sequenceCpb = licensing.Capability{Id: "cpb:sequence", DisplayName: "Squence", HasCapacityLimit: false, CapacityLimit: 0, CapacityLimitUnit: "NotApplicable"}
var calenderingCpb = licensing.Capability{Id: "cpb:calendaring", DisplayName: "Calendering", HasCapacityLimit: false, CapacityLimit: 0, CapacityLimitUnit: "NotApplicable"}
var basicReportingCpb = licensing.Capability{Id: "cpb:basic-reporting", DisplayName: "Basic Reporting", HasCapacityLimit: false, CapacityLimit: 0, CapacityLimitUnit: "NotApplicable"}
var basicOppViewCpb = licensing.Capability{Id: "cpb:basic-opportunity-view", DisplayName: "Basic Opportunity View", HasCapacityLimit: false, CapacityLimit: 0, CapacityLimitUnit: "NotApplicable"}

var sentimentCpb = licensing.Capability{Id: "cpb:sentiment", DisplayName: "ML Driven Sentiment", HasCapacityLimit: false, CapacityLimit: 0, CapacityLimitUnit: "NotApplicable"}
var advancedReportingCpb = licensing.Capability{Id: "cpb:advanced-reporting", DisplayName: "Advanced Reporting", HasCapacityLimit: false, CapacityLimit: 0, CapacityLimitUnit: "NotApplicable"}

var successPlanCpb = licensing.Capability{Id: "cpb:success-plan", DisplayName: "Success Plan", HasCapacityLimit: false, CapacityLimit: 0, CapacityLimitUnit: "NotApplicable"}
var kaiaMeetingPlanCpb = licensing.Capability{Id: "cpb:kaia-meeting", DisplayName: "Kaia Meeting Assistant", HasCapacityLimit: false, CapacityLimit: 0, CapacityLimitUnit: "NotApplicable"}

func newAccelerateVersion2022Package() *licensing.Package {
	crmCpb := licensing.Capability{Id: "cpb:crm-sync", DisplayName: "CRM Sync", HasCapacityLimit: true, CapacityLimit: 10000, CapacityLimitUnit: "CallsPerDay"}
	includedCapabilities := []licensing.Capability{
		sequenceCpb,
		calenderingCpb,
		basicOppViewCpb,
		crmCpb,
	}
	return &licensing.Package{
		Id:                   "pkg:base-accelerate-2022",
		Name:                 "Accelerate",
		IncludedCapabilities: includedCapabilities}
}

func newOptimizeVersion2022Package() *licensing.Package {
	crmCpb := licensing.Capability{Id: "cpb:crm-sync", DisplayName: "CRM Sync", HasCapacityLimit: true, CapacityLimit: 250000, CapacityLimitUnit: "CallsPerDay"}
	includedCapabilities := []licensing.Capability{
		sequenceCpb,
		calenderingCpb,
		basicOppViewCpb,
		sentimentCpb,
		advancedReportingCpb,
		crmCpb,
	}
	return &licensing.Package{
		Id:                   "pkg:base-optimize-2022",
		Name:                 "Optimize",
		IncludedCapabilities: includedCapabilities}
}

func newOchestrateVersion2022Package() *licensing.Package {
	crmCpb := licensing.Capability{Id: "cpb:crm-sync", DisplayName: "CRM Sync", HasCapacityLimit: true, CapacityLimit: 1000000, CapacityLimitUnit: "CallsPerDay"}
	includedCapabilities := []licensing.Capability{
		sequenceCpb,
		calenderingCpb,
		basicOppViewCpb,
		sentimentCpb,
		advancedReportingCpb,
		successPlanCpb,
		kaiaMeetingPlanCpb,
		crmCpb,
	}
	return &licensing.Package{
		Id:                   "pkg:base-ochestrate-2022",
		Name:                 "Ochestrate",
		IncludedCapabilities: includedCapabilities}
}
