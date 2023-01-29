from textblob import TextBlob
from core.tweet import Tweet
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


def classify_tweet(tweet):
    tweet_text = tweet['data']['text']
    blob = TextBlob(clean_tweet(tweet_text))
    # create tweet object 
    tweet_object = Tweet(tweet['data']['id'],tweet['data']['text'],tweet['includes']['users'][0]['name'],tweet['includes']['users'][0]['username'],tweet['data']['created_at'],tweet['matching_rules'][0]['tag'],blob.polarity,blob.subjectivity)
    tweet_object.print()
    
