package main

import (
	"context"

	"github.com/Sigumaa/lfu"
)

func friends(client *lfu.Client) error {
	// Get Friends List
	friends, err := client.Friends(context.TODO(), 1)
	if err != nil {
		return err
	}
	if err := printJSON(friends); err != nil {
		return err
	}
	return nil
}
