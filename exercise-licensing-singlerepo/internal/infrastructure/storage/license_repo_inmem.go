package storage

import (
	"fmt"

	"github.com/jyangorch/hello-go/exercise-licensing-singlerepo/internal/domain/licensing"
)

type LicenseRepoInMem struct {
	storage map[string]*licensing.License
}

func NewLicenseRepoInMem() *LicenseRepoInMem {
	r := LicenseRepoInMem{}
	r.storage = make(map[string]*licensing.License)
	return &r
}

func (r *LicenseRepoInMem) CreateLicense(lic *licensing.License) error {
	r.storage[lic.Id()] = lic
	return nil
}

func (r *LicenseRepoInMem) UpdateLicense(licId string, newLic *licensing.License) error {
	r.storage[newLic.Id()] = newLic
	return nil
}

func (r *LicenseRepoInMem) GetLicenseById(licId string) (*licensing.License, error) {
	if result, ok := r.storage[licId]; ok {
		return result, nil
	}
	return nil, fmt.Errorf("license not found for id=%s", licId)
}

func (r *LicenseRepoInMem) FindLicensesByAssignedLicenseeId(licenseeId string) ([]*licensing.License, error) {
	results := make([]*licensing.License, 0)
	for _, elem := range r.storage {
		if elem.IsAssigned() && elem.AssignedToLicensee().LicenseeId() == licenseeId {
			results = append(results, elem)
			return results, nil
		}
	}
	return nil, fmt.Errorf("no license found assigned to licenseeId=%s", licenseeId)
}

func (r *LicenseRepoInMem) FindNextUnassignedLicenseOfPackage(accId string, pkgId string) (*licensing.License, error) {
	for _, elem := range r.storage {
		if elem.PossessingCustomerAccountId() == accId && elem.LicensedPackage().Id == pkgId && !elem.IsAssigned() {
			return elem, nil
		}
	}
	return nil, fmt.Errorf("no more unassigned license for pkgId=%s", pkgId)
}

func (r *LicenseRepoInMem) CountTotalUnassignedLicensesOfPackage(accId string, pkgId string) (int, error) {
	count := 0
	for _, elem := range r.storage {
		if elem.PossessingCustomerAccountId() == accId && elem.LicensedPackage().Id == pkgId && !elem.IsAssigned() {
			count++
		}
	}
	return count, nil
}
