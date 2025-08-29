from typing import List

class Solution:
    def countComponents(self, n: int, edges: List[List[int]]) -> int:
        if n == 0:
            return 0

        adj = {i:[] for i in range(n)}
        for i, j in edges:
            adj[i].append(j)
            adj[j].append(i)

        visited = set()
        def dfs(vertex: int, prev: int):
            if vertex in visited:
                return

            visited.add(vertex)
            for vertexToVisit in adj[vertex]:
                if vertexToVisit == prev:
                    continue
                dfs(vertexToVisit, vertex)

            return
        
        total = 0
        for v in range(n):
            if v not in visited:
                dfs(v, -1)
                total+=1

        return total

print(Solution().countComponents(
    n=3,
    edges=[[0,1], [0,2]]
))

print(Solution().countComponents(
    n=6,
    edges=[[0,1], [1,2], [2,3], [4,5]]
))