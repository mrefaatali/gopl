package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	fmt.Println("Commencing count down, press return to abort")

	ticker := time.NewTicker(1*time.Second)

  for countdown:=10; countdown >0 ; countdown--{
    fmt.Println(countdown)
    select {
      case <-ticker.C:
      case <-abort:
        fmt.Println("ABORTED")
        ticker.Stop()
        return
    }
  }

}

/*
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		<-tick
	}
	select {
	case <-time.After(10 * time.Second):
		//do nothing
	case <-abort:
		fmt.Println("Launch aborted!")
		return
	}

	fmt.Println("LAUNCHED")
}

/*
func main() {
	fmt.Println("Commencing count down, press return to abort")
	tick := time.Tick(1 * time.Second)
  for countdown := 10; countdown>0; countdown-- {
    fmt.Println(countdown)
    <-tick
  }
  //launch()
}
*/
