CREATE TABLE scoreboards (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    createdAt TIMESTAMP DEFAULT NOW(),
    updatedAt TIMESTAMP DEFAULT NOW()
);