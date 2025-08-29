from typing import List

class Solution:
    def islandsAndTreasure(self, grid: List[List[int]]) -> None:
        visited = set()
        directions = [[1, 0], [-1, 0], [0, 1], [0, -1]]
        for i in range(len(grid)):
            for j in range(len(grid[0])):
                if grid[i][j] == -1 or (i, j) in visited:
                    continue
                if grid[i][j] == 0:
                    stack = [(i, j, 1)]
                    while len(stack) != 0:
                        cur = stack.pop()
                        visited.add((cur[0], cur[1]))
                        for d in directions:
                            newI, newJ = cur[0]+d[0], cur[1]+d[1]
                            if newI >= 0 and newI < len(grid) and newJ >= 0 and newJ < len(grid[0]) and grid[newI][newJ] > grid[cur[0]][cur[1]]:
                                grid[newI][newJ] = cur[2]
                                stack.append((newI, newJ, cur[2]+1))



p = lambda matrix: print('\n'.join(['\t'.join([str(cell) for cell in row]) for row in matrix]))

p(Solution().islandsAndTreasure([
    [2147483647,-1,0,2147483647],
    [2147483647,2147483647,2147483647,-1],
    [2147483647,-1,2147483647,-1],
    [0,-1,2147483647,2147483647]
]))
