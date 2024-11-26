from library import Library

n_shelves=5 
m_books_per_shelf=10
topics=["Science", "Fiction", "History"]


library = Library(n_shelves, m_books_per_shelf, topics)
print(f"Initial library:")
library.display()

print("\nShuffled library:")
library.shuffle()
library.display()

print("\nSorted library:")
library.sort()
library.display()