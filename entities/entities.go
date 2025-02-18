package entities

import (
	"github.com/hnpatil/gochat/entities/roommember"
	"github.com/hnpatil/gochat/pkg/id"
	"github.com/hnpatil/gochat/pkg/metadata"
	"time"
)

type Entity struct {
	CreatedAt  time.Time `json:"createdAt,omitempty" example:"2025-02-08T14:13:39.080551Z"`  //Time when entity was created
	ModifiedAt time.Time `json:"modifiedAt,omitempty" example:"2025-02-08T14:13:39.080551Z"` //Time when entity was modified
}

type User struct {
	Entity
	ID       string            `json:"id,omitempty" example:"89e46f30"` //External identifier of the user
	Metadata metadata.Metadata `json:"metadata,omitempty"`              //Metadata associated with the user
}

type Message struct {
	Entity
	ID       id.ID   `json:"id,omitempty" example:"89e48f30"`       //Unique identifier of the message
	RoomID   string  `json:"roomID,omitempty" example:"89e47f30"`   //Unique identifier of the room that message is added to
	SenderID *string `json:"senderID,omitempty" example:"89e46f30"` //Unique identifier of the user that created the message
	Content  string  `json:"content,omitempty" example:"Hello"`     //Message content
}

type Room struct {
	Entity
	ID       string            `json:"id,omitempty" example:"89e47f30"` //Unique identifier of the room
	Members  []*RoomMember     `json:"members,omitempty,omitempty"`     //List of room members
	Metadata metadata.Metadata `json:"metadata,omitempty"`              //Room meta data
}

type RoomMember struct {
	Entity
	RoomID string          `json:"roomID,omitempty" example:"89e47f30"` //Unique identifier of the room
	UserID string          `json:"userID,omitempty" example:"89e46f30"` //External identifier of the user
	Role   roommember.Role `json:"role,omitempty"  example:"ADMIN"`     //Role defines permissions of user on the room.
	User   *User           `json:"user,omitempty"`                      //User object associated with member
}
