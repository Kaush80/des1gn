package main

import "fmt"

func newPublication(pubType string, name string, pg int, pub string) (iPublication, error) {
	if pubType == "newspaper" {
		return createNewspaper(name, pg, pub), nil
	}
	if pubType == "magazine" {
		return createMagazine(name, pg, pub), nil
	}
	return nil, fmt.Errorf("No such Publication type")
}
