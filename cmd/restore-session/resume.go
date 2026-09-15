package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/zzusec/restore-session/internal/agent"
	"github.com/zzusec/restore-session/internal/agent/claude"
	"github.com/zzusec/restore-session/internal/agent/codex"
	"github.com/zzusec/restore-session/internal/i18n"
	"github.com/zzusec/restore-session/internal/session"
)

// resume hands the terminal to the agent's own command line so the picked
// session loads and the conversation continues. It returns an error only when
// the agent is one the browser cannot resume or its binary is missing; once
// the replacement process has taken over, control does not return here.
func resume(target agent.Agent, s session.Session, print *i18n.Printer) error {
	meta := target.Meta()

	binary, args, err := resumeCommand(meta.ID, s.ID)
	if err != nil {
		return err
	}

	path, lookupErr := exec.LookPath(binary)
	if lookupErr != nil {
		// Report in the user's language and stop: there is nothing to exec.
		return i18n.Errorf(i18n.CLINoBinary, i18n.Args{
			"agent": meta.Label, "error": lookupErr.Error(),
		})
	}

	env := os.Environ()
	// Codex locates its sessions through $CODEX_HOME. The listing may have been
	// read from an overridden home, so resume must point at the same tree.
	if meta.ID == codex.ID {
		env = append(env, "CODEX_HOME="+meta.Home)
	}

	fmt.Fprintln(os.Stderr, print.T(i18n.Resuming, i18n.Args{"session_id": s.ID}))
	return execReplace(path, append([]string{binary}, args...), env)
}

// resumeCommand builds the command line an agent uses to resume a session.
func resumeCommand(agentID, sessionID string) (binary string, args []string, err error) {
	switch agentID {
	case claude.ID:
		return claude.Binary, []string{"--resume", sessionID}, nil
	case codex.ID:
		return codex.Binary, []string{"resume", sessionID}, nil
	default:
		label := strings.ReplaceAll(agentID, "_", " ")
		return "", nil, i18n.Errorf(i18n.CLIResumeUnsupported, i18n.Args{"agent": label})
	}
}
