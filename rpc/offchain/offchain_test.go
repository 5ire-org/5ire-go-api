package offchain

import (
	"os"
	"testing"

	"github.com/5ire-org/5ire-go-api/v4/client"
	"github.com/5ire-org/5ire-go-api/v4/config"
)

var offchain *Offchain

func TestMain(m *testing.M) {
	cl, err := client.Connect(config.Default().RPCURL)
	if err != nil {
		panic(err)
	}
	offchain = NewOffchain(cl)
	os.Exit(m.Run())
}
