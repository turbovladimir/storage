package request

type Find struct {
	Table  string                 `json:"table" binding:"required"`
	Filter map[string]interface{} `json:"filter" binding:"required"`
}

type Update struct {
	Create
}

type Create struct {
	Rows interface{} `json:"rows" binding:"required"`
}

type CheckPhone struct {
	Phone string `json:"phone" binding:"required"`
}
