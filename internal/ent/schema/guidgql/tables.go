package guidgql

type Table int

const (
	User Table = iota
	Player
	PlayerSupervisionRequest
	PlayerSupervisionRequestApproval
	Group
	GroupMembership
	GroupMembershipApplication
	GroupSettings
	Game
	GameFavorites
	StatDescription
	Match
	Statistic
	GameVersion
)

var TableNames = map[Table]string{
	User:                             "users",
	Player:                           "players",
	PlayerSupervisionRequest:         "player_supervision_requests",
	PlayerSupervisionRequestApproval: "player_supervision_request_approvals",
	Group:                            "groups",
	GroupMembership:                  "group_memberships",
	GroupMembershipApplication:       "group_membership_applications",
	GroupSettings:                    "group_settings",
	Game:                             "games",
	GameFavorites:                    "game_favorites",
	StatDescription:                  "stat_descriptions",
	Match:                            "matches",
	Statistic:                        "statistics",
	GameVersion:                      "game_versions",
}
