package api

import (
	"context"
	"net/http"
	"strconv"
	"tantan-demo/lib/log"
	"tantan-demo/storage"

	"github.com/go-pg/pg"
)

// ResUser represents user resource object in the request or response body.
type ResUser struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

// CreateUser creates a user.
func CreateUser(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	req := &ResUser{}
	if err := ReadJSON(ctx, r, req); err != nil {
		log.Errorfx(ctx, "parse request body err(%v)", err)
		return WriteJSON(ctx, w, http.StatusBadRequest, nil)
	}
	user := &storage.UserInfo{
		Name: req.Name,
	}
	us := storage.UserStorage{}
	status := http.StatusOK
	if err := us.CreateUser(user); err != nil {
		if err == storage.ErrUserExist {
			status = http.StatusCreated
		} else {
			log.Errorfx(ctx, "us.CreateUser err(%v)", err)
			return WriteJSON(ctx, w, http.StatusInternalServerError, nil)
		}
	}
	resp := &ResUser{
		ID:   strconv.FormatUint(user.ID, 10),
		Name: user.Name,
		Type: ResTypeUser,
	}
	return WriteJSON(ctx, w, status, resp)
}

// GetAllUsers lists all users.
func GetAllUsers(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	us := storage.UserStorage{}
	list, err := us.ListAllUsers()
	if err != nil {
		if err == pg.ErrNoRows {
			log.Warningfx(ctx, "us.ListAllUsers result empty", err)
			return WriteJSON(ctx, w, http.StatusNotFound, nil)
		}
		log.Errorfx(ctx, "parse request body err(%v)", err)
		return WriteJSON(ctx, w, http.StatusInternalServerError, nil)
	}
	var resp []*ResUser
	for _, user := range list {
		item := &ResUser{
			ID:   strconv.FormatUint(user.ID, 10),
			Name: user.Name,
			Type: ResTypeUser,
		}
		resp = append(resp, item)
	}
	if len(resp) == 0 {
		log.Warningfx(ctx, "no user found")
		return WriteJSON(ctx, w, http.StatusNotFound, nil)
	}
	return WriteJSON(ctx, w, http.StatusOK, resp)
}
