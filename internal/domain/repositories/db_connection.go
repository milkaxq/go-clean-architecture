package repositories

import "context"

type DBConnection interface {
	Connect(ctx context.Context) error
	Close(ctx context.Context) error
}
