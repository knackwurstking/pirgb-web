module github.com/knackwurstking/pirgb-web

go 1.18

require (
	github.com/knackwurstking/alice v0.0.0-00010101000000-000000000000
	github.com/sirupsen/logrus v1.9.0
	nhooyr.io/websocket v1.8.7
)

require (
	github.com/klauspost/compress v1.10.3 // indirect
	github.com/stretchr/testify v1.8.0 // indirect
	golang.org/x/sys v0.0.0-20220818161305-2296e01440c6 // indirect
)

replace github.com/knackwurstking/alice => ../alice
