import React, { useRef, useEffect } from "react";
import t8ns from "../../translations";
import { Lang, D, S } from "../../etc";
import Button from "../basic/Button";
import TextInput from "../basic/TextInput";

type Props = {
  lang: Lang;
  d: D;
  send: S;
  valid: boolean;
};

export default ({ lang, d, send, valid }: Props) => {
  const nameRef = useRef<HTMLInputElement>(null);
  const { NameChoosing: t8n, submit, back } = t8ns[lang];
  useEffect(() => nameRef.current!.focus(), []);
  return (
    <div className="NameChoosing">
      {t8n.title}
      <form
        onSubmit={e => {
          e.preventDefault();
          send({ t: "NameChoose", Name: nameRef.current!.value });
        }}
      >
        <TextInput ref={nameRef} />
        {valid ? null : t8n.invalid}
        <Button type="submit" value={submit} />
      </form>
      <Button value={back} onClick={() => d({ t: "GoWelcome" })} />
    </div>
  );
};
