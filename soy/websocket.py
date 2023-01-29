from core.token import get_token
from core.livesource import LiveSource
import core.db as db
from dotenv import load_dotenv
from classifier import process_tweet

def main():
    load_dotenv()
    connection = db.create_connection()
    token = get_token()
    ls = LiveSource(token)
    # https://developer.twitter.com/en/docs/twitter-api/tweets/filtered-stream/integrate/build-a-rule
    ls.add_rule("-is:retweet -from:elonmusk (Elon Musk OR #ElonMusk OR @elonmusk)", "Elon Musk")
    ls.start(process_tweet,connection)

if __name__ == "__main__":
    main()
