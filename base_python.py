def number2base(n, base):
    digs = "0123456789abcdefghijklmnopqrstuvwxyz"
    if n == 0:
        return digs[0]
    digits = []
    while n:
        digits.insert(0, digs[n % base])
        n = n // base
    return "".join(digits)


def palindrome_sum(limit):
    count = 0
    total = 0
    for i in range(0, limit):
        base22 = number2base(i, 22)
        if base22 == base22[::-1]:  # [::-1] is array/string reversal magic
            count += 1
            total += i

    return (count, total)


limit = 1_000_000_000
(count, total) = palindrome_sum(limit)
print(f"Count: {count} or {number2base(count, 22)} in base22")
print(f"Sum:   {total} or {number2base(total, 22)} in base22")
print(f"Limit: {limit}")
