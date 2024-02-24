package repository

type CRUDDTO interface {
	GetError() error
}

type Error error

type FindParams struct {
	TableName string
	Filter
	Error
}

type Filter map[string]interface{}

type CreateParams struct {
	Models interface{}
	Error
}

type UpdateParams struct {
	Error
}
