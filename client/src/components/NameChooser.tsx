import React, { useRef, useEffect } from "react";
import { Send } from "../types";
import Button from "./Button";

type Props = {
  send: Send | null;
};

export default ({ send }: Props): JSX.Element => {
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
          send({ T: "NameChoose", Name: nameRef.current!.value });
        }}
      >
        <label htmlFor="name">Player name</label>
        <input type="text" id="name" autoCorrect="off" ref={nameRef} />
        <Button value="submit" submit disabled={send === null} />
      </form>
    </div>
  );
};
