class Tweet:
    def __init__(self, id, body, author_name, author_username, created_at, rule_tag, score, certainty):
        self.id = id
        self.body = body
        self.author_name = author_name
        self.author_username = author_username
        self.post_time = created_at
        self.rule_tag = rule_tag
        self.score = score
        self.certainty = certainty
    def print(self):
        print("-----------")
        print("ID: " + self.id)
        print("Body: " + self.body)
        print("Author Name: " + self.author_name)
        print("Author Username: " + self.author_username)
        print("Created At: " + self.post_time)
        print("Matching Rule Tag: " + self.rule_tag)
        print("Sentiment Score: " + str(self.score))
        print("Sentiment Certainty: " + str(self.certainty))