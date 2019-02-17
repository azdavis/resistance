import React, { useReducer, useEffect, useRef } from "react";
import { State, Send } from "../types";
import Closed from "./Closed";
import NameChooser from "./NameChooser";
import RoomChooser from "./RoomChooser";

const reducer = (a: State, b: State): State => ({ ...a, ...b });

const App = (): JSX.Element => {
  const [s, d] = useReducer(reducer, { t: "nameChoosing" });
  const sendRef = useRef<Send | null>(null);
  const { current: send } = sendRef;
  useEffect(() => {
    const ws = new WebSocket("wss://echo.websocket.org");
    ws.onopen = () => {
      sendRef.current = msg => ws.send(JSON.stringify(msg));
    };
    ws.onmessage = e => d(JSON.parse(e.data));
    ws.onclose = () => d({ t: "closed" });
    return ws.close.bind(ws);
  }, []);
  switch (s.t) {
    case "closed":
      return <Closed />;
    case "nameChoosing":
      return <NameChooser send={send} />;
    case "roomChoosing":
      return <RoomChooser send={send!} />;
  }
};

export default App;
