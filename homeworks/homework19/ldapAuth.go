package homework19

//import "strings"

func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

type ldap interface {
	ldapCheck(username string) bool
}

type ldapUser struct{}

func (r *ldapUser) ldapCheck(username string) bool {
	users := []string{"test1", "test2", "test3"}
	return Contains(users, username)
}

func checkUsername(username string, r ldap) bool {
	return r.ldapCheck(username)
}
