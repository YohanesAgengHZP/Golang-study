package cassconnect

import (
	"log"
	"fmt"
	"github.com/gocql/gocql"
)

const (
	//Credential for Cassandra
	CassandraUsername     = "test"
	CassandraPassword     = "testcassandra"
	CassandraKeyspace     = "smr"
	CassandraContactPoint = "127.0.0.1:9042"
)

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