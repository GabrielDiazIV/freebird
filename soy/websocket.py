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
    ls.add_rule("-is:retweet -from:elonmusk (\"Elon Musk\" OR #ElonMusk OR @elonmusk)", "Elon Musk")
    ls.add_rule("-is:retweet -from:SBF_FTX (\"Sam Bankman-Fried\" OR @SBF_FTX)", "Sam Bankman-Fried")
    ls.add_rule("-is:retweet (\"Kanye West\" OR #KanyeWest)", "Kanye West")
    ls.add_rule("-is:retweet -from:cobratate (\"Andrew Tate\" OR #AndrewTate OR @cobratate)", "Andrew Tate")
    ls.add_rule("-is:retweet (\"Mark Zuckerberg\" OR #MarkZuckerberg)", "Mark Zuckerberg")
    ls.add_rule("-is:retweet -from:JeffBezos (\"Jeff Bezos\" OR #JeffBezos OR @JeffBezos)", "Jeff Bezos")
    ls.add_rule("-is:retweet -from:jack (\"Jack Dorsey\" OR #JackDorsey OR @jack)", "Jack Dorsey")
    ls.add_rule("-is:retweet -from:tim_cook (\"Tim Cook\" OR #TimCook OR @tim_cook)", "Tim Cook")
    ls.add_rule("-is:retweet -from:lisasu (\"Lisa Su\" OR #LisaSu OR @lisasu)", "Lisa Su")
    ls.add_rule("-is:retweet -from:sama (\"Sam Altman\" OR #SamAltman OR @sama)", "Sam Altman")
    ls.add_rule("-is:retweet -from:sundarpichai (\"Sundar Pichai\" OR #SundarPichai OR @sundarpichai)", "Sundar Pichai")
    ls.add_rule("-is:retweet -from:BillGates (\"Bill Gates\" OR #BillGates OR @BillGates)", "Bill Gates")
    ls.add_rule("-is:retweet -from:joerogan (\"Joe Rogan\" OR #JoeRogan OR @joerogan)", "Joe Rogan")
    ls.add_rule("-is:retweet -from:Tesla (\"Tesla\" OR #Tesla OR @Tesla)", "Tesla")
    ls.add_rule("-is:retweet -from:FTX_Official (\"FTX\" OR #FTX OR @FTX_Official)", "FTX")
    ls.add_rule("-is:retweet -from:Meta (\"Meta\" OR #Meta OR @Meta)", "Meta")
    ls.add_rule("-is:retweet -from:amazon (\"Amazon\" OR #Amazon OR @Amazon)", "Amazon")
    ls.add_rule("-is:retweet -from:Apple (\"Apple\" OR #Apple OR @Apple)", "Apple")
    ls.add_rule("-is:retweet -from:AMD (\"AMD\" OR #AMD OR @AMD)", "AMD")
    ls.add_rule("-is:retweet -from:OpenAI (\"OpenAI\" OR #OpenAI OR @OpenAI)", "OpenAI")
    ls.add_rule("-is:retweet -from:Google (\"Google\" OR #Google OR @Google)", "Google")
    ls.add_rule("-is:retweet -from:Microsoft (\"Microsoft\" OR #Microsoft OR @Microsoft)", "Microsoft")
    ls.add_rule("-is:retweet -from:goblinx64 (\"Daniel Kasabov\" OR #DanielKasabov OR @goblinx64)", "Daniel Kasabov")
    ls.start(process_tweet,connection)

if __name__ == "__main__":
    main()
