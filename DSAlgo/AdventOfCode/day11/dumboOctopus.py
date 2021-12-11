class Solution:
    flashed = []
    def resetFlashes(self, octopuses, i, j):
        if i-1 >= 0 and j >= 0 and i-1 < len(octopuses) and j < len(octopuses[i]) and (i-1, j) not in self.flashed:
            octopuses[i-1][j] += 1
        if i-1 >= 0 and j+1 >= 0 and i-1 < len(octopuses) and j+1 < len(octopuses[i]) and (i-1, j+1) not in self.flashed:
            octopuses[i-1][j+1] += 1
        if i >= 0 and j+1 >= 0 and i < len(octopuses) and j+1 < len(octopuses[i]) and (i, j+1) not in self.flashed:
            octopuses[i][j+1] += 1
        if i+1 >= 0 and j+1 >= 0 and i+1 < len(octopuses) and j+1 < len(octopuses[i]) and (i+1, j+1) not in self.flashed:
            octopuses[i+1][j+1] += 1
        if i+1 >= 0 and j >= 0 and i+1 < len(octopuses) and j < len(octopuses[i]) and (i+1, j) not in self.flashed:
            octopuses[i+1][j] += 1
        if i+1 >= 0 and j-1 >= 0 and i+1 < len(octopuses) and j-1 < len(octopuses[i]) and (i+1, j-1) not in self.flashed:
            octopuses[i+1][j-1] += 1
        if i >= 0 and j-1 >= 0 and i < len(octopuses) and j-1 < len(octopuses[i]) and (i, j-1) not in self.flashed:
            octopuses[i][j-1] += 1
        if i-1 >= 0 and j-1 >= 0 and i-1 < len(octopuses) and j-1 < len(octopuses[i]) and (i-1, j-1) not in self.flashed:
            octopuses[i-1][j-1] += 1
        octopuses[i][j] = 0

    def countFlashes(self, octopuses):
        flashes = 0
        for i in range(len(octopuses)):
            for j in range(len(octopuses[i])):
                if octopuses[i][j] > 9:
                    flashes += 1
                    self.flashed.append((i,j))
                    self.resetFlashes(octopuses, i, j)
                    flashes += self.countFlashes(octopuses)
        return flashes

    def countTotalFlashes(self, octopuses, totalSteps):
        steps = 0
        totalFlashes = 0
        while steps < totalSteps:
            for i in range(len(octopuses)):
                for j in range(len(octopuses[i])):
                    octopuses[i][j] += 1
            steps += 1
            totalFlashes += self.countFlashes(octopuses)
            self.flashed = []
        return totalFlashes

    def countTotalFlashes(self, octopuses):
        steps = 0
        totalFlashes = 0
        while True:
            for i in range(len(octopuses)):
                for j in range(len(octopuses[i])):
                    octopuses[i][j] += 1
            steps += 1
            flashes = self.countFlashes(octopuses)
            totalFlashes += flashes
            self.flashed = []
            if flashes == 100:
                return steps

if __name__ == '__main__':
    f = open("./dist/input.txt", "r")
#     f = open("./dist/test.txt", "r")
    input = f.read()
    f.close()
    lines = input.strip().split('\n')

    octopuses = []
    for line in lines:
        row = []
        for ch in line:
            row.append(int(ch))
        octopuses.append(row)

    obj = Solution()
    print('total flashes: ', obj.countTotalFlashes(octopuses, 100))
    print('step at which all octopuses flash: ', obj.countTotalFlashes(octopuses))