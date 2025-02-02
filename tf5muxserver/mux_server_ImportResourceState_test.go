package tf5muxserver_test

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-mux/internal/tf5testserver"
	"github.com/hashicorp/terraform-plugin-mux/tf5muxserver"
)

func TestMuxServerImportResourceState(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	servers := []func() tfprotov5.ProviderServer{
		(&tf5testserver.TestServer{
			ResourceSchemas: map[string]*tfprotov5.Schema{
				"test_resource_server1": {},
			},
		}).ProviderServer,
		(&tf5testserver.TestServer{
			ResourceSchemas: map[string]*tfprotov5.Schema{
				"test_resource_server2": {},
			},
		}).ProviderServer,
	}

	muxServer, err := tf5muxserver.NewMuxServer(ctx, servers...)

	if err != nil {
		t.Fatalf("unexpected error setting up factory: %s", err)
	}

	_, err = muxServer.ProviderServer().ImportResourceState(ctx, &tfprotov5.ImportResourceStateRequest{
		TypeName: "test_resource_server1",
	})

	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if !servers[0]().(*tf5testserver.TestServer).ImportResourceStateCalled["test_resource_server1"] {
		t.Errorf("expected test_resource_server1 ImportResourceState to be called on server1")
	}

	if servers[1]().(*tf5testserver.TestServer).ImportResourceStateCalled["test_resource_server1"] {
		t.Errorf("unexpected test_resource_server1 ImportResourceState called on server2")
	}

	_, err = muxServer.ProviderServer().ImportResourceState(ctx, &tfprotov5.ImportResourceStateRequest{
		TypeName: "test_resource_server2",
	})

	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if servers[0]().(*tf5testserver.TestServer).ImportResourceStateCalled["test_resource_server2"] {
		t.Errorf("unexpected test_resource_server2 ImportResourceState called on server1")
	}

	if !servers[1]().(*tf5testserver.TestServer).ImportResourceStateCalled["test_resource_server2"] {
		t.Errorf("expected test_resource_server2 ImportResourceState to be called on server2")
	}
}
