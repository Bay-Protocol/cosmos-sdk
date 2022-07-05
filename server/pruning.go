package server

import (
	"fmt"
	"strings"

	"github.com/spf13/cast"

	"github.com/cosmos/cosmos-sdk/server/types"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
)

// GetPruningOptionsFromFlags parses command flags and returns the correct
// PruningOptions. If a pruning strategy is provided, that will be parsed and
// returned, otherwise, it is assumed custom pruning options are provided.
func GetPruningOptionsFromFlags(appOpts types.AppOptions) (storetypes.PruningOptions, error) {
	strategy := strings.ToLower(cast.ToString(appOpts.Get(FlagPruning)))

	switch strategy {
	case storetypes.PruningOptionDefault, storetypes.PruningOptionNothing, storetypes.PruningOptionEverything:
		return storetypes.NewPruningOptionsFromString(strategy), nil

	case storetypes.PruningOptionCustom:
		opts := storetypes.NewPruningOptions(
			cast.ToUint64(appOpts.Get(FlagPruningKeepRecent)),
			1,
			cast.ToUint64(appOpts.Get(FlagPruningInterval)),
		)

		if err := opts.Validate(); err != nil {
			return opts, fmt.Errorf("invalid custom pruning options: %w", err)
		}

		return opts, nil

	default:
		return storetypes.PruningOptions{}, fmt.Errorf("unknown pruning strategy %s", strategy)
	}
}
