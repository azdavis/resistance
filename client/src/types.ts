import { Dispatch } from "react";

export type State =
  | { T: "Closed" }
  | { T: "NameChoosing" }
  | { T: "RoomChoosing"; Name: string; Rooms: Array<string> };

export type Msg = { T: "NameChoose"; P: { Name: string } };

export type Send = Dispatch<Msg>;
