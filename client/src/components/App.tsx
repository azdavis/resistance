import React, { useReducer, useState, useEffect } from "react";
import { S } from "../etc";
import { reducer, init } from "../state";
import { setLang } from "../storage";
import useTriggerEffect from "../hooks/useTriggerEffect";
import Fatal from "./states/Fatal";
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

export default (): JSX.Element => {
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
  useEffect(() => {
    document.documentElement.lang = s.lang;
    setLang(s.lang);
  }, [s.lang]);
  // eslint-disable-next-line
  switch (s.t) {
    case "Fatal":
      return <Fatal {...s} />;
    case "Disconnected":
      return <Disconnected lang={s.lang} reconnect={reconnect} />;
    case "Disbanded":
      return <Disbanded lang={s.lang} d={d} />;
    case "Welcome":
      return (
        <Welcome lang={s.lang} d={d} loading={send === null || s.me === 0} />
      );
    case "HowTo":
      return <HowTo lang={s.lang} d={d} />;
    case "LangChoosing":
      return <LangChoosing d={d} lang={s.lang} />;
    case "NameChoosing":
      return <NameChoosing d={d} send={send!} {...s} />;
    case "LobbyChoosing":
      return <LobbyChoosing send={send!} {...s} />;
    case "LobbyWaiting":
      return <LobbyWaiting d={d} send={send!} {...s} />;
    case "GamePlaying":
      return <GamePlaying send={send!} {...s} />;
    case "GameEnded":
      return <GameEnded d={d} {...s} />;
  }
};
