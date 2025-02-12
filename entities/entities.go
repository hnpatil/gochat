package entities

import (
	"github.com/google/uuid"
	"github.com/hnpatil/gochat/entities/message"
	"github.com/hnpatil/gochat/entities/roommember"
	"time"
)

type Entity struct {
	CreatedAt  time.Time  `json:"createdAt,omitempty" example:"2025-02-08T14:13:39.080551Z"`  //Time when entity was created
	ModifiedAt time.Time  `json:"modifiedAt,omitempty" example:"2025-02-08T14:13:39.080551Z"` //Time when entity was modified
	DeletedAt  *time.Time `json:"deletedAt,omitempty" example:"2025-02-08T14:13:39.080551Z"`  //Time when entity was deleted
}

type User struct {
	Entity
	ID   string `json:"id,omitempty" example:"89e46f30"`   //External identifier of the user
	Name string `json:"name,omitempty" example:"John Doe"` //Name of the user
}

type Message struct {
	Entity
	ID       uuid.UUID      `json:"id,omitempty"`
	RoomID   string         `json:"roomID,omitempty"`
	SenderID string         `json:"senderID,omitempty"`
	SentAt   *time.Time     `json:"sentAt,omitempty"`
	Status   message.Status `json:"status,omitempty"`
	Content  string         `json:"content,omitempty"`
}

type Room struct {
	Entity
	ID      string        `json:"id,omitempty" example:"89e47f30"`  //Unique identifier of the room
	Name    string        `json:"name,omitempty" example:"Friends"` //Name of the room
	IsGroup bool          `json:"isGroup,omitempty" example:"true"` //Indicates if the room is a group. Is false for chats.
	Members []*RoomMember `json:"members,omitempty,omitempty"`      //List of room members
}

type RoomMember struct {
	Entity
	RoomID string          `json:"roomID,omitempty" example:"89e47f30"`                 //Unique identifier of the room
	UserID string          `json:"userID,omitempty" example:"89e46f30"`                 //External identifier of the user
	Role   roommember.Role `json:"role,omitempty" enums:"ADMIN,MEMBER" example:"ADMIN"` //Role defines permissions of user on the room.
}
