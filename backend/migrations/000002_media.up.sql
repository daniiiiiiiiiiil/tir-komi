CREATE TYPE media_type AS ENUM ('photo', 'video');

CREATE TABLE media
(
    id          INT GENERATED ALWAYS AS IDENTITY NOT NULL,
    type        media_type                       NOT NULL,
    title       VARCHAR(200) CHECK ( CHAR_LENGTH(title) BETWEEN 1 AND 200 ),
    description VARCHAR,
    file_data   BYTEA                            NOT NULL,
    file_name   VARCHAR(255)                     NOT NULL,
    mime_type   VARCHAR(150)                     NOT NULL,
    file_size   BIGINT                           NOT NULL CHECK ( file_size > 0 ),
    sort_order  INT                              NOT NULL DEFAULT 0,
    created_at  TIMESTAMPTZ                      NOT NULL DEFAULT NOW(),

    CONSTRAINT PK_media_id PRIMARY KEY (id)
);

ALTER TABLE posts DROP CONSTRAINT IF EXISTS valid_type;
ALTER TABLE posts ADD CONSTRAINT valid_type CHECK (type IN (
                                                            'information',
                                                            'structures_and_bodies',
                                                            'document',
                                                            'education',
                                                            'educational_standards',
                                                            'management',
                                                            'materials',
                                                            'scholarships',
                                                            'paid_services',
                                                            'financial_and_economic',
                                                            'accessible_environment',
                                                            'international_cooperation',
                                                            'about'
    ));

INSERT INTO posts (title, description, type, date)
VALUES ('О нас', '[]', 'about', NOW())
    ON CONFLICT DO NOTHING;


CREATE INDEX idx_media_type ON media (type);
CREATE INDEX idx_media_sort_order ON media (sort_order);