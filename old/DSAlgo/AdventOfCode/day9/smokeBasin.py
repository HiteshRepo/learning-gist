rows = [-1, 0, 0, 1]
cols = [0, -1, 1, 0]

class Solution:
    def isValid(self, x, y, X, Y):
        return (0 <= x <= X) and (0 <= y <= Y)

    def isLowest(self, elem, elems):
        return all(i > elem for i in elems)

    def findLowPoints(self, matrix):
        row = len(matrix)
        column = len(matrix[0])
        lowPoints = []

        i = 0

        while i < row:
            j = 0
            while j < column:
                adjs = []
                for k in range(len(rows)):
                    if self.isValid(i+rows[k], j+cols[k], row-1, column-1):
                        elem = matrix[i+rows[k]][j+cols[k]]
                        adjs.append(elem)

                if self.isLowest(matrix[i][j], adjs):
                    lowPoints.append((i,j))

                j += 1
            i += 1
        return lowPoints

    def calcRiskLevel(self, matrix):
        lowPoints = self.findLowPoints(matrix)
        riskLevel = 0
        for p in lowPoints:
            riskLevel += matrix[p[0]][p[1]]+1
        return riskLevel

    def calcCurrBasinLevel(self, pos, prev, matrix, lvl, visited):
        if pos[0] < 0 or pos[1] < 0 or pos[0] > len(matrix) or pos[1] > len(matrix[0]) or matrix[pos[0]][pos[1]] == 9:
            return lvl

        if (pos[0], pos[1]) not in visited:
            lvl += 1
        visited.append(pos)

        for k in range(len(rows)):
            if (pos[0]+rows[k], pos[1]+cols[k]) not in visited and self.isValid(pos[0]+rows[k], pos[1]+cols[k], len(matrix)-1, len(matrix[0])-1):
                nextPos = (pos[0]+rows[k], pos[1]+cols[k])
                currVal = matrix[pos[0]][pos[1]]
                lvl = self.calcCurrBasinLevel(nextPos, currVal, matrix, lvl, visited)

        return lvl

    def calcBasinsLevel(self, matrix):
        lowPoints = self.findLowPoints(matrix)
        print(lowPoints)
        maxLvl = 1
        import queue
        q =  queue.PriorityQueue()
        for p in lowPoints:
            lvl = self.calcCurrBasinLevel(p, -1, matrix, 0, [])
            if q.qsize() == 3:
                q.get()
            q.put(lvl)

        while not(q.empty()):
            l = q.get()
            print(l)
            maxLvl *= l

        return maxLvl



if __name__ == '__main__':
#     f = open("./dist/input.txt", "r")
    f = open("./dist/test2.txt", "r")
    input = f.read()
    f.close()
    data = input.strip().split('\n')

    matrix = []
    for r in data:
        matrix.append(list(map(int, [char for char in r.strip()])))

    obj = Solution()
#     print("risk level: ", obj.calcRiskLevel(matrix))
    print("basin level: ", obj.calcBasinsLevel(matrix))