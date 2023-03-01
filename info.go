package main

import (
	"context"

	"github.com/Sigumaa/lfu"
)

func info(client *lfu.Client) error {
	// Get User Info
	info, err := client.Info(context.TODO(), 1)
	if err != nil {
		return err
	}
	if err := printJSON(info); err != nil {
		return err
	}
	return nil
}
