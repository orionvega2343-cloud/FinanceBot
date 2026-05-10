CREATE TABLE IF NOT EXISTS categories(
                                         id SERIAL PRIMARY KEY,
                                         name TEXT NOT NULL,
                                         user_id INTEGER REFERENCES users(id),
    is_default BOOLEAN NOT NULL DEFAULT FALSE
    )

