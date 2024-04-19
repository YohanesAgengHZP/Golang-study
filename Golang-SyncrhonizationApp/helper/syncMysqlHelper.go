package helperSync

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"sync"
	Connection "sync-databases-cass-mysql/connection"
	pb "sync-databases-cass-mysql/protobuf"
	ph "sync-databases-cass-mysql/stubs"
	_ "github.com/go-sql-driver/mysql"
)

const BATCH_SIZE = 10000 //Batching per 10000

func FetchMySQLData() ([]*pb.Number, []string, error) {
	
	db, err := Connection.ConnectionMySql()
	if err != nil {
		return nil, nil, fmt.Errorf("Error connecting to MySQL: %v", err)
	}
	defer db.Close()

	// Fetch data Number
	numbers, err := GetNumber(db)
	if err != nil {
		return nil, nil, fmt.Errorf("Error fetching numbers from MySQL: %v", err)
	}

	// Convert number to protobuf
	protoNumbers, err := ph.ConvertToProtoNumbers(numbers)
	if err != nil {
		return nil, nil, fmt.Errorf("Error converting numbers to protobuf format: %v", err)
	}

	// Getch Data Avoid Brands
	avoidBrands, err := GetAvoid(db)
	if err != nil {
		return nil, nil, fmt.Errorf("Error fetching avoid brands from MySQL: %v", err)
	}

	return protoNumbers, avoidBrands, nil
}

func SyncDataAndAvoidParallel(ctx context.Context, client pb.SmrSyncClient) error {
	
	numbers, avoidBrands, err := FetchMySQLData()
	if err != nil {
		return fmt.Errorf("Error fetching data from MySQL: %v", err)
	}

	pbAvoidBrands := make([]*pb.Avoid, len(avoidBrands))
	for i, brand := range avoidBrands {
		pbAvoidBrands[i] = &pb.Avoid{Name: brand}
	}

	// Prepare gRPC request message
	request := &pb.DataRequest{
		Number: numbers,
		Avoid:  pbAvoidBrands, 
	}

	var wg sync.WaitGroup

	for i := 0; i < len(request.Number); i += BATCH_SIZE {
		end := i + BATCH_SIZE
		if end > len(request.Number) {
			end = len(request.Number)
		}

		wg.Add(1) 

		go func(start, end int) {
			defer wg.Done() 

			batchRequest := &pb.DataRequest{
				Number: request.Number[start:end],
				Avoid:  request.Avoid,
			}

			// Call gRPC method to synchronize data
			err := syncData(ctx, client, batchRequest)
			if err != nil {
				log.Printf("Error synchronizing data: %v", err)
			}
		}(i, end)
	}

	//channel wait for goroutines to cancel or completed
	wg.Wait()

	return nil
}

func syncData(ctx context.Context, client pb.SmrSyncClient, request *pb.DataRequest) error {
	// Call SyncData gRPC method
	dataResponse, err := client.SyncData(ctx, request)
	if err != nil {
		return fmt.Errorf("SyncData failed: %v", err)
	}
	fmt.Println("SyncData response:", dataResponse)
	return nil
}

func GetNumber(db *sql.DB) ([]map[string]interface{}, error) {
	
	query := `SELECT
                tmr.msisdn AS number,
                COALESCE(tgr.target_id, 0) AS target_a2p,
                COALESCE(tgr2.target_id, 0) AS target_p2p,
                tmr.status
              FROM
                t_msisdn_routing tmr
                LEFT JOIN t_groups_routes tgr ON tgr.id = tmr.sms_a2p
                LEFT JOIN t_groups_routes tgr2 ON tgr2.id = tmr.sms_p2p`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]map[string]interface{}, 0)

	for rows.Next() {
		var number string
		var targetA2P, targetP2P int
		var status sql.NullString
		if err := rows.Scan(&number, &targetA2P, &targetP2P, &status); err != nil {
			return nil, err
		}
		row := map[string]interface{}{
			"number":     number,
			"target_a2p": targetA2P,
			"target_p2p": targetP2P,
			"status":     status,
		}
		result = append(result, row)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func GetAvoid(db *sql.DB) ([]string, error) {
	query := "SELECT name FROM t_avoid_brands"

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var avoidBrands []string

	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		avoidBrands = append(avoidBrands, name)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return avoidBrands, nil
}