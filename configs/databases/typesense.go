package databases

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/dhanyalvian/go-fiber-qurandb-restapi/configs"
	"github.com/typesense/typesense-go/v3/typesense"
)

var Ts *typesense.Client

func InitTypesense() {
	cfg := configs.Cfg
	cfgTs := cfg.Typesense
	var nodes []string

	for _, node := range cfgTs.Nodes {
		nodes = append(nodes, fmt.Sprintf("%s://%s:%s", node.Protocol, node.Hostname, node.Port))
	}

	Ts = typesense.NewClient(
		// typesense.WithServer(server),
		typesense.WithNodes(nodes),
		typesense.WithAPIKey(cfgTs.ApiKey),
		typesense.WithNumRetries(5),
		typesense.WithRetryInterval(1*time.Second),
		typesense.WithHealthcheckInterval(2*time.Minute),
	)
}

func GetTs() *typesense.Client {
	return Ts
}

type DebugTransport struct {
	Transport http.RoundTripper
}

func (dt *DebugTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Dump and print the request
	requestDump, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		fmt.Println("Request dump error:", err)
	} else {
		fmt.Println("HTTP Request:")
		fmt.Println(string(requestDump))
	}

	resp, err := dt.Transport.RoundTrip(req)
	if err != nil {
		return nil, err
	}

	// Dump and print the response
	responseDump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		fmt.Println("Response dump error:", err)
	} else {
		fmt.Println("HTTP Response:")
		fmt.Println(string(responseDump))
	}

	return resp, nil
}
