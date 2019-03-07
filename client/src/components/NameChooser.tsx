import React, { useRef, useEffect } from "react";
import { Send } from "../types";
import Button from "./Button";
import "./NameChooser.css";

type Props = {
  send: Send | null;
  valid: boolean;
};

export default ({ send, valid }: Props): JSX.Element => {
  const nameRef = useRef<HTMLInputElement>(null);
  useEffect(() => nameRef.current!.focus(), []);
  return (
    <div className="NameChooser">
      <h1>Resistance™</h1>
      <form
        onSubmit={e => {
          e.preventDefault();
          if (send === null) {
            return;
          }
          send({ t: "NameChoose", Name: nameRef.current!.value });
        }}
      >
        <label htmlFor="name">Player name{!valid && <b> invalid</b>}</label>
        <input type="text" id="name" autoCorrect="off" ref={nameRef} />
        <Button value="Submit" submit disabled={send === null} />
      </form>
    </div>
  );
};
