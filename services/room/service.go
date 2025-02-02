package room

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/hnpatil/gochat/entities"
	"github.com/hnpatil/gochat/entities/roommember"
	"github.com/hnpatil/gochat/repos"
	"github.com/hnpatil/gochat/services"
	"gofr.dev/pkg/gofr"
	"slices"
)

type svc struct {
	repo repos.Room
}

func New(repo repos.Room) services.Room {
	return &svc{repo}
}

func (s *svc) Create(ctx *gofr.Context, req *services.CreateRoom) (*entities.Room, error) {
	room := &entities.Room{
		ID:      req.RoomID,
		Name:    req.Name,
		IsGroup: false,
	}

	if room.ID == "" {
		room.ID = uuid.New().String()
	}

	memberMap := make(map[string]roommember.Role, len(req.Members))
	for _, m := range req.Members {
		memberMap[m] = roommember.RoleMember
	}

	memberMap[req.UserID] = roommember.RoleAdmin

	if len(memberMap) > 2 {
		room.IsGroup = true
	}

	members := make([]*entities.RoomMember, 0, len(memberMap))
	for k, v := range memberMap {
		members = append(members, &entities.RoomMember{RoomID: room.ID, UserID: k, Role: v})
	}

	room.Members = members

	return s.repo.Create(ctx, room)
}

func (s *svc) Update(ctx *gofr.Context, req *services.UpdateRoom) (*entities.Room, error) {
	err := s.ValidateRole(ctx, req.RoomID, req.UserID, roommember.RoleAdmin)
	if err != nil {
		return nil, err
	}

	return s.repo.Update(ctx, &entities.Room{ID: req.RoomID}, &entities.Room{Name: req.Name})
}

func (s *svc) Get(ctx *gofr.Context, req *services.GetRoom) (*entities.Room, error) {
	err := s.ValidateRole(ctx, req.RoomID, req.UserID, roommember.RoleAdmin, roommember.RoleMember)
	if err != nil {
		return nil, err
	}

	return s.repo.Get(ctx, &entities.Room{ID: req.RoomID})
}

func (s *svc) List(ctx *gofr.Context, req *services.ListRoom) ([]*entities.Room, error) {
	return s.repo.List(ctx, &repos.RoomFilter{UserID: req.UserID, Page: req.Page, Size: req.Size, Include: req.Include})
}

func (s *svc) Delete(ctx *gofr.Context, req *services.DeleteRoom) error {
	err := s.ValidateRole(ctx, req.RoomID, req.UserID, roommember.RoleAdmin)
	if err != nil {
		return err
	}

	return s.repo.Delete(ctx, &entities.Room{ID: req.RoomID})
}

func (s *svc) ValidateRole(ctx *gofr.Context, roomID string, userID string, roles ...roommember.Role) error {
	room, err := s.repo.Get(ctx, &entities.Room{ID: roomID})
	if err != nil {
		return err
	}

	for _, member := range room.Members {
		if member.UserID == userID && slices.Contains(roles, member.Role) {
			return nil
		}
	}

	return services.UnAuthorisedError(fmt.Sprintf("user lacks %v role", roles))
}
