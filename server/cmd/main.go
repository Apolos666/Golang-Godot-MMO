package main

import (
	"fmt"
	"server/pkg/packets"

	"google.golang.org/protobuf/proto"
)

func main() {
	packet := &packets.Packet{
		SenderId: 69,
		Msg:      packets.NewChat("Hello, World!"),
	}

	data, err := proto.Marshal(packet)

	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
