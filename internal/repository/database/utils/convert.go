package utils

import (
	domain2 "samsamoohooh-go-api/internal/application/domain"
	"samsamoohooh-go-api/internal/repository/database/ent"
)

func ConvertDomainUser(user *ent.User) *domain2.User {
	return &domain2.User{
		ID:         user.ID,
		Name:       user.Name,
		Resolution: user.Resolution,
		Role:       domain2.UserRoleType(user.Role),
		Social:     domain2.UserSocialType(user.Social),
		SocialSub:  user.SocialSub,
		CreatedAt:  user.CreatedAt,
		UpdatedAt:  user.UpdatedAt,
		DeletedAt:  user.DeleteAt,
	}
}

func ConvertDomainUsers(users []*ent.User) []*domain2.User {
	var domainUsers []*domain2.User
	for _, user := range users {
		domainUsers = append(domainUsers, ConvertDomainUser(user))
	}

	return domainUsers
}

func ConvertDomainGroup(group *ent.Group) *domain2.Group {
	return &domain2.Group{
		ID:          group.ID,
		BookTitle:   group.BookTitle,
		Author:      group.Author,
		MaxPage:     group.MaxPage,
		Publisher:   group.Publisher,
		Bookmark:    group.BookMark,
		Description: group.Description,
		CreatedAt:   group.CreatedAt,
		UpdatedAt:   group.UpdatedAt,
		DeletedAt:   group.DeleteAt,
	}
}

func ConvertDomainGroups(group []*ent.Group) []*domain2.Group {
	var domainGroups []*domain2.Group
	for _, group := range group {
		domainGroups = append(domainGroups, ConvertDomainGroup(group))
	}

	return domainGroups
}

func ConvertDomainPost(post *ent.Post) *domain2.Post {
	return &domain2.Post{
		ID:        post.ID,
		Title:     post.Title,
		Content:   post.Content,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
		DeletedAt: post.DeleteAt,
	}
}

func ConvertDomainPosts(posts []*ent.Post) []*domain2.Post {
	var domainPosts []*domain2.Post
	for _, post := range posts {
		domainPosts = append(domainPosts, ConvertDomainPost(post))
	}

	return domainPosts
}

func ConvertDomainComment(comment *ent.Comment) *domain2.Comment {
	return &domain2.Comment{
		ID:        comment.ID,
		Content:   comment.Content,
		CreatedAt: comment.CreatedAt,
		UpdatedAt: comment.UpdatedAt,
		DeletedAt: comment.DeleteAt,
	}
}

func ConvertDomainComments(comments []*ent.Comment) []*domain2.Comment {
	var domainComments []*domain2.Comment
	for _, comment := range comments {
		domainComments = append(domainComments, ConvertDomainComment(comment))
	}

	return domainComments
}

func ConvertDomainTask(task *ent.Task) *domain2.Task {
	return &domain2.Task{
		ID:        task.ID,
		Deadline:  task.Deadline,
		Range:     task.Range,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
		DeletedAt: task.DeleteAt,
	}
}

func ConvertDomainTasks(tasks []*ent.Task) []*domain2.Task {
	var domainTasks []*domain2.Task
	for _, task := range tasks {
		domainTasks = append(domainTasks, ConvertDomainTask(task))
	}

	return domainTasks
}

func ConvertDomainTopic(topic *ent.Topic) *domain2.Topic {
	return &domain2.Topic{
		ID:        topic.ID,
		Topic:     topic.Topic,
		Feeling:   topic.Feeling,
		CreatedAt: topic.CreatedAt,
		UpdatedAt: topic.UpdatedAt,
		DeletedAt: topic.DeleteAt,
	}
}

func ConvertDomainTopics(topics []*ent.Topic) []*domain2.Topic {
	var domainTopics []*domain2.Topic
	for _, topic := range topics {
		domainTopics = append(domainTopics, ConvertDomainTopic(topic))
	}

	return domainTopics
}
