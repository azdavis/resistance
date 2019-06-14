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
  useEffect(() => nameRef.current!.focus(), []);
  return (
    <div className="NameChoosing">
      <h1>{t.playerName}</h1>
      <Button value={t.back} onClick={() => d({ t: "GoWelcome" })} />
      <form
        onSubmit={e => {
          e.preventDefault();
          send({ t: "NameChoose", Name: nameRef.current!.value });
        }}
      >
        <TextInput ref={nameRef} />
        {valid ? null : t.invalid}
        <Button type="submit" value={t.submit} />
      </form>
    </div>
  );
};
