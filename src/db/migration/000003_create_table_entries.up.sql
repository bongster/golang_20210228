CREATE TABLE IF NOT EXISTS entries
(
    id         SERIAL
        CONSTRAINT entries_pk
            PRIMARY KEY,
    user_id    INTEGER                                NOT NULL
        CONSTRAINT entries_users_id_fk
            REFERENCES users
            ON UPDATE CASCADE ON DELETE CASCADE,
    amount     BIGINT                                 NOT NULL,
    currency   VARCHAR(3)                             NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL
);