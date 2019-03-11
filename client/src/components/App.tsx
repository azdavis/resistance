import React, { useReducer, useEffect, useState } from "react";
import { Send } from "../types";
import { reducer, init } from "../state";
import Fatal from "./Fatal";
import Disbanded from "./Disbanded";
import HowTo from "./HowTo";
import NameChooser from "./NameChooser";
import LobbyChooser from "./LobbyChooser";
import LobbyWaiter from "./LobbyWaiter";
import MemberChooser from "./MemberChooser";
import MemberWaiter from "./MemberWaiter";

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
    case "MemberChoosing":
      return s.me === s.captain ? (
        <MemberChooser send={send!} {...s} />
      ) : (
        <MemberWaiter
          captain={s.clients.find(x => x.CID === s.captain)!.Name}
          isSpy={s.isSpy}
          numMembers={s.numMembers}
        />
      );
  }
};
