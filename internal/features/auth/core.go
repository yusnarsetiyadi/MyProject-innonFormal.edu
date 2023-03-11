package auth

import "time"

/*

	Struct Core For Data Transfer Object

*/

type UserCore struct {
	ID        uint
	Name      string
	Email     string
	Password  string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
