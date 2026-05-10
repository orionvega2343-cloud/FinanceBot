CREATE TABLE IF NOT EXISTS transactions(
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    type VARCHAR,
    category_id INTEGER  REFERENCES categories(id),
    sum DECIMAL,
    comment TEXT,
    created_at TIMESTAMP DEFAULT NOW()

)