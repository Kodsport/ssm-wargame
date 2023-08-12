BEGIN;

CREATE TABLE schools (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    geographical_area_code TEXT NOT NULL,
    municipality_name TEXT NOT NULL,
    skolverket_id INT UNIQUE,
    raw_skolverket_data JSONB,
    is_high_school BOOLEAN NOT NULL,
    is_elementary_school BOOLEAN NOT NULL,
    is_university BOOLEAN NOT NULL,
    latitude FLOAT,
    longitude FLOAT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ
);

CREATE INDEX schools_name_idx ON schools ( name );

CREATE TABLE authors (
    id UUID PRIMARY KEY,
    slug TEXT NOT NULL UNIQUE,
    full_name TEXT NOT NULL,
    sponsor BOOLEAN NOT NULL,
    description TEXT NOT NULL,
    image_url TEXT,
    publish BOOLEAN NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ
);

CREATE TABLE users (
    id UUID PRIMARY KEY,
    discord_id TEXT UNIQUE,
    full_name TEXT NOT NULL,
    email TEXT NOT NULL,
    role TEXT NOT NULL,
    school_id UUID REFERENCES schools(id) ON DELETE SET NULL,
    author_id UUID UNIQUE REFERENCES authors(id) ON DELETE SET NULL,
    onboarding_done BOOLEAN NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ
);

CREATE INDEX users_school_id_indx ON users ( school_id );

CREATE TABLE categories (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
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
    publish_at TIMESTAMPTZ,
    ctf_event_id UUID REFERENCES ctf_events(id) ON DELETE SET NULL,
    category_id UUID NOT NULL REFERENCES categories(id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ
);

CREATE INDEX challenges_ctf_event_id_idx ON challenges ( ctf_event_id );
CREATE INDEX challenges_category_id_idx ON challenges ( category_id );

CREATE TABLE challenge_services (
    id UUID PRIMARY KEY,
    challenge_id UUID NOT NULL REFERENCES challenges(id) ON DELETE CASCADE,
    user_display TEXT NOT NULL,
    hyperlink BOOLEAN NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ
);

CREATE INDEX challenge_services_challenge_id_idx ON challenge_services ( challenge_id );

CREATE TABLE challenge_files (
    id UUID PRIMARY KEY,
    challenge_id UUID REFERENCES challenges(id) ON DELETE SET NULL,
    friendly_name TEXT NOT NULL,
    url TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ
);

CREATE INDEX challenge_files_challenge_id_idx ON challenge_files ( challenge_id );

CREATE TABLE flags (
    id UUID PRIMARY KEY,
    challenge_id UUID NOT NULL REFERENCES challenges(id) ON DELETE CASCADE,
    type TEXT NOT NULL,
    flag TEXT NOT NULL,
    flag_prefix TEXT NOT NULL,
    flag_suffix TEXT NOT NULL,
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
    author_id UUID REFERENCES authors(id) ON DELETE CASCADE,

    PRIMARY KEY (challenge_id, author_id)
);

CREATE TABLE user_solves (
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    challenge_id UUID REFERENCES challenges(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    PRIMARY KEY (user_id, challenge_id)
);

CREATE TABLE challtools_import_token (
    id UUID NOT NULL PRIMARY KEY,
    token TEXT NOT NULL,
    expires_at TIMESTAMPTZ NOT NULL,
    ctf_event_id UUID REFERENCES ctf_events(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);


END;
