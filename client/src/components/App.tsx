import React, { useReducer, useEffect, useState } from "react";
import { Send } from "../types";
import { reducer, init } from "../state";
import useTriggerableEffect from "../hooks/useTriggerableEffect";
import Fatal from "./states/Fatal";
import Disconnected from "./states/Disconnected";
import Disbanded from "./states/Disbanded";
import Welcome from "./states/Welcome";
import HowTo from "./states/HowTo";
import NameChoosing from "./states/NameChoosing";
import LobbyChoosing from "./states/LobbyChoosing";
import LobbyWaiting from "./states/LobbyWaiting";
import GamePlaying from "./states/GamePlaying";
import GameEnded from "./states/GameEnded";

export default (): JSX.Element => {
  const [s, d] = useReducer(reducer, init);
  const [send, setSend] = useState<Send | null>(null);
  const reconnect = useTriggerableEffect(() => {
    const ws = new WebSocket("ws://localhost:8080/ws");
    const newSend: Send = ({ t, ...P }) => {
      ws.send(JSON.stringify({ T: t, P }));
    };
    ws.onopen = () => {
      newSend({ t: "Connect" });
      setSend(() => newSend);
    };
    ws.onmessage = e => {
      const { T, P } = JSON.parse(e.data);
      d({ t: T, ...P });
    };
    ws.onclose = () => d({ t: "Close" });
    return ws.close.bind(ws);
  }, []);
  switch (s.t) {
    case "Fatal":
      return <Fatal {...s} />;
    case "Disconnected":
      return <Disconnected reconnect={reconnect} />;
    case "Disbanded":
      return <Disbanded d={d} />;
    case "Welcome":
      return <Welcome d={d} loading={send === null || s.me === 0} />;
    case "HowTo":
      return <HowTo d={d} />;
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
