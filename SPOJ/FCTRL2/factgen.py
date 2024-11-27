#!/usr/bin/python

def main():
    f = open("factorials.txt", 'w')
    fc = open("factorials_cpp_array.txt", 'w')
    fc.write("const std::string factorials[] = {\n")

    factorial = 1;
    counter = 0;
    f.write(str(factorial) + "\n")
    fc.write("  \"" + str(factorial) + "\",\n")

    while(counter < 101):
        counter += 1
        factorial *= counter
        f.write(str(factorial) + "\n")
        fc.write("  \"" + str(factorial) + "\",\n")


    fc.write("};\n")
    fc.close()
    f.close()

if(__name__ == "__main__"):
    main()
