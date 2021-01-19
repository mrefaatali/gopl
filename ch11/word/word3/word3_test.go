package word3

import (
	"math/rand"
	"testing"
	"time"
)

func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25)
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000)) //random rune up to \u0999
		runes[i] = r
		runes[n-1-i] = r
	}
	return string(runes)

}

func TestRandomPalindrome(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
  t.Logf("Random seed: %d", seed)
  rng := rand.New(rand.NewSource(seed))

  for i:=0; i<20;i++{
    p := randomPalindrome(rng)
    if !IsPalindrome(p){
      t.Errorf("IsPalindrome(%q) = false", p)
    }
  }
}
