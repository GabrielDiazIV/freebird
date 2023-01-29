CREATE TABLE birds(
	id SERIAL PRIMARY KEY,
	name TEXT NOT NULL,
	entity_type INT NOT NULL,
	bird_fk INT NOT NULL REFERENCES birds(id),
	score FLOAT NOT NULL,
	n_positive INT NOT NULL,
	n_negative INT NOT NULL,
	img_url TEXT NOT NULL
);

CREATE TABLE tweets(
	id SERIAL PRIMARY KEY,
	body TEXT NOT NULL,
	bird_fk INT NOT NULL REFERENCES birds(id),
	author_name TEXT NOT NULL,
	author_username TEXT NOT NULL,
	post_time TIMESTAMP WITH TIME ZONE NOT NULL,
	score FLOAT NOT NULL
);

CREATE OR REPLACE FUNCTION notify_tweet() RETURNS TRIGGER AS $$

	DECLARE
		data json;

	BEGIN
		
		-- create row to json
		data = row_to_json(NEW);

		-- notify tweets stream
		PERFORM pg_notify('tweet_stream', data::text);

		RETURN NULL;
	END;

$$ LANGUAGE plpgsql;

CREATE TRIGGER tweets_notify_event
	AFTER UPDATE ON birds
	FOR EACH ROW EXECUTE PROCEDURE notify_tweet();

LISTEN tweet_stream;
