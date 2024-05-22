package models

import db "github.com/maxzhovtyj/live/internal/pkg/db/sqlc"

type Context struct {
	User db.User
}
