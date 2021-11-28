# gopixel

gopixel is a Go library for interacting with the [Hypixel API](https://api.hypixel.net/).

This software is alpha software and is subject to change, including but not limited to:

- The addition of new features,
- removal of old features,
- major changes to existing features,
- or minor changes to existing features.

Do not rely on this software in a crucial production environment, as breaking changes can be made to the source at any
moment, and what is currently presented is for lack of a better word, lackluster.

## Quick Start

```go
package main

import (
	"context"
	"fmt"
	"github.com/Gophigure/gopixel/client"
	"github.com/Gophigure/gopixel/httputil"
	"github.com/Gophigure/gopixel/hypixel"
)

const (
	APIKey hypixel.APIKey = "key"
	UUID   hypixel.UUID   = "uuid"
)

func main() {
	hypixelClient := client.New(context.Background(), APIKey, httputil.NewClient())

	player, err := hypixelClient.Player(UUID)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Player Name:", player.DisplayName)
}
```

## Support

For support, consider waiting a little bit, as there is no server ready for support regarding this project explicitly.
