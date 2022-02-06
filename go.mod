module github.com/PeernetOfficial/Backend

go 1.16

require (
	github.com/PeernetOfficial/core v0.0.0-20220205200523-b76601e11c56
	github.com/PeernetOfficial/kcp v0.0.0-00010101000000-000000000000
	github.com/google/uuid v1.3.0
	github.com/gorilla/websocket v1.4.3-0.20210424162022-e8629af678b7
)

replace github.com/PeernetOfficial/kcp => C:\Data\Peernet\kcp\kcp

replace github.com/PeernetOfficial/core => C:\Data\Peernet\kcp\core_kcp
