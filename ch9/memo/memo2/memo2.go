//memo package
package memo2

import "sync"

//entry
type entry struct{
  res result
  ready chan struct{} //closed when res is ready
}

//memo
type Memo struct {
	f     Func
	mu    sync.Mutex
	cache map[string]*entry
}

//function
type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

//create new memo
func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

//memo the result of f
func (memo *Memo) Get(key string) (interface{}, error) {
	memo.mu.Lock()
  e := memo.cache[key]
	if e == nil {
    e = &entry{ready: make(chan struct{})}
    memo.cache[key] = e
    memo.mu.Unlock()

    e.res.value, e.res.err = memo.f(key)

    close(e.ready) //broadcast ready
  } else {
    memo.mu.Unlock()
    <-e.ready //wait for ready
  }
  return e.res.value, e.res.err

  /*
  res, ok := memo.cache[key]
	memo.mu.Unlock()
	if !ok {
		res.value, res.err = memo.f(key)
		memo.mu.Lock()
		memo.cache[key] = res
		memo.mu.Unlock()
	}
	return res.value, res.err
  */
}
