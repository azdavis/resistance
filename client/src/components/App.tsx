import React, { useReducer, useEffect, useRef } from "react";
import { reducer, init } from "../state";
import { Send } from "../send";
import Closed from "./Closed";
import NameChooser from "./NameChooser";

const App = (): JSX.Element => {
  const [s, d] = useReducer(reducer, init);
  const sendRef = useRef<Send | null>(null);
  const { current: send } = sendRef;
  useEffect(() => {
    const ws = new WebSocket("wss://echo.websocket.org");
    ws.onopen = () => {
      sendRef.current = msg => ws.send(JSON.stringify(msg));
    };
    ws.onmessage = e => {
      try {
        d(JSON.parse(e.data));
      } catch (err) {
        ws.close();
      }
    };
    ws.onclose = () => d({ t: "close" });
    return ws.close.bind(ws);
  }, []);
  switch (s.t) {
    case "closed":
      return <Closed />;
    case "nameChoosing":
      return <NameChooser send={send} />;
  }
};

export default App;
