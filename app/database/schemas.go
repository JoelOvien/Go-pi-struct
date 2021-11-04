package database

const createSchema = `
CREATE TABLE IF NOT EXISTS posts
{
	id SERIAL PRIMARY KEY,
	title TEXT.
	content TEXT,
	authort TEXT,
}
`