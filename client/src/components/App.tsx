import React, { useReducer, useEffect, useState } from "react";
import { State, Msg, Send } from "../types";
import Closed from "./Closed";
import NameChooser from "./NameChooser";
import RoomChooser from "./RoomChooser";

// a might actually be a "partial state", which is why we merge it with s, but a
// takes precedence. a: Partial<State> didn't quite seem to work.
const reducer = (s: State, a: State): State => ({ ...s, ...a });
const init: State = { T: "nameChoosing" };

export default (): JSX.Element => {
  const [s, d] = useReducer(reducer, init);
  const [send, setSend] = useState<Send | null>(null);
  useEffect(() => {
    const ws = new WebSocket("ws://localhost:8080");
    ws.onopen = () => setSend(() => (m: Msg) => ws.send(JSON.stringify(m)));
    ws.onmessage = e => d(JSON.parse(e.data));
    ws.onclose = () => d({ T: "closed" });
    return ws.close.bind(ws);
  }, []);
  switch (s.T) {
    case "closed":
      return <Closed />;
    case "nameChoosing":
      return <NameChooser send={send} />;
    case "roomChoosing":
      return <RoomChooser send={send!} />;
  }
};
