package licensing

// A bundle of capabilities as a licensable unit to customers.
// DDD classification: Aggregate
type Package struct {

	// Id of the package
	Id string

	// Readable name of the package, e.g., Accelerate, Optimize
	// The package with Ochestrate name in 2021 can be actually different from that same Ochestrate name in 2022
	Name string

	// TODO: add link to its owning package plan
	// owningPackagePlanId string

	// Capability included in this package
	IncludedCapabilities []Capability
}

// Checks whether this package includes the given capability id
func (p *Package) IncludesCapability(cpbId string) bool {
	for _, includedCpb := range p.IncludedCapabilities {
		if includedCpb.Id == cpbId {
			return true
		}
	}
	return false
}
