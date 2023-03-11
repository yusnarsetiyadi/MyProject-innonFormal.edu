package superadmin

import "time"

/*

	Struct Core For Data Transfer Object

*/

type SuperAdminCore struct {
	ID        uint
	Name      string
	Email     string
	Password  string
	Role      string
	Image     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
