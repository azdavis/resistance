import { Reducer } from "react";
import { State, Action, CID, GID, Client, CurrentGame, Lang } from "./shared";
import { getLang } from "./storage";

export const init: State = {
  lang: getLang() || "en",
  t: "Welcome",
  me: 0,
};

const mkGamePlaying = (
  lang: Lang,
  me: CID,
  gid: GID,
  clients: Array<Client>,
  a: CurrentGame,
): State => ({
  t: "GamePlaying",
  lang,
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
  if (s.t === "Fatal") {
    return s;
  }
  if (a.t === "SetLang") {
    return { ...s, lang: a.lang };
  }
  switch (a.t) {
    case "Close":
      return {
        t: "Disconnected",
        lang: s.lang,
        me: s.me,
        game: s.t === "GamePlaying" ? { gid: s.gid, clients: s.clients } : null,
      };
    case "SetMe":
      return s.t === "Welcome" || s.t === "Disconnected"
        ? { t: "Welcome", lang: s.lang, me: a.Me }
        : s.t === "HowTo"
        ? { t: "HowTo", lang: s.lang, me: a.Me }
        : s.t === "LangChoosing"
        ? { t: "LangChoosing", lang: s.lang, me: a.Me }
        : { t: "Fatal", lang: s.lang, s, a };
    case "GoLobbies":
      return s.t === "Disbanded" || s.t === "GameEnded"
        ? { ...s, t: "LobbyChoosing", lang: s.lang }
        : s.t === "LobbyWaiting"
        ? { ...s, didLeave: true }
        : { t: "Fatal", lang: s.lang, s, a };
    case "GoWelcome":
      return s.t === "HowTo" || s.t === "NameChoosing" || s.t === "LangChoosing"
        ? { t: "Welcome", lang: s.lang, me: s.me }
        : { t: "Fatal", lang: s.lang, s, a };
    case "GoNameChoose":
      return s.t === "Welcome" && s.me !== 0
        ? { t: "NameChoosing", lang: s.lang, me: s.me, valid: true }
        : { t: "Fatal", lang: s.lang, s, a };
    case "GoLangChoose":
      return s.t === "Welcome"
        ? { t: "LangChoosing", lang: s.lang, me: s.me }
        : { t: "Fatal", lang: s.lang, s, a };
    case "GoHowTo":
      return s.t === "Welcome"
        ? { t: "HowTo", lang: s.lang, me: s.me }
        : { t: "Fatal", lang: s.lang, s, a };
    case "NameReject":
      return s.t === "NameChoosing"
        ? { ...s, valid: false }
        : { t: "Fatal", lang: s.lang, s, a };
    case "LobbyChoices":
      return s.t === "LobbyChoosing" ||
        s.t === "NameChoosing" ||
        (s.t === "LobbyWaiting" && s.didLeave)
        ? { t: "LobbyChoosing", lang: s.lang, me: s.me, lobbies: a.Lobbies }
        : s.me === 0
        ? { t: "Fatal", lang: s.lang, s, a }
        : { t: "Disbanded", lang: s.lang, me: s.me, lobbies: a.Lobbies };
    case "CurrentLobby":
      return s.t === "LobbyChoosing" || s.t === "LobbyWaiting"
        ? {
            t: "LobbyWaiting",
            lang: s.lang,
            me: s.me,
            gid: a.GID,
            clients: a.Clients,
            leader: a.Leader,
            didLeave: false,
          }
        : { t: "Fatal", lang: s.lang, s, a };
    case "CurrentGame":
      return s.t === "LobbyWaiting" || s.t === "GamePlaying"
        ? mkGamePlaying(s.lang, s.me, s.gid, s.clients, a)
        : s.t === "Disconnected" && s.game !== null
        ? mkGamePlaying(s.lang, s.me, s.game.gid, s.game.clients, a)
        : { t: "Fatal", lang: s.lang, s, a };
    case "EndGame":
      return s.t === "GamePlaying"
        ? {
            t: "GameEnded",
            lang: s.lang,
            me: s.me,
            resPts: a.ResPts,
            spyPts: a.SpyPts,
            lobbies: a.Lobbies,
          }
        : { t: "Fatal", lang: s.lang, s, a };
  }
};
