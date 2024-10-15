package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cdm command) error {
	err := s.db.Reset(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't reset user: %w", err)
	}
	fmt.Println("Users reset")
	return nil
}
