import { Dispatch } from "react";

export type PID = number;
export type CID = number;

export type PartyInfo = {
  PID: PID;
  Leader: string;
};

export type ClientInfo = {
  CID: CID;
  Name: string;
};

export type State = { T: "NameChoosing" } | { T: "Closed" } | ToClient;

export type ToClient =
  | { T: "PartyChoosing"; Parties: Array<PartyInfo> }
  | { T: "PartyWaiting"; Leader: CID; Clients: Array<ClientInfo> };

export type ToServer =
  | { T: "NameChoose"; Name: string }
  | { T: "PartyChoose"; PID: PID }
  | { T: "PartyLeave" }
  | { T: "PartyCreate" };

export type Send = Dispatch<ToServer>;
