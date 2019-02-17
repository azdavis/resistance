export type State =
  | { t: "closed" }
  | { t: "nameChoosing" }
  | { t: "roomChoosing"; name: string; rooms: Array<string> };
