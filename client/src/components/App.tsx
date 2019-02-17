import React, { useReducer, useEffect } from "react";
import useWebSocket from "../hooks/useWebSocket";
import { reducer, init } from "../state";
import Closed from "./Closed";
import NameChooser from "./NameChooser";

const App = (): JSX.Element => {
  const [s, d] = useReducer(reducer, init);
  const ws = useWebSocket("wss://echo.websocket.org");
  useEffect(() => {
    if (ws === null) {
      return;
    }
    ws.onmessage = e => {
      try {
        d(JSON.parse(e.data));
      } catch (err) {
        ws.close();
      }
    };
    ws.onclose = () => d({ t: "close" });
  }, [ws]);
  switch (s.t) {
    case "closed":
      return <Closed />;
    case "nameChoosing":
      return <NameChooser ws={ws} />;
  }
};

export default App;
