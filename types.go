package main

import "time"

type CasperResponse struct {
	APIVersion            string `json:"api_version"`
	ChainspecName         string `json:"chainspec_name"`
	StartingStateRootHash string `json:"starting_state_root_hash"`
	Peers                 []struct {
		NodeID  string `json:"node_id"`
		Address string `json:"address"`
	} `json:"peers"`
	LastAddedBlockInfo struct {
		Hash          string    `json:"hash"`
		Timestamp     time.Time `json:"timestamp"`
		EraID         int       `json:"era_id"`
		Height        int       `json:"height"`
		StateRootHash string    `json:"state_root_hash"`
		Creator       string    `json:"creator"`
	} `json:"last_added_block_info"`
	OurPublicSigningKey string      `json:"our_public_signing_key"`
	RoundLength         interface{} `json:"round_length"`
	NextUpgrade         struct {
		ActivationPoint int    `json:"activation_point"`
		ProtocolVersion string `json:"protocol_version"`
	} `json:"next_upgrade"`
	BuildVersion string `json:"build_version"`
}

type casperStructure struct {
	APIVersion            string
	ChainspecName         string
	StartingStateRootHash string
	Hash                  string
	Timestamp             time.Time
	EraID                 int
	Height                int
	OurPublicSigningKey   string
	BuildVersion          string
	ActivationPoint       int
	ProtocolVersion       string
	StateRootHash         string
	Creator               string
}

type aleoStructure struct {
	IsBootnode bool
	IsMiner    bool
	IsSyncing  bool
	Launched   time.Time
	Version    string
}

type AleoResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	Result  struct {
		IsBootnode    bool      `json:"is_bootnode"`
		IsMiner       bool      `json:"is_miner"`
		IsSyncing     bool      `json:"is_syncing"`
		Launched      time.Time `json:"launched"`
		ListeningAddr string    `json:"listening_addr"`
		Version       string    `json:"version"`
	} `json:"result"`
	ID string `json:"id"`
}
