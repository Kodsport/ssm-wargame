CREATE TABLE ctfs (
    id UUID NOT NULL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    slug TEXT NOT NULL,
    private BOOLEAN NOT NULL,
    password TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    start_time TIMESTAMPTZ NOT NULL,
    end_time TIMESTAMPTZ NOT NULL,
    UNIQUE (slug)
);

CREATE TABLE ctf_challenges (
    ctf_id UUID NOT NULL REFERENCES ctfs(id) ON DELETE CASCADE,
    challenge_id UUID NOT NULL REFERENCES challenges(id) ON DELETE CASCADE,
    custom_score INT,
    PRIMARY KEY (ctf_id, challenge_id)
);

CREATE TABLE ctf_users (
    id UUID NOT NULL PRIMARY KEY,
    ctf_id UUID NOT NULL REFERENCES ctfs(id) ON DELETE CASCADE,
    username TEXT NOT NULL,
    password TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE ctf_solves (
    id UUID NOT NULL PRIMARY KEY,
    ctf_id UUID NOT NULL REFERENCES ctfs(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES ctf_users(id) ON DELETE CASCADE,
    challenge_id UUID NOT NULL REFERENCES challenges(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);


CREATE TABLE challenge_groups (
    id UUID NOT NULL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE challenge_group_challenges (
    challenge_group_id UUID NOT NULL REFERENCES challenge_groups(id) ON DELETE CASCADE,
    challenge_id UUID NOT NULL REFERENCES challenges(id) ON DELETE CASCADE,
    custom_score INT,
    PRIMARY KEY (challenge_group_id, challenge_id)
);