package mysql

import (
	"agricoladb/repository"
	"database/sql"
)

type MysqlRepository struct {
	cards []interface{}
}

type CardRecord struct {
	ID                              uint           `db:"id"`
	LiteralID                       string         `db:"literal_id"`
	PrintedID                       sql.NullString `db:"printed_id"`
	RevisionID                      int            `db:"revision_id"`
	PlayAgricolaCardID              sql.NullInt64  `db:"play_agricola_card_id"`
	CardTypeID                      uint           `db:"card_type_id"`
	DeckID                          uint           `db:"deck_id"`
	NameJa                          string         `db:"name_ja"`
	NameEn                          string         `db:"name_en"`
	MinPlayersNumber                sql.NullInt64  `db:"min_players_number"`
	Prerequisite                    sql.NullString `db:"prerequisite"`
	Cost                            sql.NullString `db:"cost"`
	Description                     string         `db:"description"`
	IsOfficialJa                    bool           `db:"is_official_ja"`
	VictoryPoint                    sql.NullInt64  `db:"victory_point"`
	IsMutableVictoryPoint           bool           `db:"is_mutable_victory_point"`
	HasArrow                        bool           `db:"has_arrow"`
	HasBonus                        bool           `db:"has_bonus"`
	HasNegativeBonus                bool           `db:"has_negative_bonus"`
	HasPanIcon                      bool           `db:"has_pan_icon"`
	HasBreadIcon                    bool           `db:"has_bread_icon"`
	HasFarmPlannerIcon              bool           `db:"has_farm_planner_icon"`
	HasActionsBoosterIcon           bool           `db:"has_actions_booster_icon"`
	HasPointsProviderIcon           bool           `db:"has_points_provider_icon"`
	HasGoodsProviderIcon            bool           `db:"has_goods_provider_icon"`
	HasFoodProviderIcon             bool           `db:"has_food_provider_icon"`
	HasCropProviderIcon             bool           `db:"has_crop_provider_icon"`
	HasBuildingResourceProviderIcon bool           `db:"has_building_resource_provider_icon"`
	HasLivestockProviderIcon        bool           `db:"has_livestock_provider_icon"`
	HasCutPeatIcon                  bool           `db:"has_cut_peat_icon"`
	HasFellTreesIcon                bool           `db:"has_fell_trees_icon"`
	HasSlashAndBurnIcon             bool           `db:"has_slash_and_burn_icon"`
	HasHiringFareIcon               bool           `db:"has_hiring_fare_icon"`
	SpecialColor                    sql.NullInt64  `db:"special_color_id"`
}

type CardProductRecord struct {
}

// typecheck
var _ repository.Repository = &MysqlRepository{}

func NewMysqlRepository() *MysqlRepository {
	return &MysqlRepository{}
}

func (r *MysqlRepository) GetCards() {}
