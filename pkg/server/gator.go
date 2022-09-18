package server

import "fmt"

func Hello(name string) (string, error) {
  if name == "" {
    return "", fmt.Errorf("name cannot be empty")
  }

  message := fmt.Sprintf("Hello %s", name)
  return message, nil
}
