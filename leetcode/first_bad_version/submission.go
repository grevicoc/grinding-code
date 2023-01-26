/** 
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad 
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

 func firstBadVersion(n int) int {
    first := 1
    last := n
    middle := last-first/2

    for !isBadVersion(middle) {
        first = middle + 1
        middle = (last-first)/2 + first 
    }

    for isBadVersion(middle) {
        middle -= 1
    }

    return middle + 1
}