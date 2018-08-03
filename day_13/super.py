from abc import ABCMeta, abstractmethod

class Book(object, metaclass=ABCMeta):
    def __init__(self,title,author):
        self.title=title
        self.author=author   

    @abstractmethod
    def display():
        pass

class MyBook(Book):
    def __init__(self, title, author, price):
        super(MyBook, self).__init__(title, author)
        self.price = price
        
    def display(self):
        ret = "Title: {}\n".format(self.title)
        ret += "Author: {}\n".format(self.author)
        ret += "Price: {}\n".format(self.price)
        print(ret)

title=input()
author=input()
price=int(input())
new_novel=MyBook(title,author,price)
new_novel.display()
