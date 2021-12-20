class Solution:
    def traverse(self, src, adj, psf, visited, allPaths, allowed):
        if src == 'end':
            allPaths.append(psf)
            return

        if src not in visited:
            visited[src] = 0
        visited[src] += 1
        for nbr in adj[src]:
            if nbr == 'start' or ((nbr in visited and visited[nbr] > allowed-1) and nbr.islower() and nbr not in ["start", "end"]):
                continue
            self.traverse(nbr, adj, psf + ',' + nbr ,visited, allPaths, allowed)
        visited[src] -= 1

    def findAllPaths(self, adj, reslines):
        allPaths = []
        visited = {}
        psf = "start"
        src = "start"
        allowed = 2
        self.traverse(src, adj, psf, visited, allPaths, allowed)
#         for l in allPaths:
#             if l not in reslines:
#                 print(l)
        return len(allPaths)


if __name__ == '__main__':
    f = open("./dist/input.txt", "r")
#     f = open("./dist/test.txt", "r")
    input = f.read()
    f.close()
    lines = input.strip().split('\n')
    f = open("./dist/testres.txt", "r")
    input = f.read()
    f.close()
    reslines = input.strip().split('\n')

    adj = {}
    for line in lines:
        vertices = line.strip().split('-')
        if vertices[0] in adj:
            adj[vertices[0]].append(vertices[1])
        else:
            adj[vertices[0]] = [vertices[1]]

        if vertices[1] in adj:
            adj[vertices[1]].append(vertices[0])
        else:
            adj[vertices[1]] = [vertices[0]]

    obj = Solution()
    print("no. of paths: ", obj.findAllPaths(adj,reslines))