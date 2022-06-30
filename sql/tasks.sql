CREATE TABLE tasks
(
    id            SERIAL       NOT NULL
        CONSTRAINT tasks_pk
            PRIMARY KEY,
    name          VARCHAR(64)  NOT NULL,
    description   VARCHAR(256) NOT NULL,
    creator_email VARCHAR(320) NOT NULL,
    created_at    INT          NOT NULL,
    finished_at   INT          NOT NULL
);
