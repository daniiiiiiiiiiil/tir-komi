DROP INDEX IF EXISTS idx_media_type;
DROP INDEX IF EXISTS idx_media_sort_order;

DROP TABLE IF EXISTS media;

DROP TYPE IF EXISTS media_type;

DELETE FROM posts WHERE type = 'about';
ALTER TABLE posts DROP CONSTRAINT valid_type;
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
                                                            'international_cooperation'
    ));