package main

import ()

type DataRate struct {
	Modulation   string  `json:"modulation"`
	Bandwith     float64 `json:"bandwith"`
	SpreadFactor int64   `json:"spreadFactor"`
}

type TxInfo struct {
	Frequency float64 `json:"frequency"`
	DataRate  `json:"dataRate"`
	Adr       bool   `json:"adr"`
	CodeRate  string `json:"codeRate"`
}

type RxInfo []struct {
	Mac     string  `json:"mac"`
	Time    string  `json:"time"`
	Rssi    float64 `json:"rssi"`
	LoRaSNR float64 `json:"loRaSNR"`
}

type (
	LoraMessage struct {
		DevEUI string `json:"devEUI"`
		RxInfo `json:"rxInfo"`
		TxInfo `json:"txInfo"`
		FCnt   int    `json:"fCnt"`
		FPort  int    `json:"fPort"`
		Data   string `json:"data"`
	}
)
