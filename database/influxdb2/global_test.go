package influxdb

import (
	"github.com/FreifunkBremen/yanic/data"
	"github.com/FreifunkBremen/yanic/runtime"
)

const (
	TEST_SITE   = "ffhb"
	TEST_DOMAIN = "city"
)

func createTestNodes() *runtime.Nodes {
	nodes := runtime.NewNodes(&runtime.NodesConfig{})

	nodeData := &runtime.Node{
		Online: true,
		Statistics: &data.Statistics{
			Clients: data.Clients{
				Total: 23,
			},
		},
		Nodeinfo: &data.Nodeinfo{
			NodeID: "abcdef012345",
			Hardware: data.Hardware{
				Model: "TP-Link 841",
			},
			System: data.System{
				SiteCode: TEST_SITE,
			},
		},
	}
	nodeData.Nodeinfo.Software.Firmware = &struct {
		Base    string `json:"base,omitempty"`
		Release string `json:"release,omitempty"`
	}{
		Release: "2016.1.6+entenhausen1",
	}
	nodeData.Nodeinfo.Software.Autoupdater = &struct {
		Enabled bool   `json:"enabled,omitempty"`
		Branch  string `json:"branch,omitempty"`
	}{
		Enabled: true,
		Branch:  "stable",
	}
	nodes.AddNode(nodeData)

	nodes.AddNode(&runtime.Node{
		Online: true,
		Statistics: &data.Statistics{
			Clients: data.Clients{
				Total: 2,
			},
		},
		Nodeinfo: &data.Nodeinfo{
			NodeID: "112233445566",
			Hardware: data.Hardware{
				Model: "TP-Link 841",
			},
		},
	})

	nodes.AddNode(&runtime.Node{
		Online: true,
		Nodeinfo: &data.Nodeinfo{
			NodeID: "0xdeadbeef0x",
			VPN:    true,
			Hardware: data.Hardware{
				Model: "Xeon Multi-Core",
			},
			System: data.System{
				SiteCode:   TEST_SITE,
				DomainCode: TEST_DOMAIN,
			},
		},
	})

	return nodes
}
