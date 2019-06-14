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
  errorWithCode: code => `Error ${code}`,
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
  groupSize:
    "Groups of at least " +
    String(minN) +
    " and at most " +
    String(maxN) +
    " players may play.",
  groupNames: "Some players are spies. The rest are members of the resistance.",
  decideWinner:
    "The first of the spies and resistance to get " +
    String(maxPts) +
    " points wins the game.",
  rounds:
    "The game is played in rounds. " +
    "In every round of the game, a captain is chosen. " +
    "The captain chooses the mission members for this round.",
  occurVote:
    "When the captain has finished choosing, " +
    "all players vote on whether the mission occurs.",
  noOccur: "If the mission does not occur, the next round is started.",
  tooManyNoOccur:
    "If too many missions do not occur in a row, the spies get 1 point.",
  yesOccur:
    "If the mission does occur, " +
    "the members of the mission vote on whether the mission succeeds.",
  succeedPt: "If the mission succeeds, the resistance gets 1 point.",
  failPt: "If the mission fails, the spies get 1 point.",
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
};

export default en;
