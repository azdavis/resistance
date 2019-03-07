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

export type ToClient =
  | { t: "Close" } // server will not actually ever send this.
  | { t: "RejectName" }
  | { t: "LobbyChoices"; Lobbies: Array<Lobby> }
  | { t: "CurrentLobby"; Self: CID; Leader: CID; Clients: Array<Client> };

export type State =
  | { t: "Closed" }
  | { t: "NameChoosing"; valid: boolean }
  | { t: "LobbyChoosing"; lobbies: Array<Lobby> }
  | { t: "LobbyWaiting"; self: CID; leader: CID; clients: Array<Client> };
