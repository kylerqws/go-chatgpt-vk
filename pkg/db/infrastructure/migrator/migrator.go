package migrator

import (
	"context"
	"fmt"

	"github.com/uptrace/bun/migrate"

	ctrcli "github.com/kylerqws/chatbot/pkg/db/contract/client"
	ctrmig "github.com/kylerqws/chatbot/pkg/db/contract/migrator"
)

type migrator struct {
	client ctrcli.Client
}

func New(cl ctrcli.Client) ctrmig.Migrator {
	return &migrator{client: cl}
}

func (m *migrator) Migrate(ctx context.Context, mgs *migrate.Migrations) error {
	migrator, err := m.newMigrator(ctx, mgs)
	if err != nil {
		return fmt.Errorf("migrator: failed to create migrator: %w", err)
	}

	_, err = migrator.Migrate(ctx)
	if err != nil {
		return fmt.Errorf("migrator: failed to execute migrations: %w", err)
	}

	return nil
}

func (m *migrator) Rollback(ctx context.Context, mgs *migrate.Migrations) error {
	migrator, err := m.newMigrator(ctx, mgs)
	if err != nil {
		return fmt.Errorf("migrator: failed to create migrator: %w", err)
	}

	_, err = migrator.Rollback(ctx)
	if err != nil {
		return fmt.Errorf("migrator: failed to rollback migrations: %w", err)
	}

	return nil
}

func (m *migrator) newMigrator(ctx context.Context, mgs *migrate.Migrations) (*migrate.Migrator, error) {
	migrator := migrate.NewMigrator(m.client.DB(), mgs)

	if err := migrator.Init(ctx); err != nil {
		return nil, fmt.Errorf("migrator: failed to init migrator: %w", err)
	}

	return migrator, nil
}
