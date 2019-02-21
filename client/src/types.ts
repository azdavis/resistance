import { Dispatch } from "react";

export type State =
  | { T: "closed" }
  | { T: "nameChoosing" }
  | { T: "roomChoosing"; Name: string; Rooms: Array<string> };

export type Msg = { T: "nameChoose"; P: { Name: string } };

export type Send = Dispatch<Msg>;
