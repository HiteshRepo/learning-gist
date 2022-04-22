## Bit Pre-requisites for Trie

### How is a number represent in bits
1. 9 (integer 32 bits) = 00000000......1001
2. Rightmost index is 0.
3. Leftmost index is 31.
4. Had this been a long (64 bits).
   1. Rightmost index would have been 0.
   2. Leftmost index would have been 63.
5. Presence of 1 towards the left makes it bigger.

### XOR
1. 1 ^ 1 = 0
2. 1 ^ 0 = 1
3. 0 ^ 1 = 1
4. 0 ^ 0 = 2
5. XOR of different bits = 1
6. XOR of same bits = 0
7. for number of bits more than 2
   1. even number of 1's = 0
   2. odd number of 1's = 1

### Check if a bit is set or not
1. 9 = 00000000......1001
2. 3rd bit is 1 or not (set or not)?
3. 9 >> 3 (9 right shift 3) = 00000000...0001
4. 3rd bit set or not = (num >> 3) & 1
5. 00000000...0001 (9 >> 3) & 00000....00001 (1) = 00000....00000
   1. if the resultant last bit i 0, then 3rd bit is set
   2. else not set

### Turn on a bit
1. 9 = 00000000......1001
2. Turn on 2nd bit = 9 | (1 << 2)
3. 1 << 2 (1 left shift 2) = 00000000...00100
4. 00000000......1001 (9) | 00000000...00100 (1 << 2) = 000000....1101

