CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        first_name TEXT,
        last_name TEXT
);

CREATE TABLE IF NOT EXISTS bankaccounts (
        id SERIAL PRIMARY KEY,
        user_id INT FOREIGN KEY REFERRENCES users(user_id),
        account_number INT UNIQUE,
        name TEXT,
        balance INT
);

CREATE TABLE IF NOT EXISTS keys (
        id SERIAL PRIMARY KEY,
        key TEXT
);
