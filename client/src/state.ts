import { Reducer } from "react";
import { State, Action } from "./types";

export const init: State = { t: "Welcome", me: 0 };

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
    case "GoLobbies":
      return s.t === "Disbanded"
        ? { ...s, t: "LobbyChoosing" }
        : s.t === "LobbyWaiting"
        ? { ...s, didLeave: true }
        : { t: "Fatal", s, a };
    case "GoWelcome":
      return s.t === "HowTo" || s.t === "NameChoosing"
        ? { t: "Welcome", me: s.me }
        : { t: "Fatal", s, a };
    case "GoNameChoose":
      return s.t === "Welcome" && s.me !== 0
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
        (s.t === "LobbyWaiting" && s.didLeave)
        ? { t: "LobbyChoosing", me: s.me, lobbies: a.Lobbies }
        : s.me === 0
        ? { t: "Fatal", s, a }
        : { t: "Disbanded", me: s.me, lobbies: a.Lobbies };
    case "CurrentLobby":
      return s.t === "LobbyChoosing" || s.t === "LobbyWaiting"
        ? {
            t: "LobbyWaiting",
            me: s.me,
            gid: a.GID,
            clients: a.Clients,
            leader: a.Leader,
            didLeave: false,
          }
        : { t: "Fatal", s, a };
    case "CurrentGame":
      return s.t === "LobbyWaiting" || s.t === "GamePlaying"
        ? {
            t: "GamePlaying",
            me: s.me,
            gid: s.gid,
            clients: s.clients,
            isSpy: a.IsSpy,
            resPts: a.ResPts,
            spyPts: a.SpyPts,
            captain: a.Captain,
            members: a.Members,
            active: a.Active,
          }
        : { t: "Fatal", s, a };
    case "EndGame":
      return s.t === "GamePlaying"
        ? {
            t: "GameEnded",
            me: s.me,
            resPts: a.ResPts,
            spyPts: a.SpyPts,
            lobbies: a.Lobbies,
          }
        : { t: "Fatal", s, a };
  }
};
