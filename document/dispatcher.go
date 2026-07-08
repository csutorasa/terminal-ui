package document

import "sync"

// Dispatcher callback function.
type DispatcherCallback func()

// Dispatcher allows to execute callbacks after the event processing is done.
type Dispatcher struct {
	callbacks []DispatcherCallback
	mutex     sync.Mutex
}

// Creates a new [Dispatcher].
func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		callbacks: []DispatcherCallback{},
	}
}

// Adds a new callback to the queue.
func (d *Dispatcher) Dispatch(callback DispatcherCallback) {
	if callback == nil {
		panic("callback is nil")
	}
	d.mutex.Lock()
	defer d.mutex.Unlock()
	d.callbacks = append(d.callbacks, callback)
}

// Executes all callbacks in the queue.
func (d *Dispatcher) Run() {
	for {
		callback := d.findFirst()
		if callback == nil {
			break
		}
		callback()
	}
}

func (d *Dispatcher) findFirst() DispatcherCallback {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	if len(d.callbacks) == 0 {
		return nil
	}
	first := d.callbacks[0]
	d.callbacks = d.callbacks[1:]
	return first
}
