package main

import (
	"testing"

	"github.com/zzusec/restore-session/internal/agent/claude"
	"github.com/zzusec/restore-session/internal/agent/codex"
)

// TestResumeCommand pins the command line each agent expects when one of its
// sessions is resumed. Claude and Codex take a bare session id under different
// flags; anything else is refused, which is what keeps the browser from trying
// to exec an agent it cannot resume.
func TestResumeCommand(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		agentID   string
		sessionID string
		binary    string
		args      []string
		wantErr   bool
	}{
		{
			name:      "claude uses --resume",
			agentID:   claude.ID,
			sessionID: "11111111-2222-3333-4444-555555555555",
			binary:    claude.Binary,
			args:      []string{"--resume", "11111111-2222-3333-4444-555555555555"},
		},
		{
			name:      "codex uses resume subcommand",
			agentID:   codex.ID,
			sessionID: "abcdef00-0000-0000-0000-000000000000",
			binary:    codex.Binary,
			args:      []string{"resume", "abcdef00-0000-0000-0000-000000000000"},
		},
		{
			name:      "an unsupported agent is refused",
			agentID:   "opencode",
			sessionID: "irrelevant",
			wantErr:   true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			binary, args, err := resumeCommand(test.agentID, test.sessionID)
			if test.wantErr {
				if err == nil {
					t.Fatalf("resumeCommand(%q): want error, got nil", test.agentID)
				}
				return
			}
			if err != nil {
				t.Fatalf("resumeCommand(%q): unexpected error: %v", test.agentID, err)
			}
			if binary != test.binary {
				t.Errorf("binary = %q, want %q", binary, test.binary)
			}
			if len(args) != len(test.args) {
				t.Fatalf("args = %v, want %v", args, test.args)
			}
			for i := range test.args {
				if args[i] != test.args[i] {
					t.Errorf("args[%d] = %q, want %q", i, args[i], test.args[i])
				}
			}
		})
	}
}
