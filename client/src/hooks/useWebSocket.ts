import { useState, useEffect } from "react";

const useWebSocket = (url: string): WebSocket | null => {
  const [ws, setWs] = useState<WebSocket | null>(null);
  useEffect(() => {
    const w = new WebSocket(url);
    w.onopen = () => setWs(w);
    return w.close.bind(w);
  }, [url]);
  return ws;
};

export default useWebSocket;
