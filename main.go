package main

import (
	"log"
	"os"
	"simple-wal/protobuf"
	"time"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func WriteMsginProto(filename string, message proto.Message) error {
	data, err := proto.Marshal(message)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}
func main() {
	person := &protobuf.Person{
		Name:  "Arafat",
		Id:    1234,
		Email: "arafat@gmail.com",
		Phones: []*protobuf.Person_PhoneNumber{
			{
				Number: "555-4321",
				Type:   protobuf.PhoneType_PHONE_TYPE_HOME,
			},
		},
		LastUpdated: timestamppb.New(time.Now()),
	}
	book := &protobuf.AddressBook{
		People: []*protobuf.Person{person},
	}

	err := WriteMsginProto("person.bin", book)

	if err != nil {
		panic(err)
	}

	log.Println("Person Write into person.bin")

}
