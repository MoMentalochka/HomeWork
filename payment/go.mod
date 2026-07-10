module github.com/MoMentalochka/HomeWork/payment

require (
	github.com/MoMentalochka/HomeWork/shared v0.0.0-00010101000000-000000000000
	github.com/google/uuid v1.6.0
	google.golang.org/grpc v1.82.0
	github.com/stretchr/testify v1.10.0
)

require (
	go.opentelemetry.io/otel v1.44.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.44.0 // indirect
	golang.org/x/net v0.56.0 // indirect
	golang.org/x/sys v0.46.0 // indirect
	golang.org/x/text v0.38.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260630182238-925bb5da69e7 // indirect
	google.golang.org/protobuf v1.36.11 // indirect
)

go 1.26.4

replace github.com/MoMentalochka/HomeWork/shared => ./shared
