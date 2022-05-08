package proxy

import (
	"net"

	"github.com/pion/ice/v2"
	"go.uber.org/zap"

	"riasc.eu/wice/internal/log"
)

func CreateUDPMux(listenPort int) (ice.UDPMux, error) {
	conn, err := createFilteredSTUNConnection(listenPort)
	if err != nil {
		return nil, err
	}

	return ice.NewUDPMuxDefault(ice.UDPMuxParams{
		UDPConn: conn,
		Logger:  log.NewPionLoggerFactory(zap.L()).NewLogger("udpmux"),
	}), nil
}

func CreateUDPMuxSrflx() (ice.UniversalUDPMux, int, error) {
	// We do not need a filtered connection here as we anyway need to redirect
	// the non-STUN traffic via nftables

	conn, err := net.ListenUDP("udp", nil)
	if err != nil {
		return nil, 0, err
	}

	lAddr := conn.LocalAddr().(*net.UDPAddr)

	mux := ice.NewUniversalUDPMuxDefault(ice.UniversalUDPMuxParams{
		UDPConn: conn,
		Logger:  log.NewPionLoggerFactory(zap.L()).NewLogger("udpmuxsrflx"),
	})

	return mux, lAddr.Port, nil
}
