package main

import (
	trippb "coolcar/server/proto/gen/go"
	"fmt"
	"log"
	"net"

	"google.golang.org/protobuf/proto"
)

func main() {
	lis,err:= net.Listen("tcp",":8081")
	if err!=nil {
		log.Fatalf("failed to listen:%v",err)
	}
	
}

func other() {
	trip := trippb.Trip{
		Start:       "aba",
		End:         "def",
		DurationSec: 3600,
		FeeCent:     10000,
		StartPos: &trippb.Location{
			Latitude:  30,
			Longitude: 120,
		},
		EndPos: &trippb.Location{
			Latitude:  35,
			Longitude: 115,
		},
		PathLocations: []*trippb.Location{
			{
				Latitude:  31,
				Longitude: 119,
			},
			{
				Latitude:  32,
				Longitude: 118,
			},
		},
	}
	fmt.Println(&trip)
	s, err := proto.Marshal(&trip)
	if err != nil {
		panic(err)
	}
	fmt.Printf("protobuf:%X\n", s)

	var trip2 trippb.Trip
	err = proto.Unmarshal(s, &trip2)
	if err != nil {
		panic(err)
	}
	fmt.Println(&trip2)_
}
