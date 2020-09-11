package bmoc

import (
	"github.com/system-pclub/GCatch/analysis/pointer"
	"github.com/system-pclub/GCatch/syncgraph"
)

func Detect() {
	stPtrResult, vecStOpValue := pointer.AnalyzeAllSyncOp()
	vecChannel := pointer.WithdrawAllChan(stPtrResult, vecStOpValue)
	pointer.WithdrawAllTraditionals(stPtrResult, vecStOpValue)

	mapDependency := syncgraph.GenDMap(vecChannel)
	_ = mapDependency // Use dependency map and vecChannel to build syncGraph
}
