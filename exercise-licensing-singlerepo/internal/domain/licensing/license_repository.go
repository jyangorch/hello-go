package licensing

// Definition: Repository for License.
// DDD Classification: Repository
type LicenseRepository interface {

	// Create license
	CreateLicense(lic *License) error

	// Update license
	UpdateLicense(licId string, newLic *License) error

	// Get license by id
	GetLicenseById(licId string) (*License, error)

	// Find licenses by licensee id
	FindLicensesByAssignedLicenseeId(licenseeId string) ([]*License, error)

	// Find next unassigned license of the given package id under the customer account id
	FindNextUnassignedLicenseOfPackage(accId string, pkgId string) (*License, error)

	// Count the total unassigned license of the given package id under the customer account id
	CountTotalUnassignedLicensesOfPackage(accId string, pkgId string) (int, error)
}
