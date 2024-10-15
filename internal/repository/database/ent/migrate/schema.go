// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CommentsColumns holds the columns for the "comments" table.
	CommentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "delete_time", Type: field.TypeTime, Nullable: true},
		{Name: "content", Type: field.TypeString},
		{Name: "post_comments", Type: field.TypeInt, Nullable: true},
		{Name: "user_comments", Type: field.TypeInt, Nullable: true},
	}
	// CommentsTable holds the schema information for the "comments" table.
	CommentsTable = &schema.Table{
		Name:       "comments",
		Columns:    CommentsColumns,
		PrimaryKey: []*schema.Column{CommentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "comments_posts_comments",
				Columns:    []*schema.Column{CommentsColumns[5]},
				RefColumns: []*schema.Column{PostsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "comments_users_comments",
				Columns:    []*schema.Column{CommentsColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// GroupsColumns holds the columns for the "groups" table.
	GroupsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "delete_time", Type: field.TypeTime, Nullable: true},
		{Name: "book_title", Type: field.TypeString},
		{Name: "author", Type: field.TypeString},
		{Name: "max_page", Type: field.TypeInt},
		{Name: "publisher", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "book_mark", Type: field.TypeInt},
	}
	// GroupsTable holds the schema information for the "groups" table.
	GroupsTable = &schema.Table{
		Name:       "groups",
		Columns:    GroupsColumns,
		PrimaryKey: []*schema.Column{GroupsColumns[0]},
	}
	// PostsColumns holds the columns for the "posts" table.
	PostsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "delete_time", Type: field.TypeTime, Nullable: true},
		{Name: "title", Type: field.TypeString},
		{Name: "content", Type: field.TypeString},
		{Name: "group_posts", Type: field.TypeInt, Nullable: true},
		{Name: "user_posts", Type: field.TypeInt, Nullable: true},
	}
	// PostsTable holds the schema information for the "posts" table.
	PostsTable = &schema.Table{
		Name:       "posts",
		Columns:    PostsColumns,
		PrimaryKey: []*schema.Column{PostsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "posts_groups_posts",
				Columns:    []*schema.Column{PostsColumns[6]},
				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "posts_users_posts",
				Columns:    []*schema.Column{PostsColumns[7]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// TasksColumns holds the columns for the "tasks" table.
	TasksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "delete_time", Type: field.TypeTime, Nullable: true},
		{Name: "deadline", Type: field.TypeTime},
		{Name: "range", Type: field.TypeInt},
		{Name: "group_tasks", Type: field.TypeInt, Nullable: true},
	}
	// TasksTable holds the schema information for the "tasks" table.
	TasksTable = &schema.Table{
		Name:       "tasks",
		Columns:    TasksColumns,
		PrimaryKey: []*schema.Column{TasksColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "tasks_groups_tasks",
				Columns:    []*schema.Column{TasksColumns[6]},
				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// TopicsColumns holds the columns for the "topics" table.
	TopicsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "delete_time", Type: field.TypeTime, Nullable: true},
		{Name: "field", Type: field.TypeString},
		{Name: "feeling", Type: field.TypeString},
		{Name: "task_topics", Type: field.TypeInt, Nullable: true},
		{Name: "user_topics", Type: field.TypeInt, Nullable: true},
	}
	// TopicsTable holds the schema information for the "topics" table.
	TopicsTable = &schema.Table{
		Name:       "topics",
		Columns:    TopicsColumns,
		PrimaryKey: []*schema.Column{TopicsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "topics_tasks_topics",
				Columns:    []*schema.Column{TopicsColumns[6]},
				RefColumns: []*schema.Column{TasksColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "topics_users_topics",
				Columns:    []*schema.Column{TopicsColumns[7]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "delete_time", Type: field.TypeTime, Nullable: true},
		{Name: "name", Type: field.TypeString},
		{Name: "resolution", Type: field.TypeString},
		{Name: "role", Type: field.TypeEnum, Enums: []string{"ADMIN", "GUEST"}},
		{Name: "social", Type: field.TypeEnum, Enums: []string{"KAKAO", "APPLE", "GOOGLE"}},
		{Name: "social_sub", Type: field.TypeString, Unique: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// UserGroupsColumns holds the columns for the "user_groups" table.
	UserGroupsColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeInt},
		{Name: "group_id", Type: field.TypeInt},
	}
	// UserGroupsTable holds the schema information for the "user_groups" table.
	UserGroupsTable = &schema.Table{
		Name:       "user_groups",
		Columns:    UserGroupsColumns,
		PrimaryKey: []*schema.Column{UserGroupsColumns[0], UserGroupsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_groups_user_id",
				Columns:    []*schema.Column{UserGroupsColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_groups_group_id",
				Columns:    []*schema.Column{UserGroupsColumns[1]},
				RefColumns: []*schema.Column{GroupsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CommentsTable,
		GroupsTable,
		PostsTable,
		TasksTable,
		TopicsTable,
		UsersTable,
		UserGroupsTable,
	}
)

func init() {
	CommentsTable.ForeignKeys[0].RefTable = PostsTable
	CommentsTable.ForeignKeys[1].RefTable = UsersTable
	PostsTable.ForeignKeys[0].RefTable = GroupsTable
	PostsTable.ForeignKeys[1].RefTable = UsersTable
	TasksTable.ForeignKeys[0].RefTable = GroupsTable
	TopicsTable.ForeignKeys[0].RefTable = TasksTable
	TopicsTable.ForeignKeys[1].RefTable = UsersTable
	UserGroupsTable.ForeignKeys[0].RefTable = UsersTable
	UserGroupsTable.ForeignKeys[1].RefTable = GroupsTable
}
