import { Reducer } from "react";
import { CID, GID, Client, CurrentGame } from "./shared";
import { State, Action } from "./etc";

export const init: State = { t: "Welcome", me: 0 };

const mkGamePlaying = (
  me: CID,
  gid: GID,
  clients: Array<Client>,
  a: CurrentGame,
): State => ({
  t: "GamePlaying",
  me,
  gid,
  clients,
  isSpy: a.IsSpy,
  resPts: a.ResPts,
  spyPts: a.SpyPts,
  captain: a.Captain,
  members: a.Members === null ? a.NumMembers : a.Members,
  active: a.Active,
});

export const reducer: Reducer<State, Action> = (s, a) => {
  if (s.t === "Invalid" || s.t === "LangChooseFail") {
    return s;
  }
  switch (a.t) {
    case "Close":
      return {
        t: "Disconnected",
        me: s.me,
        code: a.code,
        game: s.t === "GamePlaying" ? { gid: s.gid, clients: s.clients } : null,
      };
    case "GoLangChooseFail":
      return { t: "LangChooseFail", msg: a.msg };
    case "SetMe":
      return s.t === "Welcome" || s.t === "Disconnected"
        ? { t: "Welcome", me: a.Me }
        : s.t === "HowTo"
        ? { t: "HowTo", me: a.Me }
        : s.t === "LangChoosing"
        ? { t: "LangChoosing", me: a.Me }
        : { t: "Invalid", s, a };
    case "GoLobbies":
      return s.t === "Disbanded" || s.t === "GameEnded"
        ? { ...s, t: "LobbyChoosing" }
        : s.t === "LobbyWaiting"
        ? { ...s, didLeave: true }
        : { t: "Invalid", s, a };
    case "GoWelcome":
      return s.t === "HowTo" || s.t === "NameChoosing" || s.t === "LangChoosing"
        ? { t: "Welcome", me: s.me }
        : { t: "Invalid", s, a };
    case "GoNameChoose":
      return s.t === "Welcome" && s.me !== 0
        ? { t: "NameChoosing", me: s.me, valid: true }
        : { t: "Invalid", s, a };
    case "GoLangChoose":
      return s.t === "Welcome"
        ? { t: "LangChoosing", me: s.me }
        : { t: "Invalid", s, a };
    case "GoHowTo":
      return s.t === "Welcome"
        ? { t: "HowTo", me: s.me }
        : { t: "Invalid", s, a };
    case "NameReject":
      return s.t === "NameChoosing"
        ? { ...s, valid: false }
        : { t: "Invalid", s, a };
    case "LobbyChoices":
      return s.t === "LobbyChoosing" ||
        s.t === "NameChoosing" ||
        (s.t === "LobbyWaiting" && s.didLeave)
        ? { t: "LobbyChoosing", me: s.me, lobbies: a.Lobbies }
        : s.me === 0
        ? { t: "Invalid", s, a }
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
        : { t: "Invalid", s, a };
    case "CurrentGame":
      return s.t === "LobbyWaiting" || s.t === "GamePlaying"
        ? mkGamePlaying(s.me, s.gid, s.clients, a)
        : s.t === "Disconnected" && s.game !== null
        ? mkGamePlaying(s.me, s.game.gid, s.game.clients, a)
        : { t: "Invalid", s, a };
    case "EndGame":
      return s.t === "GamePlaying"
        ? {
            t: "GameEnded",
            me: s.me,
            resPts: a.ResPts,
            spyPts: a.SpyPts,
            lobbies: a.Lobbies,
          }
        : { t: "Invalid", s, a };
  }
};
