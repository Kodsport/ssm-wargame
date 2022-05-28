BEGIN;

CREATE TABLE schools (
    id INT PRIMARY KEY,
    name TEXT NOT NULL,
    geographical_area_code INT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ
);

CREATE TABLE users (
    id UUID PRIMARY KEY,
    discord_id TEXT UNIQUE,
    first_name TEXT,
    last_name TEXT,
    email TEXT NOT NULL,
    role TEXT NOT NULL,
    school_id INT REFERENCES schools(id) ON DELETE SET NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ
);

CREATE INDEX users_school_id_indx ON users ( school_id );

CREATE TABLE categories (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ
);

-- för att visa vilket event en chall kommer ifrån
CREATE TABLE ctf_events (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    url TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ
);

CREATE TABLE challenges (
    id UUID PRIMARY KEY,
    slug TEXT NOT NULL UNIQUE,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    score INT NOT NULL CHECK (score >= 0),
    published BOOLEAN NOT NULL,
    ctf_event_id UUID REFERENCES ctf_events(id) ON DELETE SET NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ
);

CREATE INDEX challenges_ctf_event_id_idx ON challenges ( ctf_event_id );

CREATE TABLE challenge_services (
    id UUID PRIMARY KEY,
    challenge_id UUID NOT NULL REFERENCES challenges(id) ON DELETE CASCADE,
    type TEXT NOT NULL,
    url TEXT,
    host TEXT,
    port INT,
    name TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ
);

CREATE INDEX challenge_services_challenge_id_idx ON challenge_services ( challenge_id );

CREATE TABLE challenge_files (
    id UUID PRIMARY KEY,
    challenge_id UUID REFERENCES challenges(id) ON DELETE SET NULL,
    friendly_name TEXT NOT NULL,
    bucket TEXT NOT NULL,
    key TEXT NOT NULL,
    md5 TEXT NOT NULL,
    uploaded BOOLEAN NOT NULL,
    size BIGINT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ,

    UNIQUE (bucket, key)
);

CREATE INDEX challenge_files_challenge_id_idx ON challenge_files ( challenge_id );

CREATE TABLE flags (
    id UUID PRIMARY KEY,
    challenge_id UUID NOT NULL REFERENCES challenges(id) ON DELETE CASCADE,
    flag TEXT NOT NULL UNIQUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ
);

CREATE INDEX flags_challenge_id_idx ON flags ( challenge_id );

CREATE TABLE monthly_challenges (
    challenge_id UUID PRIMARY KEY REFERENCES challenges(id) ON DELETE CASCADE,
    start_date TIMESTAMPTZ NOT NULL,
    end_date TIMESTAMPTZ NOT NULL,
    display_month TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ,

    CHECK (start_date < end_date)
);

CREATE TABLE submissions (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    challenge_id UUID NOT NULL REFERENCES challenges(id) ON DELETE CASCADE,
    successful BOOLEAN NOT NULL,
    input TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX submissions_challenge_id_idx ON submissions ( challenge_id );
CREATE INDEX submissions_user_id_idx ON submissions ( user_id );

CREATE TABLE challenge_authors (
    challenge_id UUID REFERENCES challenges(id) ON DELETE CASCADE,
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,

    PRIMARY KEY (challenge_id, user_id)
);

CREATE TABLE challenge_categories (
    challenge_id UUID REFERENCES challenges(id) ON DELETE CASCADE,
    category_id UUID REFERENCES categories(id) ON DELETE CASCADE,

    PRIMARY KEY (challenge_id, category_id)
);

CREATE TABLE user_solves (
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    challenge_id UUID REFERENCES challenges(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    PRIMARY KEY (user_id, challenge_id)
);

END;
