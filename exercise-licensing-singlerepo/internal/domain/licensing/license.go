package licensing

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

// License represents a permit that grants a user the permission (aka entitlement) to use software features.
//
// See [TODO: insert link] for more information about License.
//
// Some examples to help understand how licenses will be represented at Outreach:
//
// Example scenario 1: a customer purchases 100 user licenses of the Optimize package.
//
// Example scenario 2: a customer purcahses 100 user licenses of the Optimize package, and separately 70
// user licenses of the Kaia package as Add-On to the Optimize package.
//
// Example scenario 3: a customer purchases 100 user licenses of the Optimize package, and separately 40
// user licenses of the Ochestrate package.
//
// DDD Classification: Aggregate
type License struct {

	// License ID
	id string

	// Package being licensed by this license.
	licensedPackage *Package

	// The customer account to which this license is issued.
	possessingCustomerAccountId string

	// License's lifecycle is managed through syncing with its governing subscription.
	governingSubscriptionId string

	// Expiration detail of this license. Nil if this license is not issued.
	issuanceDetail *LicenseIssuanceDetail

	// Expiration detail of this license. Nil if this license is not expired.
	expirationDetail *LicenseExpirationDetail

	// Cancellation detail of this license. Nil if this license is not cancelled.
	cancellationDetail *LicenseCancellationDetail

	// Renewal detail of this license. Nil if this license has not being renewed.
	renewalDetail *LicenseRenewalDetail

	// Current license assignment
	currentAssignment *LicenseAssignment

	// Previous license assignment
	previousAssignments []*LicenseAssignment

	// True if this license is a trial license
	isTrial bool
}

type LicenseIssuanceDetail struct {

	// Time when this licenses was issued. Nil if not expired.
	IssuedAt time.Time

	// Issuance reason (e.g., due to new logo, expansion, is a renewal from another)
	IssuanceReason string
}

type LicenseExpirationDetail struct {

	// Time when this licenses was expired. Nil if not expired.
	ExpiredAt time.Time
}

type LicenseCancellationDetail struct {

	// The time this license was cancelled. Nil if not cancelled.
	CancelledAt time.Time
}

type LicenseRenewalDetail struct {

	// License ID to which this license renews to
	RenewedToLicenseId string

	// Time when this license is renewed
	RenewedAt time.Time

	// Renewal reason
	RenewalReason string
}

func NewIssuedLicense(accId string, subId string, pkg *Package) *License {
	lic := &License{
		id:                          uuid.NewString(),
		possessingCustomerAccountId: accId,
		licensedPackage:             pkg,
		governingSubscriptionId:     subId,
		issuanceDetail:              &LicenseIssuanceDetail{IssuedAt: time.Now(), IssuanceReason: "New Logo (FIXME)"},
	}
	return lic
}

func (lic *License) String() string {
	return fmt.Sprintf(
		`{id=%s, possessingCustomerAccountId=%s, governingSubscriptionId=%s, licensedPackageId=%s, currentAssignment=%+v, previousAssignmentCount=%d}`,
		lic.id,
		lic.possessingCustomerAccountId,
		lic.governingSubscriptionId,
		lic.licensedPackage.Id,
		lic.currentAssignment,
		len(lic.previousAssignments))
}

func (lic *License) Id() string {
	return lic.id
}

func (lic *License) LicensedPackage() *Package {
	return lic.licensedPackage
}

func (lic *License) PossessingCustomerAccountId() string {
	return lic.possessingCustomerAccountId
}

func (lic *License) GoverningSubscriptionId() string {
	return lic.governingSubscriptionId
}

func (lic *License) IsActive() bool {
	return lic.cancellationDetail == nil && lic.expirationDetail == nil && lic.renewalDetail == nil
}

func (lic *License) AssignedToLicensee() Licensee {
	return lic.currentAssignment.Assignee
}

// Whether this license is currently assigned to any licensee
func (lic *License) IsAssigned() bool {
	return lic.currentAssignment != nil
}

func (lic *License) Assign(licensee Licensee) {
	if lic.currentAssignment != nil {
		lic.previousAssignments = append(lic.previousAssignments, lic.currentAssignment)
	}
	lic.currentAssignment = NewCurrentLicenseAssignment(licensee)
}

func (lic *License) Unassign() {
	if lic.currentAssignment != nil {
		lic.previousAssignments = append(lic.previousAssignments, lic.currentAssignment)
	}
	lic.currentAssignment = nil
}

func (lic *License) IsTrial() bool {
	return lic.isTrial
}
