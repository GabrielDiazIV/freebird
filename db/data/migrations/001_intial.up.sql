CREATE TABLE birds(
	id SERIAL PRIMARY KEY,
	name TEXT,
	entity_type INT,
	bird_fk INT REFERENCES birds(id),
	score FLOAT,
	n_positive INT,
	n_negative INT,
	img_url TEXT
);

CREATE TABLE tweets(
	body TEXT,
	bird_fk INT REFERENCES birds(id),
	author_name TEXT,
	author_username TEXT,
	post_time DATETIME,
	score FLOAT,
	certainty FLOAT
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
	AFTER INSERT ON tweets
	FOR EACH ROW EXECUTE PROCEDURE notify_tweet();

LISTEN tweet_stream;
