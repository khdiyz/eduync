package constants

type RoleType string

const (
	RoleSuperAdmin RoleType = "SUPER ADMIN"
	RoleAdmin      RoleType = "ADMIN"
	RoleMentor     RoleType = "MENTOR"
)

type EnrollmentStatus string

const (
	EnrollmentStatusActive   EnrollmentStatus = "ACTIVE"
	EnrollmentStatusInactive EnrollmentStatus = "INACTIVE"
	EnrollmentStatusFrozen   EnrollmentStatus = "FROZEN"
)
