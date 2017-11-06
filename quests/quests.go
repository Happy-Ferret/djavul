// Package quests implements quest functions.
package quests

import (
	"log"

	"github.com/sanctuary/djavul/gendung"
	"github.com/sanctuary/djavul/multi"
)

// isActive reports whether the given quest is active.
//
// NOTE: quest_num and quest_id are equivalent, as indicated by this function.
//
// PSX ref: 0x80067B70
// PSX def: unsigned char QuestStatus__Fi(int i)
//
// ref: 0x451831
func isActive(questNum QuestID) bool {
	if !Quests[questNum].Active {
		return false
	}
	if *gendung.DLvl != Quests[questNum].DLvl {
		return false
	}
	if *gendung.IsQuestLevel {
		return false
	}
	// Multiplayer quests:
	// * The Butcher
	// * The Curse of King Leoric
	// * Archbishop Lazarus
	// * Diablo
	if *multi.MaxPlayers != 1 && !QuestsData[questNum].Multiplayer {
		return false
	}
	return true
}

// initTheButcherArea initializes the quest area of The Butcher.
//
// PSX ref: 0x8015ED8C
// PSX def: void DrawButcher__Fv()
//
// ref: 0x451BEA
func initTheButcherArea() {
	// TODO: Implement initTheButcherArea.
	log.Print("note: initTheButcherArea not yet implemented.")
}

// initTheCurseOfKingLeoricArea initializes the quest area of The Curse of King
// Leoric.
//
// PSX ref: 0x8015EDD0
// PSX def: void DrawSkelKing__Fiii(int q, int x, int y)
//
// ref: 0x451C11
func initTheCurseOfKingLeoricArea(questID QuestID, xx, yy int32) {
	// TODO: Implement initTheCurseOfKingLeoricArea.
	log.Print("note: initTheCurseOfKingLeoricArea not yet implemented.")
}

// initWarlordOfBloodArea initializes the quest area of Warlord of Blood.
//
// PSX ref: 0x8015EE64
// PSX def: void DrawWarLord__Fii(int x, int y)
//
// ref: 0x451C32
func initWarlordOfBloodArea(xx, yy int32) {
	// TODO: Implement initWarlordOfBloodArea.
	log.Print("note: initWarlordOfBloodArea not yet implemented.")
}

// initTheChamberOfBoneArea initializes the quest area of The Chamber of Bone.
//
// PSX ref: 0x8015EF60
// PSX def: void DrawSChamber__Fiii(int q, int x, int y)
//
// ref: 0x451CC2
func initTheChamberOfBoneArea(questID QuestID, xx, yy int32) {
	// TODO: Implement initTheChamberOfBoneArea.
	log.Print("note: initTheChamberOfBoneArea not yet implemented.")
}

// initOdgensSignArea initializes the quest area of Odgen's Sign.
//
// PSX ref: 0x8015F09C
// PSX def: void DrawLTBanner__Fii(int x, int y)
//
// ref: 0x451D7C
func initOdgensSignArea(xx, yy int32) {
	// TODO: Implement initOdgensSignArea.
	log.Print("note: initOdgensSignArea not yet implemented.")
}

// initHallsOfTheBlindArea initializes the quest area of Halls of the Blind.
//
// PSX ref: 0x8015F178
// PSX def: void DrawBlind__Fii(int x, int y)
//
// ref: 0x451E08
func initHallsOfTheBlindArea(xx, yy int32) {
	// TODO: Implement initHallsOfTheBlindArea.
	log.Print("note: initHallsOfTheBlindArea not yet implemented.")
}

// initValorArea initializes the quest area of Valor.
//
// PSX ref: 0x8015F254
// PSX def: void DrawBlood__Fii(int x, int y)
//
// ref: 0x451E94
func initValorArea(xx, yy int32) {
	// TODO: Implement initValorArea.
	log.Print("note: initValorArea not yet implemented.")
}

// initQuestArea initializes the given quest area.
//
// PSX ref: 0x8015F334
// PSX def: void DRLG_CheckQuests__Fii(int x, int y)
//
// ref: 0x451F20
func initQuestArea(xx, yy int32) {
	questID := QuestID(0)
	for _, quest := range Quests {
		if IsActive(questID) {
			switch quest.ID {
			case TheButcher:
				InitTheButcherArea()
			case OgdensSign:
				InitOdgensSignArea(xx, yy)
			case HallsOfTheBlind:
				InitHallsOfTheBlindArea(xx, yy)
			case Valor:
				InitValorArea(xx, yy)
			case WarlordOfBlood:
				InitWarlordOfBloodArea(xx, yy)
			case TheCurseOfKingLeoric:
				InitTheCurseOfKingLeoricArea(questID, xx, yy)
			case TheChamberOfBone:
				InitTheChamberOfBoneArea(questID, xx, yy)
			}
		}
		questID++
	}
}
