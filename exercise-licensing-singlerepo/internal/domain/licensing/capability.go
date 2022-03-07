package licensing

// Definition: A unit of software functionality at which user’s entitlement is evaluated.
// DDD Classification: Entity
type Capability struct {

	// ID of the capability
	Id string

	// Customer-facing display name of the capability
	DisplayName string

	// Whether this capability has capacity limit (upper bound)
	HasCapacityLimit bool

	// Capacity limit
	CapacityLimit int

	// Unit of the capacity limit
	CapacityLimitUnit string
}
