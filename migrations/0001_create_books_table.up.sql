CREATE TABLE IF NOT EXISTS books (
		id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
		title TEXT NOT NULL,
		author TEXT NOT NULL,
		year INT NOT NULL,
		publisher TEXT,
		description TEXT,
		UNIQUE(title, author, year)
	);