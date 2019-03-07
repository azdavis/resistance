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
  | { T: "NameChoosing"; Valid: boolean }
  | { T: "LobbyChoosing"; Lobbies: Array<Lobby> }
  | { T: "LobbyWaiting"; Self: CID; Leader: CID; Clients: Array<Client> };

export type State = { T: "Closed" } | ToClient;
