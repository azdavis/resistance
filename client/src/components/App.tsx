import React, { useReducer, useEffect, useState } from "react";
import { State, Action, Send } from "../types";
import Fatal from "./Fatal";
import NameChooser from "./NameChooser";
import LobbyChooser from "./LobbyChooser";
import LobbyWaiter from "./LobbyWaiter";
import MissionMemberChooser from "./MissionMemberChooser";
import MissionMemberWaiter from "./MissionMemberWaiter";

const reducer = (s: State, a: Action): State => {
  switch (a.t) {
    case "Close":
      return { t: "Fatal", s, a };
    case "RejectName":
      return { t: "NameChoosing", valid: false };
    case "LobbyChoices":
      return { t: "LobbyChoosing", lobbies: a.Lobbies };
    case "CurrentLobby":
      return {
        t: "LobbyWaiting",
        me: a.Me,
        leader: a.Leader,
        clients: a.Clients,
        isSpy: false,
      };
    case "SetIsSpy":
      return s.t === "LobbyWaiting"
        ? { ...s, isSpy: a.IsSpy }
        : { t: "Fatal", s, a };
    case "NewMission":
      return s.t === "LobbyWaiting"
        ? { ...s, t: "MissionMemberChoosing", captain: a.Captain }
        : { t: "Fatal", s, a };
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
    case "Fatal":
      return <Fatal {...s} />;
    case "NameChoosing":
      return <NameChooser send={send} {...s} />;
    case "LobbyChoosing":
      return <LobbyChooser send={send!} {...s} />;
    case "LobbyWaiting":
      return <LobbyWaiter send={send!} {...s} />;
    case "MissionMemberChoosing":
      return s.me === s.captain ? (
        <MissionMemberChooser send={send!} {...s} />
      ) : (
        <MissionMemberWaiter
          captain={s.clients.find(x => x.CID === s.captain)!.Name}
          isSpy={s.isSpy}
        />
      );
  }
};
