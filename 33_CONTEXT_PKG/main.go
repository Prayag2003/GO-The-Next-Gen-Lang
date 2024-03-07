package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

type User struct {
	name   string
	userID int
	score  int
}

func callingThirdPartyAPI(ctx context.Context, userID int) (User, error) {
	time.Sleep(300 * time.Millisecond)

	user := User{}
	if ctx.Err() == context.DeadlineExceeded {
		return user, errors.New("Context Timeout exceeded")
	}

	name := "Prayag"
	score := 100
	user = User{
		name:   name,
		userID: userID,
		score:  score,
	}
	return user, nil
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 400*time.Millisecond)
	defer cancel() // to prevent leaking of context

	uID := 1
	user, err := callingThirdPartyAPI(ctx, uID)

	if err != nil {
		fmt.Println("\nTook too long to fetch the 3rd party data\n")
		return
	}
	fmt.Printf("\n { User name : %s, userID : %d, score : %d }\n", user.name, user.userID, user.score)

}
