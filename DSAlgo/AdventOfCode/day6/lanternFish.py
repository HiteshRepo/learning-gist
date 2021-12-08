days = -1

def decreaseBy1(fishInternal):
    if fishInternal == 0:
        return 6
    return fishInternal-1

def updateFishData(currentFishCount):
    i = 0
    addFish = 0
    while i < len(currentFishCount):
        if currentFishCount[i] == 0:
            addFish += 1
        i += 1

    while addFish > 0:
        currentFishCount.append(9)
        addFish -= 1

    decreasedFishData = list(map(decreaseBy1, currentFishCount))
    return decreasedFishData


# import multiprocessing as mp
def predictFishCount(currentFishCount):
    dayCnt = 0
    while dayCnt < days:
#         fishDataChunked = chunkIt(currentFishCount, 10)
#         pool = mp.Pool(mp.cpu_count())
#         results = [pool.apply(updateFishData, args=(row,)) for row in fishDataChunked]
        currentFishCount = updateFishData(currentFishCount)
        dayCnt += 1

    return currentFishCount

def chunkIt(seq, num):
    avg = len(seq) / float(num)
    out = []
    last = 0.0

    while last < len(seq):
        out.append(seq[int(last):int(last + avg)])
        last += avg

    return out

if __name__ == '__main__':
    f = open("./dist/input.txt", "r")
#     f = open("./dist/test.txt", "r")
    input = f.read()
    f.close()
    fishData = list(map(int, input.strip().split(',')))

#     days = 18
#     print("fish count after 18 days: ", len(predictFishCount(fishData)))
#     days = 80
#     print(len(fishData))
#     print("fish count after 80 days: ", len(predictFishCount(fishData)))
    days = 256
    fishData = [3]

    print("fish count after 256 days: ", )