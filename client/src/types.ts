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
  | { t: "NewMission"; Captain: CID };

export type Action = SelfAction | ToClient;

export type State =
  | { t: "Fatal"; s: State; a: Action }
  | { t: "NameChoosing"; valid: boolean }
  | { t: "LobbyChoosing"; lobbies: Array<Lobby> }
  | {
      t: "LobbyWaiting";
      self: CID;
      leader: CID;
      clients: Array<Client>;
      isSpy: boolean;
    }
  | {
      t: "MissionMemberChoosing";
      captain: CID;
      self: CID;
      leader: CID;
      clients: Array<Client>;
      isSpy: boolean;
    };
