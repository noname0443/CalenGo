package internal

const (
	readNoteByName      string = `SELECT name, description, start_date as start_time, end_date as end_time FROM calen_go.notes WHERE name = ?;`
	createNoteByName    string = `INSERT INTO calen_go.notes (name, description, start_date, end_date, author_id) VALUES (?, ?, ?, ?, ?);`
	readNotesByTime     string = `SELECT name, description, start_date as start_time, end_date as end_time FROM calen_go.notes WHERE start_date >= CAST(? AS DATE) AND end_date <= CAST(? AS DATE) AND author_id = ?;`
	createUser          string = `INSERT INTO calen_go.users (name, password) VALUES (:name, :password);`
	readUser            string = `SELECT id, name, password FROM calen_go.users WHERE name = ? AND password = ?;`
	readConflictedNotes string = `SELECT T1.name, T1.description, T1.start_date, T1.end_date, T2.name, T2.description, T2.start_date, T2.end_date FROM calen_go.notes T1 JOIN calen_go.notes T2 ON T1.id != T2.id AND (T1.start_date >= T2.start_date AND T1.start_date <= T2.end_date OR T1.end_date >= T2.start_date AND T1.end_date <= T2.end_date) WHERE T1.author_id = T2.author_id AND T1.author_id = ? ORDER BY T1.name, T2.name;`
)
