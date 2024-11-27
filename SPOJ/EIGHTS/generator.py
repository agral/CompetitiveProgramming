def main():
    f = open('cubes.txt', 'w')
    f888 = open('cubes_888.txt', 'w')

    for c in range(1000000):
        cube = c * c * c
        f.write(repr(c).rjust(6) + " --> " + str(cube) + "\n")
        if(cube % 1000 == 888):
            f888.write(repr(c).rjust(6) + " --> " + str(cube) + "\n")

    f888.close()
    f.close()

if(__name__ == '__main__'):
    main()
