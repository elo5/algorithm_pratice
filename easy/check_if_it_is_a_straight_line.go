package easy

/**
1232. Check If It Is a Straight Line
You are given an array coordinates, coordinates[i] = [x, y], where [x, y] represents the coordinate of a point. Check if these points make a straight line in the XY plane.

Example 1:
Input: coordinates = [[1,2],[2,3],[3,4],[4,5],[5,6],[6,7]]
Output: true

Example 2:
Input: coordinates = [[1,1],[2,2],[3,4],[4,5],[5,6],[7,7]]
Output: false

Constraints:
2 <= coordinates.length <= 1000
coordinates[i].length == 2
-10^4 <= coordinates[i][0], coordinates[i][1] <= 10^4
coordinates contains no duplicate point.
 */

func CheckStraightLine(coordinates [][]int) bool {

	if len(coordinates) == 2 {
		return true
	}

	straight := true
	for i := 2 ; i < len(coordinates); i++ {
		x1 := coordinates[i - 2][0]
		y1 := coordinates[i - 2][1]

		x2 := coordinates[i - 1][0]
		y2 := coordinates[i - 1][1]

		x3 := coordinates[i][0]
		y3 := coordinates[i][1]

		if x3 * (y2 - y1) - y3 * (x2 - x1) != x1 * y2 - x2 * y1 {
			straight = false
			break
		}
	}
	return straight
}
