package main

import (
	"encoding/json"
	"os"
)

// Config for server
type Config struct {
	Listen        string `json:"listen"`
	Target        string `json:"target"`
	Key           string `json:"key"`
	Crypt         string `json:"crypt"`
	NoFullEncrypt bool   `json:"nofullenc"`
	Mode          string `json:"mode"`
	MTU           int    `json:"mtu"`
	SndWnd        int    `json:"sndwnd"`
	RcvWnd        int    `json:"rcvwnd"`
	DataShard     int    `json:"datashard"`
	ParityShard   int    `json:"parityshard"`
	DSCP          int    `json:"dscp"`
	NoComp        bool   `json:"nocomp"`
	AckNodelay    bool   `json:"acknodelay"`
	NoDelay       int    `json:"nodelay"`
	Interval      int    `json:"interval"`
	Resend        int    `json:"resend"`
	NoCongestion  int    `json:"nc"`
	SndBuf        int    `json:"sndbuf"`
	RcvBuf        int    `json:"rcvbuf"`
	SmuxBuf       int    `json:"smuxbuf"`
	StreamBuf     int    `json:"streambuf"`
	IOCopyBuf     int    `json:"iocopybuf"`
	SmuxVer       int    `json:"smuxver"`
	KeepAlive     int    `json:"keepalive"`
	Log           string `json:"log"`
	SnmpLog       string `json:"snmplog"`
	SnmpPeriod    int    `json:"snmpperiod"`
	Pprof         bool   `json:"pprof"`
	Quiet         bool   `json:"quiet"`
	TCP           bool   `json:"tcp"`
	IRTOBackOff   uint8  `json:"irtobackoff"`
	IRTOBThresh   int    `json:"irtobthresh"`
	NoEarRetran   bool   `json:"noearlyretran"`
}

func parseJSONConfig(config *Config, path string) error {
	file, err := os.Open(path) // For read access.
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewDecoder(file).Decode(config)
}
