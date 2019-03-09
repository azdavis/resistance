import { Reducer } from "react";
import { State, Action } from "./types";

const reducer: Reducer<State, Action> = (s, a) => {
  switch (a.t) {
    case "Close":
      return { t: "Fatal", s, a };
    case "RejectName":
      return { t: "NameChoosing", valid: false };
    case "LobbyChoices":
      return { t: "LobbyChoosing", lobbies: a.Lobbies };
    case "CurrentLobby":
      return {
        t: "LobbyWaiting",
        me: a.Me,
        leader: a.Leader,
        clients: a.Clients,
        isSpy: false,
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
