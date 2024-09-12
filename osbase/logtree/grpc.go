package logtree

import (
	"google.golang.org/grpc/grpclog"

	"source.monogon.dev/go/logging"
)

// GRPCify turns a LeveledLogger into a go-grpc compatible logger.
func GRPCify(logger logging.Leveled) grpclog.LoggerV2 {
	lp, ok := logger.(*leveledPublisher)
	if !ok {
		// Fail fast, as this is a programming error.
		panic("Expected *leveledPublisher in LeveledLogger from supervisor")
	}

	lp2 := *lp
	lp2.depth += 1

	return &leveledGRPCV2{
		lp: &lp2,
	}
}

type leveledGRPCV2 struct {
	lp *leveledPublisher
}

func (g *leveledGRPCV2) Info(args ...interface{}) {
	g.lp.Info(args...)
}

func (g *leveledGRPCV2) Infoln(args ...interface{}) {
	g.lp.Info(args...)
}

func (g *leveledGRPCV2) Infof(format string, args ...interface{}) {
	g.lp.Infof(format, args...)
}

func (g *leveledGRPCV2) Warning(args ...interface{}) {
	g.lp.Warning(args...)
}

func (g *leveledGRPCV2) Warningln(args ...interface{}) {
	g.lp.Warning(args...)
}

func (g *leveledGRPCV2) Warningf(format string, args ...interface{}) {
	g.lp.Warningf(format, args...)
}

func (g *leveledGRPCV2) Error(args ...interface{}) {
	g.lp.Error(args...)
}

func (g *leveledGRPCV2) Errorln(args ...interface{}) {
	g.lp.Error(args...)
}

func (g *leveledGRPCV2) Errorf(format string, args ...interface{}) {
	g.lp.Errorf(format, args...)
}

func (g *leveledGRPCV2) Fatal(args ...interface{}) {
	g.lp.Fatal(args...)
}

func (g *leveledGRPCV2) Fatalln(args ...interface{}) {
	g.lp.Fatal(args...)
}

func (g *leveledGRPCV2) Fatalf(format string, args ...interface{}) {
	g.lp.Fatalf(format, args...)
}

func (g *leveledGRPCV2) V(l int) bool {
	return g.lp.V(logging.VerbosityLevel(l)).Enabled()
}
