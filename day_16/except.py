#!/bin/python3


data = input().strip()

try:
    print(int(data))
except ValueError:
    print('Bad String')
