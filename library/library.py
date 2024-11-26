import random
from book import Book



class Library:
    def __init__(self, n_shelves: int, m_books_per_shelf: int, topics: list[str]):
        self.n_shelves = n_shelves
        self.m_books_per_shelf = m_books_per_shelf
        self.topics = topics
        self.shelves = self._create_library()

    def _create_library(self):
        books = []
        for topic in self.topics:
            authors = [f"Author_{random.randint(1, 100)}" for _ in range(random.randint(5, 20))]
            authors.sort()  
            books.extend([Book(author, topic) for author in authors])

        random.shuffle(books)  

        shelves = []
        for i in range(0, len(books), self.m_books_per_shelf):
            shelves.append(books[i:i + self.m_books_per_shelf])
        return shelves

    def display(self):
        for i, shelf in enumerate(self.shelves, start=1):
            print(f"Shelf {i}: {shelf}")

    def shuffle(self):
        all_books = [book for shelf in self.shelves for book in shelf]
        random.shuffle(all_books)
        self.shelves = []
        for i in range(0, len(all_books), self.m_books_per_shelf):
            self.shelves.append(all_books[i:i + self.m_books_per_shelf])

    def sort(self):
        all_books = [book for shelf in self.shelves for book in shelf]
        # all_books.sort(key=lambda book: (book.topic, book.author))
        all_books = self.merge_sort(all_books)
        self.shelves = []
        for i in range(0, len(all_books), self.m_books_per_shelf):
            self.shelves.append(all_books[i:i + self.m_books_per_shelf])

    def merge_sort(self, books):
        if len(books) <= 1:
            return books

        mid = len(books) // 2
        left_half = self.merge_sort(books[:mid])
        right_half = self.merge_sort(books[mid:])

        return self._merge(left_half, right_half)
    
    def _merge(self, left, right):
        sorted_books = []
        i = j = 0

        while i < len(left) and j < len(right):
            if (left[i].topic, left[i].author) <= (right[j].topic, right[j].author):
                sorted_books.append(left[i])
                i += 1
            else:
                sorted_books.append(right[j])
                j += 1

        sorted_books.extend(left[i:])
        sorted_books.extend(right[j:])
        return sorted_books


