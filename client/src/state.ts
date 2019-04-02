import { Reducer } from "react";
import {
  State,
  Action,
  CID,
  GID,
  Client,
  CurrentGame,
  LangState,
  LangAction,
} from "./types";

export const init: LangState = { lang: "en", t: "Welcome", me: 0 };

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

const inner: Reducer<State, Action> = (s, a) => {
  if (s.t === "Fatal") {
    return s;
  }
  switch (a.t) {
    case "Close":
      return {
        t: "Disconnected",
        me: s.me,
        game: s.t === "GamePlaying" ? { gid: s.gid, clients: s.clients } : null,
      };
    case "SetMe":
      return s.t === "Welcome" || s.t === "Disconnected"
        ? { t: "Welcome", me: a.Me }
        : s.t === "HowTo"
        ? { t: "HowTo", me: a.Me }
        : { t: "Fatal", s, a };
    case "GoLobbies":
      return s.t === "Disbanded" || s.t === "GameEnded"
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
        ? mkGamePlaying(s.me, s.gid, s.clients, a)
        : s.t === "Disconnected" && s.game !== null
        ? mkGamePlaying(s.me, s.game.gid, s.game.clients, a)
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

export const reducer: Reducer<LangState, LangAction> = (s, a) => {
  const { lang, ...rest } = s;
  if (a.t === "SetLang") {
    return { lang: a.lang, ...rest };
  }
  return { lang, ...inner(rest, a) };
};
