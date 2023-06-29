package main

import "fmt"

func main() {
	fmt.Println(numUniqueEmails([]string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"}))
	fmt.Println(numUniqueEmails([]string{"a@leetcode.com", "b@leetcode.com", "c@leetcode.com"}))
	fmt.Println(numUniqueEmails([]string{"a@e+c.com", "a@e+c+f.com"}))
}

func numUniqueEmails(emails []string) int {
	uniqueReceivers := make(map[string]struct{})

	for i := 0; i < len(emails); i++ {
		uniqueReceivers[getReceiver(emails[i])] = struct{}{}
	}

	return len(uniqueReceivers)
}

func getReceiver(email string) string {
	receiver := make([]byte, 0, len(email))
	var shouldSkip bool
	var dotAfterAt bool
	var plusAfterAt bool

	for i := 0; i < len(email); i++ {
		char := email[i]

		if char == '.' && !dotAfterAt {
			continue
		}

		if char == '+' && !plusAfterAt {
			shouldSkip = true
		}

		if char == '@' {
			shouldSkip = false
			dotAfterAt = true
			plusAfterAt = true
		}

		if !shouldSkip {
			receiver = append(receiver, char)
		}
	}

	return string(receiver)
}
