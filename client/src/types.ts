// These should be kept in sync with types.go.
import { Dispatch } from "react";

export type GID = number;
export type CID = number;

type ToServer =
  | { t: "Connect" }
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

export type Lobby = { GID: GID; Leader: string };
export type Client = { CID: CID; Name: string };

type SelfAction =
  | { t: "Close" }
  | { t: "AckDisbanded" }
  | { t: "LobbyLeave" }
  | { t: "GoWelcome" }
  | { t: "GoNameChoose" }
  | { t: "GoHowTo" }
  | { t: "AckRole" }
  | { t: "AckMissionResult" }
  | { t: "GameLeave" };

type ToClient =
  | { t: "SetMe"; Me: CID }
  | { t: "NameReject" }
  | { t: "LobbyChoices"; Lobbies: Array<Lobby> }
  | { t: "CurrentLobby"; GID: GID; Leader: CID; Clients: Array<Client> }
  | { t: "BeginGame"; IsSpy: boolean; Captain: CID; NumMembers: number }
  | { t: "MemberPropose"; Members: Array<CID> }
  | { t: "MemberAccept" }
  | { t: "MemberReject"; Captain: CID; Members: number; SpyWin: boolean }
  | { t: "MissionResult"; Success: boolean; Captain: CID; Members: number };

export type Action = SelfAction | ToClient;
export type D = Dispatch<Action>;

export type State =
  | { t: "Fatal"; s: State; a: Action }
  | { t: "Disbanded"; me: CID; lobbies: Array<Lobby> }
  | { t: "Welcome"; me: CID }
  | { t: "HowTo"; me: CID }
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
      t: "RoleViewing";
      me: CID;
      gid: GID;
      clients: Array<Client>;
      isSpy: boolean;
      captain: CID;
      members: number | Array<CID>;
    }
  | {
      t: "MemberChoosing";
      me: CID;
      gid: GID;
      clients: Array<Client>;
      resWin: number;
      spyWin: number;
      captain: CID;
      members: number;
    }
  | {
      t: "MemberVoting";
      me: CID;
      gid: GID;
      clients: Array<Client>;
      resWin: number;
      spyWin: number;
      captain: CID;
      members: Array<CID>;
    }
  | {
      t: "MissionVoting";
      me: CID;
      gid: GID;
      clients: Array<Client>;
      resWin: number;
      spyWin: number;
      canVote: boolean;
    }
  | {
      t: "MissionResultViewing";
      me: CID;
      gid: GID;
      clients: Array<Client>;
      resWin: number;
      spyWin: number;
      success: boolean;
      captain: CID;
      members: number | Array<CID>;
      didLeave: boolean;
    };
