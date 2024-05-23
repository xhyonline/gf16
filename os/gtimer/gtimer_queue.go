// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gtimer

import (
	"container/heap"
	"github.com/gogf/gf/container/gtype"
	"math"
	"sync"
)

// priorityQueue is an abstract data type similar to a regular queue or stack data structure in which
// each element additionally has a "priority" associated with it. In a priority queue, an element with
// high priority is served before an element with low priority.
// priorityQueue is based on heap structure.
type priorityQueue struct {
	mu             sync.RWMutex
	heap           *priorityQueueHeap // the underlying queue items manager using heap.
	latestPriority *gtype.Int64       // latestPriority stores the most priority value of the heap, which is used to check if necessary to call the Pop of heap by Timer.
}

// priorityQueueHeap is a heap manager, of which the underlying `array` is a array implementing a heap structure.
type priorityQueueHeap struct {
	array []priorityQueueItem
}

// priorityQueueItem stores the queue item which has a `priority` attribute to sort itself in heap.
type priorityQueueItem struct {
	value    interface{}
	priority int64
}

// newPriorityQueue creates and returns a priority queue.
func newPriorityQueue() *priorityQueue {
	queue := &priorityQueue{
		heap: &priorityQueueHeap{
			array: make([]priorityQueueItem, 0),
		},
		latestPriority: gtype.NewInt64(math.MaxInt64),
	}
	heap.Init(queue.heap)
	return queue
}

// Len retrieves and returns the length of the queue.
func (q *priorityQueue) Len() int {
	q.mu.RLock()
	defer q.mu.RUnlock()
	return q.heap.Len()
}

// LatestPriority retrieves and returns the minimum and the most priority value of the queue.
func (q *priorityQueue) LatestPriority() int64 {
	return q.latestPriority.Val()
}

// Push pushes a value to the queue.
// The `priority` specifies the priority of the value.
// The lesser the `priority` value the higher priority of the `value`.
func (q *priorityQueue) Push(value interface{}, priority int64) {
	q.mu.Lock()
	heap.Push(q.heap, priorityQueueItem{
		value:    value,
		priority: priority,
	})
	q.mu.Unlock()
	// Update the minimum priority using atomic operation.
	for {
		latestPriority := q.latestPriority.Val()
		if priority >= latestPriority {
			break
		}
		if q.latestPriority.Cas(latestPriority, priority) {
			break
		}
	}
}

// Pop retrieves, removes and returns the most high priority value from the queue.
func (q *priorityQueue) Pop() interface{} {
	q.mu.Lock()
	if v := heap.Pop(q.heap); v != nil {
		item := v.(priorityQueueItem)
		q.mu.Unlock()
		// Update the minimum priority using atomic operation.
		for {
			latestPriority := q.latestPriority.Val()
			if item.priority >= latestPriority {
				break
			}
			if q.latestPriority.Cas(latestPriority, item.priority) {
				break
			}
		}
		return item.value
	} else {
		q.mu.Unlock()
	}
	return nil
}
