/*
Example 1
If a = 3, b = 5, and margin = 3, then close_compare(a, b, margin) should return 0.

This is because a and b are no more than 3 numbers apart.

Example 2
If a = 3, b = 5, and margin = 0, then close_compare(a, b, margin) should return -1.

This is because the distance between a and b is greater than 0, and a is less than b.
*/

// /
func CloseCompare(a, b, margin float64) int {
	x := 0
	z := margin
	switch {
	case a < b && z == 0 || z != 0 && b-z > a:
		x = -1

	case a > b && b+z < a:
		x = 1

	default:
		x = 0
	}
	return x
}

/*
func CloseCompare(a, b, margin float64) int {
	switch {
	  case a-b > margin:
		  return 1
	  case b-a > margin:
		  return -1
	  default:
		  return 0
	  }
  }

*/

/*
import "math"

func CloseCompare(a, b, margin float64) int {
  if math.Abs(a-b) <= margin { return 0 }
  if a > b { return 1 }
  return -1
}

*/