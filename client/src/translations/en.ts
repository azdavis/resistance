import { Translation } from "../etc";
import { minN, maxN, maxPts } from "../shared";

const en: Translation = {
  code: "en",
  langName: "English",
  resName: "Resistance",
  spyName: "Spies",
  submit: "Submit",
  leave: "Leave",
  back: "Back",
  Disbanded: {
    title: "Disbanded",
    body: "The game or lobby you were in was disbanded.",
  },
  Disconnected: {
    title: "Disconnected",
    reconnect: "Reconnect",
  },
  Fatal: {
    title: "Fatal error",
    body: "An error occurred from which the application cannot recover.",
  },
  GamePlaying: {
    viewAllegiance: "View allegiance",
    captain: (x: string) => `Captain: ${x}`,
    members: (n: number) => `Members (${n}):`,
    beingChosen: "(being chosen)",
    succeedPrompt: "Should the mission succeed?",
    succeed: "Succeed",
    fail: "Fail",
    beingVotedOn: "(being voted on)",
    occurPrompt: "Should the mission occur?",
    occur: "Occur",
    notOccur: "Not occur",
  },
  HowTo: {
    title: "How to play",
    groupSize:
      "Groups of at least " +
      String(minN) +
      " and at most " +
      String(maxN) +
      " players may play.",
    groupNames:
      "Some players are spies. The rest are members of the resistance.",
    decideWinner:
      "The first of the spies and resistance to get " +
      String(maxPts) +
      " points wins the game.",
    captain:
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
    succeed: "If the mission succeeds, the resistance gets 1 point.",
    fail: "If the mission fails, the spies get 1 point.",
  },
  LangChoosing: {
    title: "Set language",
  },
  LobbyChoosing: {
    title: "Lobbies",
    create: "Create new",
    existing: (n: number) => `Existing lobbies (${n})`,
  },
  LobbyWaiting: {
    title: (n: number) => `Lobby (${n})`,
    start: "Start",
  },
  NameChoosing: {
    title: "Player name",
    invalid: "Invalid",
  },
  Welcome: {
    play: "Play",
    learnHow: "Learn how to play",
    setLang: "Set language",
    viewCode: "View source code",
  },
};

export default en;
