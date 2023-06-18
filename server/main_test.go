package main_test

import (
	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
	"github.com/stretchr/testify/require"
	svr "pact-contract-go/server"
	"testing"
)

const PACTS_PATH = "../client/post/pacts/myconsumer-myprovider.json"

func TestProviderIndexLocal(t *testing.T) {
	pact := &dsl.Pact{
		Provider: "MyProvider",
	}

	go svr.StartServer()

	resp, err := pact.VerifyProvider(t, types.VerifyRequest{
		ProviderBaseURL: "http://localhost:3000",
		PactURLs:        []string{PACTS_PATH},
	})
	require.NoError(t, err)
	require.NotNil(t, resp)

}
