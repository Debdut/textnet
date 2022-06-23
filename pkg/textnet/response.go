package textnet

import st "github.com/debdut/textnet/pkg/state"

func Response(identifier string, message string) string {
	reply := "unknown command"
	state, ok := STATES[identifier]
	if !ok {
		state = st.State{
			Type:  "Null",
			State: &st.NullState{},
		}
	}

	// update state
	STATES[identifier] = state

	return reply
}
