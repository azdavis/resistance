import { Dispatch } from "react";

export type GID = number;
export type CID = number;

// These should be kept in sync with types.go.
export type ToServer =
  // Client "sends" a Close by closing the connection.
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

type SelfAction = { t: "Close" } | { t: "AckDisbanded" } | { t: "LobbyLeave" };

// These should be kept in sync with types.go.
type ToClient =
  | { t: "RejectName" }
  | { t: "LobbyChoices"; Lobbies: Array<Lobby> }
  | { t: "CurrentLobby"; Me: CID; Leader: CID; Clients: Array<Client> }
  | { t: "SetIsSpy"; IsSpy: boolean }
  | { t: "NewMission"; Captain: CID; NumClients: number };

export type Action = SelfAction | ToClient;
export type D = Dispatch<Action>;

export type State =
  | { t: "Fatal"; s: State; a: Action }
  | { t: "Disbanded"; lobbies: Array<Lobby> }
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
      t: "MissionChoosing";
      captain: CID;
      me: CID;
      leader: CID;
      clients: Array<Client>;
      isSpy: boolean;
      numClients: number;
    };
