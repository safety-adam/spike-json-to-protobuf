package main

import (
	"fmt"
	"log"
	"spike-json-to-protobuf/proto"

	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	jsonExample := []byte(`{"response": [{"question":"Example question","response_type":"Good,Fair,Poor,N/A"}, {"question":"Example question 2","response_type":"Yes,No"}]}`)

	var resp proto.Response

	if err := protojson.Unmarshal(jsonExample, &resp); err != nil {
		log.Fatalf("failed to unmarshal JSON: %v", err)
	}

	for _, r := range resp.GetResponse() {
		fmt.Printf("Question: %s\n", r.GetFields()["question"].GetStringValue())
		fmt.Printf("Response Type: %s\n", r.GetFields()["response_type"].GetStringValue())
		fmt.Println()
	}
}
