CREATE TYPE user_role AS ENUM ('user', 'admin');

CREATE TABLE advertisement
(
    id         INT GENERATED ALWAYS AS IDENTITY NOT NULL,
    title      VARCHAR                          NOT NULL,
    description VARCHAR,
    image      BYTEA,
    pdf        BYTEA,
    url        VARCHAR(250) CHECK ( CHAR_LENGTH(url) BETWEEN 1 AND 250),
    created_at TIMESTAMPTZ                      NOT NULL DEFAULT NOW(),

    CONSTRAINT PK_advertisement_id  PRIMARY KEY (id)
);

CREATE TABLE vacant_positions
(
    id          INT GENERATED ALWAYS AS IDENTITY NOT NULL,
    title       VARCHAR(200)                     NOT NULL CHECK ( CHAR_LENGTH(title) BETWEEN 1 AND 200),
    description VARCHAR,
    date        TIMESTAMPTZ,

    CONSTRAINT PK_vacant_positions_id PRIMARY KEY (id)
);

CREATE TABLE reviews
(
    id     INT GENERATED ALWAYS AS IDENTITY NOT NULL,
    name   VARCHAR(100)                     NOT NULL CHECK ( CHAR_LENGTH(name) BETWEEN 1 AND 100),
    email  VARCHAR(128) UNIQUE CHECK (email ~* '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$') ,
    description VARCHAR,
    rating INT NOT NULL CHECK (rating BETWEEN 1 AND 5),
    created_at  TIMESTAMPTZ     NOT NULL,

    CONSTRAINT PK_reviews_id PRIMARY KEY (id)
);

CREATE TABLE methodological_material
(
    id          INT GENERATED ALWAYS AS IDENTITY NOT NULL,
    title       VARCHAR(200)                     NOT NULL CHECK ( CHAR_LENGTH(title) BETWEEN 1 AND 200),
    description VARCHAR,
    date        TIMESTAMPTZ,
    pdf         BYTEA,

    CONSTRAINT PK_methodological_material_id PRIMARY KEY (id)
);

CREATE TABLE users
(
    id    INT GENERATED ALWAYS AS IDENTITY NOT NULL,
    email VARCHAR(128) NOT NULL UNIQUE CHECK (email ~* '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$') ,
    password    VARCHAR(256),
    role        user_role       NOT NULL DEFAULT 'user',
    created_at  TIMESTAMPTZ     NOT NULL,

    CONSTRAINT PK_users_id PRIMARY KEY (id)
);

CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_advertisement_created_at ON advertisement(created_at DESC);
CREATE INDEX idx_reviews_rating ON reviews(rating);
CREATE INDEX idx_vacant_positions_date ON vacant_positions(date DESC);
CREATE INDEX idx_methodological_material_date ON methodological_material(date DESC);