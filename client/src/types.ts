import { Dispatch } from "react";

export type GID = number;
export type CID = number;

export type ToServer =
  | { T: "NameChoose"; Name: string }
  | { T: "LobbyChoose"; GID: GID }
  | { T: "LobbyLeave" }
  | { T: "LobbyCreate" }
  | { T: "GameStart" };

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
  | { T: "Close" } // server will not actually ever send this.
  | { T: "RejectName" }
  | { T: "LobbyChoices"; Lobbies: Array<Lobby> }
  | { T: "CurrentLobby"; Self: CID; Leader: CID; Clients: Array<Client> };

export type State =
  | { T: "Closed" }
  | { T: "NameChoosing"; valid: boolean }
  | { T: "LobbyChoosing"; lobbies: Array<Lobby> }
  | { T: "LobbyWaiting"; self: CID; leader: CID; clients: Array<Client> };
