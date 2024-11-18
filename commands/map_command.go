package commands

import (
	"errors"
	"fmt"
)

func mapCommand(c *config, args ...string) error {
	la, err := c.client.LocationAreas(c.next)
	if err != nil {
		return err
	}

	c.next = la.Next
	c.previous = la.Previous

	for _, area := range la.Results {
		c.write(fmt.Sprintf("%s\n", area.Name))
	}

	c.write("\n")
	return nil
}

func mapbCommand(c *config, args ...string) error {
	if c.previous == nil {
		return errors.New("cannot go back, you're on the first page")
	}

	la, err := c.client.LocationAreas(c.previous)
	if err != nil {
		return err
	}

	c.next = la.Next
	c.previous = la.Previous

	for _, area := range la.Results {
		c.write(fmt.Sprintf("%s\n", area.Name))
	}

	c.write("\n")
	return nil
}
