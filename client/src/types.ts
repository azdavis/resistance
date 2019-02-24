import { Dispatch } from "react";

export type State =
  | { T: "Closed" }
  | { T: "NameChoosing" }
  | { T: "PartyChoosing"; Name: string; Parties: Array<string> };

export type Msg = { T: "NameChoose"; P: { Name: string } };

export type Send = Dispatch<Msg>;
