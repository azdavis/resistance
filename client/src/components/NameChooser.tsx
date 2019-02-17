import React, { useEffect } from "react";
import useUncontrolledInput from "../hooks/useUncontrolledInput";

type Props = {
  ws: WebSocket | null;
};

const NameChooser = ({ ws }: Props): JSX.Element => {
  const name = useUncontrolledInput();
  useEffect(() => name.ref.current!.focus(), []);
  return (
    <>
      <h1>resistance™</h1>
      <form
        onSubmit={e => {
          e.preventDefault();
          ws && ws.send(name.get());
        }}
      >
        <label htmlFor="name">player name</label>
        <input type="text" id="name" ref={name.ref} />
        <input type="submit" value="submit" disabled={ws === null} />
      </form>
    </>
  );
};

export default NameChooser;
