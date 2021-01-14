//memo package
package memo3

//entry
type entry struct {
	res   result
	ready chan struct{} //closed when res is ready
}

//memo
type Memo struct {
	requests chan request
}

//request
type request struct {
	key      string
	response chan<- result //the client wants a single result
}

//function
type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

//create new memo
func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.server(f)
	return memo
}

//memo the result of f
func (memo *Memo) Get(key string) (interface{}, error) {
	response := make(chan result)
	memo.requests <- request{key, response}
	res := <-response
	return res.value, res.err
}

func (memo *Memo) Close() { close(memo.requests) }

func (memo *Memo) server(f Func){
  cache := make(map[string]*entry)
  for req := range memo.requests {
    e := cache[req.key]
    if e == nil {
      e = &entry{ready: make(chan struct{})}
      cache[req.key] = e
      go e.call(f, req.key)
    }
    go e.deliver(req.response)
  }
}

func (e *entry) call(f Func, key string){
  e.res.value, e.res.err = f(key)
  close(e.ready) //broadcast the ready
}

func (e *entry) deliver(response chan<- result){
  <-e.ready //wait for the ready
  response <- e.res
}