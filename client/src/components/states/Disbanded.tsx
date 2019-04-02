import React from "react";
import { Lang, D } from "../../types";
import Button from "../basic/Button";

type Props = {
  lang: Lang;
  d: D;
};

const text = {
  title: {
    en: <h1>Disbanded</h1>,
  },
  body: {
    en: <p>The game or lobby you were in was disbanded.</p>,
  },
  leave: {
    en: "Leave",
  },
};

export default ({ lang, d }: Props) => (
  <div className="Disbanded">
    {text.title[lang]}
    {text.body[lang]}
    <Button value={text.leave[lang]} onClick={() => d({ t: "GoLobbies" })} />
  </div>
);
