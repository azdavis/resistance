import { Translation } from "../etc";
import { minN, maxN, maxPts } from "../shared";

const en: Translation = {
  lang: "en",
  resName: "Resistance",
  spyName: "Spies",
  submit: "Submit",
  leave: "Leave",
  back: "Back",
  disbanded: "Disbanded",
  disconnected: "Disconnected",
  reconnect: "Reconnect",
  invalid: "Invalid",
  invalidStateTransition: "An invalid state transition occurred.",
  viewAllegiance: "View allegiance",
  captain: x => `Captain: ${x}`,
  members: n => `Members (${n}):`,
  beingChosen: "(being chosen)",
  succeedPrompt: "Should the mission succeed?",
  succeed: "Succeed",
  fail: "Fail",
  beingVotedOn: "(being voted on)",
  occurPrompt: "Should the mission occur?",
  occur: "Occur",
  notOccur: "Not occur",
  howToPlay: "How to play",
  setLang: "Set language",
  lobbies: "Lobbies",
  createNew: "Create new",
  existingLobbies: n => `Existing lobbies (${n})`,
  lobbyWaiting: n => `Lobby (${n})`,
  start: "Start",
  playerName: "Player name",
  play: "Play",
  learnHow: "Learn how to play",
  viewCode: "View source code",
  howTo: [
    `Groups of at least ${minN} and at most ${maxN} players may play.`,
    "Some players are spies. The rest are members of the resistance. " +
      `The first of the spies and resistance to get ${maxPts} wins the game.`,
    "The game is played in rounds. " +
      "In every round of the game, a captain is chosen. " +
      "The captain chooses the mission members for this round. " +
      "When the captain has finished choosing, " +
      "all players vote on whether the mission occurs.",
    "If the mission does not occur, the next round is started. " +
      "If too many missions do not occur in a row, the spies get 1 point. " +
      "If the mission does occur, " +
      "the members of the mission vote on whether the mission succeeds.",
    "If the mission succeeds, the resistance gets 1 point. " +
      "If the mission fails, the spies get 1 point.",
  ],
};

export default en;
