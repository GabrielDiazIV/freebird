from core.token import get_token
from core.livesource import LiveSource
from dotenv import load_dotenv
from classifier import classify_tweet

def main():
    load_dotenv()
    token = get_token()
    ls = LiveSource(token)
    # https://developer.twitter.com/en/docs/twitter-api/tweets/filtered-stream/integrate/build-a-rule
    ls.add_rule("ryan kutz")
    ls.start(classify_tweet)

if __name__ == "__main__":
    main()
