class Difference:
    def __init__(self, a):
        self.__elements = a

    def computeDifference(self):
        self.__elements.sort()
        e = self.__elements[0] - self.__elements[len(self.__elements) - 1]
        self.maximumDifference = e if e > 0 else e * -1

_ = input()
a = [int(e) for e in input().split(' ')]

d = Difference(a)
d.computeDifference()

print(d.maximumDifference)
