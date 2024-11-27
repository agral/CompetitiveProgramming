def main():
    # prepare list of factorials:
    factorials = [1,]
    factorials[0] = 1
    factorial = 1
    counter = 0
    while(counter < 100):
        counter += 1
        factorial *= counter
        factorials.append(factorial)

    count = int(input())
    for i in range(count):
        fact = int(input())
        print(factorials[fact])

if(__name__ == '__main__'):
    main()
