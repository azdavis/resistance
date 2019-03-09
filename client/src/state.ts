import { Reducer } from "react";
import { State, Action } from "./types";

const reducer: Reducer<State, Action> = (s, a) => {
  switch (a.t) {
    case "Close":
      return { t: "Fatal", s, a };
    case "AckDisbanded":
      return s.t === "Disbanded"
        ? { ...s, t: "LobbyChoosing" }
        : { t: "Fatal", s, a };
    case "LobbyLeave":
      return s.t === "LobbyWaiting"
        ? { ...s, didLeave: true }
        : { t: "Fatal", s, a };
    case "GoNameChoose":
      return { t: "NameChoosing", valid: true };
    case "GoHowTo":
      return { t: "HowTo" };
    case "RejectName":
      return { t: "NameChoosing", valid: false };
    case "LobbyChoices":
      return s.t === "LobbyChoosing" ||
        s.t === "NameChoosing" ||
        (s.t === "LobbyWaiting" && s.didLeave)
        ? { t: "LobbyChoosing", lobbies: a.Lobbies }
        : { t: "Disbanded", lobbies: a.Lobbies };
    case "CurrentLobby":
      return {
        t: "LobbyWaiting",
        me: a.Me,
        leader: a.Leader,
        clients: a.Clients,
        isSpy: false,
        didLeave: false,
      };
    case "SetIsSpy":
      return s.t === "LobbyWaiting"
        ? { ...s, isSpy: a.IsSpy }
        : { t: "Fatal", s, a };
    case "NewMission":
      return s.t === "LobbyWaiting"
        ? {
            ...s,
            t: "MissionChoosing",
            captain: a.Captain,
            numClients: a.NumClients,
          }
        : { t: "Fatal", s, a };
  }
};

const init: State = { t: "NameChoosing", valid: true };

export { reducer, init };
