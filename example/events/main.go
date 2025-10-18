package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Nadim147c/go-hyprland"
)

func main() {
	client := hyprland.NewEventListener()
	client.OnAllEvents(func(ctx *hyprland.EventContext) {
		fmt.Println(ctx.Time.Format(time.TimeOnly), ctx.RawEvent)
	})
	log.Fatal(client.Listen(context.Background()))
}
