package main

import (
	"github.com/mikerybka/web/types"
	"github.com/mikerybka/webgen"
)

func main() {
	admin := struct {
		Users []types.User
	}{
		Users: []types.User{
			{
				FirstName: "Mike",
				LastName:  "Rybka",
				Email:     "",
			},
		},
	}
	webgen.Write("cmd/test/output", admin)
}
