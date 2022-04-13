package authproviders

import (
	"context"

	"github.com/stackrox/stackrox/generated/storage"
)

// Store provides storage functionality for auth providers.
type Store interface {
	GetAllAuthProviders() ([]*storage.AuthProvider, error)

	AddAuthProvider(ctx context.Context, authProvider *storage.AuthProvider) error
	UpdateAuthProvider(ctx context.Context, authProvider *storage.AuthProvider) error
	RemoveAuthProvider(ctx context.Context, id string) error
}
