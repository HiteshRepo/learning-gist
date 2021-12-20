class Solution:
    def defineMatrix(self, matrix, dataMatrix):
        for coord in dataMatrix:
            matrix[coord[0]][coord[1]] = '#'

    def display(self, matrix):
        for r in matrix:
            print(r)

    def fold(self, cmdCol, matrix):
        for cmd in cmdCol:
            if cmd[0] == 'y':
                inc = cmd[1]
                dec = cmd[1] - 1
                while inc < len(matrix) and dec >= 0:
                    for i in range(len(matrix[0])):
                        if matrix[inc][i] == '#':
                            print(inc, i)
                            print(dec, i)
                            matrix[dec][i] = '#'
                    dec -= 1
                    inc += 1
                matrix = matrix[:cmd[1]]

            if cmd[0] == 'x':
                inc = cmd[1]
                dec = cmd[1] - 1
                while inc < len(matrix[0]) and dec >= 0:
                    for i in range(len(matrix)):
                        if matrix[i][inc] == '#':
                            matrix[i][dec] = '#'
                    dec -= 1
                    inc += 1
                for i in range(len(matrix)):
                    matrix[i] = matrix[i][:cmd[1]]


if __name__ == '__main__':
#     f = open("./dist/input.txt", "r")
    f = open("./dist/test.txt", "r")
    input = f.read()
    f.close()
    lines = input.strip().split('\n')

    dataMatrix = []
    cmdCol = []
    maxRow = 0
    maxCol = 0
    startInstructions = False
    for l in lines:
        if l == '':
            startInstructions = True
            continue
        if not startInstructions:
            rc = list(map(int, l.strip().split(',')))
            if rc[1] > maxRow:
                maxRow = rc[1]
            if rc[0] > maxCol:
                maxCol = rc[0]
            dataMatrix.append((rc[1], rc[0]))
        else:
            insts = l.strip().split(' ')
            inst = insts[-1]
            dir = inst.strip().split('=')
            if len(dir) > 0:
                cmdCol.append((dir[0], int(dir[1])))

    matrix = [['.' for i in range(maxCol+1)] for i in range(maxRow+1)]

    obj = Solution()
    obj.defineMatrix(matrix, dataMatrix)
    obj.fold(cmdCol, matrix)
    obj.display(matrix)