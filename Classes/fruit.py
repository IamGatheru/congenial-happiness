#!/usr/bin/python3

class Fruit:
    def __init__(self):
        self.name = "kiwi"
        self.color = "#000000"
        self.texture = "rough"
        self.taste = "sour"

    def Eat(self):
        print(f"my fruit is {self.taste}")

my_fruit = Fruit()

class Watermelon(Fruit):
    def __init__(self):
        super().__init__()
        self.name = "watermelon"
        self.texture = "smooth"
        self.taste = "sweet"

Watermelon().Eat()
