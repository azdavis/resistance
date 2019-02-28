import { Dispatch } from "react";

export type State =
  | { T: "Closed" }
  | { T: "NameChoosing" }
  | { T: "PartyChoosing"; Name: string; Parties: Array<string> };

export type PID = number;

export type Action =
  | { T: "NameChoose"; Name: string }
  | { T: "PartyChoose"; PID: PID }
  | { T: "PartyCreate"; Name: string };

export type Send = Dispatch<Action>;
