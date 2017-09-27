package data

// Neighbours struct
type Neighbours struct {
	Batadv         map[string]BatadvNeighbours `json:"batadv"`
	LLDP           map[string]LLDPNeighbours   `json:"lldp"`
	Babel          map[string]BabelNeighbours  `json:"babel"`
	WifiNeighbours map[string]WifiNeighbours   `json:"wifi"`
	NodeID         string                      `json:"node_id"`
}

// WifiLink struct
type WifiLink struct {
	Inactive int `json:"inactive"`
	Noise    int `json:"nois"`
	Signal   int `json:"signal"`
}

// BatmanLink struct
type BatmanLink struct {
	Lastseen float64 `json:"lastseen"`
	Tq       int     `json:"tq"`
}

// BabelLink struct
type BabelLink struct {
	Address string `json:"address"`
}

// LLDPLink struct
type LLDPLink struct {
	Name        string `json:"name"`
	Description string `json:"descr"`
}

// BatadvNeighbours struct
type BatadvNeighbours struct {
	Neighbours map[string]BatmanLink `json:"neighbours"`
}

// BabelNeighbours struct
type BabelNeighbours []BabelLink

// WifiNeighbours struct
type WifiNeighbours struct {
	Neighbours map[string]WifiLink `json:"neighbours"`
}

// LLDPNeighbours struct
type LLDPNeighbours map[string]LLDPLink
