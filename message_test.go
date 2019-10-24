// package rcss

// import (
// 	"bytes"
// 	"testing"
// )

// func TestMessage_MarshalBinary(t *testing.T) {
// 	fixtures := map[string]Message{
// 		"(init)": {name: "init"},

// 		"(init teamname)": {
// 			name: "init",
// 			values: []string{
// 				"teamname",
// 			},
// 		},

// 		"(init teamname (version 15))": {
// 			name: "init",
// 			values: []string{
// 				"teamname",
// 			},
// 			submessages: []Message{
// 				{name: "version", values: []string{"15"}},
// 			},
// 		},

// 		"(init teamname (goalie))": {
// 			name: "init",
// 			values: []string{
// 				"teamname",
// 			},
// 			submessages: []Message{
// 				{name: "goalie"},
// 			},
// 		},

// 		"(init teamname (version 15) (goalie))": {
// 			name: "init",
// 			values: []string{
// 				"teamname",
// 			},
// 			submessages: []Message{
// 				{name: "version", values: []string{"15"}},
// 				{name: "goalie"},
// 			},
// 		},

// 		"(hello Pouyan Heyratpour (from 15 (fullname AliAsghar Abniki) (when today (at noon))) (goalie))": {
// 			name: "hello",
// 			values: []string{
// 				"Pouyan",
// 				"Heyratpour",
// 			},
// 			submessages: []Message{
// 				{name: "from", values: []string{"15"}, submessages: []Message{
// 					{name: "fullname", values: []string{"AliAsghar", "Abniki"}},
// 					{name: "when", values: []string{"today"}, submessages: []Message{
// 						{name: "at", values: []string{"noon"}},
// 					}},
// 				}},
// 				{name: "goalie"},
// 			},
// 		},
// 	}

// 	for expected, msg := range fixtures {
// 		if b, err := msg.MarshalBinary(); err != nil {
// 			t.Errorf("msg to binary error: %s", err)
// 		} else if s := bytes.NewBufferString(expected); !bytes.Equal(s.Bytes(), b) {
// 			t.Errorf("expected %s got %s", s, b)
// 		} else {
// 			t.Logf("got %s", b)
// 		}
// 	}
// }

// func TestMessage_UnmarshalBinary(t *testing.T) {
// 	fixtures := []string{
// 		"(init)",
// 		"(init teamname)",
// 		"(init teamname (version 15))",
// 		"(init teamname (version 15) (goalie))",
// 		"(init teamname (goalie))",
// 		"(hello Pouyan Heyratpour (from 15 (fullname AliAsghar Abniki) (when today (at noon))) (goalie))",
// 	}

// 	for _, str := range fixtures {
// 		b := bytes.NewBufferString(str).Bytes()

// 		var msg Message
// 		if err := msg.UnmarshalBinary(b); err != nil {
// 			t.Errorf("msg decode error: %s", err)
// 		} else if out, err := msg.MarshalBinary(); err != nil {
// 			t.Errorf("msg encode error: %s", err)
// 		} else if !bytes.Equal(b, out) {
// 			t.Errorf("expected %s got %s", b, out)
// 		} else {
// 			t.Logf("got %s", out)
// 		}
// 	}
// }
package rcss
