#User function template for Python

class Solution:
	def shortest_distance(self, matrix):
		for i in range(len(matrix)):
		    self.floydWarshall(i, matrix)
    def floydWarshall(self, fixed, matrix):
        for i in range(len(matrix)):
            for j in range(len(matrix[i])):
                if i == fixed and j == fixed:
                    continue
                res = matrix[i][fixed] + matrix[fixed][j]
                if res < matrix[i][j]:
                    matrix[i][j] = res
#{
#  Driver Code Starts
#Initial template for Python

if __name__ == '__main__':
    f = open("./dist/floyd-warshall/fileInput.txt", "r")
    input = f.read()
    f.close()
    T=input.split('\n')
    for i in range(int(T[0])):
		n = int(T[0])
		matrix = []
		for i in range(n):
			matrix.append(list(map(int, T[i+1].split())))
    obj = Solution()
    obj.shortest_distance(matrix)
    for _ in range(n):
        for __ in range(n):
            print(matrix[_][__], " ")
        print()
# } Driver Code Ends