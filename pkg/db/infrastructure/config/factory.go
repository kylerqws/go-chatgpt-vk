package config

import (
	"context"
	"fmt"

	ctrcfg "github.com/kylerqws/chatbot/pkg/db/contract/config"
	"github.com/kylerqws/chatbot/pkg/db/infrastructure/config/source"
)

const (
	SourceTypeKey     = "sourceType"
	DefaultSourceType = "env"
)

func New(ctx context.Context) (ctrcfg.Config, error) {
	st, ok := ctx.Value(SourceTypeKey).(string)
	if !ok || st == "" {
		st = DefaultSourceType
	}

	switch st {
	case "env":
		return source.NewEnvConfig(ctx)
	default:
		return nil, fmt.Errorf("unsupported database configuration source: %q", st)
	}
}
