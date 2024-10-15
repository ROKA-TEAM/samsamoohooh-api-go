package utils

import (
	"samsamoohooh-go-api/internal/domain"
	"samsamoohooh-go-api/internal/repository/database/ent"
)

func ConvertDomainUser(user *ent.User) *domain.User {
	return &domain.User{
		ID:         user.ID,
		Name:       user.Name,
		Resolution: user.Resolution,
		Social:     domain.UserSocialType(user.Social),
		SocialSub:  user.SocialSub,
		CreatedAt:  user.CreatedAt,
		UpdatedAt:  user.UpdatedAt,
		DeletedAt:  user.DeleteAt,
	}
}

func ConvertDomainUsers(users []*ent.User) []*domain.User {
	var domainUsers []*domain.User
	for _, user := range users {
		domainUsers = append(domainUsers, ConvertDomainUser(user))
	}

	return domainUsers
}

func ConvertDomainGroup(group *ent.Group) *domain.Group {
	return &domain.Group{
		ID:        group.ID,
		BookTitle: group.BookTitle,
		Author:    group.Author,
		MaxPage:   group.MaxPage,
		Publisher: group.Publisher,
		Bookmark:  group.BookMark,
		CreatedAt: group.CreatedAt,
		UpdatedAt: group.UpdatedAt,
		DeletedAt: group.DeleteAt,
	}
}

func ConvertDomainGroups(group []*ent.Group) []*domain.Group {
	var domainGroups []*domain.Group
	for _, group := range group {
		domainGroups = append(domainGroups, ConvertDomainGroup(group))
	}

	return domainGroups
}

func ConvertDomainPost(post *ent.Post) *domain.Post {
	return &domain.Post{
		ID:        post.ID,
		Title:     post.Title,
		Content:   post.Content,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
		DeletedAt: post.DeleteAt,
	}
}

func ConvertDomainPosts(posts []*ent.Post) []*domain.Post {
	var domainPosts []*domain.Post
	for _, post := range posts {
		domainPosts = append(domainPosts, ConvertDomainPost(post))
	}

	return domainPosts
}

func ConvertDomainComment(comment *ent.Comment) *domain.Comment {
	return &domain.Comment{
		ID:        comment.ID,
		Content:   comment.Content,
		CreatedAt: comment.CreatedAt,
		UpdatedAt: comment.UpdatedAt,
		DeletedAt: comment.DeleteAt,
	}
}

func ConvertDomainComments(comments []*ent.Comment) []*domain.Comment {
	var domainComments []*domain.Comment
	for _, comment := range comments {
		domainComments = append(domainComments, ConvertDomainComment(comment))
	}

	return domainComments
}

func ConvertDomainTask(task *ent.Task) *domain.Task {
	return &domain.Task{
		ID:        task.ID,
		Deadline:  task.Deadline,
		Range:     task.Range,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
		DeletedAt: task.DeleteAt,
	}
}

func ConvertDomainTasks(tasks []*ent.Task) []*domain.Task {
	var domainTasks []*domain.Task
	for _, task := range tasks {
		domainTasks = append(domainTasks, ConvertDomainTask(task))
	}

	return domainTasks
}

func ConvertDomainTopic(topic *ent.Topic) *domain.Topic {
	return &domain.Topic{
		ID:        topic.ID,
		Topic:     topic.Topic,
		Feeling:   topic.Feeling,
		CreatedAt: topic.CreatedAt,
		UpdatedAt: topic.UpdatedAt,
		DeletedAt: topic.DeleteAt,
	}
}

func ConvertDomainTopics(topics []*ent.Topic) []*domain.Topic {
	var domainTopics []*domain.Topic
	for _, topic := range topics {
		domainTopics = append(domainTopics, ConvertDomainTopic(topic))
	}

	return domainTopics
}
