package licensing

import (
	"testing"

	"github.com/jyangorch/hello-go/exercise-licensing-singlerepo/internal/domain/licensing"
	"github.com/jyangorch/hello-go/exercise-licensing-singlerepo/internal/infrastructure/storage"
	"gotest.tools/v3/assert"
)

func TestIssueLicenses(t *testing.T) {

	var licRepo licensing.LicenseRepository = storage.NewLicenseRepoInMem()
	var pkgRepo licensing.PackageRepository = storage.NewPackageRepoInMem()
	ls := NewLicensingService(&licRepo, &pkgRepo)

	accId := "acc-1"
	subId := "sub-1"
	pkgId := "pkg:base-optimize-2022"

	t.Run("happy case", func(t *testing.T) {
		licenses, err := ls.IssueLicenses(accId, subId, pkgId, 3)
		assert.NilError(t, err)
		assert.Equal(t, 3, len(licenses))
		// print out for debugging
		for _, lic := range licenses {
			t.Log(lic)
		}
	})
}

func TestAssignAvailableLicenseOfPackage(t *testing.T) {

	var licRepo licensing.LicenseRepository = storage.NewLicenseRepoInMem()
	var pkgRepo licensing.PackageRepository = storage.NewPackageRepoInMem()
	ls := NewLicensingService(&licRepo, &pkgRepo)

	accId := "acc-1"
	subId := "sub-1"
	pkgId := "pkg:base-optimize-2022"
	insId := "ins-101"
	insUsrIdAlice := "usr-alice"
	insUsrIdBob := "usr-bob"
	insUsrIdCharles := "usr-charles"
	insUsrIdDaniel := "usr-daniel"

	_, err := ls.IssueLicenses(accId, subId, pkgId, 3)
	assert.NilError(t, err)

	t.Run("assign 2 licenses should suceed", func(t *testing.T) {
		lic, err := ls.AssignAvailableLicenseOfPackage(pkgId, accId, insId, insUsrIdAlice)
		t.Log(lic)
		assert.NilError(t, err)
		assert.Check(t, lic != nil)
		lic, err = ls.AssignAvailableLicenseOfPackage(pkgId, accId, insId, insUsrIdBob)
		t.Log(lic)
		assert.NilError(t, err)
		assert.Check(t, lic != nil)
		unassignedLicensesCount, err := ls.CountTotalUnassignedLicensesOfPackage(accId, pkgId)
		assert.NilError(t, err)
		assert.Equal(t, unassignedLicensesCount, 1)
	})

	t.Run("assign 1 more licenses should succeed", func(t *testing.T) {
		ls.AssignAvailableLicenseOfPackage(pkgId, accId, insId, insUsrIdCharles)
		unassignedLicensesCount, err := ls.CountTotalUnassignedLicensesOfPackage(accId, pkgId)
		assert.NilError(t, err)
		assert.Equal(t, unassignedLicensesCount, 0)
	})

	t.Run("assign 1 more licenses should fail", func(t *testing.T) {
		_, err := ls.AssignAvailableLicenseOfPackage(pkgId, accId, insId, insUsrIdDaniel)
		assert.Error(t, err, "no more unassigned license for pkgId=pkg:base-optimize-2022")
	})
}

func TestAssignSpecificLicense(t *testing.T) {

	var licRepo licensing.LicenseRepository = storage.NewLicenseRepoInMem()
	var pkgRepo licensing.PackageRepository = storage.NewPackageRepoInMem()
	ls := NewLicensingService(&licRepo, &pkgRepo)

	accId := "acc-1"
	subId := "sub-1"
	pkgId := "pkg:base-optimize-2022"
	insId := "ins-101"
	insUsrIdAlice := "usr-alice"
	insUsrIdBob := "usr-bob"

	_, err := ls.IssueLicenses(accId, subId, pkgId, 3)
	assert.NilError(t, err)

	t.Run("reassign should succeed", func(t *testing.T) {
		// assign a random license to alice
		lic, err := ls.AssignAvailableLicenseOfPackage(pkgId, accId, insId, insUsrIdAlice)
		t.Log(lic)
		assert.NilError(t, err)
		assert.Check(t, lic != nil)
		// reassign the same license to bob
		lic, err = ls.AssignSpecificLicense(lic.Id(), accId, insId, insUsrIdBob)
		t.Log(lic)
		assert.NilError(t, err)
		assert.Check(t, lic != nil)
		unassignedLicensesCount, err := ls.CountTotalUnassignedLicensesOfPackage(accId, pkgId)
		assert.NilError(t, err)
		assert.Equal(t, unassignedLicensesCount, 2)
	})
}

func TestVerifyEntitlement(t *testing.T) {

	var licRepo licensing.LicenseRepository = storage.NewLicenseRepoInMem()
	var pkgRepo licensing.PackageRepository = storage.NewPackageRepoInMem()
	ls := NewLicensingService(&licRepo, &pkgRepo)

	accId := "acc-1"
	subId := "sub-1"
	pkgId := "pkg:base-optimize-2022"
	cpbIdSeq := "sequence"
	cpbIdKaia := "kaia"
	insId := "ins-101"
	insUsrIdAlice := "usr-alice"
	insUsrIdBob := "usr-bob"

	_, err := ls.IssueLicenses(accId, subId, pkgId, 3)
	assert.NilError(t, err)
	_, err = ls.AssignAvailableLicenseOfPackage(pkgId, accId, insId, insUsrIdAlice)
	assert.NilError(t, err)

	t.Run("alice is entitled to sequence", func(t *testing.T) {
		entitlement, err := ls.VerifyEntitlement(accId, insId, insUsrIdAlice, cpbIdSeq)
		assert.NilError(t, err)
		assert.Equal(t, entitlement.IsEntitled, true)
	})

	t.Run("bob is not entitled to sequence", func(t *testing.T) {
		entitlement, err := ls.VerifyEntitlement(accId, insId, insUsrIdBob, cpbIdKaia)
		assert.NilError(t, err)
		assert.Equal(t, entitlement.IsEntitled, false)
	})

}
