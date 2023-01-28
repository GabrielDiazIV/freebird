from metasource import Source

# class for ryan kutz's sentiment analysis warlord code
class Wrapper:
    def __init__(self, source):
        if not issubclass(source, Source):
            raise Exception("source object does not inherit from Source metaclass")
        self._source = source
