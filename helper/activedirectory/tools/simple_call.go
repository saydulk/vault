package tools

import (
	"fmt"
	"os"

	"github.com/hashicorp/vault/helper/activedirectory"
)

var (
	// ex. "ldap://138.91.247.105"
	rawURL = os.Getenv("TEST_LDAP_URL")

	// these can be left blank if the operation performed doesn't require them
	username = os.Getenv("TEST_LDAP_USERNAME")
	password = os.Getenv("TEST_LDAP_PASSWORD")
)

// main executes one call using a simple client pointed at the given instance.
func main() {

	config := newInsecureConfig()
	client := activedirectory.NewClient(config)

	baseDN := []string{"example, com"}
	filters := map[*activedirectory.Field][]string{
		activedirectory.FieldRegistry.GivenName: {"Sara", "Sarah"},
	}

	entries, err := client.Search(baseDN, filters)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("found %d entries:\n", len(entries))
	for _, entry := range entries {
		fmt.Printf("%s\n", entry)
	}
}

func newInsecureConfig() *activedirectory.Configuration {
	return &activedirectory.Configuration{
		Certificate:   "",
		InsecureTLS:   true,
		Password:      password,
		StartTLS:      false,
		TLSMinVersion: 771,
		TLSMaxVersion: 771,
		URL:           rawURL,
		Username:      username,
	}
}