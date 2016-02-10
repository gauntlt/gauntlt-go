package accounts

import (
	. "github.com/gauntlt/gauntlt-go"
	"github.com/stretchr/testify/assert"
)

func userIsLoggedIn() bool {
	return true
}

func generatePassword() {
	return
}

func init() {
	user, pass := "", ""

	Before("@login", func() {
		// runs before every feature or scenario tagged with @login
		generatePassword()
	})

	Given(`^I have user/pass "(.+?)" / "(.+?)"$`, func(u, p string) {
		user, pass = u, p
	})

	// ...

	Then(`^the user should be successfully logged in$`, func() {
		if !userIsLoggedIn() {
			assert.Exactly(nil, true, false, "Wrong")
		}
	})
}
