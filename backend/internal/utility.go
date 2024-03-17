package internal

const (
	readNoteByName   string = `SELECT name, description, start_date as start_time, end_date as end_time FROM calen_go.notes WHERE name = ?;`
	createNoteByName string = `INSERT INTO calen_go.notes (name, description, start_date, end_date) VALUES (:name, :description, CAST(:start_time AS DATE), CAST(:end_time AS DATE));`
)
