// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"

	"github.com/open-boardgame-stats/backend/internal/ent"
	"github.com/open-boardgame-stats/backend/internal/ent/enums"
	"github.com/open-boardgame-stats/backend/internal/ent/groupsettings"
	"github.com/open-boardgame-stats/backend/internal/ent/schema/guidgql"
	"github.com/open-boardgame-stats/backend/internal/ent/schema/stat"
)

// This type is exposed for type safety on client side
type AggregateMetadata struct {
	Type    AggregateMetadataType `json:"type"`
	StatIds []*guidgql.GUID       `json:"statIds"`
}

type AggregateMetadataInput struct {
	Type             AggregateMetadataType `json:"type"`
	StatOrderNumbers []int                 `json:"statOrderNumbers"`
}

type CreateGameInput struct {
	Name             string                  `json:"name"`
	MinPlayers       int                     `json:"minPlayers"`
	MaxPlayers       int                     `json:"maxPlayers"`
	Description      *string                 `json:"description,omitempty"`
	BoardgamegeekURL *string                 `json:"boardgamegeekURL,omitempty"`
	StatDescriptions []*StatDescriptionInput `json:"statDescriptions"`
}

type CreateMatchInput struct {
	GameVersionID guidgql.GUID    `json:"gameVersionId"`
	PlayerIds     []*guidgql.GUID `json:"playerIds"`
	Stats         []*StatInput    `json:"stats"`
}

type CreateOrUpdateGroupInput struct {
	ID          *guidgql.GUID       `json:"id,omitempty"`
	Name        string              `json:"name"`
	Description *string             `json:"description,omitempty"`
	LogoURL     string              `json:"logoUrl"`
	Settings    *GroupSettingsInput `json:"settings"`
}

type CreatePlayerInput struct {
	Name string `json:"name"`
}

// This type is exposed for type safety on client side
type EnumMetadata struct {
	PossibleValues []string `json:"possibleValues"`
}

type EnumMetadataInput struct {
	PossibleValues []string `json:"possibleValues"`
}

type Favorites struct {
	Total int         `json:"total"`
	Users []*ent.User `json:"users"`
}

type GroupApplicationInput struct {
	GroupID guidgql.GUID `json:"groupId"`
	Message *string      `json:"message,omitempty"`
}

type GroupSettingsInput struct {
	Visibility          groupsettings.Visibility `json:"visibility"`
	JoinPolicy          groupsettings.JoinPolicy `json:"joinPolicy"`
	MinimumRoleToInvite *enums.Role              `json:"minimumRoleToInvite,omitempty"`
}

type Header struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type RequestPlayerSupervisionInput struct {
	PlayerID guidgql.GUID `json:"playerId"`
	Message  *string      `json:"message,omitempty"`
}

type ResolvePlayerSupervisionRequestInput struct {
	RequestID guidgql.GUID `json:"requestId"`
	Approved  bool         `json:"approved"`
}

type StatDescriptionInput struct {
	Type        stat.StatType      `json:"type"`
	Name        string             `json:"name"`
	Description *string            `json:"description,omitempty"`
	Metadata    *StatMetadataInput `json:"metadata,omitempty"`
	OrderNumber int                `json:"orderNumber"`
}

type StatInput struct {
	// The StatDescription ID of the stat to be created
	StatID   guidgql.GUID `json:"statId"`
	Value    string       `json:"value"`
	PlayerID guidgql.GUID `json:"playerId"`
}

type StatMetadataInput struct {
	// Once input unions are in graphql, this will be one
	EnumMetadata      *EnumMetadataInput      `json:"enumMetadata,omitempty"`
	AggregateMetadata *AggregateMetadataInput `json:"aggregateMetadata,omitempty"`
}

type UploadURL struct {
	URL     string    `json:"url"`
	Headers []*Header `json:"headers"`
}

type AggregateMetadataType string

const (
	// Sum of all values
	AggregateMetadataTypeSum AggregateMetadataType = "sum"
)

var AllAggregateMetadataType = []AggregateMetadataType{
	AggregateMetadataTypeSum,
}

func (e AggregateMetadataType) IsValid() bool {
	switch e {
	case AggregateMetadataTypeSum:
		return true
	}
	return false
}

func (e AggregateMetadataType) String() string {
	return string(e)
}

func (e *AggregateMetadataType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = AggregateMetadataType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid AggregateMetadataType", str)
	}
	return nil
}

func (e AggregateMetadataType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
