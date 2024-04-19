package helperNumber

import (
	"strconv"
	"fmt"
	pb "sync-databases-cass-mysql/protobuf"
)

func ConvertToProtoNumbers(numbers []map[string]interface{}) ([]*pb.Number, error) {
	var protoNumbers []*pb.Number

	for _, num := range numbers {
		targetA2P, _ := strconv.ParseUint(fmt.Sprintf("%v", num["target_a2p"]), 10, 32)
		targetP2P, _ := strconv.ParseUint(fmt.Sprintf("%v", num["target_p2p"]), 10, 32)
		status, _ := strconv.ParseUint(fmt.Sprintf("%v", num["status"]), 10, 32)

		pbNumber := &pb.Number{
			Number:    num["number"].(string),
			TargetA2P: uint32(targetA2P),
			TargetP2P: uint32(targetP2P),
			Status:    uint32(status),
		}
		protoNumbers = append(protoNumbers, pbNumber)
	}

	return protoNumbers, nil
}