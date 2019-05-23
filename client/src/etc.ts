import { Dispatch } from "react";
import { CID, GID, Client, Lobby, ToServer, ToClient } from "./shared";

export type S = Dispatch<ToServer>;

export type Lang = "en" | "ja";
export const langs: Array<Lang> = ["en", "ja"];

type SelfAction =
  | { t: "Close" }
  | { t: "GoLobbies" }
  | { t: "GoWelcome" }
  | { t: "GoNameChoose" }
  | { t: "GoLangChoose" }
  | { t: "GoHowTo" }
  | { t: "SetLang"; lang: Lang };

export type Action = SelfAction | ToClient;
export type D = Dispatch<Action>;

type StateNoLang =
  | { t: "Fatal"; s: State; a: Action }
  | {
      t: "Disconnected";
      me: CID;
      game: { gid: GID; clients: Array<Client> } | null;
    }
  | { t: "Disbanded"; me: CID; lobbies: Array<Lobby> }
  | { t: "Welcome"; me: CID }
  | { t: "HowTo"; me: CID }
  | { t: "LangChoosing"; me: CID }
  | { t: "NameChoosing"; me: CID; valid: boolean }
  | { t: "LobbyChoosing"; me: CID; lobbies: Array<Lobby> }
  | {
      t: "LobbyWaiting";
      me: CID;
      gid: GID;
      clients: Array<Client>;
      leader: CID;
      didLeave: boolean;
    }
  | {
      t: "GamePlaying";
      me: CID;
      gid: GID;
      clients: Array<Client>;
      isSpy: boolean;
      resPts: number;
      spyPts: number;
      captain: CID;
      members: number | Array<CID>;
      active: boolean;
    }
  | {
      t: "GameEnded";
      me: CID;
      resPts: number;
      spyPts: number;
      lobbies: Array<Lobby>;
    };

export type State = { lang: Lang } & StateNoLang;

export type Translation = {
  code: string;
  langName: string;
  resName: string;
  spyName: string;
  submit: string;
  leave: string;
  back: string;
  Disbanded: {
    title: string;
    body: string;
  };
  Disconnected: {
    title: string;
    reconnect: string;
  };
  Fatal: {
    title: string;
    body: string;
  };
  GamePlaying: {
    viewAllegiance: string;
    captain: (x: string) => string;
    members: (n: number) => string;
    beingChosen: string;
    succeedPrompt: string;
    succeed: string;
    fail: string;
    beingVotedOn: string;
    occurPrompt: string;
    occur: string;
    notOccur: string;
  };
  HowTo: {
    title: string;
    groupSize: string;
    groupNames: string;
    decideWinner: string;
    captain: string;
    occurVote: string;
    noOccur: string;
    tooManyNoOccur: string;
    yesOccur: string;
    succeed: string;
    fail: string;
  };
  LangChoosing: {
    title: string;
  };
  LobbyChoosing: {
    title: string;
    create: string;
    existing: (n: number) => string;
  };
  LobbyWaiting: {
    title: (n: number) => string;
    start: string;
  };
  NameChoosing: {
    title: string;
    invalid: string;
  };
  Welcome: {
    play: string;
    learnHow: string;
    setLang: string;
    viewCode: string;
  };
};
