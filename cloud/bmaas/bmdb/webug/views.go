package webug

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"k8s.io/klog/v2"

	"source.monogon.dev/cloud/bmaas/bmdb/model"
	"source.monogon.dev/cloud/bmaas/bmdb/reflection"
)

// baseParams are passed to all rendered templates, and are consumed by tags in
// templates/base.html.
type baseParams struct {
	// Address to display in page header.
	BMDBAddress string
	// Schema version to display in page header.
	BMDBSchema string
}

// makeBase builds baseParams from information about the current connection.
func (s *server) makeBase() baseParams {
	address := fmt.Sprintf("%s@%s", s.conn.DatabaseName, s.conn.Address)
	if s.conn.InMemory {
		address += " (in memory)"
	}
	return baseParams{
		BMDBAddress: address,
		BMDBSchema:  s.curSchema().Version,
	}
}

// viewMachines renders a list of all machines in the BMDB.
func (s *server) viewMachines(w http.ResponseWriter, r *http.Request, args ...string) {
	start := time.Now()
	res, err := s.curSchema().GetMachines(r.Context(), &reflection.GetMachinesOpts{Strict: true})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "could not dump BMDB: %v", err)
		return
	}
	duration := time.Since(start)

	type params struct {
		Base       baseParams
		Query      string
		Machines   []*reflection.Machine
		NMachines  int
		RenderTime time.Duration
	}
	err = templates.ExecuteTemplate(w, "machines.html", &params{
		Base:       s.makeBase(),
		Query:      res.Query,
		Machines:   res.Data,
		NMachines:  len(res.Data),
		RenderTime: duration,
	})
	if err != nil {
		klog.Errorf("Template rendering failed: %v", err)
	}
}

// viewMachineDetail renders a detailed page for a single machine.
func (s *server) viewMachineDetail(w http.ResponseWriter, r *http.Request, args ...string) {
	mid, err := uuid.Parse(args[0])
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, "invalid machine UUID")
		return
	}

	opts := reflection.GetMachinesOpts{
		FilterMachine:   &mid,
		Strict:          true,
		ExpiredBackoffs: true,
	}
	res, err := s.curSchema().GetMachines(r.Context(), &opts)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "could not dump BMDB: %v", err)
		return
	}
	if len(res.Data) == 0 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "machine not found")
		return
	}
	machine := res.Data[0]

	// Params to pass to template.
	type sessionOrError struct {
		Session *model.Session
		Error   string
	}
	type params struct {
		Base    baseParams
		Machine *reflection.Machine

		HistoryError string
		History      []model.WorkHistory

		Sessions map[string]sessionOrError
	}
	p := params{
		Base:     s.makeBase(),
		Machine:  machine,
		Sessions: make(map[string]sessionOrError),
	}

	// History retrieval is performed with strict consistency guarantees, and thus
	// might block. Make sure we don't block the entire page.
	subQueriesCtx, subQueriesCtxC := context.WithTimeout(r.Context(), time.Second)
	defer subQueriesCtxC()
	history, err := s.conn.ListHistoryOf(subQueriesCtx, mid)
	if err != nil {
		p.HistoryError = err.Error()
	}

	// Same for sessions.
	for name, work := range machine.Work {
		sessions, err := s.conn.GetSession(subQueriesCtx, work.SessionID)
		if err != nil {
			p.Sessions[name] = sessionOrError{Error: err.Error()}
		} else {
			if len(sessions) == 0 {
				// This can happen if the session literally just disappeared.
				//
				// TODO(q3k): put all of these operations in a DB TX so that we don't end up with
				// possible inconsistencies?
				p.Sessions[name] = sessionOrError{Error: "not found"}
				continue
			}
			p.Sessions[name] = sessionOrError{Session: &sessions[0]}
		}
	}

	p.History = make([]model.WorkHistory, len(history))
	for i := 0; i < len(history); i += 1 {
		p.History[i] = history[len(history)-(i+1)]
	}
	if err := templates.ExecuteTemplate(w, "machine.html", &p); err != nil {
		klog.Errorf("Template rendering failed: %v", err)
	}
}

// viewProviderRedirects redirects a given provider and provider_id into a
// provider's web portal for more detailed information about an underlying
// machine.
func (s *server) viewProviderRedirect(w http.ResponseWriter, r *http.Request, args ...string) {
	providerUrls := map[string]string{
		"Equinix": "https://console.equinix.com/devices/%s/overview",
	}
	if providerUrls[args[0]] == "" {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Usage: /provider/Equinix/<id>")
		return
	}
	url := fmt.Sprintf(providerUrls[args[0]], args[1])
	http.Redirect(w, r, url, http.StatusFound)
}

// viewSession shows detailed information about a BMDB session.
func (s *server) viewSession(w http.ResponseWriter, r *http.Request, args ...string) {
	// TODO(q3k): implement this once we add session info to work history so that
	// this can actually display something useful.
	fmt.Fprintf(w, "underconstruction.gif")
}
