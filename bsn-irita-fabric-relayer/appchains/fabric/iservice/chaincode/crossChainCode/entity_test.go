package crossChainCode

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJson(t *testing.T) {

	ser := &serviceRequest{
		ServiceName: "test",
		Input:       `{"type":"object","properties":{"id":{"type":"string"}}}`,
		Timeout:     100,
	}

	sb, _ := json.Marshal(ser)

	fmt.Println(string(sb))

	var sts []string

	sts = append(sts, string(sb))

	sb1, _ := json.Marshal(sts)

	fmt.Println(string(sb1))

}
