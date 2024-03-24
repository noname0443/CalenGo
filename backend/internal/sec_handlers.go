package internal

import (
	"context"

	"github.com/noname0443/CalenGo/backend/api"
)

type SenHandler struct{}

func (sh *SenHandler) HandleBasicAuth(ctx context.Context, operationName string, t api.BasicAuth) (context.Context, error) {
	return context.WithValue(ctx, "credentials", t), nil
}
