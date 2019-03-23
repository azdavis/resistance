import { Reducer } from "react";
import { State, Action } from "./types";

export const init: State = { t: "Welcome" };

export const reducer: Reducer<State, Action> = (s, a) => {
  if (s.t === "Fatal") {
    return s;
  }
  switch (a.t) {
    case "Close":
      return { t: "Fatal", s, a };
    case "SetMe":
      return s;
    case "AckDisbanded":
      return s.t === "Disbanded"
        ? { ...s, t: "LobbyChoosing" }
        : { t: "Fatal", s, a };
    case "LobbyLeave":
      return s.t === "LobbyWaiting"
        ? { ...s, didLeave: true }
        : { t: "Fatal", s, a };
    case "GoWelcome":
      return { t: "Welcome" };
    case "GoNameChoose":
      return { t: "NameChoosing", valid: true };
    case "GoHowTo":
      return { t: "HowTo" };
    case "NameReject":
      return { t: "NameChoosing", valid: false };
    case "LobbyChoices":
      return s.t === "LobbyChoosing" ||
        s.t === "NameChoosing" ||
        (s.t === "LobbyWaiting" && s.didLeave) ||
        (s.t === "MissionResultViewing" && s.didLeave)
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
    case "FirstMission":
      return s.t === "LobbyWaiting"
        ? {
            t: "RoleViewing",
            me: s.me,
            clients: s.clients,
            isSpy: a.IsSpy,
            captain: a.Captain,
            members: a.Members,
          }
        : { t: "Fatal", s, a };
    case "AckRole":
      return s.t === "RoleViewing"
        ? typeof s.members === "number"
          ? {
              t: "MemberChoosing",
              me: s.me,
              clients: s.clients,
              resWin: 0,
              spyWin: 0,
              captain: s.captain,
              members: s.members,
            }
          : {
              t: "MemberVoting",
              me: s.me,
              clients: s.clients,
              resWin: 0,
              spyWin: 0,
              captain: s.captain,
              members: s.members,
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
        : s.t === "RoleViewing"
        ? { ...s, members: a.Members }
        : s.t === "MissionResultViewing"
        ? { ...s, members: a.Members }
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
    case "MemberReject":
      return s.t === "MemberVoting"
        ? {
            t: "MemberChoosing",
            me: s.me,
            clients: s.clients,
            resWin: s.resWin,
            spyWin: s.spyWin + (a.SpyWin ? 1 : 0),
            captain: a.Captain,
            members: a.Members,
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
            captain: a.Captain,
            members: a.Members,
            didLeave: false,
          }
        : { t: "Fatal", s, a };
    case "AckMissionResult":
      return s.t === "MissionResultViewing"
        ? typeof s.members === "number"
          ? {
              t: "MemberChoosing",
              me: s.me,
              clients: s.clients,
              resWin: s.resWin,
              spyWin: s.spyWin,
              captain: s.captain,
              members: s.members,
            }
          : {
              t: "MemberVoting",
              me: s.me,
              clients: s.clients,
              resWin: s.resWin,
              spyWin: s.spyWin,
              captain: s.captain,
              members: s.members,
            }
        : { t: "Fatal", s, a };
    case "GameLeave":
      return s.t === "MissionResultViewing"
        ? { ...s, didLeave: true }
        : { t: "Fatal", s, a };
  }
};
