import React, { useReducer, useEffect, useState } from "react";
import { Send } from "../types";
import { reducer, init } from "../state";
import Fatal from "./Fatal";
import Disbanded from "./Disbanded";
import NameChooser from "./NameChooser";
import LobbyChooser from "./LobbyChooser";
import LobbyWaiter from "./LobbyWaiter";
import MissionChooser from "./MissionChooser";
import MissionWaiter from "./MissionWaiter";

export default (): JSX.Element => {
  const [s, d] = useReducer(reducer, init);
  const [send, setSend] = useState<Send | null>(null);
  useEffect(() => {
    const ws = new WebSocket("ws://localhost:8080/ws");
    const newSend: Send = ({ t, ...P }) => {
      ws.send(JSON.stringify({ T: t, P }));
    };
    ws.onopen = () => setSend(() => newSend);
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
    case "Disbanded":
      return <Disbanded d={d} />;
    case "NameChoosing":
      return <NameChooser send={send} {...s} />;
    case "LobbyChoosing":
      return <LobbyChooser send={send!} {...s} />;
    case "LobbyWaiting":
      return <LobbyWaiter send={send!} d={d} {...s} />;
    case "MissionChoosing":
      return s.me === s.captain ? (
        <MissionChooser send={send!} {...s} />
      ) : (
        <MissionWaiter
          captain={s.clients.find(x => x.CID === s.captain)!.Name}
          isSpy={s.isSpy}
        />
      );
  }
};
