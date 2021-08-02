class Solution:
    
    #Function to return a list of integers denoting spiral traversal of matrix.
    def spirallyTraverse(self,matrix, r, c): 
        dir = 0
        top = 0
        down = r-1
        right = c-1
        left = 0
        lst = []
        while top <= down and left <= right:
            if dir == 0:
                i = left
                while i<=right:
                    lst.append(matrix[top][i])
                    i = i+1
                top += 1
            elif dir == 1:
                i = top
                while i<=down:
                    lst.append(matrix[i][right])
                    i = i+1
                right -= 1
            elif dir == 2:
                i = right
                while i>=left:
                    lst.append(matrix[down][i])
                    i = i-1
                down -= 1
            elif dir == 3:
                i = down
                while i>=top:
                    lst.append(matrix[i][left])
                    i -= 1
                left += 1
            dir = (dir+1)%4
        return lst