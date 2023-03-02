package initdb

import (
	"context"
	_ "embed"

	"github.com/AgricolaDevJP/agricoladb-server/ent"

	"github.com/jszwec/csvutil"
)

//go:embed "masterdata/agricoladb-master - revisions.csv"
var RevisonsBytes []byte

const (
	RevisionsCapacity = 8
)

type Revision struct {
	ID     int    `csv:"id"`
	Key    string `csv:"key"`
	NameJa string `csv:"name_ja"`
	NameEn string `csv:"name_en"`
}

func initRevisions(ctx context.Context, tx *ent.Tx) error {
	revisions := []*Revision{}
	if err := csvutil.Unmarshal(RevisonsBytes, &revisions); err != nil {
		return err
	}

	builders := make([](*ent.RevisionCreate), 0, RevisionsCapacity)
	for _, revision := range revisions {
		builder := tx.Revision.Create().SetID(revision.ID).SetKey(revision.Key).SetNameJa(revision.NameJa).SetNameEn(revision.NameEn)
		builders = append(builders, builder)
	}

	return tx.Revision.CreateBulk(builders...).OnConflict().UpdateNewValues().Exec(ctx)
}
