# Note to self: it's shit code. It works.
CH_WALL =  "#"
CH_BOX =   "O"
CH_BOX_LEFT = "["
CH_BOX_RIGHT = "]"
CH_ROBOT = "@"
CH_TILE =  "."
DO_EXAMPLE = True
DO_EXAMPLE = False

MOVES = {
    "<": [0, -1],
    "^": [-1, 0],
    ">": [0, 1],
    "v": [1, 0]
}

def main():
    A = 0
    B = 0
    map = []
    map2 = []
    moves = ""
    H, W = 0, 0
    ry, rx = 0, 0 # this holds current robot's position
    RY, RX = 0, 0 # this is for part2
    is_map_loaded = False

    def print_map():
        for h in range(H):
            for w in range(W):
                if ry == h and rx == w:
                    print(CH_ROBOT, end="")
                else:
                    print(map[h][w], end="")
            print()

    def print_map2():
        for h in range(H):
            for w in range(2 * W):
                if ry == h and rx == w:
                    print(CH_ROBOT, end="")
                else:
                    print(map2[h][w], end="")
            print()

    with open("example.txt" if DO_EXAMPLE else "input.txt", "r") as file:
        for line in file:
            line = line.strip()
            if len(line) == 0:
                is_map_loaded = True
            elif not is_map_loaded:
                mapline = []
                mapline2 = []
                for w in range(len(line)):
                    if line[w] == "@":
                        ry, rx = H, w
                        RY, RX = H, 2 * w
                        mapline.append(CH_TILE)
                        mapline2.append(CH_TILE)
                        mapline2.append(CH_TILE)
                    else:
                        mapline.append(line[w])
                        # for mapline2, double the empty tiles and the walls;
                        # but represent boxes as two distinct chars.
                        mapline2.append(CH_BOX_LEFT if line[w] == CH_BOX else line[w])
                        mapline2.append(CH_BOX_RIGHT if line[w] == CH_BOX else line[w])
                H += 1
                map.append(mapline)
                map2.append(mapline2)
            else:
                moves += (line)
        W = len(map[0])

        for move in moves:
            offset = MOVES[move]
            ny, nx = ry + offset[0], rx + offset[1]
            moved_boxes = 0
            while map[ny][nx] == CH_BOX:
                ny += offset[0]
                nx += offset[1]
                moved_boxes += 1
            if map[ny][nx] == CH_TILE:
                # This move is possible.
                # First move the robot to the tile it intended to go; it was either an empty space
                # or a moved box (possibly part of entire chain of boxes):
                ry += offset[0]
                rx += offset[1]

                # then move the entire line of boxes (if there are any):
                if moved_boxes > 0:
                    # by marking this final empty tile as a box:
                    map[ny][nx] = CH_BOX
                    # and unmarking the box where the robot is now back to the empty tile.
                    # This is equivalent to moving the entire line of `moved_boxes`.
                    map[ry][rx] = CH_TILE
            # else there's a wall in the end, and no way to perform this move
        for h in range(H):
            for w in range(W):
                if map[h][w] == CH_BOX:
                    A += 100 * h + w;
        print(f"Answer A: {A}")

        ry, rx = RY, RX
        for move in moves:
            offset = MOVES[move]
            ny, nx = ry + offset[0], rx + offset[1]
            if offset[0] == 0: # horizontal movements are easier - only one line to process.
                moved_boxes_count = 0
                while map2[ny][nx] == CH_BOX_LEFT or map2[ny][nx] == CH_BOX_RIGHT:
                    moved_boxes_count += 1
                    nx += 2 * offset[1]
                if map2[ny][nx] == CH_TILE:
                    # Similar to case A, this move is possible. But this time need to move all the boxes.
                    for dx in range(2 * moved_boxes_count + 1): # the +1 takes into account robot's empty space too.
                        map2[ny][nx - offset[1] * dx] = map2[ny][nx - offset[1] * (dx + 1)]
                    rx += offset[1]
                # else there's a wall in the end, and again no way to perform this move. Do nothing.
            else: # vertical movement; harder but doable ;-)
                if map2[ny][nx] == CH_TILE:
                    ry, rx = ny, nx # simply move the robot to the empty space.
                elif map2[ny][nx] == CH_WALL:
                    print("Lol, robot tried to move directly into wall") # make fun of pajeetware whenever applicable
                elif map2[ny][nx] == CH_BOX_LEFT or map2[ny][nx] == CH_BOX_RIGHT:
                    # The robot's maybe moving boxes!
                    is_move_ok = True # so far, will be set to false if any box can't be moved.

                    # note to self: noticed that parentheses below?
                    # these are important. These cost ~15 minutes of debugging,
                    # and the errorneous behavior here is hilarious as fuck.
                    currently_moving = [[ny, nx - (0 if map2[ny][nx] == CH_BOX_LEFT else 1)]]
                    all_moving = currently_moving[:]
                    while is_move_ok and len(currently_moving) > 0:
                        new_moved_boxes = []
                        for box in currently_moving:
                            # box can be moved if no wall directly in front of left-side and right-side.
                            if map2[box[0] + offset[0]][box[1]] == CH_WALL or map2[box[0] + offset[0]][box[1] + 1] == CH_WALL:
                                is_move_ok = False

                            # box can move other boxes if these boxes have left-side to the left, directly, or to the right
                            # of this box's left-side. This works for both "up" and "down" directions.
                            for dx in range(-1, 2): # delta-x: -1, 0, 1.
                                if map2[box[0] + offset[0]][box[1] + dx] == CH_BOX_LEFT:
                                    new_moved_boxes.append([box[0] + offset[0], box[1] + dx])
                                    all_moving.append([box[0] + offset[0], box[1] + dx])
                        # after all currently_moving boxes are processed, move to process the next row.
                        currently_moving = new_moved_boxes[:]

                    if is_move_ok:
                        # OK, now we know the position of left-sides of all the boxes that need to move.
                        # Let's remove them all from the map first, 
                        for box in all_moving:
                            map2[box[0]][box[1]] =     CH_TILE
                            map2[box[0]][box[1] + 1] = CH_TILE
                        # Then make them appear in their target positions:
                        for box in all_moving:
                            map2[box[0] + offset[0]][box[1]] =     CH_BOX_LEFT
                            map2[box[0] + offset[0]][box[1] + 1] = CH_BOX_RIGHT
                        # Finally, the robot moves too:
                        ry += offset[0]
                    # otherwise (if not is_move_ok) nothing moves; the boxes and the robot all stay still.

        print_map2() # print final state for quick verification
        for h in range(H):
            for w in range(2 * W):
                if map2[h][w] == CH_BOX_LEFT:
                    B += 100 * h + w
        print(f"Answer B: {B}")

if __name__ == "__main__":
    main()
