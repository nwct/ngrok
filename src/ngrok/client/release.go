// +build release

package client

var (
	rootCrtPaths = []string{"assets/client/tls/ngrokroot.crt"}
)

func useInsecureSkipVerify() bool {
	return true
}
