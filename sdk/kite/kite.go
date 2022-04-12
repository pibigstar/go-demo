package main

import "github.com/koding/kite"

func main() {
	// Create a kite
	k := kite.New("math", "1.0.0")

	// Add our handler method with the name "square"
	k.HandleFunc("square", func(r *kite.Request) (interface{}, error) {
		a := r.Args.One().MustFloat64()
		result := a * a    // calculate the square
		return result, nil // send back the result
	}).DisableAuthentication()

	// Attach to a server with port 3636 and run it
	k.Config.Port = 3636
	k.Run()
}
