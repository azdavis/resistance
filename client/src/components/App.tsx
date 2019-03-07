import React, { useReducer, useEffect, useState } from "react";
import { State, Action, Send } from "../types";
import Invalid from "./Invalid";
import NameChooser from "./NameChooser";
import LobbyChooser from "./LobbyChooser";
import LobbyWaiter from "./LobbyWaiter";

const reducer = (s: State, a: Action): State => {
  switch (a.t) {
    case "Close":
      return { t: "Invalid", s, a };
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
        isSpy: false,
      };
    case "SetIsSpy":
      return s.t === "LobbyWaiting"
        ? { ...s, isSpy: a.IsSpy }
        : { t: "Invalid", s, a };
    case "NewMission":
      return { t: "Invalid", s, a };
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
    case "NameChoosing":
      return <NameChooser send={send} {...s} />;
    case "LobbyChoosing":
      return <LobbyChooser send={send!} {...s} />;
    case "LobbyWaiting":
      return <LobbyWaiter send={send!} {...s} />;
  }
};
