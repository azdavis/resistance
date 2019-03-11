// These should be kept in sync with types.go.
import { Dispatch } from "react";

export type GID = number;
export type CID = number;

type ToServer =
  // Client "sends" a Close by closing the connection.
  | { t: "NameChoose"; Name: string }
  | { t: "LobbyChoose"; GID: GID }
  | { t: "LobbyLeave" }
  | { t: "LobbyCreate" }
  | { t: "GameStart" }
  | { t: "MemberChoose"; Members: Array<CID> }
  | { t: "MemberVote"; Vote: boolean };

export type Send = Dispatch<ToServer>;

export type Lobby = { GID: GID; Leader: string };
export type Client = { CID: CID; Name: string };

type SelfAction =
  | { t: "Close" }
  | { t: "AckDisbanded" }
  | { t: "LobbyLeave" }
  | { t: "GoNameChoose" }
  | { t: "GoHowTo" }
  | { t: "AckRole" };

type ToClient =
  | { t: "RejectName" }
  | { t: "LobbyChoices"; Lobbies: Array<Lobby> }
  | { t: "CurrentLobby"; Me: CID; Leader: CID; Clients: Array<Client> }
  | { t: "SetIsSpy"; IsSpy: boolean }
  | { t: "NewMission"; Captain: CID; NumMembers: number }
  | { t: "MemberPropose"; Members: Array<CID> }
  | { t: "MemberResult"; Vote: boolean };

export type Action = SelfAction | ToClient;
export type D = Dispatch<Action>;

export type State =
  | { t: "Fatal"; s: State; a: Action }
  | { t: "Disbanded"; lobbies: Array<Lobby> }
  | { t: "HowTo" }
  | { t: "NameChoosing"; valid: boolean }
  | { t: "LobbyChoosing"; lobbies: Array<Lobby> }
  | {
      t: "LobbyWaiting";
      me: CID;
      leader: CID;
      clients: Array<Client>;
      isSpy: boolean;
      didLeave: boolean;
    }
  | {
      t: "MemberChoosing";
      me: CID;
      captain: CID;
      clients: Array<Client>;
      isSpy: boolean;
      numMembers: number;
    }
  | {
      t: "MemberVoting";
      me: CID;
      captain: CID;
      clients: Array<Client>;
      isSpy: boolean;
      members: Array<CID>;
    };
