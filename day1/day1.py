import os

DATA_PATH = os.path.join(os.getcwd(), "day1", "data.txt")


def part_one():
    a = []
    b = []
    with open(os.path.join(DATA_PATH), "r") as f:
        for line in f:
            numbers = line.split("   ")
            a.append(int(numbers[0]))
            b.append(int(numbers[1]))
    a.sort()
    b.sort()
    total = 0
    for i in range(len(a)):
        total += abs(a[i] - b[i])
    print(f"Total distance: {total}")


def part_two():
    left = []
    right = []
    right_count = {}
    with open(DATA_PATH, "r") as f:
        for line in f:
            numbers = line.split("   ")
            left.append(int(numbers[0]))

            right_number = int(numbers[1])

            if right_count.get(right_number) is None:
                right_count[right_number] = 0
            right_count[right_number] += 1

            right.append(right_number)
    sim_score = 0
    for left_num in left:
        sim_score += left_num * right_count.get(left_num, 0)
    print(f"Similarity Score: {sim_score}")


if __name__ == "__main__":
    part_one()
    part_two()
