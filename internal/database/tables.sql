CREATE TABLE IF NOT EXISTS books (
    id SERIAL PRIMARY KEY,
    title CHAR(250) NOT NULL,
    description TEXT NOT NULL,
	genre CHAR(250)[] NOT NULL,
	authors CHAR(250)[] NOT NULL,
	publish_date DATE NOT NULL,
	publisher CHAR(250) NOT NULL,
	pages INTEGER NOT NULL,
	created_at TIMESTAMP DEFAULT NOW(),
	updated_at TIMESTAMP NULL,
	deleted_at TIMESTAMP NULL
);

