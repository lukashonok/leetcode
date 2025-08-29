from typing import List

class Interval(object):
    def __init__(self, start, end):
        self.start = start
        self.end = end

class Solution:
    def canAttendMeetings(self, intervals: List[Interval]) -> bool:
        if len(intervals) == 0:
            return True
        intervals.sort(key=lambda x: x.start)
        prevEnd = intervals[0].end
        for interval in intervals[1:]:
            if interval.start < prevEnd:
                return False
            else:
                prevEnd = max(interval.end, prevEnd)
        return True

solution = Solution()
# solution.canAttendMeetings([(0,30),(15,20),(5,10)])
solution.canAttendMeetings([Interval(5,8),Interval(9,15)])

