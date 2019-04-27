// This should be kept in sync with server/shared.go.
import { Dispatch } from "react";

export const minN = 5;
export const maxN = 7;
export const okGameSize = (n: number): boolean => minN <= n && n <= maxN;
export const maxPts = 3;
export const maxSkip = 3;

export type GID = number;
export type CID = number;

type ToServer =
  | { t: "Connect" }
  | { t: "Reconnect"; Me: CID; GID: GID }
  | { t: "NameChoose"; Name: string }
  | { t: "LobbyChoose"; GID: GID }
  | { t: "LobbyLeave" }
  | { t: "LobbyCreate" }
  | { t: "GameStart" }
  | { t: "MemberChoose"; Members: Array<CID> }
  | { t: "MemberVote"; Vote: boolean }
  | { t: "MissionVote"; Vote: boolean }
  | { t: "GameLeave" };

export type Send = Dispatch<ToServer>;

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

export type CurrentGame = {
  t: "CurrentGame";
  IsSpy: boolean;
  ResPts: number;
  SpyPts: number;
  Captain: CID;
  NumMembers: number;
  Members: Array<CID> | null;
  Active: boolean;
};

export type Lobby = { GID: GID; Leader: string };
export type Client = { CID: CID; Name: string };

type ToClient =
  | { t: "SetMe"; Me: CID }
  | { t: "NameReject" }
  | { t: "LobbyChoices"; Lobbies: Array<Lobby> }
  | { t: "CurrentLobby"; GID: GID; Leader: CID; Clients: Array<Client> }
  | CurrentGame
  | {
      t: "EndGame";
      ResPts: number;
      SpyPts: number;
      Lobbies: Array<Lobby>;
    };

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
