package examples

import (
	"fmt"

	"github.com/Implex-ltd/crapsolver/crapsolver"
)

func main() {
	Crap := crapsolver.NewSolver()

	token, err := Crap.Solve(&crapsolver.TaskConfig{
		SiteKey: "4c672d35-0701-42b2-88c3-78380b0db560",
		Domain:  "discord.com",
	})

	if err != nil {
		panic(err)
	}

	fmt.Println("solved:", token)
}
