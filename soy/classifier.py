from textblob import TextBlob
from core.tweet import Tweet
import core.db as db
import re

def clean_tweet(tweet):
    # put tweet in lowercase
    temp = tweet.lower()
    # regex to remove mentions
    temp = re.sub("@[A-Za-z0-9_]+","",temp)
    # regex to remove hashtags
    temp = re.sub("#[A-Za-z0-9_]+","",temp)
    # regex to remove links
    temp = re.sub(r"http\S+", "", temp)
    temp = re.sub(r"www.\S+", "", temp)
    return temp


def process_tweet(tweet, connection):
    tweet_text = tweet['data']['text']
    blob = TextBlob(clean_tweet(tweet_text))
    if (blob.polarity == 0.0):
        return
    # create tweet object 
    tweet_object = Tweet(tweet['data']['id'],tweet['data']['text'],tweet['includes']['users'][0]['name'],tweet['includes']['users'][0]['username'],tweet['data']['created_at'],tweet['matching_rules'][0]['tag'],blob.polarity,blob.subjectivity)
    tweet_object.print()
    # get bird from rule tag
    bird = db.query(f"SELECT * FROM birds WHERE name = '{tweet_object.rule_tag}';",connection)[0]
    # sanitize single quotes from tweet body
    tweet_object.body = tweet_object.body.replace("'","''")
    tweet_object.author_name = tweet_object.author_name.replace("'","''")
    # insert tweet into tweets table    
    db.update(f"INSERT INTO tweets(body,bird_fk,author_name,author_username,post_time,score,certainty) VALUES ('{tweet_object.body}', '{bird['id']}', '{tweet_object.author_name}', '{tweet_object.author_username}', '{tweet_object.post_time}',{tweet_object.score},{tweet_object.certainty});", connection)
    # increment positive/negative sums
    if tweet_object.score > 0.0:
        db.update(f"UPDATE birds SET n_positive = n_positive + 1 WHERE name = '{bird['name']}';",connection)
    elif tweet_object.score < 0.0:
        db.update(f"UPDATE birds SET n_negative = n_negative + 1 WHERE name = '{bird['name']}';",connection)
    #increment total sentiment score
    db.update(f"UPDATE birds SET score = score + {tweet_object.score} WHERE name = '{bird['name']}';",connection)
    
