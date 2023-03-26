package initdb

import (
	"context"
	_ "embed"

	"github.com/AgricolaDevJP/agricoladb-server/ent"

	"github.com/jszwec/csvutil"
)

//go:embed "masterdata/agricoladb-master - products.csv"
var productsBytes []byte

const (
	ProductsCapacity = 64
)

type Product struct {
	ID            int    `csv:"id"`
	IsOfficialJa  bool   `csv:"is_official_ja"`
	NameJa        string `csv:"name_ja"`
	NameEn        string `csv:"name_en"`
	PublishedYear *int   `csv:"published_year"`
	RevisionID    int    `csv:"revision_id"`
}

func initProducts(ctx context.Context, tx *ent.Tx) error {
	products := []*Product{}
	if err := csvutil.Unmarshal(productsBytes, &products); err != nil {
		return err
	}

	builders := make([](*ent.ProductCreate), 0, ProductsCapacity)
	for _, product := range products {
		builder := tx.Product.Create().SetID(product.ID).SetIsOfficialJa(product.IsOfficialJa).SetNameJa(product.NameJa).SetNameEn(product.NameEn).SetNillablePublishedYear(product.PublishedYear).SetRevisionID(product.RevisionID)
		builders = append(builders, builder)
	}

	return tx.Product.CreateBulk(builders...).OnConflict().UpdateNewValues().Exec(ctx)
}
