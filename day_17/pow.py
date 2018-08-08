class Calculator(object):
    def power(self, n, p):
        if p < 0 or n < 0:
            raise Exception('n and p should be non-negative')
        ret = 1
        for i in range(0, p):
            ret *= n
        return ret

myCalculator=Calculator()
T=int(input())
for i in range(T):
    n,p = map(int, input().split())
    try:
        ans=myCalculator.power(n,p)
        print(ans)
    except Exception as e:
        print(e)
