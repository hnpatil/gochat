package room

import (
	"github.com/google/uuid"
	"github.com/hnpatil/gochat/entities"
	"github.com/hnpatil/gochat/entities/roommember"
	"github.com/hnpatil/gochat/entities/user"
	"github.com/hnpatil/gochat/errors"
	"github.com/hnpatil/gochat/pkg/metadata"
	"github.com/hnpatil/gochat/repos"
	"github.com/hnpatil/gochat/services"
	"gofr.dev/pkg/gofr"
	"slices"
)

type svc struct {
	repo  repos.Room
	users repos.User
}

func New(repo repos.Room, users repos.User) services.Room {
	return &svc{repo: repo, users: users}
}

func (s *svc) Create(ctx *gofr.Context, req *services.CreateRoom) (*entities.Room, error) {
	room := &entities.Room{
		ID:       req.RoomID,
		Metadata: req.Metadata,
	}

	if room.ID == "" {
		room.ID = uuid.New().String()
	}

	memberMap := make(map[string]roommember.Role, len(req.Members))
	for _, m := range req.Members {
		memberMap[m] = roommember.RoleMember
	}

	memberMap[req.UserID] = roommember.RoleAdmin

	userIDs := make([]string, 0, len(memberMap))

	for k, _ := range memberMap {
		userIDs = append(userIDs, k)
	}

	usrs, err := s.users.List(ctx, &repos.UserFilter{UserID: userIDs})
	if err != nil {
		return nil, err
	}

	if len(usrs) != len(userIDs) {
		return nil, errors.EntityNotFound(user.Entity)
	}

	if len(memberMap) < 3 {
		for k, _ := range memberMap {
			memberMap[k] = roommember.RoleAdmin
		}
	}

	members := make([]*entities.RoomMember, 0, len(memberMap))
	for k, v := range memberMap {
		members = append(members, &entities.RoomMember{RoomID: room.ID, UserID: k, Role: v})
	}

	room.Members = members

	return s.repo.Create(ctx, room)
}

func (s *svc) Update(ctx *gofr.Context, req *services.UpdateRoom) (*entities.Room, error) {
	room, err := s.ValidateRole(ctx, req.RoomID, req.UserID, roommember.RoleAdmin)
	if err != nil {
		return nil, err
	}

	updatedMeta, err := metadata.ApplyPatch(room.Metadata, req.Metadata)
	if err != nil {
		return nil, err
	}

	return s.repo.Update(ctx, &entities.Room{ID: req.RoomID}, &entities.Room{Metadata: updatedMeta})
}

func (s *svc) Get(ctx *gofr.Context, req *services.GetRoom) (*entities.Room, error) {
	room, err := s.ValidateRole(ctx, req.RoomID, req.UserID, roommember.RoleAdmin, roommember.RoleMember)
	if err != nil {
		return nil, err
	}

	return room, nil
}

func (s *svc) List(ctx *gofr.Context, req *services.ListRoom) ([]*entities.Room, error) {
	return s.repo.List(ctx, &repos.RoomFilter{UserID: req.UserID, Page: req.Page, Size: req.Size, Include: req.Include})
}

func (s *svc) ValidateRole(ctx *gofr.Context, roomID string, userID string, roles ...roommember.Role) (*entities.Room, error) {
	room, err := s.repo.Get(ctx, &entities.Room{ID: roomID})
	if err != nil {
		return nil, err
	}

	for _, member := range room.Members {
		if member.UserID == userID && slices.Contains(roles, member.Role) {
			return room, nil
		}
	}

	return nil, errors.MissingRoles(roles)
}
