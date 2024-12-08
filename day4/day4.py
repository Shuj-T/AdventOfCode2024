import os

DATA_PATH = os.path.join(os.getcwd(), "day4", "data.txt")


def check_grid(i, j, grid):
    if 0 <= i < len(grid) and 0 <= j < len(grid[0]):
        return grid[i][j]
    # print("invalid", i, j)
    return ""


def check_xmas(i, j, grid, xi=0, dir=(0, 0)):
    # print(dir) if dir != (0, 0) else 0
    count = 0
    xmas = "XMAS"
    dirs = [(1, 1), (1, 0), (0, 1), (-1, 0), (0, -1), (-1, -1), (-1, 1), (1, -1)]
    if xi == 0:
        # print("START", i, j)
        for dir in dirs:
            if xmas[xi] == check_grid(i, j, grid):
                count += check_xmas(i + dir[0], j + dir[1], grid, xi + 1, dir)
        return count
    else:
        if xmas[xi] == check_grid(i, j, grid):
            if xi + 1 == len(xmas):
                # print("XMAS!!!")
                return 1
            return check_xmas(i + dir[0], j + dir[1], grid, xi + 1, dir)
        else:
            # print(f"{xmas[xi]} == {check_grid(i, j, grid)}", "FAILED")
            return 0


def part_one(grid):
    xmas_count = 0
    x_loc = []
    xmas_loc = []
    for i, row in enumerate(grid):
        for j, char in enumerate(row):
            if char == "X":
                x_loc.append((i, j))
                before = xmas_count
                xmas_count += check_xmas(i, j, grid)
                if before < xmas_count:
                    xmas_loc.append((i, j))
    # print(xmas_loc)
    # print(x_loc)
    print(xmas_count)


def check_adj(a, b):
    if a != b and a in ["M", "S"] and b in ["M", "S"]:
        return True
    return False


def check_x_mas(i, j, grid):
    dir = [(1, 1), (-1, -1), (1, -1), (-1, 1)]
    tl = check_grid(i - 1, j - 1, grid)
    tr = check_grid(i - 1, j + 1, grid)
    bl = check_grid(i + 1, j - 1, grid)
    br = check_grid(i + 1, j + 1, grid)

    if check_adj(tl, br) and check_adj(tr, bl):
        return True
    return False


def part_two(grid):
    x_mas_count = 0
    for i, row in enumerate(grid):
        for j, char in enumerate(row):
            if char == "A":
                if check_x_mas(i, j, grid):
                    x_mas_count += 1
    print(x_mas_count)


if "__main__" == __name__:
    grid = []
    with open(DATA_PATH) as lines:
        for line in lines:
            row = [char for char in line if not char == "\n"]
            grid.append(row)
    # part_one(grid)
    part_two(grid)
