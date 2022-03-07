package licensing

import (
	"github.com/jyangorch/hello-go/exercise-licensing-singlerepo/internal/domain/licensing"
)

// Licensing API Facade, an entry point to Licensing business logic. Each API represents a business use case.
//
// This Facade can be invoked by Driving Adapter in any form using any technology framework
// (e.g., CLI, RPC Service Activity, Queue Consumer Handler, Temporal Activity).
//
// DDD classification: Application Service
type LicensingService interface {

	// ------------------------------------------------------------------------------------------
	// Below are use cases for Outreach License Adminstration managing license lifecyles
	// ------------------------------------------------------------------------------------------
	// Issue X new licenses of the given package, to the given customer account, under the given subscription
	IssueLicenses(accId string, subId string, pkgId string, licenseCount int) ([]*licensing.License, error)

	// TODO: ExpireLicenses(accId string, subId string)
	// TODO: RenewLicenses(accId string, subId string)

	// ------------------------------------------------------------------------------------------
	// Below are use cases for Customer Admin managing user assignment
	// ------------------------------------------------------------------------------------------
	// Assign specific license id to user
	AssignSpecificLicense(licId string, accId string, insId string, insUsrId string) (*licensing.License, error)

	// Assign a random available license of a given package to a given user
	AssignAvailableLicenseOfPackage(pkgId string, accId string, insId string, insUsrId string) (*licensing.License, error)

	// Count the total unassigned licenses, possessed by the given customer account
	// FIXME: replace with a more generalize method like GatherLicenseAssignmentSummary returning total assigneds and unassigneds across all packages
	CountTotalUnassignedLicensesOfPackage(accId string, pkgId string) (int, error)

	// ------------------------------------------------------------------------------------------
	// Below are use cases for Outreach Application
	// ------------------------------------------------------------------------------------------
	// Verify an instance user has entitlement to the given capability
	VerifyEntitlement(accId string, insId string, insUsrId string, cpbId string) (licensing.Entitlement, error)
}

type licensingService struct {

	// underlying license repository interface to access licenses
	licRepo *licensing.LicenseRepository

	// underlying package repository interface to access packages
	pkgRepo *licensing.PackageRepository
}

func NewLicensingService(
	licRepo *licensing.LicenseRepository,
	pkgRepo *licensing.PackageRepository) *licensingService {
	return &licensingService{licRepo, pkgRepo}
}

func (ls *licensingService) IssueLicenses(accId string, subId string, pkgId string, licenseCount int) ([]*licensing.License, error) {
	pkg, err := (*ls.pkgRepo).GetPackageById(pkgId)
	if err != nil {
		return nil, err
	}
	return ls.IssueLicensesOfPackage(accId, subId, pkg, licenseCount)
}

func (ls *licensingService) IssueLicensesOfPackage(accId string, subId string, pkg *licensing.Package, licenseCount int) ([]*licensing.License, error) {
	results := make([]*licensing.License, licenseCount)
	for i := 0; i < licenseCount; i++ {
		lic := licensing.NewIssuedLicense(accId, subId, pkg)
		(*ls.licRepo).CreateLicense(lic)
		results[i] = lic
	}
	// This is where we trigger Application Events
	return results, nil
}

func (ls *licensingService) AssignAvailableLicenseOfPackage(pkgId string, accId string, insId string, insUsrId string) (*licensing.License, error) {
	availableLic, err := (*ls.licRepo).FindNextUnassignedLicenseOfPackage(accId, pkgId)
	if err != nil {
		return nil, err
	}
	return ls.assignSpecificLicenseHelper(availableLic, accId, insId, insUsrId)
}

func (ls *licensingService) AssignSpecificLicense(licId string, accId string, insId string, insUsrId string) (*licensing.License, error) {
	specificLic, err := (*ls.licRepo).GetLicenseById(licId)
	if err != nil {
		return nil, err
	}
	return ls.assignSpecificLicenseHelper(specificLic, accId, insId, insUsrId)
}

func (ls *licensingService) assignSpecificLicenseHelper(specificLic *licensing.License, accId string, insId string, insUsrId string) (*licensing.License, error) {
	specificLic.Assign(licensing.NewInstanceUser(insId, insUsrId))
	err := (*ls.licRepo).UpdateLicense(specificLic.Id(), specificLic)
	if err != nil {
		return nil, err
	}
	return specificLic, nil
}

func (ls *licensingService) VerifyEntitlement(accId string, insId string, insUsrId string, cpbId string) (licensing.Entitlement, error) {

	insUsr := licensing.NewInstanceUser(insId, insUsrId)
	licenses, err := (*ls.licRepo).FindLicensesByAssignedLicenseeId(insUsr.LicenseeId())

	if err != nil {
		return licensing.Entitlement{
			IsEntitled:            false,
			EvaluatedUserId:       insUsr.LicenseeId(),
			EvaluatedCapabilityId: cpbId}, nil
	}
	for _, lic := range licenses {
		pkg := lic.LicensedPackage()
		if pkg.IncludesCapability(cpbId) {
			return licensing.Entitlement{
				IsEntitled:            true,
				EvaluatedUserId:       insUsr.LicenseeId(),
				EvaluatedCapabilityId: cpbId}, nil
		}
	}
	return licensing.Entitlement{
		IsEntitled:            true,
		EvaluatedUserId:       insUsr.LicenseeId(),
		EvaluatedCapabilityId: cpbId}, nil
}

func (ls *licensingService) CountTotalUnassignedLicensesOfPackage(accId string, pkgId string) (int, error) {
	return (*ls.licRepo).CountTotalUnassignedLicensesOfPackage(accId, pkgId)
}
