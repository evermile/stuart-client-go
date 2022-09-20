package main

import (
	"context"
	"fmt"
	stuartclient "github.com/evermile/stuart-client-go"
)

func main() {
	ctx := context.Background()
	client := stuartclient.NewClient(ctx, stuartclient.ProdEnv,
		"76abbb24859c310df25e3d6f943b7148c412c39b592949d0527392276140aaf4",
		"35697b6d9eea4a4d1be3fd01022030a4a28957f1c8f4865c6cca015748e4f08")

	coverage, err := client.GetZoneCoverage(ctx, "london", stuartclient.PickupType)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", coverage.Features)
}
