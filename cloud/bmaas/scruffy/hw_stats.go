package scruffy

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/protobuf/proto"
	"k8s.io/klog/v2"

	"source.monogon.dev/cloud/bmaas/bmdb/model"
	"source.monogon.dev/cloud/bmaas/server/api"
)

// hwStatsRunner collects metrics from the machine hardware inventory in BMDB and
// exposes them as Prometheus metrics via a registry passed to newHWStatsRunner.
type hwStatsRunner struct {
	s *Server

	nodesPerRegion      *prometheus.GaugeVec
	memoryPerRegion     *prometheus.GaugeVec
	cpuThreadsPerRegion *prometheus.GaugeVec
}

// newHWStatsRunner builds a hwStatsRunner. The hwStatsRunner then has the
// given's Server BMDB connection bound to it and can perform actual database
// statistic gathering.
func newHWStatsRunner(s *Server, reg *prometheus.Registry) *hwStatsRunner {
	hwsr := &hwStatsRunner{
		s: s,

		nodesPerRegion: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "bmdb_hwstats_region_nodes",
		}, []string{"provider", "location"}),

		memoryPerRegion: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "bmdb_hwstats_region_ram_bytes",
		}, []string{"provider", "location"}),

		cpuThreadsPerRegion: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "bmdb_hwstats_region_cpu_threads",
		}, []string{"provider", "location"}),
	}
	reg.MustRegister(hwsr.nodesPerRegion, hwsr.memoryPerRegion, hwsr.cpuThreadsPerRegion)
	return hwsr
}

func (h *hwStatsRunner) run(ctx context.Context) {
	klog.Infof("Starting stats runner...")

	ti := time.NewTicker(time.Minute)

	for {
		err := h.runOnce(ctx)
		if err != nil {
			if errors.Is(err, ctx.Err()) {
				return
			}
			klog.Errorf("Stats run failed: %v", err)
		}
		select {
		case <-ti.C:
		case <-ctx.Done():
			klog.Infof("Exiting stats runner (%v)...", ctx.Err())
			return
		}
	}
}

// statsPerRegion are gathered and aggregated (summed) per region.
type statsPerRegion struct {
	nodes      uint64
	ramBytes   uint64
	numThreads uint64
}

// add a given AgentHardwareReport to this region's data.
func (s *statsPerRegion) add(hwrep *api.AgentHardwareReport) {
	s.nodes++
	s.ramBytes += uint64(hwrep.Report.MemoryInstalledBytes)
	for _, cpu := range hwrep.Report.Cpu {
		s.numThreads += uint64(cpu.HardwareThreads)
	}
}

// regionKey is used to uniquely identify each region per each provider.
type regionKey struct {
	provider model.Provider
	location string
}

func (r *regionKey) String() string {
	return fmt.Sprintf("%s/%s", r.provider, r.location)
}

func (h *hwStatsRunner) runOnce(ctx context.Context) error {
	sess, err := h.s.session(ctx)
	if err != nil {
		return err
	}

	var start uuid.UUID

	perRegion := make(map[regionKey]*statsPerRegion)
	var total statsPerRegion

	for {
		var res []model.ListMachineHardwareRow
		err = sess.Transact(ctx, func(q *model.Queries) error {
			res, err = q.ListMachineHardware(ctx, model.ListMachineHardwareParams{
				Limit:     100,
				MachineID: start,
			})
			return err
		})
		if err != nil {
			return err
		}
		klog.Infof("Machines: %d chunk", len(res))
		if len(res) == 0 {
			break
		}
		for _, row := range res {
			var hwrep api.AgentHardwareReport
			err = proto.Unmarshal(row.HardwareReportRaw.([]byte), &hwrep)
			if err != nil {
				klog.Warningf("Could not decode hardware report from %s: %v", row.MachineID, err)
				continue
			}

			if !row.ProviderLocation.Valid {
				klog.Warningf("%s has no provider location, skipping", row.MachineID)
				continue
			}

			key := regionKey{
				provider: row.Provider,
				location: row.ProviderLocation.String,
			}
			if _, ok := perRegion[key]; !ok {
				perRegion[key] = &statsPerRegion{}
			}
			perRegion[key].add(&hwrep)
			total.add(&hwrep)

			start = row.MachineID
		}
	}

	for k, st := range perRegion {
		labels := prometheus.Labels{
			"provider": string(k.provider),
			"location": k.location,
		}

		h.nodesPerRegion.With(labels).Set(float64(st.nodes))
		h.memoryPerRegion.With(labels).Set(float64(st.ramBytes))
		h.cpuThreadsPerRegion.With(labels).Set(float64(st.numThreads))
	}
	return nil
}
