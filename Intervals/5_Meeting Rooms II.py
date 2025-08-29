from typing import List

class Interval(object):
    def __init__(self, start, end):
        self.start = start
        self.end = end
    def __repr__(self):
        return self.start + " " + self.end

class Solution:
    def minMeetingRooms(self, intervals: List[Interval]) -> int:
        starts = sorted(i.start for i in intervals)
        ends = sorted(i.end for i in intervals)
        res, count, s, e = 0, 0, 0, 0
        while s < len(starts):
            curMin = min(starts[s], ends[e])
            if curMin == ends[e] or starts[s] == ends[e]:
                count-=1
                e+=1
            else:
                count+=1
                s+=1
            res = max(res, count)
        return res
    

solution = Solution()
# solution.minMeetingRooms([(0,30),(15,20),(5,10)])
print(solution.minMeetingRooms([Interval(0,40),Interval(5,10),Interval(15,20)]))
print(solution.minMeetingRooms([Interval(0,40),Interval(5,10),Interval(15,20),Interval(10,15)]))
print(solution.minMeetingRooms([Interval(4,9)]))
print(solution.minMeetingRooms([Interval(1,5),Interval(2,6),Interval(3,7),Interval(4,8),Interval(5,9)]))
print(solution.minMeetingRooms([]))
