import { Dispatch } from "react";

export type PID = number;
export type CID = number;

export type PartyInfo = {
  PID: PID;
  Name: string;
};

export type ClientInfo = {
  CID: CID;
  Name: string;
};

export type State = { T: "NameChoosing" } | ToClient;

export type ToClient =
  | { T: "Closed" }
  | { T: "PartyChoosing"; Parties: Array<PartyInfo> }
  | { T: "PartyDisbanded"; Parties: Array<PartyInfo> }
  | { T: "PartyWaiting"; Clients: Array<ClientInfo> };

export type ToServer =
  | { T: "NameChoose"; Name: string }
  | { T: "PartyChoose"; PID: PID }
  | { T: "PartyCreate"; Name: string };

export type Send = Dispatch<ToServer>;
