import React, { useReducer, useState, useEffect } from "react";
import { S, Translation } from "../etc";
import { reducer, init } from "../state";
import Storage from "../storage";
import useTriggerEffect from "../hooks/useTriggerEffect";
import Invalid from "./states/Invalid";
import SetLangFail from "./states/SetLangFail";
import Disconnected from "./states/Disconnected";
import Disbanded from "./states/Disbanded";
import Welcome from "./states/Welcome";
import HowTo from "./states/HowTo";
import LangChoosing from "./states/LangChoosing";
import NameChoosing from "./states/NameChoosing";
import LobbyChoosing from "./states/LobbyChoosing";
import LobbyWaiting from "./states/LobbyWaiting";
import GamePlaying from "./states/GamePlaying";
import GameEnded from "./states/GameEnded";

const defaultLang = Storage.getLang() || "en";

export default (): JSX.Element | null => {
  const [s, d] = useReducer(reducer, init);
  const [send, setSend] = useState<S | null>(null);
  const reconnect = useTriggerEffect(() => {
    const ws = new WebSocket("ws://localhost:8080");
    const newSend: S = ({ t, ...P }) => {
      ws.send(JSON.stringify({ T: t, P }));
    };
    ws.onopen = () => {
      newSend(
        s.t === "Disconnected" && s.me !== 0 && s.game !== null
          ? { t: "Reconnect", Me: s.me, GID: s.game.gid }
          : { t: "Connect" },
      );
      setSend(() => newSend);
    };
    ws.onmessage = e => {
      const { T, P } = JSON.parse(e.data);
      d({ t: T, ...P });
    };
    ws.onclose = () => d({ t: "Close" });
    return ws.close.bind(ws);
  }, []);
  const [lang, setLang] = useState(defaultLang);
  const [t, setTrans] = useState<Translation | null>(null);
  useEffect(() => {
    import(`../translations/${lang}`)
      .then(res => {
        document.documentElement.lang = lang;
        Storage.setLang(lang);
        setTrans(res.default);
      })
      .catch(e => d({ t: "GoSetLangFail", msg: String(e.message || e) }));
  }, [lang]);
  if (s.t === "SetLangFail") {
    return <SetLangFail msg={s.msg} />;
  }
  if (t === null) {
    return null;
  }
  switch (s.t) {
    case "Invalid":
      return <Invalid t={t} s={s.s} a={s.a} />;
    case "Disconnected":
      return <Disconnected t={t} reconnect={reconnect} />;
    case "Disbanded":
      return <Disbanded t={t} d={d} />;
    case "Welcome":
      return <Welcome t={t} d={d} loading={send === null || s.me === 0} />;
    case "HowTo":
      return <HowTo t={t} d={d} />;
    case "LangChoosing":
      return <LangChoosing t={t} setLang={setLang} d={d} />;
    case "NameChoosing":
      return <NameChoosing t={t} d={d} send={send!} valid={s.valid} />;
    case "LobbyChoosing":
      return <LobbyChoosing t={t} send={send!} lobbies={s.lobbies} />;
    case "LobbyWaiting":
      return (
        <LobbyWaiting
          t={t}
          d={d}
          send={send!}
          me={s.me}
          leader={s.leader}
          clients={s.clients}
        />
      );
    case "GamePlaying":
      return (
        <GamePlaying
          t={t}
          send={send!}
          me={s.me}
          clients={s.clients}
          isSpy={s.isSpy}
          resPts={s.resPts}
          spyPts={s.spyPts}
          captain={s.captain}
          members={s.members}
          active={s.active}
        />
      );
    case "GameEnded":
      return <GameEnded t={t} d={d} resPts={s.resPts} spyPts={s.spyPts} />;
  }
};
