package middleware

type contextKey string

const (
	ClaimsKey   contextKey = "claims"
	UserIdKey   contextKey = "userId"
	UserRoleKey contextKey = "userRole"
)
