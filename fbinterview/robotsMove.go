/* https://www.careercup.com/question?id=5095457003929600 */

package fbinterview

import (
	"fmt"
	"strings"
)

var ALPH_NUM = 26

// 1..26

// kw -- keyboard width
// x = 1..26
// y = 0..
func PrintSentence(text string, kw int) string {

	cx, cy := 0, 0
	x, y := 0, 0

	moves := []string{}

	for _, r := range text {
		num := int(r - 'A' + 1)

		x = num % kw
		y = num / kw

		move := ""
		dy := cy - y
		if dy < 0 {
			move += fmt.Sprintf("D%d", -dy)
		} else if dy > 0 {
			move += fmt.Sprintf("U%d", dy)
		}
		if move != "" {
			moves = append(moves, move)
		}

		move = ""
		dx := cx - x

		if dx < 0 {
			move = fmt.Sprintf("R%d", -dx)
		} else if dx > 0 {
			move = fmt.Sprintf("L%d", dx)
		}

		if move != "" {
			moves = append(moves, move)
		}

		moves = append(moves, "T")

		cx = x
		cy = y

	}

	result := strings.Join(moves, ", ")
	fmt.Println("Result: ", result)

	return result
}
