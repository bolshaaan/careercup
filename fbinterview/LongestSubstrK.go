package fbinterview

/*
https://www.careercup.com/question?id=5158272756613120

Given a string, find the longest substring with k distinct characters.

e.g - “aaaabbbb”, k = 2, “aaaabbbb”
“asdfrttt” k = 3, “asd”, “frttt”
*/

func LongestSubstring(str string, uniqChars int) string {

	if str == "" {
		return ""
	}

	chrs := make(map[byte]interface{})
	var substr, maxstr string

	var i, j int = 0, 0

	for j = 0; j <= len(str); j++ {

		//fmt.Println("Substr: ", substr, " i = ", i, " j = ", j )

		if j < len(str) {
			chrs[str[j]] = nil // new char in set
		}

		if len(chrs) > uniqChars || j >= len(str) { // signal to check
			if len(substr) > len(maxstr) {
				maxstr = substr
			}

			//fmt.Println("len susbtr ", len(substr))
			delete(chrs, substr[0])

			// remove char from substring
			for ; i < len(str) && str[i] == substr[0]; i++ {
			}
		}

		if i < len(str) && j < len(str) {
			substr = str[i : j+1]
		}

		//fmt.Println("after substr: ", substr)
	}

	return maxstr
}
