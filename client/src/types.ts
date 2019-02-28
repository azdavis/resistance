import { Dispatch } from "react";

export type PID = number;

export type PartyInfo = {
  PID: PID;
  Name: string;
};

export type State =
  | { T: "Closed" }
  | { T: "NameChoosing" }
  | { T: "PartyChoosing"; Name: string; Parties: Array<PartyInfo> };

export type Action =
  | { T: "NameChoose"; Name: string }
  | { T: "PartyChoose"; PID: PID }
  | { T: "PartyCreate"; Name: string };

export type Send = Dispatch<Action>;
