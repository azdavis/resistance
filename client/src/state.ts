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
        clients: a.Clients,
        leader: a.Leader,
        didLeave: false,
      };
    case "SetIsSpy":
      return s.t === "LobbyWaiting"
        ? {
            t: "RoleViewing",
            me: s.me,
            clients: s.clients,
            isSpy: a.IsSpy,
            mission: null,
          }
        : { t: "Fatal", s, a };
    case "NewMission":
      // TODO get mission when not first mission?
      return s.t === "RoleViewing"
        ? {
            ...s,
            mission: { captain: a.Captain, numMembers: a.NumMembers },
          }
        : { t: "Fatal", s, a };
    case "AckRole":
      return s.t === "RoleViewing" && s.mission !== null
        ? {
            t: "MemberChoosing",
            me: s.me,
            clients: s.clients,
            resWin: 0,
            spyWin: 0,
            captain: s.mission.captain,
            numMembers: s.mission.numMembers,
          }
        : { t: "Fatal", s, a };
    case "MemberPropose":
      return s.t === "MemberChoosing"
        ? {
            t: "MemberVoting",
            me: s.me,
            clients: s.clients,
            resWin: s.resWin,
            spyWin: s.spyWin,
            captain: s.captain,
            members: a.Members,
          }
        : { t: "Fatal", s, a };
    case "MemberAccept":
      return s.t === "MemberVoting"
        ? {
            t: "MissionVoting",
            me: s.me,
            clients: s.clients,
            resWin: s.resWin,
            spyWin: s.spyWin,
            canVote: s.members.includes(s.me),
          }
        : { t: "Fatal", s, a };
    case "MissionResult":
      return s.t === "MissionVoting"
        ? {
            t: "MissionResultViewing",
            me: s.me,
            clients: s.clients,
            resWin: s.resWin + (a.Success ? 1 : 0),
            spyWin: s.spyWin + (a.Success ? 0 : 1),
            success: a.Success,
          }
        : { t: "Fatal", s, a };
    case "AckMissionResult":
      return s;
  }
};

export { init, reducer };
