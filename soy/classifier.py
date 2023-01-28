from textblob import TextBlob
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
    print(clean_tweet(tweet_text))
    print(blob.sentiment.polarity)
    print(blob.sentiment.subjectivity)
