import { Dispatch } from "react";
import { CID, GID, Client, Lobby, ToServer, ToClient } from "./shared";

export type S = Dispatch<ToServer>;

type SelfAction =
  | { t: "Close" }
  | { t: "GoLangChooseFail"; msg: string }
  | { t: "GoLobbies" }
  | { t: "GoWelcome" }
  | { t: "GoNameChoose" }
  | { t: "GoLangChoose" }
  | { t: "GoHowTo" };

export type Action = SelfAction | ToClient;
export type D = Dispatch<Action>;

export type State =
  | { t: "LangChooseFail"; msg: string }
  | { t: "Invalid"; s: State; a: Action }
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

export type Lang = "en" | "ja";
export const langs: Array<[Lang, string]> = [
  ["en", "English"],
  ["ja", "日本語"],
];

export type Translation = {
  lang: Lang;
  resName: string;
  spyName: string;
  submit: string;
  leave: string;
  back: string;
  disbanded: string;
  disconnected: string;
  reconnect: string;
  invalid: string;
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
  lobbies: string;
  createNew: string;
  existingLobbies: (n: number) => string;
  lobbyWaiting: (n: number) => string;
  start: string;
  playerName: string;
  play: string;
  learnHow: string;
  setLang: string;
  viewCode: string;
  howToPlay: string;
  howTo: Array<string>;
};
