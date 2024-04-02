package Connection

import (
	"database/sql"
	"fmt"
	"log"
    "github.com/gocql/gocql"
	_ "github.com/go-sql-driver/mysql"
)

const (
    // Credential for MySQL
	MysqlUsername = "test_user"
	password = "test_password"
	hostName = "localhost:3306"
    dbName = "test_db"

    //Credential for Cassandra
    CassandraUsername = "test"
    CassandraPassword = "testcassandra"
    CassandraKeyspace = "smr"
    CassandraContactPoint = "127.0.0.1:9042"
)

//===================================== MYSQL CONNECTION ================================================
var (
	db *sql.DB
)

func ConnectionMySql() {
    err := InitMySQL()
	if err != nil {
		log.Fatal(err)
	}
    fmt.Printf("Success MySQL")
	defer Close()
}

func InitMySQL() error {
	var err error
	db, err = sql.Open("mysql", credentialMySQL())
	if err != nil {
		return fmt.Errorf("error opening DB: %s", err)
	}

	// Verify the connection
	err = db.Ping()
	if err != nil {
		return fmt.Errorf("error connecting to DB: %s", err)
	}

	return nil
}

func credentialMySQL() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", MysqlUsername, password, hostName, dbName)
}

func Close() {
	if db != nil {
		db.Close()
	}
}

//===================================== Cassandra CONNECTION ================================================

var CassandraSession *gocql.Session

func ConnectionCassandra() {
	cluster := gocql.NewCluster(CassandraContactPoint)
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: CassandraUsername,
		Password: CassandraPassword,
	}
	cluster.Keyspace = CassandraKeyspace
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatalf("error connecting to Cassandra: %s", err)
	}
    fmt.Printf("Success Cassandra")
	CassandraSession = session
}

func CloseCassandra() {
	if CassandraSession != nil {
		CassandraSession.Close()
	}
}