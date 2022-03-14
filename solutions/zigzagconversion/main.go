package main

import "fmt"

func main() {
	fmt.Println(convert("PAYPALISHIRING", 3))
	fmt.Println(convert("PAYPALISHIRING", 4))
	fmt.Println(convert("A", 1))
	fmt.Println(convert("AB", 1))
	fmt.Println(convert("hjouvsuyoypayulyeimuotehzriicfskpggkbbipzzrzucxamludfykgruowzgiooobppleqlwphapjnadqhdcnvwdtxjbmyppphauxnspusgdhiixqmbfjxjcvudjsuyibyebmwsiqyoygyxymzevypzvjegebeocfuftsxdixtigsieehkchzdflilrjqfnxztqrsvbspkyhsenbppkqtpddbuotbbqcwivrfxjujjddntgeiqvdgaijvwcyaubwewpjvygehljxepbpiwuqzdzubdubzvafspqpqwuzifwovyddwyvvburczmgyjgfdxvtnunneslsplwuiupfxlzbknhkwppanltcfirjcddsozoyvegurfwcsfmoxeqmrjowrghwlkobmeahkgccnaehhsveymqpxhlrnunyfdzrhbasjeuygafoubutpnimuwfjqsjxvkqdorxxvrwctdsneogvbpkxlpgdirbfcriqifpgynkrrefx", 503))
}

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	containers := make([][]byte, numRows)

	var recentContainersIndex uint16
	var isAdd bool
	for i := 0; i < len(s); i++ {
		containers[recentContainersIndex] = append(containers[recentContainersIndex], s[i])

		if recentContainersIndex == 0 {
			isAdd = true
		} else if int(recentContainersIndex) == len(containers)-1 {
			isAdd = false
		}

		if isAdd {
			recentContainersIndex++
		} else {
			recentContainersIndex--
		}
	}

	var res string
	for _, v := range containers {
		res += string(v)
	}

	return res
}
