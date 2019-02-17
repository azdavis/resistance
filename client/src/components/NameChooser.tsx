import React, { useEffect } from "react";
import { Send } from "../logic/send";
import useUncontrolledInput from "../hooks/useUncontrolledInput";

type Props = {
  send: Send | null;
};

const NameChooser = ({ send }: Props): JSX.Element => {
  const name = useUncontrolledInput();
  useEffect(() => name.ref.current!.focus(), []);
  return (
    <>
      <h1>resistance™</h1>
      <form
        onSubmit={e => {
          e.preventDefault();
          send && send(name.get());
        }}
      >
        <label htmlFor="name">player name</label>
        <input type="text" id="name" ref={name.ref} />
        <input type="submit" value="submit" disabled={send === null} />
      </form>
    </>
  );
};

export default NameChooser;
