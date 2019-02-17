type State = { t: "closed" } | { t: "nameChoosing" };

type Action = { t: "close" };

const init: State = { t: "nameChoosing" };

const reducer = (s: State, a: Action): State => {
  switch (a.t) {
    case "close":
      return { t: "closed" };
  }
};

export { init, reducer };
