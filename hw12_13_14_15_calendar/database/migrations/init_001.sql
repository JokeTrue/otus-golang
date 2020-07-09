CREATE TABLE IF NOT EXISTS events
(
    id              UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id         NUMERIC   NOT NULL,
    title           TEXT      NOT NULL,
    description     TEXT,
    start_date      TIMESTAMP NOT NULL,
    end_date        TIMESTAMP NOT NULL,
    notify_interval NUMERIC CHECK (notify_interval > 0)
);