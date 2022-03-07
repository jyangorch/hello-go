package licensing

import (
	"time"
)

// Represents an license assignment fact
//
// DDD Classification: Value Object
type LicenseAssignment struct {

	// Assigned licensee
	Assignee Licensee

	// When the licensee was assigned
	AssignedAt time.Time

	// When the license was unassigned; zero value if not yet unassigned
	UnassignedAt time.Time
}

func NewCurrentLicenseAssignment(licensee Licensee) *LicenseAssignment {
	return &LicenseAssignment{
		Assignee:     licensee,
		AssignedAt:   time.Now(),
		UnassignedAt: time.Time{}, // zero value
	}
}
