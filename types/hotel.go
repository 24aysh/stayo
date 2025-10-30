package types

import "go.mongodb.org/mongo-driver/v2/bson"

type Hotel struct {
	ID       bson.ObjectID   `bson:"_id,omitempty" json:"id,omitempty"` // these are called struct tags, it will omit the ID from json if found empty
	Name     string          `bson:"name" json:"name"`
	Location string          `bson:"location" json:"location"`
	Rooms    []bson.ObjectID `bson:"rooms" json:"rooms"`
	Rating   int             `bson:"rating" json:"rating"`
}

type RoomType int

const (
	_ RoomType = iota
	SinglePersonRoomType
	DoubleRoomType
	SeaSideRoomType
	DeluxeRoomType
)

type Room struct {
	ID      bson.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Type    RoomType      `bson:"type" json:"type"`
	Size    string        `bson:"size" json:"size"`
	Seaside bool          `bson:"seaside" json:"seaside"`
	Price   float64       `bson:"price" json:"price"`
	HotelID bson.ObjectID `bson:"hotelId" json:"hotelId"`
}
