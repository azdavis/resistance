import { Dispatch } from "react";

export type PID = number;

export type PartyInfo = {
  PID: PID;
  Name: string;
};

export type State = {
  T:
    | "Closed"
    | "NameChoosing"
    | "PartyChoosing"
    | "PartyDisbanded"
    | "PartyWaiting";
  Parties: Array<PartyInfo>;
};

export type ToClient = Partial<State>;

export type ToServer =
  | { T: "NameChoose"; Name: string }
  | { T: "PartyChoose"; PID: PID }
  | { T: "PartyCreate"; Name: string };

export type Send = Dispatch<ToServer>;
