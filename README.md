# distance-util
Library to convert coordinates to distance(in Kilometers)

## Installation
`go get -u github.com/Vishalj32/go-distance-util`

**test.go:**

```
package main

import (
	"fmt"
	"github.com/Vishalj32/go-distance-util/distance"
	"log"
)

func main() {

	origin := distance.Coordinates{Latitude: 19.2183, Longitude: 72.9781}
	destination := distance.Coordinates{Latitude: 19.229473114013672, Longitude: 72.9699935913086}

	result, err := distance.GetDistanceFromCoordinates(origin, destination)
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("Distance is %f kms", result)
}

```

**>> go run test.go**
```
Distance is 1.505903 kms
```

