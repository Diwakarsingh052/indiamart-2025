package db

var conn string

func NewConnection(db string) {
	conn = db
}

func GetConnection() string {
	return conn
}
