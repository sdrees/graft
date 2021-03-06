// Copyright 2013-2017 Apcera Inc. All rights reserved.

package graft

// ChanHandler is a convenience handler when a user wants to simply use
// channels for the async handling of errors and state changes.
type ChanHandler struct {
	StateMachineHandler

	// Chan to receive state changes.
	stateChangeChan chan<- StateChange
	// Chan to receive errors.
	errorChan chan<- error
}

// StateChange captures "from" and "to" States for the ChanHandler.
type StateChange struct {
	// From is the previous state.
	From State

	// To is the new state.
	To State
}

// NewChanHandler returns a Handler implementation which uses channels for
// handling errors and state changes.
func NewChanHandler(scCh chan<- StateChange, errCh chan<- error) *ChanHandler {
	return NewChanHandlerWithStateMachine(new(defaultStateMachineHandler), scCh, errCh)
}

// NewChanHandlerWithStateMachine returns a Handler implementation which uses
// channels for handling errors and state changes and a StateMachineHandler for
// hooking into external state. The external state machine influences leader
// election votes.
func NewChanHandlerWithStateMachine(
	stateHandler StateMachineHandler,
	scCh chan<- StateChange,
	errCh chan<- error) *ChanHandler {
	return &ChanHandler{
		StateMachineHandler: stateHandler,
		stateChangeChan:     scCh,
		errorChan:           errCh,
	}
}

// Queue the state change onto the channel
func (chand *ChanHandler) StateChange(from, to State) {
	chand.stateChangeChan <- StateChange{From: from, To: to}
}

// Queue the error onto the channel
func (chand *ChanHandler) AsyncError(err error) {
	chand.errorChan <- err
}
