package licensing

// Definition: The licensing right for a user to access a capability.
// DDD Classification: Value Object
type Entitlement struct {

	// True if the evaluated user id is entitled to the evaluated capability id
	IsEntitled bool

	// True if the evaluated user id is entitled to the feature part the evaluated capability
	// but not entitled to the capability due to exceeding capability
	IsEntitledToFeatureButExceedCapability bool

	// User id evalauted for entitlement
	EvaluatedUserId string

	// Capability id evaluated for entitlement
	EvaluatedCapabilityId string
}
