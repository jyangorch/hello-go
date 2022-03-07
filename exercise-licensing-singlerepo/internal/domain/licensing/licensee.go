package licensing

import "fmt"

//
// License type "enum"
//
type LicenseeType int

const (
	INSTANCE_USER LicenseeType = iota
	ORGANIZATION_USER
	GROUP
)

func (lt LicenseeType) String() string {
	return [...]string{"INSTANCE_USER", "ORGANIZATION_USER", "GROUP"}[lt]
}

// Generic abstraction of the holder of a license
// The holder can be a "user", or can be a "group.
// A "user" can also be "instance user", a user defined within the scope (namespace) within instance
// A "user" can also be "organization user", a user defined within the scope within an customer organization and can cross instances
// Today we have to live with "instance user", but we want to move towards "organization user" as we want a user assigned a license
// can be granted the same access across instances (production or sandbox) provisioned to the organization
type Licensee interface {

	// Global unique identifier of the license
	LicenseeId() string

	// Type of licensee
	LicenseeType() LicenseeType

	// Represents a user (either a Instance User or a Organization User) but not a group
	IsUserIdentity() bool
}

//
// Instance User, representing a user within the scope of an instance
//
type InstanceUser struct {
	InstanceScopeUserId string
	Type                LicenseeType
	InstanceId          string
}

func NewInstanceUser(insId string, instanceScopeUserId string) InstanceUser {
	return InstanceUser{
		InstanceId:          insId,
		InstanceScopeUserId: instanceScopeUserId,
		Type:                INSTANCE_USER,
	}
}

func (is InstanceUser) LicenseeId() string {
	return is.encodeLicenseeId()
}

func (is InstanceUser) LicenseeType() LicenseeType {
	return INSTANCE_USER
}

func (is InstanceUser) IsUserIdentity() bool {
	return true
}

func (is InstanceUser) encodeLicenseeId() string {
	return fmt.Sprintf("%s:%s/%s", is.LicenseeType(), is.InstanceId, is.InstanceScopeUserId)
}

//
// Organization User, representing a user within a customer organization.
// The same Organization User can be presented in multiple instances for that organization.
//
type OrganizationUser struct {
	OrganizationScopeUserId string
	Type                    LicenseeType
	OrganizationId          string
	EmailAddress            string
}

func NewOrganizationUser(organizationId string, organizationScopeUserId string, emailAddress string) OrganizationUser {
	return OrganizationUser{
		OrganizationId:          organizationId, // equivalent to the customer account id?
		OrganizationScopeUserId: organizationScopeUserId,
		Type:                    ORGANIZATION_USER,
		EmailAddress:            emailAddress,
	}
}

func (os OrganizationUser) LicenseeId() string {
	return os.encodeLicenseeId()
}

func (os OrganizationUser) LicenseeType() LicenseeType {
	return ORGANIZATION_USER
}

func (os OrganizationUser) IsUserIdentity() bool {
	return true
}

func (os OrganizationUser) encodeLicenseeId() string {
	return fmt.Sprintf("%s:%s", os.LicenseeType(), os.OrganizationScopeUserId)
}
