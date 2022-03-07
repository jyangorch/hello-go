package licensing

// repository interface for package
type PackageRepository interface {

	// Get package by id
	GetPackageById(pkgId string) (*Package, error)
}
