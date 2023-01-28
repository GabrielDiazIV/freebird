class SourceMeta(type):
    def __instancecheck__(cls, instance):
        return cls.__subclasscheck__(type(instance))

    def __subclasscheck__(cls, subclass):
        return hasattr(subclass, 'start') and callable(subclass.start)

class Source(metaclass=SourceMeta):
    def start(self, classify_func):
        raise Exception(f"classify not implemented by {type(self).mro()}")

