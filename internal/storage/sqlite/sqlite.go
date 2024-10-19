package sqlite

import ("fmt" 
"database/sql"
_ "github.com/mattn/go-sqlite3") 

type Storage struct {

	database *sql.DB
}

func New(storagePath string) (*Storage,error){ 
const op = "storage.sqlite.New"

database, err := sql.Open("sqlite3",storagePath)
if err != nil{
return nil, fmt.Errorf("%s: %w", op, err)
}

statement, err := database.Prepare("CREATE TABLE IF NOT EXISTS url (
id INTEGER PRIMARY KEY,
 alias TEXT NOT NULL UNIQUE,
  url TEXT NOT NULL);
   CREATE INDEX IF NOT EXISTS idx_alias ON url(alias);
   ")

if err != nil {
	return nil, fmt.Errorf("#{op}: #{err}")
}

_, err = statement.Exec()
if err != nil {
	return nil, fmt.Errorf("%s: %w", op, err)
}

return &Storage{database: database}, nil
}