#!/usr/bin/python3

class guitar:
    def __init__(self):
        self.color = ("#000000", "#FFFFFF")
        self.n_strings = 6
        self.__cost = 50
        self.quality = "Excellent"

    def Play(self):
        print("pam pam pam pam pam")
my_guitar = guitar()

my_guitar.Play()

class bassGuitar(guitar):
    pass

print(bassGuitar().n_strings)
print(my_guitar._guitar__cost,"only")
print(bassGuitar().color)
