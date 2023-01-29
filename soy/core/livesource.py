import requests
import json

class LiveSource:
    def __init__(self, token):
        self._bearer = token
        self.__delete_rules(self.__rules())
        self._rules = []
        pass

    def __auth(self):
        return { "Authorization": f"Bearer {self._bearer}" }

    def stream(self):
        params = {'expansions' : 'author_id', 'tweet.fields' : 'text,created_at'}
        res = requests.get("https://api.twitter.com/2/tweets/search/stream", headers=self.__auth(), stream=True, params=params)
        if res.status_code != 200:
            raise Exception(f"stream did not connect: {res.text}")
        return res

    def __rules(self):
        res = requests.get("https://api.twitter.com/2/tweets/search/stream/rules", headers=self.__auth())
        if res.status_code != 200:
            raise Exception(f"failed to get rules: {res.text}")
        res_json = res.json()
        if res_json["meta"]["result_count"] == 0:
            return []
        return [ rule["id"] for rule in res_json["data"]]

    def __delete_rules(self, ids):
        if len(ids) == 0:
            return
        delete_payload = {
                "delete": {
                    "ids": ids
                }
        }
        res = requests.post("https://api.twitter.com/2/tweets/search/stream/rules", json=delete_payload, headers=self.__auth())
        if res.status_code != 200:
            raise Exception(f"failed to delete previous rules: {res.text}")

    def __submit_rules(self):
        if len(self._rules) == 0:
            raise Exception("there must be at least one rule to enable streaming")
        
        rules_payload = {
                "add": self._rules,
        }
        res = requests.post("https://api.twitter.com/2/tweets/search/stream/rules", json=rules_payload, headers=self.__auth())
        print(res.status_code)
        if res.status_code != 201:
            raise Exception(f"failed to upload new rules: {res.text}")

    def add_rule(self, rule, tag):
        if len(self._rules) >= 25:
            raise Exception("a maximum of 25 rules may exist on a stream")
        if len(rule) > 512:
            raise Exception("rules must not exceed 512 characters")
        if not isinstance(rule, str):
            raise Exception("rules must be strings")

        self._rules.append({ "value": rule, "tag": tag})

    def start(self, classify, connection):
        self.__submit_rules()
        print(f"a total of {len(self._rules)} were submitted")

        stream = self.stream()
        for res in stream.iter_lines():
            if res:
                res_json = json.loads(res)
                classify(res_json, connection)


