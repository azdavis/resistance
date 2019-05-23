import React from "react";
import { minN, maxN, maxPts } from "../shared";

export default {
  resName: "Resistance",
  spyName: "Spies",
  submit: "Submit",
  leave: "Leave",
  back: "Back",
  Disbanded: {
    title: <h1>Disbanded</h1>,
    body: <p>The game or lobby you were in was disbanded.</p>,
  },
  Disconnected: {
    title: <h1>Disconnected</h1>,
    reconnect: "Reconnect",
  },
  Fatal: {
    title: <h1>Fatal error</h1>,
    body: <p>An error occurred from which the application cannot recover.</p>,
  },
  GamePlaying: {
    viewAllegiance: "View allegiance",
    captain: (x: string) => <div>Captain: {x}</div>,
    members: (n: number) => <div>Members ({n}):</div>,
    beingChosen: <div>(being chosen)</div>,
    succeedPrompt: "Should the mission succeed?",
    succeed: "Succeed",
    fail: "Fail",
    beingVotedOn: <div>(being voted on)</div>,
    occurPrompt: "Should the mission occur?",
    occur: "Occur",
    notOccur: "Not occur",
  },
  HowTo: {
    title: <h1>How to play</h1>,
    groupSize: (
      <p>
        Groups of at least {minN} and at most {maxN} players may play.
      </p>
    ),
    groupNames: (
      <p>Some players are spies. The rest are members of the resistance.</p>
    ),
    decideWinner: (
      <p>
        The first of the spies and resistance to get {maxPts} points wins the
        game.
      </p>
    ),
    captain: (
      <p>
        The game is played in rounds. In every round of the game, a captain is
        chosen. The captain chooses the mission members for this round.
      </p>
    ),
    occurVote: (
      <p>
        When the captain has finished choosing, all players vote on whether the
        mission occurs.
      </p>
    ),
    noOccur: <p>If the mission does not occur, the next round is started.</p>,
    tooManyNoOccur: (
      <p>If too many missions do not occur in a row, the spies get 1 point.</p>
    ),
    yesOccur: (
      <p>
        If the mission does occur, the members of the mission vote on whether
        the mission succeeds.
      </p>
    ),
    succeed: <p>If the mission succeeds, the resistance gets 1 point.</p>,
    fail: <p>If the mission fails, the spies get 1 point.</p>,
  },
  LangChoosing: {
    title: <h1>Set language</h1>,
    langNames: "English",
  },
  LobbyChoosing: {
    title: <h1>Lobbies</h1>,
    create: "Create new",
    existing: (n: number) => <h2>Existing lobbies ({n})</h2>,
  },
  LobbyWaiting: {
    title: (n: number) => <h1>Lobby ({n})</h1>,
    start: "Start",
  },
  NameChoosing: {
    title: <h1>Player name</h1>,
    invalid: "Invalid",
  },
  Welcome: {
    play: "Play",
    learnHow: "Learn how to play",
    setLang: "Set language",
    viewCode: "View source code",
  },
};
