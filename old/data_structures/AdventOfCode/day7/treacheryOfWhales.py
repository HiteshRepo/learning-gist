import sys

class Solution:

    fuelCostMem = {}
    minFuelCost = sys.maxsize

    def calcCostForPosition(self, pos, crabPositions):
        if pos in self.fuelCostMem:
            return self.fuelCostMem[pos]
        cost = 0
        for p in crabPositions:
            if p > pos:
                cost += p-pos
            else:
                cost += pos-p
        return cost

    def getVarCost(self, diff):
        return ((diff * (diff + 1)) / 2)

    def calcCostForPositionWithVariableStep(self, pos, crabPositions):
        if pos in self.fuelCostMem:
            return self.fuelCostMem[pos]
        cost = 0
        for p in crabPositions:
            if p > pos:
                cost += p-pos+self.getVarCost(p-pos-1)
            else:
                cost += pos-p+self.getVarCost(pos-p-1)
        return cost

    def findMinFuelCostWithConstFuelExpensePerStep(self, crabPositions):
        self.fuelCostMem = {}
        self.minFuelCost = sys.maxsize
        maxPos = max(crabPositions)
        for pos in range(maxPos):
            cost = self.calcCostForPosition(pos, crabPositions)
            if pos not in self.fuelCostMem:
                self.fuelCostMem[pos] = cost
            if cost < self.minFuelCost:
                self.minFuelCost = cost
        return self.minFuelCost

    def findMinFuelCostWithVariableFuelExpensePerStep(self, crabPositions):
        self.fuelCostMem = {}
        self.minFuelCost = sys.maxsize
        maxPos = max(crabPositions)
        for pos in range(maxPos):
            cost = self.calcCostForPositionWithVariableStep(pos, crabPositions)
            if pos not in self.fuelCostMem:
                self.fuelCostMem[pos] = cost
            if cost < self.minFuelCost:
                finalPos = pos
                self.minFuelCost = cost
        return self.minFuelCost



if __name__ == '__main__':
    f = open("./dist/input.txt", "r")
#     f = open("./dist/test.txt", "r")
    input = f.read()
    f.close()
    crabPositions = list(map(int, input.strip().split(',')))

    obj = Solution()
    print("min fuel cost for crab submarines: ", obj.findMinFuelCostWithConstFuelExpensePerStep(crabPositions))
    print("min fuel cost for crab submarines: ", obj.findMinFuelCostWithVariableFuelExpensePerStep(crabPositions))