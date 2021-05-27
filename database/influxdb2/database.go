package influxdb

import (
	"context"

	influxdb "github.com/influxdata/influxdb-client-go/v2"
	influxdbAPI "github.com/influxdata/influxdb-client-go/v2/api"

	"github.com/FreifunkBremen/yanic/database"
)

const (
	MeasurementLink               = "link"        // Measurement for per-link statistics
	MeasurementNode               = "node"        // Measurement for per-node statistics
	MeasurementDHCP               = "dhcp"        // Measurement for DHCP server statistics
	MeasurementGlobal             = "global"      // Measurement for summarized global statistics
	CounterMeasurementFirmware    = "firmware"    // Measurement for firmware statistics
	CounterMeasurementModel       = "model"       // Measurement for model statistics
	CounterMeasurementAutoupdater = "autoupdater" // Measurement for autoupdater
	batchMaxSize                  = 1000
)

type Connection struct {
	database.Connection
	config   Config
	client   influxdb.Client
	writeAPI influxdbAPI.WriteAPI
}

type Config map[string]interface{}

func (c Config) Address() string {
	return c["address"].(string)
}
func (c Config) Token() (string, bool) {
	if d, ok := c["token"]; ok {
		return d.(string), true
	}
	return "", false
}
func (c Config) Username() string {
	return c["username"].(string)
}
func (c Config) Password() string {
	return c["password"].(string)
}
func (c Config) Organization() string {
	if d, ok := c["organization"]; ok {
		return d.(string)
	}
	return ""
}
func (c Config) Bucket() string {
	if d, ok := c["bucket"]; ok {
		return d.(string)
	}
	return ""
}
func (c Config) Tags() map[string]string {
	if c["tags"] != nil {
		return c["tags"].(map[string]string)
	}
	return nil
}

func init() {
	database.RegisterAdapter("influxdb2", Connect)
}
func Connect(configuration map[string]interface{}) (database.Connection, error) {
	var config Config
	config = configuration

	token, tokenOK := config.Token()
	// Make client
	client := influxdb.NewClientWithOptions(config.Address(), token, influxdb.DefaultOptions().SetBatchSize(batchMaxSize))
	if !tokenOK {
		ctx := context.Background()
		// The first call must be signIn
		err := client.UsersAPI().SignIn(ctx, config.Username(), config.Password())
		if err != nil {
			return nil, err
		}
	}

	writeAPI := client.WriteAPI(config.Organization(), config.Bucket())

	db := &Connection{
		config:   config,
		client:   client,
		writeAPI: writeAPI,
	}

	return db, nil
}

// Close all connection and clean up
func (conn *Connection) Close() {
	conn.writeAPI.Flush()
	conn.client.Close()
}
