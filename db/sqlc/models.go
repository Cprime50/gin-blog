// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package db

import (
	"time"
)

type Article struct {
	ID          string    `json:"id"`
	AuthorID    string    `json:"author_id"`
	Slug        string    `json:"slug"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Body        string    `json:"body"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ArticleTag struct {
	ArticleID string `json:"article_id"`
	TagID     string `json:"tag_id"`
}

type Comment struct {
	ID        string    `json:"id"`
	AuthorID  string    `json:"author_id"`
	ArticleID string    `json:"article_id"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Favorite struct {
	ArticleID string `json:"article_id"`
	UserID    string `json:"user_id"`
}

type Follow struct {
	FollowerID  string `json:"follower_id"`
	FollowingID string `json:"following_id"`
}

type Tag struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Bio       *string   `json:"bio"`
	Image     *string   `json:"image"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
