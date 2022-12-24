package initdata

import (
	"agricoladb/ent"
	"context"
	"os"

	"github.com/jszwec/csvutil"
)

const (
	RevisionsFileName = "agricoladb-master - card_special_colors.csv"
	RevisionsCapacity = 8
)

type Revision struct {
	ID     int    `csv:"id"`
	Key    string `csv:"key"`
	NameJa string `csv:"name_ja"`
	NameEn string `csv:"name_en"`
}

func initRevisions(ctx context.Context, tx *ent.Tx, csvFilePath string) error {
	bytes, err := os.ReadFile(csvFilePath)
	if err != nil {
		return err
	}

	revisions := []*Revision{}
	if err := csvutil.Unmarshal(bytes, &revisions); err != nil {
		return err
	}

	builders := make([](*ent.RevisionCreate), 0, RevisionsCapacity)
	for _, revision := range revisions {
		builder := tx.Revision.Create().SetID(revision.ID).SetKey(revision.Key).SetNameJa(revision.NameJa).SetNameEn(revision.NameEn)
		builders = append(builders, builder)
	}

	return tx.Revision.CreateBulk(builders...).OnConflict().UpdateNewValues().Exec(ctx)
}
