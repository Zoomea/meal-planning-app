def f(n):
    c = 0
    i = 0
    while n >= 0:
        i += 1
        n -= 2
        c += n - 2
        print(f"c={c} n={n} i={i}")
    return c


for i in range(100):
    print(f"n={i} f({i})={f(i)}")
