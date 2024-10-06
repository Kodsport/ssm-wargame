BEGIN;

CREATE TABLE knack_koden_teams (
    id UUID PRIMARY KEY,
    teacher_full_name TEXT NOT NULL,
    teacher_email TEXT NOT NULL,
    teacher_phonenr TEXT NOT NULL,
    school_name TEXT NOT NULL,
    class_name TEXT NOT NULL,
    postal_code TEXT NOT NULL,
    password TEXT NOT NULL UNIQUE,

    ip_addr TEXT NOT NULL,
    user_agent TEXT NOT NULL, -- this is purely an anti-spam measure

    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ
);

CREATE TABLE knack_koden_solves (
    challenge_id UUID NOT NULL REFERENCES challenges(id),
    knack_koden_team_id UUID NOT NULL REFERENCES knack_koden_teams(id),

    created_at TIMESTAMPTZ NOT NULL,

    PRIMARY KEY (challenge_id, knack_koden_team_id)
);


END;
