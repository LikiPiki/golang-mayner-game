CREATE TABLE IF NOT EXISTS users (
	id INTEGER PRIMARY KEY NOT NULL,
	name TEXT,
	mayner1 INTEGER,
	mayner2 INTEGER, 
	mayner3 INTEGER,
	mayner4 INTEGER,
	score LONG, 
	money INTEGER,
	time INTEGER
	user_id INTEGER
	active INTEGER
);

CREATE TABLE IF NOT EXISTS value (
	id INTEGER PRIMARY KEY NOT NULL,
	name TEXT,
	cost INTEGER
);

INSERT INTO value (name, cost) VALUES ("Bitcoin", 1);
INSERT INTO value (name, cost) VALUES ("Ethereum", 1);
INSERT INTO value (name, cost) VALUES ("Bitcoin cash", 0);

