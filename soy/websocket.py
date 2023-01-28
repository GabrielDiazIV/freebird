from core.token import get_token
from core.livesource import LiveSource
from dotenv import load_dotenv

def main():

    load_dotenv()
    token = get_token()
    ls = LiveSource(token)
    # https://developer.twitter.com/en/docs/twitter-api/tweets/filtered-stream/integrate/build-a-rule
    ls.add_rule("cats -dogs")
    ls.stream()

if __name__ == "__main__":
    main()
