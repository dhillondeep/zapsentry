package zapsentry

import (
	"net/http"

	"github.com/getsentry/sentry-go"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	zapSentryScopeKey = "_zapsentry_scope_"
)

// ScopeFunc is a function that can be used to apply changes to the underlying scope
type ScopeFunc func(scope *sentry.Scope)

// Scope is abstraction over sentry.Scope
type Scope struct {
	scope *sentry.Scope
}

// NewScope creates a new Scope object
func NewScope() *Scope {
	return &Scope{sentry.NewScope()}
}

// SetRequest sets the request on the underlying scope
func (s *Scope) SetRequest(r *http.Request) *Scope {
	s.scope.SetRequest(r)
	return s
}

// SetUser sets the user on the underlying scope
func (s *Scope) SetUser(user sentry.User) *Scope {
	s.scope.SetUser(user)
	return s
}

// SetTag sets a tag on the underlying scope
func (s *Scope) SetTag(key, value string) *Scope {
	s.scope.SetTag(key, value)
	return s
}

// SetTags sets the tags on the underlying scope
func (s *Scope) SetTags(tags map[string]string) *Scope {
	s.scope.SetTags(tags)
	return s
}

// Apply applies direct changes to the underlying scope
func (s *Scope) Apply(sf ScopeFunc) *Scope {
	sf(s.scope)
	return s
}

// Build constructs a zapcore.Field object from the current scope
func (s *Scope) Build() zapcore.Field {
	f := zap.Skip()
	f.Interface = s.scope
	f.Key = zapSentryScopeKey

	return f
}
