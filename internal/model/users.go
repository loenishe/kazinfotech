package model

type User struct {
	Username string `db:"user_id"`
	Password string `db:"password"`
	UserRole string `db:"user_role"`
	ParentID string `db:"parent_id"`
}

type Response struct {
	ResponseCode    int
	ResponseMessage string
	Data            interface{}
}
