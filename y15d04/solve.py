import sys
import hashlib

input, i = open(sys.argv[1]).read().strip(), 0


def find_num(prefix, i=0):
    while True:
        hash = hashlib.md5()
        hash.update((input + str(i)).encode("utf-8"))
        if hash.hexdigest().startswith(prefix):
            print(i)
            break
        i += 1


find_num("00000")  # part 1
find_num("000000")  # part 2
