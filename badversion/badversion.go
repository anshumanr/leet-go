package badversion

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	lo, high := 1, n

	bad := -1
	for lo <= high {
		mid := (lo + high) / 2

		if isBadVersion(lo) {
			return lo
		}

		if isBadVersion(mid) {
			bad = mid
			high = mid - 1
		} else {
			lo = mid + 1
		}
	}

	return bad
}
