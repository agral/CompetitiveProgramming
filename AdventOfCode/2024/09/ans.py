DO_EXAMPLE = True
DO_EXAMPLE = False

FREE_SPACE_ID = -1
def print_memory(memory):
    for v in memory:
        print("." if v == FREE_SPACE_ID else str(v), end="")
    print()

def checksum(memory):
    ans = 0
    for v in range(1, len(memory)):
        if memory[v] >= 0:
            ans += v * memory[v]
    return ans

def main():
    A = 0
    B = 0
    with open("example.txt" if DO_EXAMPLE else "input.txt", "r") as file:
        data = [line.strip() for line in file][0]
        memory = []
        memptr = 0
        id = 0

        is_file = True
        for ch_block in data:
            length = int(ch_block)
            for _ in range(memptr, memptr + length):
                memory.append(id if is_file else FREE_SPACE_ID)
            memptr += length
            if is_file:
                id += 1
            is_file = not is_file
        original_memory = memory[:] # store the original memory layout for part B

        # part A: do the strange defrag thing:
        lptr = 0
        rptr = len(memory) - 1
        while True:
            # advance lptr to the next free space
            while lptr < len(memory) and memory[lptr] != FREE_SPACE_ID:
                lptr += 1
            # retract rptr to the previous occupied byte (i.e. anything that is not a free space):
            while rptr >= 0 and memory[rptr] == FREE_SPACE_ID:
                rptr -= 1
            # the defrag is done when lptr is no longer before rptr:
            if (lptr >= rptr):
                break
            # otherwise swap the bytes pointed to by lptr and rptr.
            # this moves some memory block to the left, and the free space to the right.
            memory[lptr], memory[rptr] = memory[rptr], memory[lptr]
        A = checksum(memory)
        print(f"Answer A: {A}")

        # part B: do the strange whole-file defrag thing:
        memory = original_memory[:] # first restore the original memory layout.

        rptr = len(memory) - 1
        while (rptr > 0):
            # find the beginning and the size of the file pointed to by rptr:
            fendptr = rptr
            while rptr >= 0 and memory[rptr] == memory[fendptr]:
                rptr -= 1
            # now rptr is one byte before the file; either on an empty space or at the end of next file.
            fstartptr = rptr + 1
            fsize = fendptr - fstartptr + 1

            # before moving the file, retract rptr so that it's on the next file to be defragged:
            # retract rptr until it's on a file, not an empty space:
            while rptr >= 0 and memory[rptr] == FREE_SPACE_ID:
                rptr -= 1
            # Rptr now marks the end byte of a next file intended to be moved.

            # find first contiguous block of space able to fit fsize:
            free_start_ptr = 0
            moved = False
            while free_start_ptr < fstartptr and not moved:
                # advance free_start_ptr to the next empty space block:
                while free_start_ptr < fstartptr and memory[free_start_ptr] != FREE_SPACE_ID:
                    free_start_ptr += 1
                # calculate length of this empty space block:
                length = 1
                while free_start_ptr + length < len(memory) and memory[free_start_ptr + length] == FREE_SPACE_ID:
                    length += 1
                if length >= fsize and free_start_ptr < fstartptr:
                    #print(f"Moving block #{memory[fstartptr]} ({fstartptr}-{fendptr}) to {free_start_ptr}.")
                    # Move the file to this free space, free the space occupied by moved file:
                    fid = memory[fstartptr]
                    for b in range(fsize):
                        memory[free_start_ptr + b] = fid
                        memory[fstartptr + b] = FREE_SPACE_ID
                    moved = True
                else:
                    free_start_ptr += length # move past this insufficiently-sized free memory block



        # that solves the example input, but gives WA for the real one. Bugger.
        if False:
            spaces = {} # maps space length to position.
            sptr = 0
            while sptr < len(memory):
                while sptr < len(memory) and memory[sptr] != FREE_SPACE_ID:
                    sptr += 1
                sstart = sptr
                while sptr < len(memory) and memory[sptr] == FREE_SPACE_ID:
                    sptr += 1
                ssize = sptr - sstart

                # append the found space block to spaces map:
                if (ssize > 0):
                    if ssize not in spaces.keys():
                        spaces[ssize] = []
                    spaces[ssize].append(sstart)
            skeys = sorted(spaces.keys())

            lptr = 0 # lptr points to the first known empty space from the left
            while lptr < len(memory) and memory[lptr] != FREE_SPACE_ID:
                lptr += 1
            rptr = len(memory)-1 # rptr points to the file with biggest ID.
                                 # (note that the files are initially sorted with IDs strictly increasing)
            while rptr > 0:
                # find the beginning and the size of the file pointed to by rptr:
                fendptr = rptr
                while rptr >= 0 and memory[rptr] == memory[fendptr]:
                    rptr -= 1
                # now rptr is one byte before the file; either on an empty space or at the end of next file.
                fstartptr = rptr + 1
                fsize = fendptr - fstartptr + 1

                # before moving the file, retract rptr so that it's on the next file to be defragged:
                # retract rptr until it's on a file, not an empty space:
                while rptr >= 0 and memory[rptr] == FREE_SPACE_ID:
                    rptr -= 1
                # Rptr now marks the end byte of a next file intended to be moved.

                # try moving the fstartptr file to the left, if space is available:
                for k in skeys:
                    if k >= fsize:
                        if len(spaces[k]) > 0 and spaces[k][0] < fstartptr:
                            ftargetptr = spaces[k][0]
                            remaining_free_mem = k - fsize
                            spaces[k] = spaces[k][1:] # remove the chosen free memory from spaces
                            if remaining_free_mem > 0:
                                if remaining_free_mem not in spaces.keys():
                                    spaces[remaining_free_mem] = []
                                spaces[remaining_free_mem].append(ftargetptr + fsize)
                                spaces[remaining_free_mem].sort()
                            fid = memory[fstartptr]
                            for b in range(fsize):
                                memory[ftargetptr + b] = fid
                                memory[fstartptr + b] = FREE_SPACE_ID
                            break

        B = checksum(memory)
        print(f"Answer B: {B}")

if __name__ == "__main__":
    main()
