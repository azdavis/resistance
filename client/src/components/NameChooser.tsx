import React, { useRef, useEffect } from "react";
import { D, Send } from "../types";
import Button from "./basic/Button";
import "./NameChooser.css";

type Props = {
  d: D;
  send: Send;
  valid: boolean;
};

export default ({ d, send, valid }: Props) => {
  const nameRef = useRef<HTMLInputElement>(null);
  useEffect(() => nameRef.current!.focus(), []);
  return (
    <div className="NameChooser">
      <h1>Player name</h1>
      <form
        onSubmit={e => {
          e.preventDefault();
          send({ t: "NameChoose", Name: nameRef.current!.value });
        }}
      >
        <input type="text" autoCorrect="off" ref={nameRef} />
        {!valid && "Invalid"}
        <Button type="submit" value="Submit" />
      </form>
      <Button value="Back" onClick={() => d({ t: "GoWelcome" })} />
    </div>
  );
};
