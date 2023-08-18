package common

const (
	RoleID = iota
	RoleIDAdmin
	RoleIDStudent
	RoleIDTeacher

	RoleTeacher = "teacher"
	RoleStudent = "student"
	RoleAdmin   = "admin"
)

const (
	GenderMale = iota
	GenderFemale
)

const (
	StatusBanned = iota + 1
	StatusNormal
)
