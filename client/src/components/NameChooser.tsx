import React, { useRef, useEffect } from "react";
import { Send } from "../logic/send";

type Props = {
  send: Send | null;
};

const NameChooser = ({ send }: Props): JSX.Element => {
  const nameRef = useRef<HTMLInputElement>(null);
  useEffect(() => nameRef.current!.focus(), []);
  return (
    <>
      <h1>resistance™</h1>
      <form
        onSubmit={e => {
          e.preventDefault();
          send!({ t: "nameChoose", name: nameRef.current!.value });
        }}
      >
        <label htmlFor="name">player name</label>
        <input type="text" id="name" ref={nameRef} />
        <input type="submit" value="submit" disabled={send === null} />
      </form>
    </>
  );
};

export default NameChooser;
