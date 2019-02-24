import React, { useReducer, useEffect, useState } from "react";
import { State, Msg, Send } from "../types";
import Closed from "./Closed";
import NameChooser from "./NameChooser";
import PartyChooser from "./PartyChooser";

// a might actually be a "partial state", which is why we merge it with s, but a
// takes precedence. a: Partial<State> didn't quite seem to work.
const reducer = (s: State, a: State): State => ({ ...s, ...a });
const init: State = { T: "NameChoosing" };

export default (): JSX.Element => {
  const [s, d] = useReducer(reducer, init);
  const [send, setSend] = useState<Send | null>(null);
  useEffect(() => {
    const ws = new WebSocket("ws://localhost:8080");
    const send = ({ T, ...P }: Msg) => {
      ws.send(JSON.stringify({ T, P }));
    };
    ws.onopen = () => setSend(() => send);
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
  }
};
