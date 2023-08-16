BEGIN;

CREATE TABLE courses (
    id UUID PRIMARY KEY,
    category TEXT NOT NULL CHECK (category IN ('web', 'pwn', 'rev', 'crypto')),
    difficulty TEXT NOT NULL CHECK (difficulty IN ('easy', 'difficult', 'advanced')),
    title TEXT NOT NULL,
    slug TEXT NOT NULL UNIQUE,
    description TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    publish BOOLEAN NOT NULL,
    updated_at TIMESTAMPTZ
);

CREATE TABLE course_authors (
    course_id UUID REFERENCES courses (id) ON DELETE CASCADE,
    author_id UUID REFERENCES authors (id) ON DELETE CASCADE,

    PRIMARY KEY (course_id, author_id)
);

CREATE TABLE course_items (
    id UUID PRIMARY KEY,
    position INT NOT NULL,
    course_id UUID NOT NULL REFERENCES courses (id) ON DELETE CASCADE,
    challenge_id UUID NOT NULL REFERENCES challenges (id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ
);

CREATE TABLE course_enrollments (
    id UUID PRIMARY KEY,
    course_id UUID NOT NULL REFERENCES courses (id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    finished BOOLEAN NOT NULL,
    finished_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    UNIQUE (user_id, course_id)
);

END;
