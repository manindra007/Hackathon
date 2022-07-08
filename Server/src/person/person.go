package person

type EmergencyContact struct {
	Phone string `json:"phone"`
	Email string `json:"email"`
}

type Person struct {
	Id               string             `gorm:"primary_key" json:"id"`
	Email            string             `gorm:"unique" json:"email"`
	Name             string             `json:"name"`
	Address          string             `json:"address"`
	Phone            string             `json:"phone"`
	EmergencyContact []EmergencyContact `json:"ContactDetail"`
}
type PersonContact struct {
	IdEmail string `gorm:"unique" json:"idemail"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
}

var Users []Person
