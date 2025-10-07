CREATE INDEX ctf_solves_ctf_challenge_idx ON ctf_solves (ctf_id, challenge_id);

CREATE INDEX ctf_solves_ctf_challenge_time_idx ON ctf_solves (ctf_id, challenge_id, created_at);

CREATE INDEX ctf_solves_team_time_idx ON ctf_solves (ctf_id, challenge_id, team_id, created_at) WHERE team_id IS NOT NULL;

CREATE INDEX ctf_solves_ctf_user_idx ON ctf_solves (ctf_id, user_id);

CREATE INDEX ctf_solves_ctf_team_challenge_idx ON ctf_solves (ctf_id, team_id, challenge_id, created_at) WHERE team_id IS NOT NULL;