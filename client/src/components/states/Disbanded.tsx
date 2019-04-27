import React from "react";
import { Lang, D } from "../../etc";
import { leave } from "../../text";
import Button from "../basic/Button";

type Props = {
  lang: Lang;
  d: D;
};

const text = {
  title: {
    en: <h1>Disbanded</h1>,
    ja: <h1>解散</h1>,
  },
  body: {
    en: <p>The game or lobby you were in was disbanded.</p>,
    ja: <p>あなたのいたゲームまたはロビーは解散された。</p>,
  },
};

export default ({ lang, d }: Props) => (
  <div className="Disbanded">
    {text.title[lang]}
    {text.body[lang]}
    <Button value={leave[lang]} onClick={() => d({ t: "GoLobbies" })} />
  </div>
);
