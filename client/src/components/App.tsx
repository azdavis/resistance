import React, { useReducer, useEffect, useState } from "react";
import { Send } from "../types";
import { reducer, init } from "../state";
import Fatal from "./states/Fatal";
import Disbanded from "./states/Disbanded";
import HowTo from "./states/HowTo";
import NameChooser from "./states/NameChooser";
import LobbyChooser from "./states/LobbyChooser";
import LobbyWaiter from "./states/LobbyWaiter";
import RoleViewer from "./states/RoleViewer";
import MemberChooser from "./states/MemberChooser";
import MemberWaiter from "./states/MemberWaiter";
import MemberVoter from "./states/MemberVoter";
import MissionVoter from "./states/MissionVoter";
import MissionWaiter from "./states/MissionWaiter";
import MissionResultViewer from "./states/MissionResultViewer";

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
    case "HowTo":
      return <HowTo d={d} />;
    case "NameChoosing":
      return <NameChooser d={d} send={send} {...s} />;
    case "LobbyChoosing":
      return <LobbyChooser send={send!} {...s} />;
    case "LobbyWaiting":
      return <LobbyWaiter send={send!} d={d} {...s} />;
    case "RoleViewing":
      return <RoleViewer d={d} isSpy={s.isSpy} wait={s.mission === null} />;
    case "MemberChoosing":
      return s.me === s.captain ? (
        <MemberChooser send={send!} {...s} />
      ) : (
        <MemberWaiter
          captain={s.clients.find(x => x.CID === s.captain)!.Name}
          numMembers={s.numMembers}
        />
      );
    case "MemberVoting":
      return <MemberVoter send={send!} {...s} />;
    case "MissionVoting":
      return s.canVote ? <MissionVoter send={send!} /> : <MissionWaiter />;
    case "MissionResultViewing":
      return <MissionResultViewer send={send!} d={d} {...s} />;
  }
};
