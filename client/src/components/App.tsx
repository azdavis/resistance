import React, { useReducer, useEffect, useState } from "react";
import { State, ToClient, Send } from "../types";
import Closed from "./Closed";
import NameChooser from "./NameChooser";
import LobbyChooser from "./LobbyChooser";
import LobbyWaiter from "./LobbyWaiter";

const reducer = (s: State, tc: ToClient): State => {
  switch (tc.T) {
    case "Close":
      return { T: "Closed" };
    case "RejectName":
      return { T: "NameChoosing", valid: false };
    case "LobbyChoices":
      return { T: "LobbyChoosing", lobbies: tc.Lobbies };
    case "CurrentLobby":
      return {
        T: "LobbyWaiting",
        self: tc.Self,
        leader: tc.Leader,
        clients: tc.Clients,
      };
  }
};
const init: State = { T: "NameChoosing", valid: true };

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
    ws.onclose = () => d({ T: "Close" });
    return ws.close.bind(ws);
  }, []);
  switch (s.T) {
    case "Closed":
      return <Closed />;
    case "NameChoosing":
      return <NameChooser send={send} valid={s.valid} />;
    case "LobbyChoosing":
      return <LobbyChooser send={send!} lobbies={s.lobbies} />;
    case "LobbyWaiting":
      return (
        <LobbyWaiter
          send={send!}
          self={s.self}
          leader={s.leader}
          clients={s.clients}
        />
      );
  }
};
