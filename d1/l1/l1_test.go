//+build !djavul

package l1_test

import (
	"bytes"
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"testing"

	"github.com/pkg/errors"
	"github.com/sanctuary/djavul/d1/diablo"
	"github.com/sanctuary/djavul/d1/gendung"
	"github.com/sanctuary/djavul/d1/l1"
	"github.com/sanctuary/djavul/d1/multi"
	"github.com/sanctuary/djavul/d1/quests"
)

func TestCreateDungeon(t *testing.T) {
	// Load level graphics.
	//
	//    - levels/l1data/l1.cel
	//    - levels/l1data/l1s.cel
	//    - levels/l1data/l1.til
	//    - levels/l1data/l1.min
	golden := []struct {
		// meta.
		dungeonName string
		// pre.
		dlvl    uint8
		dtype   gendung.DungeonType
		questID quests.QuestID
		seed    int32
		// post.
		tiles        string
		pieces       string
		arches       string
		transparency string
	}{
		{
			// meta.
			dungeonName: "Cathedral",
			// pre.
			dlvl:    1,
			dtype:   gendung.Cathedral,
			questID: quests.Invalid,
			seed:    123,
			// post.
			tiles:        "12a0410904ebf2507b6b7017f0ae191ae476686b",
			pieces:       "e15a7afb7505cb01b0b3d1befce5b8d4833ae1c6",
			arches:       "5438e3d7761025a2ee6f7fec155c840fc289f5dd",
			transparency: "1269467cb381070f72bc6c8e69938e88da7e58cc",
		},
		{
			// meta.
			dungeonName: "Cathedral (fix corners)",
			// pre.
			dlvl:    1,
			dtype:   gendung.Cathedral,
			questID: quests.Invalid,
			seed:    35,
			// post.
			tiles:        "69157cd5da2deab75788e134ade347f6fbf90bcc",
			pieces:       "813cfaf6eaba7bd186a1a0f6ded2e4f2c95af38b",
			arches:       "19848d30f077e077529476961d2e51080ce71c9a",
			transparency: "4ec6659fbc5c81a49741a61ad01455d926ecbcf1",
		},
		{
			// meta.
			dungeonName: "The Butcher",
			// pre.
			dlvl:    quests.QuestsData[quests.TheButcher].DLvl,
			dtype:   gendung.Cathedral,
			questID: quests.TheButcher,
			seed:    123,
			// post.
			tiles:        "659b95eec3e1c18d13b7f9932de108b88b356b9b",
			pieces:       "15f2209ff5d066cfd568a1eab77e4328d08474e8",
			arches:       "42941df3ada356ebf87ce2987d26a06c44da711a",
			transparency: "74c24e596ec57a91261bc3a559270f31d6811336",
		},
		{
			// meta.
			dungeonName: "Poisoned Water Supply",
			// pre.
			dlvl:    quests.QuestsData[quests.PoisonedWaterSupply].DLvl,
			dtype:   gendung.Cathedral,
			questID: quests.PoisonedWaterSupply,
			seed:    123,
			// post.
			tiles:        "f4525775a47ef083d85c7faf5560b9808ce203ff",
			pieces:       "cb2f37c9d04a39ec22c4171c6e95c88a260364e3",
			arches:       "87418c244b8123dbdb3439812a2e1d8af5032c21",
			transparency: "5d548a45afb50e56cd77847fd832822eee0b01e7",
		},
		{
			// meta.
			dungeonName: "Ogden's Sign",
			// pre.
			dlvl:    quests.QuestsData[quests.OgdensSign].DLvl,
			dtype:   gendung.Cathedral,
			questID: quests.OgdensSign,
			seed:    123,
			// post.
			tiles:        "3a54760d2ce39932f556dbb3ae924c8425e5f9ea",
			pieces:       "f6fcf0461dfad18da42b3d25dde5e60cdc7b4daf",
			arches:       "7e97023f45d78a37dffb569111762018e6b0c93f",
			transparency: "10156f455d85c0c4be6d26be23fc540776253aa9",
		},
	}

	for _, g := range golden {
		// Load level graphics.
		*gendung.DType = g.dtype
		diablo.LoadLevelGraphics()

		// Establish pre-conditions.
		*gendung.DLvl = g.dlvl
		*multi.MaxPlayers = 1
		for i := range quests.Quests {
			quests.Quests[i].ID = quests.QuestID(i)
			quests.Quests[i].Active = false
		}
		*gendung.IsQuestLevel = false
		if g.questID != quests.Invalid {
			quests.Quests[g.questID].Active = true
			quests.Quests[g.questID].DLvl = g.dlvl
		}
		// Generate dungeon based on the given seed.
		l1.CreateDungeon(g.seed, 0)
		if err := check(*gendung.TileIDMap, "tiles", g.seed, g.tiles); err != nil {
			t.Errorf("%s (dlvl=%d): %v", g.dungeonName, g.dlvl, errors.WithStack(err))
		}
		if err := check(*gendung.PieceIDMap, "pieces", g.seed, g.pieces); err != nil {
			t.Errorf("%s (dlvl=%d): %v", g.dungeonName, g.dlvl, errors.WithStack(err))
		}
		if err := check(*gendung.ArchNumMap, "arches", g.seed, g.arches); err != nil {
			t.Errorf("%s (dlvl=%d): %v", g.dungeonName, g.dlvl, errors.WithStack(err))
		}
		if err := check(*gendung.TransparencyMap, "transparency", g.seed, g.transparency); err != nil {
			t.Errorf("%s (dlvl=%d): %v", g.dungeonName, g.dlvl, errors.WithStack(err))
		}
	}
}

// check validates the data against the given SHA1 hashsum.
func check(data interface{}, name string, seed int32, want string) error {
	buf := &bytes.Buffer{}
	if err := binary.Write(buf, binary.LittleEndian, data); err != nil {
		return errors.WithStack(err)
	}
	sum := sha1.Sum(buf.Bytes())
	got := fmt.Sprintf("%040x", sum[:])
	if got != want {
		return errors.Errorf("SHA1 hash mismatch for %v, seed 0x%08X; expected %q, got %q", name, seed, want, got)
	}
	return nil
}
