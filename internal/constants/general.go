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

type StudentAction string

const (
	ActionJoined   StudentAction = "JOINED"
	ActionLeft     StudentAction = "LEFT"
	ActionFreeze   StudentAction = "FROZE"
	ActionUnfreeze StudentAction = "UNFROZE"
	ActionPaid     StudentAction = "PAID"
	ActionUnpaid   StudentAction = "UNPAID"
)
