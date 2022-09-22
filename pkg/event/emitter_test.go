package event

import (
	"context"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/kubeshop/testkube/pkg/api/v1/testkube"
	"github.com/kubeshop/testkube/pkg/event/bus"
	"github.com/kubeshop/testkube/pkg/event/kind/dummy"
	"github.com/stretchr/testify/assert"
)

type EventBusMock struct {
	events sync.Map
}

func (b *EventBusMock) Publish(event testkube.Event) error {
	b.events.Range(func(key, e interface{}) bool {
		e.(chan testkube.Event) <- event
		return true
	})
	return nil
}
func (b *EventBusMock) Subscribe(queue string, handler bus.Handler) error {

	ch := make(chan testkube.Event)
	b.events.Store(queue, ch)

	go func() {
		for e := range ch {
			handler(e)
		}
	}()
	return nil
}

func (b *EventBusMock) Unsubscribe(queue string) error {
	return nil

}
func (b *EventBusMock) Close() error {
	b.events.Range(func(key, e interface{}) bool {
		b.events.Delete(key)
		return true
	})
	return nil
}

var eventBus bus.Bus

func init() {
	os.Setenv("DEBUG", "true")
	eventBus = &EventBusMock{}
}

func TestEmitter_Register(t *testing.T) {

	t.Run("Register adds new listener", func(t *testing.T) {
		// given
		emitter := NewEmitter(eventBus)
		// when
		emitter.Register(&dummy.DummyListener{Id: "l1"})

		// then
		assert.Equal(t, 1, len(emitter.Listeners))

		t.Log("T1 completed")
	})
}

func TestEmitter_Listen(t *testing.T) {
	t.Run("listener handles only given events based on selectors", func(t *testing.T) {
		// given
		emitter := NewEmitter(eventBus)
		// given listener with matching selector
		listener1 := &dummy.DummyListener{Id: "l1", SelectorString: "type=listener1"}
		// and listener with second matic selector
		listener2 := &dummy.DummyListener{Id: "l2", SelectorString: "type=listener2"}

		// and emitter with registered listeners
		emitter.Register(listener1)
		emitter.Register(listener2)

		// listening emitter
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		emitter.Listen(ctx)
		// wait for listeners to start
		time.Sleep(time.Millisecond * 50)

		// events
		event1 := newExampleTestEvent1()
		event1.TestExecution.Labels = map[string]string{"type": "listener1"}
		event2 := newExampleTestEvent2()
		event2.TestExecution.Labels = map[string]string{"type": "listener2"}

		// when
		emitter.Notify(event1)
		emitter.Notify(event2)

		time.Sleep(time.Millisecond * 50)
		// then

		assert.Equal(t, 1, listener1.GetNotificationCount())
		assert.Equal(t, 1, listener2.GetNotificationCount())
	})

}

func TestEmitter_Notify(t *testing.T) {
	t.Run("notifies listeners in queue groups", func(t *testing.T) {
		// given
		emitter := NewEmitter(eventBus)
		// and 2 listeners subscribed to the same queue
		// * first on pod1
		listener1 := &dummy.DummyListener{Id: "l3"}
		// * second on pod2
		listener2 := &dummy.DummyListener{Id: "l3"}

		emitter.Register(listener1)
		emitter.Register(listener2)

		// and listening emitter
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		emitter.Listen(ctx)

		time.Sleep(time.Millisecond * 50)

		// when event sent to queue group
		emitter.Notify(newExampleTestEvent1())

		time.Sleep(time.Millisecond * 50)

		// then only one listener should be notified
		assert.Equal(t, 1, listener2.GetNotificationCount()+listener1.GetNotificationCount())
	})
}

func TestEmitter_Reconcile(t *testing.T) {

	t.Run("emitter refersh listeners in reconcile loop", func(t *testing.T) {
		// given first reconciler loop was done
		emitter := NewEmitter(eventBus)
		emitter.Loader.Register(&dummy.DummyLoader{IdPrefix: "dummy1"})
		emitter.Loader.Register(&dummy.DummyLoader{IdPrefix: "dummy2"})

		ctx, cancel := context.WithCancel(context.Background())

		go emitter.Reconcile(ctx)

		time.Sleep(100 * time.Millisecond)
		assert.Len(t, emitter.Listeners, 4)

		cancel()

		// and we'll add additional new loader
		emitter.Loader.Register(&dummy.DummyLoader{IdPrefix: "dummy1"}) // existing one
		emitter.Loader.Register(&dummy.DummyLoader{IdPrefix: "dummy3"})

		ctx, cancel = context.WithCancel(context.Background())

		// when
		go emitter.Reconcile(ctx)

		// then each reconciler (3 reconcilers) should load 2 listeners
		time.Sleep(100 * time.Millisecond)
		assert.Len(t, emitter.Listeners, 6)

		cancel()
	})

}

func newExampleTestEvent1() testkube.Event {
	return testkube.Event{
		Id:            "eventID1",
		Type_:         testkube.EventStartTest,
		TestExecution: testkube.NewExecutionWithID("executionID1", "test/test", "test"),
	}
}

func newExampleTestEvent2() testkube.Event {
	return testkube.Event{
		Id:            "eventID2",
		Type_:         testkube.EventStartTest,
		TestExecution: testkube.NewExecutionWithID("executionID2", "test/test", "test"),
	}
}
