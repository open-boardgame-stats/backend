-- create "player_supervision_requests" table
CREATE TABLE "player_supervision_requests" ("id" uuid NOT NULL, "message" character varying NULL, "player_supervision_requests" uuid NOT NULL, "user_sent_supervision_requests" uuid NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "player_supervision_requests_players_supervision_requests" FOREIGN KEY ("player_supervision_requests") REFERENCES "players" ("id") ON DELETE NO ACTION, CONSTRAINT "player_supervision_requests_users_sent_supervision_requests" FOREIGN KEY ("user_sent_supervision_requests") REFERENCES "users" ("id") ON DELETE NO ACTION);
-- create "player_supervision_request_approvals" table
CREATE TABLE "player_supervision_request_approvals" ("id" uuid NOT NULL, "approved" boolean NULL, "player_supervision_request_approvals" uuid NOT NULL, "user_supervision_request_approvals" uuid NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "player_supervision_request_app_fba9101bfc08fa71b712363a55488bd4" FOREIGN KEY ("player_supervision_request_approvals") REFERENCES "player_supervision_requests" ("id") ON DELETE NO ACTION, CONSTRAINT "player_supervision_request_app_12b494b4b73bac2c5e47619d7c7fd2f0" FOREIGN KEY ("user_supervision_request_approvals") REFERENCES "users" ("id") ON DELETE NO ACTION);