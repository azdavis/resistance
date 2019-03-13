// These should be kept in sync with consts.ts.
package main

// MinN is the minimum number of players in a game.
const MinN = 5

// MaxN is the maximum number of players in a game.
const MaxN = 7

// OkGameSize returns whether n is an acceptable number of players for the game.
func OkGameSize(n int) bool {
	return MinN <= n && n <= MaxN
}
