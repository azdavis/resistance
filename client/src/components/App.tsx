import React, { useReducer, useEffect, useState } from "react";
import { Send } from "../types";
import { reducer, init } from "../state";
import Fatal from "./Fatal";
import Disbanded from "./Disbanded";
import HowTo from "./HowTo";
import NameChooser from "./NameChooser";
import LobbyChooser from "./LobbyChooser";
import LobbyWaiter from "./LobbyWaiter";
import RoleViewer from "./RoleViewer";
import MemberChooser from "./MemberChooser";
import MemberWaiter from "./MemberWaiter";
import MemberVoter from "./MemberVoter";
import MissionVoter from "./MissionVoter";
import MissionWaiter from "./MissionWaiter";
import MissionResultViewer from "./MissionResultViewer";

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
    case "Welcome":
      throw "TODO";
    case "HowTo":
      return <HowTo d={d} />;
    case "NameChoosing":
      return <NameChooser d={d} send={send} {...s} />;
    case "LobbyChoosing":
      return <LobbyChooser send={send!} {...s} />;
    case "LobbyWaiting":
      return <LobbyWaiter d={d} send={send!} {...s} />;
    case "RoleViewing":
      return <RoleViewer d={d} isSpy={s.isSpy} />;
    case "MemberChoosing":
      return s.me === s.captain ? (
        <MemberChooser send={send!} {...s} />
      ) : (
        <MemberWaiter
          clients={s.clients}
          captain={s.captain}
          members={s.members}
        />
      );
    case "MemberVoting":
      return <MemberVoter send={send!} {...s} />;
    case "MissionVoting":
      return s.canVote ? <MissionVoter send={send!} /> : <MissionWaiter />;
    case "MissionResultViewing":
      return <MissionResultViewer d={d} send={send!} {...s} />;
  }
};
