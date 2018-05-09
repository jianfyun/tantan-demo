package storage

import (
	"time"
)

const (
	StateLiked    = "liked"
	StateDisliked = "disliked"
	StateMatched  = "matched"
)

// Relationship represents the model of table relationships.
type Relationship struct {
	ID          uint64
	UserID      uint64
	OtherUserID uint64
	State       string
	CreateTime  time.Time
}

// RelationStorage controls the read/write operations of table relationships.
type RelationStorage struct {
}

// UpsertRelation inserts or updates the relationship state from one user to another.
func (r *RelationStorage) UpsertRelation(uid, oid uint64, state string) error {
	relation := &Relationship{
		UserID:      uid,
		OtherUserID: oid,
		State:       state,
	}
	_, err := db.Model(relation).OnConflict("(user_id, other_user_id) DO UPDATE").Set("state = EXCLUDED.state").Insert()
	if err != nil {
		return err
	}
	return nil
}

// FindRelation retrieves the relationship from uid to oid.
func (r *RelationStorage) FindRelation(uid, oid uint64) (*Relationship, error) {
	relation := &Relationship{}
	_, err := db.QueryOne(relation, "SELECT * FROM relationships WHERE user_id = ? AND other_user_id = ?", uid, oid)
	if err != nil {
		return nil, err
	}
	return relation, nil
}

// ListUserAllRelations retrieves all the relationships of one user.
func (r *RelationStorage) ListUserAllRelations(uid uint64) ([]*Relationship, error) {
	var relations []*Relationship
	_, err := db.Query(&relations, "SELECT * FROM relationships WHERE user_id = ?", uid)
	if err != nil {
		return nil, err
	}
	return relations, nil
}
