package myabciapp

import "github.com/interchainio/tm-load-test/pkg/loadtest"

// MyABCIAppClientFactory creates instances of MyABCIAppClient
type MyABCIAppClientFactory struct {}

// MyABCIAppClientFactory implements loadtest.ClientFactory
var _ loadtest.ClientFactory = (*MyABCIAppClientFactory)(nil)

// MyABCIAppClient is responsible for generating transactions. Only one client
// will be created per connection to the remote Tendermint RPC endpoint, and
// each client will be responsible for maintaining its own state in a
// thread-safe manner.
type MyABCIAppClient struct {}

// MyABCIAppClient implements loadtest.Client
var _ loadtest.Client = (*MyABCIAppClient)(nil)

func (f *MyABCIAppClientFactory) ValidateConfig(cfg loadtest.Config) error {
    // Do any checks here that you need to ensure that the load test 
    // configuration is compatible with your client.
    return nil
}

func (f *MyABCIAppClientFactory) NewClient(cfg loadtest.Config) (loadtest.Client, error) {
    return &MyABCIAppClient{}, nil
}

// GenerateTx must return the raw bytes that make up the transaction for your
// ABCI app. The conversion to base64 will automatically be handled by the 
// loadtest package, so don't worry about that. Only return an error here if you
// want to completely fail the entire load test operation.
func (c *MyABCIAppClient) GenerateTx() ([]byte, error) {
    return []byte("this is my transaction"), nil
}
