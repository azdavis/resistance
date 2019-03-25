import React, { useRef, useEffect } from "react";
import { D, Send } from "../../types";
import Button from "../basic/Button";
import TextInput from "../basic/TextInput";

type Props = {
  d: D;
  send: Send;
  valid: boolean;
};

export default ({ d, send, valid }: Props) => {
  const nameRef = useRef<HTMLInputElement>(null);
  useEffect(() => nameRef.current!.focus(), []);
  return (
    <div className="NameChoosing">
      <h1>Player name</h1>
      <form
        onSubmit={e => {
          e.preventDefault();
          send({ t: "NameChoose", Name: nameRef.current!.value });
        }}
      >
        <TextInput ref={nameRef} />
        {valid ? null : "Invalid"}
        <Button type="submit" value="Submit" />
      </form>
      <Button value="Back" onClick={() => d({ t: "GoWelcome" })} />
    </div>
  );
};
