import React, { useReducer, useEffect, useState } from "react";
import { State, ToClient, Send } from "../types";
import Closed from "./Closed";
import NameChooser from "./NameChooser";
import PartyChooser from "./PartyChooser";
import PartyDisbanded from "./PartyDisbanded";
import PartyWaiter from "./PartyWaiter";

const reducer = (s: State, tc: ToClient): State => tc;
const init: State = { T: "NameChoosing" };

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
      return <NameChooser send={send} />;
    case "PartyChoosing":
      return <PartyChooser send={send!} />;
    case "PartyDisbanded":
      return <PartyDisbanded d={d} />;
    case "PartyWaiting":
      return <PartyWaiter send={send!} />;
  }
};
