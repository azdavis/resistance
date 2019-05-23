import { Dispatch } from "react";
import { CID, GID, Client, Lobby, ToServer, ToClient } from "./shared";

export type Lang = "en" | "ja";
export const langs: Array<Lang> = ["en", "ja"];
export const langNames: { [L in Lang]: string } = {
  en: "English",
  ja: "日本語",
};

export type S = Dispatch<ToServer>;

type SelfAction =
  | { t: "Close" }
  | { t: "GoSetLangFail"; msg: string }
  | { t: "GoLobbies" }
  | { t: "GoWelcome" }
  | { t: "GoNameChoose" }
  | { t: "GoLangChoose" }
  | { t: "GoHowTo" };

export type Action = SelfAction | ToClient;
export type D = Dispatch<Action>;

export type State =
  | { t: "Invalid"; s: State; a: Action }
  | { t: "SetLangFail"; msg: string }
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

export type Translation = {
  code: string;
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
  Invalid: {
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
