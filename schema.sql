-- for users, we use UUID to increase security and privacy
-- the tradeoff is that TEXT comparisons instead of INTEGER comparisons of foreign keys in join tables are minimally slower

CREATE TABLE users (
    id TEXT PRIMARY KEY,
    email TEXT UNIQUE NOT NULL
    password_hash TEXT NOT NULL
)

-- DATETIME in sessions table only signals intent, still takes any TEXT
-- YYYY-MM-DD HH:MM:SS time syntax is not enforced on db level, just suggested through naming

CREATE TABLE sessions (
    id TEXT PRIMARY KEY,
    user_id INTEGER NOT NULL,
    expires_at DATETIME NOT NULL
    FOREIGN KEY (user_id) REFERENCES users(id)
)
