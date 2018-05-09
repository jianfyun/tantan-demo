package storage

import "testing"

func TestRelationStorage_UpsertRelation(t *testing.T) {
	rs := &RelationStorage{}
	if err := rs.UpsertRelation(1, 2, StateDisliked); err != nil {
		t.Error(err)
		t.FailNow()
	}
	if err := rs.UpsertRelation(1, 2, StateLiked); err != nil {
		t.Error(err)
		t.FailNow()
	}
	if err := rs.UpsertRelation(2, 1, StateLiked); err != nil {
		t.Error(err)
		t.FailNow()
	}
	if err := rs.UpsertRelation(1, 3, StateLiked); err != nil {
		t.Error(err)
		t.FailNow()
	}
}

func TestRelationStorage_FindRelation(t *testing.T) {
	rs := &RelationStorage{}
	relation, err := rs.FindRelation(1, 2)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("relation=%v", relation)
}

func TestRelationStorage_ListUserAllRelations(t *testing.T) {
	rs := &RelationStorage{}
	relations, err := rs.ListUserAllRelations(1)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	for _, relation := range relations {
		t.Logf("relation=%v", relation)
	}
}
