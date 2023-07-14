package entities

type User struct {
	ID      string `json:"_id" bson:"_id"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
	Deleted bool   `json:"deleted"`
}
