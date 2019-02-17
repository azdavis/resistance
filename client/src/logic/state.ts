type State =
  | { t: "closed" }
  | { t: "nameChoosing" }
  | { t: "roomChoosing"; name: string; rooms: Array<string> };

const init: State = { t: "nameChoosing" };

const reducer = (oldS: State, newS: State): State => ({ ...oldS, ...newS });

export { init, reducer };
