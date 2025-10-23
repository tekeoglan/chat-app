package model

import (
	"context"
	"net/http"
	"time"
)

const COOKIE_PREFIX_SESSION = "session_id"

const CONTEXT_USER_KEY = "user_id"

type SessionService interface {
	CreateSession(c context.Context, userId string) (string, error)
	RetriveSession(c context.Context, sessionId string) (string, error)
	RemoveSession(c context.Context, sessionId string) error
	GetCokiPath() string
	GetCokiDomain() string
	GetCokiExpr() int
	GetExpr() time.Duration
	GetSameSite() http.SameSite
	IsCokiSecure() bool
	IsCokiHttpOnly() bool
}
