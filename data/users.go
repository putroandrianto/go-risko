package data

type User struct {
	Id		int 	`json:"id"`
	Name	string 	`json:"name"`
	Age		int 	`json:"age"`
	RiskProfile *RiskProfile `gorm:"foreignkey:UserId" json:"risk_profile,omitempty"` 
}