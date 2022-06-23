package textnet

import (
	"strings"

	st "github.com/debdut/textnet/pkg/state"
)

func Response(identifier string, message string) string {
	state, ok := STATES[identifier]
	if !ok {
		state = st.State{
			Type:  "Null",
			State: &st.NullState{},
		}
	}

	processedMessage := strings.Trim(message, "\n\t .")
	action := st.IdentifyAction(state, processedMessage)
	if action.Type == "Unknown" {
		if action.Payload != nil {
			reply := action.Payload.(*st.Unknown).Message
			if reply != "" {
				return reply
			}
		}

		return "unknown action"
	}

	state, err := st.Machine(state, action)
	// TODO: explain error
	if err != nil {
		return err.Error()
	}

	// TODO: handle state messages

	// update state
	STATES[identifier] = state

	return "invalid action in context"
}
