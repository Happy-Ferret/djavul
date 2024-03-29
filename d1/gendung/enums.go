package gendung

// DungeonType specifies a dungeon type.
type DungeonType uint8

// Dungeon types.
const (
	Tristram  DungeonType = iota // dlvl:       0
	Cathedral                    // dlvl:  1 -  4
	Catacombs                    // dlvl:  5 -  8
	Caves                        // dlvl:  9 - 12
	Hell                         // dlvl: 13 - 16
	DTypeNone DungeonType = 0xFF
)

// DFlag represents a set of flags used for dungeon generation.
type DFlag uint8

// Dungeon generation flags.
const (
	DFlag08 DFlag = 0x08
)
