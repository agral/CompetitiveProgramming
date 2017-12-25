import math

def isprime(x):
    for k in range(2, int(math.sqrt(x))+2):
        if x % k == 0:
            return False
    return True


offset = 108100
primes_count = 0
for c in range(1000):
    candidate = offset + (17*c)
    if isprime(candidate):
        primes_count += 1

print(primes_count)
