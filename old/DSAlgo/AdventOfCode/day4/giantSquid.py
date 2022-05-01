class Solution:
    set1 = {0,1,2,3,4}
    set2 = {0,5,10,15,20}
    set3 = {1,6,11,16,21}
    set4 = {2,7,12,17,22}
    set5 = {3,8,13,18,23}
    set6 = {4,9,14,19,24}
    set7 = {5,6,7,8,9}
    set8 = {10,11,12,13,14}
    set9 = {15,16,17,18,19}
    set10 = {20,21,22,23,24}

    def isBingo(self, hits):
        if self.set1.intersection(hits) == self.set1:
            return True
        if self.set2.intersection(hits) == self.set2:
            return True
        if self.set3.intersection(hits) == self.set3:
            return True
        if self.set4.intersection(hits) == self.set4:
            return True
        if self.set5.intersection(hits) == self.set5:
            return True
        if self.set6.intersection(hits) == self.set6:
            return True
        if self.set7.intersection(hits) == self.set7:
            return True
        if self.set8.intersection(hits) == self.set8:
            return True
        if self.set9.intersection(hits) == self.set9:
            return True
        if self.set10.intersection(hits) == self.set10:
            return True
        return False

    def createBoards(self, boards):
        dictBoards = {}
        dictHits = {}
        boardPrefix = "board"
        cnt = 1
        for board in boards:
            dictBoards[boardPrefix+str(cnt)] = board
            dictHits[boardPrefix+str(cnt)] = []
            cnt += 1
        return dictHits, dictBoards

    def playForFirst(self, randoms, boards):
        dictHits, dictBoards = self.createBoards(boards)
        boardPrefix = "board"
        hits = []
        winBoard = ""
        for i in randoms:
            for j in range(len(boards)):
                currBoard = boardPrefix + str(j+1)
                if i in dictBoards[currBoard]:
                    dictHits[currBoard].append(dictBoards[currBoard].index(i))
                    hits.append(i)
                if self.isBingo(dictHits[currBoard]):
                    winBoard = currBoard
                    break
            if winBoard != "":
                break

        sum = 0
        for i in dictBoards[winBoard]:
            if i not in hits:
                sum += i

        return sum * hits[-1]

    def playForLast(self, randoms, boards):
            dictHits, dictBoards = self.createBoards(boards)
            boardPrefix = "board"
            hits = []
            currBoard = ""
            boardTracker = []
            for i in randoms:
                for j in range(len(boards)):
                    currBoard = boardPrefix + str(j+1)
                    if currBoard in boardTracker:
                        continue
                    if i in dictBoards[currBoard]:
                        dictHits[currBoard].append(dictBoards[currBoard].index(i))
                        hits.append(i)
                    if self.isBingo(dictHits[currBoard]):
                        boardTracker.append(currBoard)
                if len(boardTracker) == len(boards):
                    break

            sum = 0
            for i in dictBoards[boardTracker[-1]]:
                if i not in hits:
                    sum += i

            return sum * hits[-1]



if __name__ == '__main__':
    f = open("./dist/input.txt", "r")
#     f = open("./dist/test.txt", "r")
    input = f.read()
    f.close()
    data = input.split('\n')
    randoms = list(map(int, data[0].strip().split(',')))
    data = data[2:]

    boards = []
    gridLen = 5
    cnt = 0
    currBoard = []
    for ln in data:
        if ln.strip() == '':
            cnt = 0
            continue
        currLn = list(map(int, ln.strip().split()))
        currBoard = currBoard + currLn
        if cnt == gridLen-1:
            boards.append(currBoard)
            currBoard = []
            cnt = 0
        cnt += 1

    obj = Solution()
    print("play score of first win: ", obj.playForFirst(randoms, boards))
    print("play score of last win: ", obj.playForLast(randoms, boards))