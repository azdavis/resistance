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

export type LobbyInfo = {
  GID: GID;
  Leader: string;
};

export type ClientInfo = {
  CID: CID;
  Name: string;
};

export type ToClient =
  | { T: "NameChoosing"; Valid: boolean }
  | { T: "LobbyChoosing"; Lobbies: Array<LobbyInfo> }
  | { T: "LobbyWaiting"; Self: CID; Leader: CID; Clients: Array<ClientInfo> };

export type State = { T: "Closed" } | ToClient;
