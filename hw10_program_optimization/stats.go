package hw10programoptimization

import (
	"bufio"
	"io"
	"strings"

	"github.com/mailru/easyjson"
)

//easyjson:json
type User struct {
	ID       int    `json:"-"`
	Name     string `json:"-"`
	Username string `json:"-"`
	Email    string
	Phone    string `json:"-"`
	Password string `json:"-"`
	Address  string `json:"-"`
}

type DomainStat map[string]int

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	return countDomains(r, domain)
}

// type users [100_000]User

func countDomains(r io.Reader, domain string) (DomainStat, error) {
	dm := "." + domain
	fileScanner := bufio.NewScanner(r)
	var user User
	result := make(DomainStat)

	for fileScanner.Scan() {
		if err := easyjson.Unmarshal(fileScanner.Bytes(), &user); err != nil {
			return nil, err
		}
		if strings.Contains(user.Email, dm) {
			result[strings.ToLower(strings.SplitN(user.Email, "@", 2)[1])]++
		}
	}
	return result, nil
}
