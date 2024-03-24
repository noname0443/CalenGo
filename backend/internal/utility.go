package internal

const (
	readNoteByName   string = `SELECT name, description, start_date as start_time, end_date as end_time FROM calen_go.notes WHERE name = ?;`
	createNoteByName string = `INSERT INTO calen_go.notes (name, description, start_date, end_date, author_id) VALUES (?, ?, ?, ?, ?);`
	readNotesByTime  string = `SELECT name, description, start_date as start_time, end_date as end_time FROM calen_go.notes WHERE start_date >= CAST(? AS DATE) AND end_date <= CAST(? AS DATE) AND author_id = ?;`
	createUser       string = `INSERT INTO calen_go.users (name, password) VALUES (:name, :password);`
	readUser         string = `SELECT id, name, password FROM calen_go.users WHERE name = ? AND password = ?;`
)
