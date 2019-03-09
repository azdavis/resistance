import { Dispatch } from "react";

export type GID = number;
export type CID = number;

// These should be kept in sync with types.go.
export type ToServer =
  | { t: "NameChoose"; Name: string }
  | { t: "LobbyChoose"; GID: GID }
  | { t: "LobbyLeave" }
  | { t: "LobbyCreate" }
  | { t: "GameStart" }
  | { t: "MissionChoose"; Members: Array<CID> };

export type Send = Dispatch<ToServer>;

export type Lobby = {
  GID: GID;
  Leader: string;
};

export type Client = {
  CID: CID;
  Name: string;
};

type SelfAction = { t: "Close" };

// These should be kept in sync with types.go.
type ToClient =
  | { t: "RejectName" }
  | { t: "LobbyChoices"; Lobbies: Array<Lobby> }
  | { t: "CurrentLobby"; Me: CID; Leader: CID; Clients: Array<Client> }
  | { t: "SetIsSpy"; IsSpy: boolean }
  | { t: "NewMission"; Captain: CID; NumClients: number };

export type Action = SelfAction | ToClient;

export type State =
  | { t: "Fatal"; s: State; a: Action }
  | { t: "NameChoosing"; valid: boolean }
  | { t: "LobbyChoosing"; lobbies: Array<Lobby> }
  | {
      t: "LobbyWaiting";
      me: CID;
      leader: CID;
      clients: Array<Client>;
      isSpy: boolean;
    }
  | {
      t: "MissionChoosing";
      captain: CID;
      me: CID;
      leader: CID;
      clients: Array<Client>;
      isSpy: boolean;
      numClients: number;
    };
