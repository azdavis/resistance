import React, { useRef, useEffect } from "react";
import { D, Send } from "../../types";
import Button from "../Button";
import "./NameChooser.css";

type Props = {
  d: D;
  send: Send | null;
  valid: boolean;
};

export default ({ d, send, valid }: Props) => {
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
        <label>
          Player name{!valid && <b> invalid</b>}
          <input type="text" autoCorrect="off" ref={nameRef} />
        </label>
        <Button type="submit" value="Submit" disabled={send === null} />
      </form>
      <Button value="How to play" onClick={() => d({ t: "GoHowTo" })} />
    </div>
  );
};
