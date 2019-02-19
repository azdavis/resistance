import { Dispatch } from "react";

export type State =
  | { t: "closed" }
  | { t: "nameChoosing" }
  | { t: "roomChoosing"; name: string; rooms: Array<string> };

export type Msg = { t: "nameChoose"; name: string };

export type Send = Dispatch<Msg>;
