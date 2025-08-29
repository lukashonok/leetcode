from typing import List

class Solution:
    def validTree(self, n: int, edges: List[List[int]]) -> bool:
        if n == 0:
            return True
        if len(edges) != n-1:
            return False

        adj = {i:[] for i in range(n)}
        for i, j in edges:
            adj[i].append(j)
            adj[j].append(i)

        visited = set()
        def dfs(vertex: int, prev: int):
            if vertex in visited or len(adj[vertex]) > 3:
                return False
            visited.add(vertex)
            for vertexToVisit in adj[vertex]:
                if prev == vertexToVisit:
                    continue
                if not dfs(vertexToVisit, vertex):
                    return False
            return True
        return dfs(0, -1) and len(visited) == n


print(Solution().validTree(
    5, [[0, 1], [0, 2], [0, 3], [1, 4]]
))

print(Solution().validTree(
    5, [[0, 1], [1, 2], [2, 3], [1, 3], [1, 4]]
))

print(Solution().validTree(
    5, [[0,1],[0,2],[1,2],[3,4]]
))