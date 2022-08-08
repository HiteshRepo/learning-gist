
class Solution:
    def findPosition(self, movements):
        horizontal = 0
        vertical = 0
        for movement in movements:
            moves = movement.split()
            direction = moves[0]
            distance = int(moves[1])

            if direction == 'forward':
                horizontal += distance

            if direction == 'up':
                vertical -= distance

            if direction == 'down':
                vertical += distance

        return horizontal*vertical

    def findAim(self, movements):
        horizontal = 0
        vertical = 0
        aim = 0
        for movement in movements:
            moves = movement.split()
            direction = moves[0]
            distance = int(moves[1])

            if direction == 'forward':
                horizontal += distance
                vertical += (aim*distance)

            if direction == 'up':
                aim -= distance

            if direction == 'down':
                aim += distance

        return horizontal*vertical



if __name__ == '__main__':
    f = open("./dist/input1.txt", "r")
    input = f.read()
    f.close()
    movements = input.split('\n')

#     Test case 1
    movements = ["forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"]

    obj = Solution()
    print("final position: ", obj.findPosition(movements))

    obj = Solution()
    print("final position by aim: ", obj.findAim(movements))