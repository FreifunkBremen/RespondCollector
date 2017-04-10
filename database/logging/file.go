package logging

/**
 * This database type is just for,
 * - debugging without a influxconn
 * - example for other developers for new databases
 */
import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/FreifunkBremen/yanic/database"
	"github.com/FreifunkBremen/yanic/runtime"
)

type Connection struct {
	database.Connection
	config Config
	file   *os.File
}

type Config map[string]interface{}

func (c Config) Enable() bool {
	return c["enable"].(bool)
}
func (c Config) Path() string {
	return c["path"].(string)
}

func init() {
	database.AddDatabaseType("logging", Connect)
}

func Connect(configuration interface{}) (database.Connection, error) {
	var config Config
	config = configuration.(map[string]interface{})
	if !config.Enable() {
		return nil, nil
	}

	file, err := os.OpenFile(config.Path(), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return nil, err
	}
	return &Connection{config: config, file: file}, nil
}

func (conn *Connection) AddNode(nodeID string, node *runtime.Node) {
	conn.log("AddNode: [", nodeID, "] clients: ", node.Statistics.Clients.Total)
}

func (conn *Connection) AddStatistics(stats *runtime.GlobalStats, time time.Time) {
	conn.log("AddStatistics: [", time.String(), "] nodes: ", stats.Nodes, ", clients: ", stats.Clients, " models: ", len(stats.Models))
}

func (conn *Connection) DeleteNode(deleteAfter time.Duration) {
	conn.log("DeleteNode")
}

func (conn *Connection) Close() {
	conn.log("Close")
	conn.file.Close()
}

func (conn *Connection) log(v ...interface{}) {
	log.Println(v)
	conn.file.WriteString(fmt.Sprintln("[", time.Now().String(), "]", v))
}
