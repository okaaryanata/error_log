package errorlog

import (
	"database/sql"
	"fmt"
	"log"
	"runtime"
	"strings"
)

// Elog stand for error log
// hold connection from database and environment
type Elog struct {
	connection  *sql.DB
	environment string
	repo        string
}

// ConnectLog try establish connection to database log
// the environment option are development, production ( default development )
func ConnectLog(dbSource, env, repo string) Elog {
	db, err := sql.Open("mysql", dbSource)
	if err != nil {
		log.Print(err.Error())
	}

	if env != "" {
		return Elog{db, env, repo}
	}

	return Elog{db, "development", repo}
}

//Close connection to database
func (l Elog) Close() {
	l.connection.Close()
}

func (l Elog) Error(err error, payload string) {
	_, fn, line, _ := runtime.Caller(1)
	log.Printf("[ERROR] %s [%s:%d]\nwith payload: %s", err, fn, line, payload)

	if l.isProd() {
		emessage := fmt.Sprintf("[ERROR] %s [%s:%d]", err, fn, line)

		_, err := l.connection.Exec(l.insertStmnt(emessage, payload))
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func (l Elog) isProd() bool {
	return l.environment == "production"
}

func (l Elog) insertStmnt(message, payload string) string {
	payload = strings.ReplaceAll(payload, "'", "\\'")
	message = strings.ReplaceAll(message, "'", "\\'")
	tmplt := `INSERT INTO error_logs (repository, error, payload)
		VALUES ('%s', '%s', '%s')`
	return fmt.Sprintf(tmplt, l.repo, message, payload)
}
