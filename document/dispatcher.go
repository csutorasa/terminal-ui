package document

import "sync"

type DispatcherCallback func()

type Dispatcher struct {
	callbacks []DispatcherCallback
	mutex     sync.Mutex
}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		callbacks: []DispatcherCallback{},
	}
}

func (d *Dispatcher) Dispatch(callback DispatcherCallback) {
	if callback == nil {
		panic("callback is nil")
	}
	d.mutex.Lock()
	defer d.mutex.Unlock()
	d.callbacks = append(d.callbacks, callback)
}

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
