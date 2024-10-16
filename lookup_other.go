//go:build !android && !ios

package main

func lookupName() string {
	return "Unknown"
}
