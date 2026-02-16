# Runtime complexity: O(1) (as each component's length is clearly determined, and the map is bounded)
# Auxiliary space complexity: O(12) == O(1)
# Subjective level: easy
# Solved on: 2026-02-16
class Solution:
    def reformatDate(self, date: str) -> str:
        MONTHS = {"Jan": "01", "Feb": "02", "Mar": "03", "Apr": "04", "May": "05", "Jun": "06",
                  "Jul": "07", "Aug": "08", "Sep": "09", "Oct": "10", "Nov": "11", "Dec": "12"}
        iday, imonth, iyear = date.split(" ")
        day = int(iday[:-2])
        m = MONTHS[imonth]
        return f"{iyear}-{m}-{day:02d}"
