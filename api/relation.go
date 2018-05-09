package api

import (
	"context"
	"net/http"
	"strconv"
	"tantan-demo/lib/log"
	"tantan-demo/storage"

	"github.com/go-pg/pg"
	"github.com/gorilla/mux"
)

// ResRelationship represents relationship resource object in the request or response body.
type ResRelationship struct {
	UserID string `json:"user_id"`
	State  string `json:"state"`
	Type   string `json:"type"`
}

// PutUserRelation creates or updates relationship state to another user.
func PutUserRelation(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["user_id"], 10, 64)
	if err != nil {
		log.Errorfx(ctx, "parse user_id err(%v)", err)
		return WriteJSON(ctx, w, http.StatusBadRequest, nil)
	}
	oid, err := strconv.ParseUint(vars["other_user_id"], 10, 64)
	if err != nil {
		log.Errorfx(ctx, "parse other_user_id err(%v)", err)
		return WriteJSON(ctx, w, http.StatusBadRequest, nil)
	}
	req := &ResRelationship{}
	if err := ReadJSON(ctx, r, req); err != nil {
		log.Errorfx(ctx, "parse request body err(%v)", err)
		return WriteJSON(ctx, w, http.StatusBadRequest, nil)
	}
	rs := &storage.RelationStorage{}
	if err := rs.UpsertRelation(uid, oid, req.State); err != nil {
		log.Errorfx(ctx, "rs.UpsertRelation err(%v)", err)
		return WriteJSON(ctx, w, http.StatusInternalServerError, nil)
	}
	resp := &ResRelationship{
		UserID: vars["other_user_id"],
		Type:   ResTypeRelationship,
	}
	relation, err := rs.FindRelation(oid, uid)
	if err != nil {
		if err != pg.ErrNoRows {
			log.Errorfx(ctx, "rs.FindRelation err(%v)", err)
			return WriteJSON(ctx, w, http.StatusInternalServerError, nil)
		}
	} else if req.State == storage.StateLiked && relation.State == storage.StateLiked {
		resp.State = storage.StateMatched
	}
	if resp.State == "" {
		resp.State = req.State
	}
	return WriteJSON(ctx, w, http.StatusOK, resp)
}

// GetUserRelations lists a user's all relationships.
func GetUserRelations(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["user_id"], 10, 64)
	if err != nil {
		log.Errorfx(ctx, "parse user_id err(%v)", err)
		return WriteJSON(ctx, w, http.StatusBadRequest, nil)
	}
	rs := &storage.RelationStorage{}
	list, err := rs.ListUserAllRelations(uid)
	if err != nil {
		if err == pg.ErrNoRows {
			log.Warningfx(ctx, "rs.ListUserAllRelations result empty", err)
			return WriteJSON(ctx, w, http.StatusNotFound, nil)
		}
		log.Errorfx(ctx, "rs.ListUserAllRelations err(%v)", err)
		return WriteJSON(ctx, w, http.StatusInternalServerError, nil)
	}
	var resp []*ResRelationship
	for _, relation := range list {
		resp = append(resp, &ResRelationship{
			UserID: strconv.FormatUint(relation.OtherUserID, 10),
			State:  relation.State,
			Type:   ResTypeRelationship,
		})
	}
	if len(resp) == 0 {
		log.Warningfx(ctx, "no relationship founded")
		return WriteJSON(ctx, w, http.StatusNotFound, nil)
	}
	return WriteJSON(ctx, w, http.StatusOK, resp)
}
