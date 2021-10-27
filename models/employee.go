package models


type DomainEmployee struct {
	FullName		string	`json:"first_name" structs:"first_name" db:"first_name"`
	FatherName		string	`json:"father_name" structs:"father_name" db:"father_name"`
	Age				int		`json:"age" structs:"age" db:"age"`
	Email			string	`json:"email" structs:"email" db:"email"`
	Password		string	`json:"password" structs:"password" db:"password"`
	NIC				string  `json:"nic" structs:"nic" db:"nic"`
	Gender			string	`json:"gender" structs:"gender" db:"gender"`

}