package main

import "time"

type Response struct {
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

type auth struct {
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
