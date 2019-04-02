import React, { useRef, useEffect } from "react";
import { Lang, D, Send } from "../../types";
import Button from "../basic/Button";
import TextInput from "../basic/TextInput";

type Props = {
  lang: Lang;
  d: D;
  send: Send;
  valid: boolean;
};

const text = {
  title: {
    en: <h1>Player name</h1>,
  },
  invalid: {
    en: "Invalid",
  },
  submit: {
    en: "Submit",
  },
  back: {
    en: "Back",
  },
};

export default ({ lang, d, send, valid }: Props) => {
  const nameRef = useRef<HTMLInputElement>(null);
  useEffect(() => nameRef.current!.focus(), []);
  return (
    <div className="NameChoosing">
      {text.title[lang]}
      <form
        onSubmit={e => {
          e.preventDefault();
          send({ t: "NameChoose", Name: nameRef.current!.value });
        }}
      >
        <TextInput ref={nameRef} />
        {valid ? null : text.invalid[lang]}
        <Button type="submit" value={text.submit[lang]} />
      </form>
      <Button value={text.back[lang]} onClick={() => d({ t: "GoWelcome" })} />
    </div>
  );
};
