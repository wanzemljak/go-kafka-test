package main

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	conn, _ := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "kraljina", 0)
	conn.SetDeadline(time.Time{})
	for {
		var message string
		fmt.Println("Enter message:")
		fmt.Scan(&message)
		conn.WriteMessages(kafka.Message{Value: []byte(message)})
	}
}
