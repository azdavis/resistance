import { Reducer } from "react";
import { State, Action } from "./types";

export const init: State = { t: "Welcome", me: null };

export const reducer: Reducer<State, Action> = (s, a) => {
  if (s.t === "Fatal") {
    return s;
  }
  switch (a.t) {
    case "Close":
      return { t: "Fatal", s, a };
    case "SetMe":
      return s.t === "Welcome"
        ? { ...s, me: a.Me }
        : s.t === "HowTo"
        ? { ...s, me: a.Me }
        : { t: "Fatal", s, a };
    case "AckDisbanded":
      return s.t === "Disbanded"
        ? { ...s, t: "LobbyChoosing" }
        : { t: "Fatal", s, a };
    case "LobbyLeave":
      return s.t === "LobbyWaiting"
        ? { ...s, didLeave: true }
        : { t: "Fatal", s, a };
    case "GoWelcome":
      return s.t === "HowTo" || s.t === "NameChoosing"
        ? { t: "Welcome", me: s.me }
        : { t: "Fatal", s, a };
    case "GoNameChoose":
      return s.t === "Welcome" && s.me !== null
        ? { t: "NameChoosing", me: s.me, valid: true }
        : { t: "Fatal", s, a };
    case "GoHowTo":
      return s.t === "Welcome"
        ? { t: "HowTo", me: s.me }
        : { t: "Fatal", s, a };
    case "NameReject":
      return s.t === "NameChoosing"
        ? { ...s, valid: false }
        : { t: "Fatal", s, a };
    case "LobbyChoices":
      return s.t === "LobbyChoosing" ||
        s.t === "NameChoosing" ||
        (s.t === "LobbyWaiting" && s.didLeave) ||
        (s.t === "MissionResultViewing" && s.didLeave)
        ? { t: "LobbyChoosing", me: s.me, lobbies: a.Lobbies }
        : s.me === null
        ? { t: "Fatal", s, a }
        : { t: "Disbanded", me: s.me, lobbies: a.Lobbies };
    case "CurrentLobby":
      return s.t === "LobbyChoosing" || s.t === "LobbyWaiting"
        ? {
            t: "LobbyWaiting",
            me: s.me,
            clients: a.Clients,
            leader: a.Leader,
            didLeave: false,
          }
        : { t: "Fatal", s, a };
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
