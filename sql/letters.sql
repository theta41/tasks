CREATE TABLE letters
(
    id          SERIAL       NOT NULL
        CONSTRAINT letters_pk
            PRIMARY KEY,
    email       VARCHAR(320) NOT NULL,
    "order"     INT          NOT NULL,
    task_id     INT          NOT NULL
        CONSTRAINT letters_tasks_id_fk
            REFERENCES tasks
            ON UPDATE CASCADE ON DELETE CASCADE,
    sent        BOOLEAN      NOT NULL,
    answered    BOOLEAN      NOT NULL,
    accepted    BOOLEAN      NOT NULL,
    accept_uuid uuid         NOT NULL,
    accepted_at INT,
    sent_at     INT          NOT NULL
);