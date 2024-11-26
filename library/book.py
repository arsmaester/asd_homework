class Book:
    def __init__(self, author: str, topic: str):
        self.author = author
        self.topic = topic

    def __repr__(self):
        return f"(author='{self.author}', topic='{self.topic}')"