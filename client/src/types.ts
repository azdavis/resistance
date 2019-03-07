import { Dispatch } from "react";

export type GID = number;
export type CID = number;

export type ToServer =
  | { t: "NameChoose"; Name: string }
  | { t: "LobbyChoose"; GID: GID }
  | { t: "LobbyLeave" }
  | { t: "LobbyCreate" }
  | { t: "GameStart" };

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

type ToClient =
  | { t: "RejectName" }
  | { t: "LobbyChoices"; Lobbies: Array<Lobby> }
  | { t: "CurrentLobby"; Self: CID; Leader: CID; Clients: Array<Client> }
  | { t: "SetIsSpy"; IsSpy: boolean }
  | { t: "NewMission"; IsCaptain: boolean };

export type Action = SelfAction | ToClient;

export type State =
  | { t: "Invalid"; s: State; a: Action }
  | { t: "NameChoosing"; valid: boolean }
  | { t: "LobbyChoosing"; lobbies: Array<Lobby> }
  | {
      t: "LobbyWaiting";
      self: CID;
      leader: CID;
      clients: Array<Client>;
      isSpy: boolean;
    };
