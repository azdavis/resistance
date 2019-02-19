import React, { useReducer, useEffect, useState } from "react";
import { State, Msg, Send } from "../types";
import Closed from "./Closed";
import NameChooser from "./NameChooser";
import RoomChooser from "./RoomChooser";

const reducer = (oldS: State, newS: State): State => ({ ...oldS, ...newS });
const init: State = { t: "nameChoosing" };

export default (): JSX.Element => {
  const [s, d] = useReducer(reducer, init);
  const [send, setSend] = useState<Send | null>(null);
  useEffect(() => {
    const ws = new WebSocket("wss://echo.websocket.org");
    ws.onopen = () => setSend(() => (msg: Msg) => ws.send(JSON.stringify(msg)));
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
