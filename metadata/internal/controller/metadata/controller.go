package metadata

import (
	"context"
	"errors"

	"golang-microservices/metadata/internal/repository"
	"golang-microservices/metadata/pkg/model"
)

// ErrNotFound I cant find it ok
var ErrNotFound = errors.New("not found")

type metadataRepository interface {
	Get(ctx context.Context, id string) (*model.Metadata, error)
}

// Controller service metadata
type Controller struct {
	repo metadataRepository
}

// New creates a metadata controller
func New(repo metadataRepository) *Controller {
	return &Controller{repo}
}

// Get returns movie metadata by id
func (c *Controller) Get(ctx context.Context, id string) (*model.Metadata, error) {
	res, err := c.repo.Get(ctx, id)
	if err != nil && errors.Is(err, repository.ErrNotFound) {
		return nil, ErrNotFound
	}
	return res, err
}
