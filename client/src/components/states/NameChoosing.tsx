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
  useEffect(() => nameRef.current!.focus(), []);
  return (
    <div className="NameChoosing">
      {t8ns[lang].NameChoosing.title}
      <form
        onSubmit={e => {
          e.preventDefault();
          send({ t: "NameChoose", Name: nameRef.current!.value });
        }}
      >
        <TextInput ref={nameRef} />
        {valid ? null : t8ns[lang].NameChoosing.invalid}
        <Button type="submit" value={t8ns[lang].submit} />
      </form>
      <Button value={t8ns[lang].back} onClick={() => d({ t: "GoWelcome" })} />
    </div>
  );
};
