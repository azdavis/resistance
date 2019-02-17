type State = { t: "closed" } | { t: "nameChoosing" };

const init: State = { t: "nameChoosing" };

const reducer = (oldS: State, newS: State): State => ({ ...oldS, ...newS });

export { init, reducer };
