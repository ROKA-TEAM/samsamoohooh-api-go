// Code generated by ent, DO NOT EDIT.

package user

import (
	"samsamoohooh-go-api/internal/infra/storage/mysql/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.User {
	return predicate.User(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.User {
	return predicate.User(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.User {
	return predicate.User(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.User {
	return predicate.User(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.User {
	return predicate.User(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUpdatedAt, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldName, v))
}

// Resolution applies equality check predicate on the "resolution" field. It's identical to ResolutionEQ.
func Resolution(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldResolution, v))
}

// SocialSub applies equality check predicate on the "socialSub" field. It's identical to SocialSubEQ.
func SocialSub(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldSocialSub, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldUpdatedAt, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldName, v))
}

// ResolutionEQ applies the EQ predicate on the "resolution" field.
func ResolutionEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldResolution, v))
}

// ResolutionNEQ applies the NEQ predicate on the "resolution" field.
func ResolutionNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldResolution, v))
}

// ResolutionIn applies the In predicate on the "resolution" field.
func ResolutionIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldResolution, vs...))
}

// ResolutionNotIn applies the NotIn predicate on the "resolution" field.
func ResolutionNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldResolution, vs...))
}

// ResolutionGT applies the GT predicate on the "resolution" field.
func ResolutionGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldResolution, v))
}

// ResolutionGTE applies the GTE predicate on the "resolution" field.
func ResolutionGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldResolution, v))
}

// ResolutionLT applies the LT predicate on the "resolution" field.
func ResolutionLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldResolution, v))
}

// ResolutionLTE applies the LTE predicate on the "resolution" field.
func ResolutionLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldResolution, v))
}

// ResolutionContains applies the Contains predicate on the "resolution" field.
func ResolutionContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldResolution, v))
}

// ResolutionHasPrefix applies the HasPrefix predicate on the "resolution" field.
func ResolutionHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldResolution, v))
}

// ResolutionHasSuffix applies the HasSuffix predicate on the "resolution" field.
func ResolutionHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldResolution, v))
}

// ResolutionEqualFold applies the EqualFold predicate on the "resolution" field.
func ResolutionEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldResolution, v))
}

// ResolutionContainsFold applies the ContainsFold predicate on the "resolution" field.
func ResolutionContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldResolution, v))
}

// RoleEQ applies the EQ predicate on the "role" field.
func RoleEQ(v Role) predicate.User {
	return predicate.User(sql.FieldEQ(FieldRole, v))
}

// RoleNEQ applies the NEQ predicate on the "role" field.
func RoleNEQ(v Role) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldRole, v))
}

// RoleIn applies the In predicate on the "role" field.
func RoleIn(vs ...Role) predicate.User {
	return predicate.User(sql.FieldIn(FieldRole, vs...))
}

// RoleNotIn applies the NotIn predicate on the "role" field.
func RoleNotIn(vs ...Role) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldRole, vs...))
}

// SocialEQ applies the EQ predicate on the "social" field.
func SocialEQ(v Social) predicate.User {
	return predicate.User(sql.FieldEQ(FieldSocial, v))
}

// SocialNEQ applies the NEQ predicate on the "social" field.
func SocialNEQ(v Social) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldSocial, v))
}

// SocialIn applies the In predicate on the "social" field.
func SocialIn(vs ...Social) predicate.User {
	return predicate.User(sql.FieldIn(FieldSocial, vs...))
}

// SocialNotIn applies the NotIn predicate on the "social" field.
func SocialNotIn(vs ...Social) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldSocial, vs...))
}

// SocialSubEQ applies the EQ predicate on the "socialSub" field.
func SocialSubEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldSocialSub, v))
}

// SocialSubNEQ applies the NEQ predicate on the "socialSub" field.
func SocialSubNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldSocialSub, v))
}

// SocialSubIn applies the In predicate on the "socialSub" field.
func SocialSubIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldSocialSub, vs...))
}

// SocialSubNotIn applies the NotIn predicate on the "socialSub" field.
func SocialSubNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldSocialSub, vs...))
}

// SocialSubGT applies the GT predicate on the "socialSub" field.
func SocialSubGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldSocialSub, v))
}

// SocialSubGTE applies the GTE predicate on the "socialSub" field.
func SocialSubGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldSocialSub, v))
}

// SocialSubLT applies the LT predicate on the "socialSub" field.
func SocialSubLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldSocialSub, v))
}

// SocialSubLTE applies the LTE predicate on the "socialSub" field.
func SocialSubLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldSocialSub, v))
}

// SocialSubContains applies the Contains predicate on the "socialSub" field.
func SocialSubContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldSocialSub, v))
}

// SocialSubHasPrefix applies the HasPrefix predicate on the "socialSub" field.
func SocialSubHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldSocialSub, v))
}

// SocialSubHasSuffix applies the HasSuffix predicate on the "socialSub" field.
func SocialSubHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldSocialSub, v))
}

// SocialSubEqualFold applies the EqualFold predicate on the "socialSub" field.
func SocialSubEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldSocialSub, v))
}

// SocialSubContainsFold applies the ContainsFold predicate on the "socialSub" field.
func SocialSubContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldSocialSub, v))
}

// HasGroups applies the HasEdge predicate on the "groups" edge.
func HasGroups() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, GroupsTable, GroupsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasGroupsWith applies the HasEdge predicate on the "groups" edge with a given conditions (other predicates).
func HasGroupsWith(preds ...predicate.Group) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newGroupsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasComments applies the HasEdge predicate on the "comments" edge.
func HasComments() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CommentsTable, CommentsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCommentsWith applies the HasEdge predicate on the "comments" edge with a given conditions (other predicates).
func HasCommentsWith(preds ...predicate.Comment) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newCommentsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPosts applies the HasEdge predicate on the "posts" edge.
func HasPosts() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PostsTable, PostsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPostsWith applies the HasEdge predicate on the "posts" edge with a given conditions (other predicates).
func HasPostsWith(preds ...predicate.Post) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newPostsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTopics applies the HasEdge predicate on the "topics" edge.
func HasTopics() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TopicsTable, TopicsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTopicsWith applies the HasEdge predicate on the "topics" edge with a given conditions (other predicates).
func HasTopicsWith(preds ...predicate.Topic) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newTopicsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.User) predicate.User {
	return predicate.User(sql.NotPredicates(p))
}
