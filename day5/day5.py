import os

DATA_PATH = os.path.join(os.getcwd(), "day5", "data.txt")


def check_before(before, before_list):
    for val in before_list:
        if val in before:
            return False
    return True


def part_one(search, updates):
    middle_count = 0
    for update in updates:
        flag = True
        before = set()
        for _, num in enumerate(update):
            if not check_before(before, search.get(num, [])):
                flag = False
                break
            before.add(num)
        if flag:
            mid_index = int((len(update) - 1) / 2)
            middle_count += update[mid_index]
    print(middle_count)


def part_two():
    pass


if "__main__" == __name__:
    search = {}
    updates = []
    with open(DATA_PATH) as lines:
        for line in lines:
            if "|" in line:
                a, b = line.split("|")
                a = int(a)
                b = int(b)
                if search.get(a, False) == False:
                    search[a] = []
                search[a].append(b)
            elif "," in line:
                updates.append([int(a) for a in line.split(",")])

    part_one(search, updates)
