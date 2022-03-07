package licensing

import "time"

// Reperesents a packaging plan that aggregates a set of supported packages
//
// Packaging Plan is an abstraction that enables proper granfathering
//
// A subscription always has subscribes to one particular packaging plan
type PackagingPlan struct {

	// e.g., "pkgplan:v2.3"
	Id string

	// e.g., v2
	MajorVersion int

	// Revision number within the major version e.g., "3" is the revision of v2.3
	Revision int

	// When this packaging plan was create
	CreatedAt time.Time

	// Packages supported in this packaging plan
	SupportedPackages []*Package
}
