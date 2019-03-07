import React, { useReducer, useEffect, useState } from "react";
import { State, Action, Send } from "../types";
import Invalid from "./Invalid";
import Closed from "./Closed";
import NameChooser from "./NameChooser";
import LobbyChooser from "./LobbyChooser";
import LobbyWaiter from "./LobbyWaiter";

const reducer = (s: State, a: Action): State => {
  switch (a.t) {
    case "Close":
      return { t: "Closed" };
    case "RejectName":
      return { t: "NameChoosing", valid: false };
    case "LobbyChoices":
      return { t: "LobbyChoosing", lobbies: a.Lobbies };
    case "CurrentLobby":
      return {
        t: "LobbyWaiting",
        self: a.Self,
        leader: a.Leader,
        clients: a.Clients,
      };
  }
};
const init: State = { t: "NameChoosing", valid: true };

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
    case "Invalid":
      return <Invalid {...s} />;
    case "Closed":
      return <Closed />;
    case "NameChoosing":
      return <NameChooser send={send} {...s} />;
    case "LobbyChoosing":
      return <LobbyChooser send={send!} {...s} />;
    case "LobbyWaiting":
      return <LobbyWaiter send={send!} {...s} />;
  }
};
