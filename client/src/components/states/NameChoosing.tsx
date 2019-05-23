import React, { useRef, useEffect } from "react";
import { Translation, D, S } from "../../etc";
import Button from "../basic/Button";
import TextInput from "../basic/TextInput";

type Props = {
  t: Translation;
  d: D;
  send: S;
  valid: boolean;
};

export default ({ t, d, send, valid }: Props) => {
  const nameRef = useRef<HTMLInputElement>(null);
  const { NameChoosing: NC, submit, back } = t;
  useEffect(() => nameRef.current!.focus(), []);
  return (
    <div className="NameChoosing">
      <h1>{NC.title}</h1>
      <form
        onSubmit={e => {
          e.preventDefault();
          send({ t: "NameChoose", Name: nameRef.current!.value });
        }}
      >
        <TextInput ref={nameRef} />
        {valid ? null : NC.invalid}
        <Button type="submit" value={submit} />
      </form>
      <Button value={back} onClick={() => d({ t: "GoWelcome" })} />
    </div>
  );
};
