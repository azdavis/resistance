import { Reducer } from "react";
import { State, Action } from "./types";

const init: State = { t: "NameChoosing", valid: true };

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
      return init;
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
            t: "MemberChoosing",
            captain: a.Captain,
            numMembers: a.NumMembers,
          }
        : { t: "Fatal", s, a };
    case "MemberPropose":
      return s.t === "MemberChoosing"
        ? {
            t: "MemberVoting",
            captain: s.captain,
            me: s.me,
            leader: s.leader,
            clients: s.clients,
            isSpy: s.isSpy,
            members: a.Members,
          }
        : { t: "Fatal", s, a };
  }
};

export { init, reducer };
