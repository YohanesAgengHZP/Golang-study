package queryHelper

import (
	"context"
	"fmt"
	// cass "grpc-server-test/connection"
	pb "grpc-server-test/protobuf"
	"log"
	"github.com/gocql/gocql"
)

var CassandraSession *gocql.Session

func CassandraQueryNumber(ctx context.Context, in *pb.DataRequest) error {
	// Ensure that Cassandra session is initialized
	if CassandraSession == nil {
		return fmt.Errorf("Cassandra session is not initialized")
	}

	// Iterate over each Number in the DataRequest
	for _, num := range in.GetNumber() {
		// Update each Number in Cassandra
		if err := UpdateNumberInCassandra(ctx, num); err != nil {
			return fmt.Errorf("failed to update number in Cassandra: %v", err)
		}
	}

	log.Println("Successfully updated numbers in Cassandra")
	return nil
}

// UpdateNumberInCassandra updates the target_a2p, target_p2p, and number_status columns in the Cassandra table for a given number.
func UpdateNumberInCassandra(ctx context.Context, num *pb.Number) error {
	// Prepare the update query
	query := "UPDATE number_routing_table SET target_a2p = ?, target_p2p = ?, number_status = ? WHERE number = ?"

	// Execute the query with context
	if err := CassandraSession.Query(
		query, num.GetTargetA2P(),
		num.GetTargetP2P(),
		num.GetStatus(),
		num.GetNumber()).WithContext(ctx).Exec();
		err != nil {
		return fmt.Errorf("failed to execute Cassandra query: %v", err)
	}

	return nil
}