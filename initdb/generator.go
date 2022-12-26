package initdb

import (
	"agricoladb/ent"
	"agricoladb/ent/migrate"
	"context"
	"fmt"
	"path/filepath"
)

type Generator struct {
	client        *ent.Client
	csvDir        string
	needsTruncate bool
}

func NewGenerator(client *ent.Client, csvDir string, needsTruncate bool) *Generator {
	return &Generator{
		client:        client,
		csvDir:        csvDir,
		needsTruncate: needsTruncate,
	}
}

func (g *Generator) Generate() error {
	ctx := context.Background()
	return g.GenerateWithContext(ctx)
}

func (g *Generator) GenerateWithContext(ctx context.Context) error {
	if g.needsTruncate {
		if err := truncate(ctx, g.client); err != nil {
			return err
		}
	}

	tx, err := g.client.Tx(ctx)
	if err != nil {
		return err
	}

	if err = initRevisions(ctx, tx, filepath.Join(g.csvDir, RevisionsFileName)); err != nil {
		return rollback(tx, fmt.Errorf("failed initRevisions: %w", err))
	}
	if err = initProducts(ctx, tx, filepath.Join(g.csvDir, ProductsFileName)); err != nil {
		return rollback(tx, fmt.Errorf("failed initProducts: %w", err))
	}
	if err = initDecks(ctx, tx, filepath.Join(g.csvDir, DecksFileName)); err != nil {
		return rollback(tx, fmt.Errorf("failed initDecks: %w", err))
	}
	if err = initCardSpecialColors(ctx, tx, filepath.Join(g.csvDir, CardSpecialColorsFileName)); err != nil {
		return rollback(tx, fmt.Errorf("failed initCardSpecialColors: %w", err))
	}
	if err = initCardTypes(ctx, tx, filepath.Join(g.csvDir, CardTypesFileName)); err != nil {
		return rollback(tx, fmt.Errorf("failed initCardTypes: %w", err))
	}
	if err = initCards(ctx, tx, filepath.Join(g.csvDir, CardsFileName)); err != nil {
		return rollback(tx, fmt.Errorf("failed initCards: %w", err))
	}

	return tx.Commit()
}

func rollback(tx *ent.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%w: %v", err, rerr)
	}
	return err
}

func truncate(ctx context.Context, client *ent.Client) error {
	if _, err := client.ExecContext(ctx, "set foreign_key_checks = 0"); err != nil {
		return err
	}
	for _, table := range migrate.Tables {
		query := fmt.Sprintf("TRUNCATE `%s`", table.Name)
		if _, err := client.ExecContext(ctx, query); err != nil {
			return err
		}
	}
	if _, err := client.ExecContext(ctx, "set foreign_key_checks = 1"); err != nil {
		return err
	}
	return nil
}
