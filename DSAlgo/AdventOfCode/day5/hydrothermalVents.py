class Coordinate:
    x = -1
    y = -1
    def __init__(self, x, y):
        self.x = x
        self.y = y

class Solution:
    def findStraightHydrothermalVents(self, matrix, lines):
        for ln in lines:
            if ln[0].x != ln[1].x and ln[0].y != ln[1].y:
                continue

            start = -1
            stop = -1
            variable = 'x'
            if ln[0].x != ln[1].x:
                start = ln[0].x
                stop = ln[1].x
                if ln[0].x > ln[1].x:
                    start = ln[1].x
                    stop = ln[0].x

            if ln[0].y != ln[1].y:
                variable = 'y'
                start = ln[0].y
                stop = ln[1].y
                if ln[0].y > ln[1].y:
                    start = ln[1].y
                    stop = ln[0].y

            if start == -1 or stop == -1:
                continue

            while start <= stop:
                if variable == 'x':
                    matrix[start][ln[0].y] += 1
                if variable == 'y':
                    matrix[ln[0].x][start] += 1
                start += 1

#         for i in matrix:
#             print(i)

        cnt = 0
        for i in matrix:
            for j in i:
                if j >= 2:
                    cnt += 1

        return cnt

    def findHydrothermalVents(self, matrix, lines):
        for ln in lines:
            if ln[0].x != ln[1].x and ln[0].y != ln[1].y:
                startx = ln[0].x
                stopx = ln[1].x
                starty = ln[0].y
                stopy = ln[1].y

#                 straight diagonal
                if startx < stopx and starty < stopy:
                    while startx < stopx and starty < stopy:
                        matrix[startx][starty] += 1
                        startx += 1
                        starty += 1

#                 opposite diagonal
                elif startx < stopx and starty > stopy:
                    while startx < stopx and starty > stopy:
                        matrix[startx][starty] += 1
                        startx += 1
                        starty -= 1

#                 reverse diagonal
                elif startx > stopx and starty < stopy:
                    while startx > stopx and starty < stopy:
                        matrix[startx][starty] += 1
                        startx -= 1
                        starty += 1

#                 reverse opposite diagonal
                elif startx > stopx and starty > stopy:
                    while startx > stopx and starty > stopy:
                        matrix[startx][starty] += 1
                        startx -= 1
                        starty -= 1

                matrix[startx][starty] += 1
#         for i in matrix:
#             print(i)

        cnt = 0
        for i in matrix:
            for j in i:
                if j >= 2:
                    cnt += 1

        return cnt



if __name__ == '__main__':
    f = open("./dist/input.txt", "r")
#     f = open("./dist/test.txt", "r")
    input = f.read()
    f.close()
    data = input.split('\n')
    lines = []
    maxRow = 0
    maxCol = 0

    for d in data:
        line = []
        ln = d.strip().split("->")
        for c in ln:
            xy = list(map(int, c.strip().split(',')))
            obj = Coordinate(int(xy[1]), int(xy[0]))
            if obj.x > maxRow:
                maxRow = obj.x
            if obj.y > maxCol:
                maxCol = obj.y
            line.append(obj)
        lines.append(line)

    matrix = [[0 for i in range(maxCol+1)] for j in range(maxRow+1)]
    obj = Solution()
    print("the number of points where at least two lines overlap: ", obj.findStraightHydrothermalVents(matrix, lines))
    print("the number of points where at least two lines overlap: ", obj.findHydrothermalVents(matrix, lines))
