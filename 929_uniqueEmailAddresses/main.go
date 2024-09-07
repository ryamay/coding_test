package uniqueemailaddresses

import "strings"

func numUniqueEmails(emails []string) int {
	mailMap := make(map[string]bool)

	for _, email := range emails {
		localName, domainName := splitEmail(email)
		mailMap[localName+"@"+domainName] = true
	}

	return len(mailMap)
}

func splitEmail(fullEmail string) (localName, domainName string) {
	splitted := strings.Split(fullEmail, "@")
	domainName = splitted[1]

	localName = splitted[0]
	localName = strings.Split(localName, "+")[0]       // "+"以後は切り捨てる
	localName = strings.ReplaceAll(localName, ".", "") // "."は削除する

	return localName, domainName
}
