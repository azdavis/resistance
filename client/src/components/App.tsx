import React, { useReducer, useEffect, useState } from "react";
import { State, Send } from "../types";
import Closed from "./Closed";
import NameChooser from "./NameChooser";
import LobbyChooser from "./LobbyChooser";
import LobbyWaiter from "./LobbyWaiter";

const reducer = (s: State, a: State): State => a;
const init: State = { T: "NameChoosing", Valid: true };

export default (): JSX.Element => {
  const [s, d] = useReducer(reducer, init);
  const [send, setSend] = useState<Send | null>(null);
  useEffect(() => {
    const ws = new WebSocket("ws://localhost:8080/ws");
    const newSend: Send = ({ T, ...P }) => {
      ws.send(JSON.stringify({ T, P }));
    };
    ws.onopen = () => setSend(() => newSend);
    ws.onmessage = e => {
      const { T, P } = JSON.parse(e.data);
      d({ T, ...P });
    };
    ws.onclose = () => d({ T: "Closed" });
    return ws.close.bind(ws);
  }, []);
  switch (s.T) {
    case "Closed":
      return <Closed />;
    case "NameChoosing":
      return <NameChooser send={send} valid={s.Valid} />;
    case "LobbyChoosing":
      return <LobbyChooser send={send!} lobbies={s.Lobbies} />;
    case "LobbyWaiting":
      return (
        <LobbyWaiter
          send={send!}
          self={s.Self}
          leader={s.Leader}
          clients={s.Clients}
        />
      );
  }
};
