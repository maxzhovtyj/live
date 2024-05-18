CREATE TABLE IF NOT EXISTS users
(
    id            SERIAL PRIMARY KEY,
    first_name    VARCHAR(128) NOT NULL,
    last_name     VARCHAR(128) NOT NULL,
    email         VARCHAR(128) NOT NULL UNIQUE,
    password_hash VARCHAR(128) NOT NULL
);

CREATE TABLE conversations
(
    id         SERIAL PRIMARY KEY,
    name       VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE messages
(
    id              SERIAL PRIMARY KEY,
    conversation_id INT  NOT NULL,
    sender_id       INT  NOT NULL,
    body            TEXT NOT NULL,
    created_at      TIMESTAMPTZ DEFAULT NOW(),
    FOREIGN KEY (conversation_id) REFERENCES conversations (id),
    FOREIGN KEY (sender_id) REFERENCES users (id)
);

CREATE TABLE conversation_participants
(
    conversation_id INT NOT NULL,
    user_id         INT NOT NULL,
    PRIMARY KEY (conversation_id, user_id),
    FOREIGN KEY (conversation_id) REFERENCES conversations (id),
    FOREIGN KEY (user_id) REFERENCES users (id)
);
