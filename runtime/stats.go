package runtime

const DISABLED_AUTOUPDATER = "disabled"

// CounterMap to manage multiple values
type CounterMap map[string]uint32

// GlobalStats struct
type GlobalStats struct {
	Clients       uint32
	ClientsWifi   uint32
	ClientsWifi24 uint32
	ClientsWifi5  uint32
	Gateways      uint32
	Nodes         uint32

	Firmwares   CounterMap
	Models      CounterMap
	Autoupdater CounterMap
}

//NewGlobalStats returns global statistics for InfluxDB
func NewGlobalStats(nodes *Nodes) (result *GlobalStats) {
	result = &GlobalStats{
		Firmwares:   make(CounterMap),
		Models:      make(CounterMap),
		Autoupdater: make(CounterMap),
	}

	nodes.RLock()
	for _, node := range nodes.List {
		if node.Online {
			result.Nodes++
			if stats := node.Statistics; stats != nil {
				result.Clients += stats.Clients.Total
				result.ClientsWifi24 += stats.Clients.Wifi24
				result.ClientsWifi5 += stats.Clients.Wifi5
				result.ClientsWifi += stats.Clients.Wifi
			}
			if node.IsGateway() {
				result.Gateways++
			}
			if info := node.Nodeinfo; info != nil {
				result.Models.Increment(info.Hardware.Model)
				result.Firmwares.Increment(info.Software.Firmware.Release)
				if info.Software.Autoupdater.Enabled {
					result.Autoupdater.Increment(info.Software.Autoupdater.Branch)
				} else {
					result.Autoupdater.Increment(DISABLED_AUTOUPDATER)
				}
			}
		}
	}
	nodes.RUnlock()
	return
}

// Increment counter in the map by one
// if the value is not empty
func (m CounterMap) Increment(key string) {
	if key != "" {
		m[key]++
	}
}
