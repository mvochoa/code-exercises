import sys

cin = iter(sys.stdin.read().split())
m = int(next(cin))
n = int(next(cin))
k = int(next(cin))
rows = {}
cols = {}

for i in range(k):
    x = next(cin)
    y = next(cin)
    if x == 'R':
        if y in rows:
            rows.pop(y)
        else:
            rows[y] = 0
    else:
        if y in cols:
            cols.pop(y)
        else:
            cols[y] = 0

total = len(rows)*n-len(cols)*len(rows)
total += len(cols)*m-len(rows)*len(cols)

print(total)