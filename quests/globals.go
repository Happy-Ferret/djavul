//+build !djavul

// Global variable wrappers for quests.cpp

package quests

// Global variables.
var (
	// QuestsData contains the data related to each quest ID.
	//
	// References:
	//    * https://github.com/sanctuary/notes/blob/master/enums.h#quest_id
	//
	// ref: 0x4A1AE0
	QuestsData = new([16]QuestData)

	// Quests contains the quests of the current game.
	//
	// PSX ref: 0x800DDA40
	// PSX def: QuestStruct quests[16];
	//
	// ref: 0x69BD10
	Quests = new([16]Quest)
)
