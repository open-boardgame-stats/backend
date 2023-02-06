package seed

import (
	"context"

	"github.com/open-boardgame-stats/backend/internal/ent"
)

func createUsers(ctx context.Context, tx *ent.Tx, player1, player2 *ent.Player) (*ent.User, *ent.User) {
	user1 := tx.User.Create().
		SetName("test-user-1").
		SetEmail("test-user-1@obgs.app").
		SetPassword("test-password").
		SetMainPlayer(player1).
		SaveX(ctx)

	user2 := tx.User.Create().
		SetName("test-user-2").
		SetEmail("test-user-2@obgs.app").
		SetPassword("test-password").
		SetMainPlayer(player2).
		SaveX(ctx)

	return user1, user2
}
