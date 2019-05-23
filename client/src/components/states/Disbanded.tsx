import React from "react";
import t8ns from "../../translations";
import { Lang, D } from "../../etc";
import Button from "../basic/Button";

type Props = {
  lang: Lang;
  d: D;
};

export default ({ lang, d }: Props) => {
  const { Disbanded: t8n, leave } = t8ns[lang];
  return (
    <div className="Disbanded">
      <h1>{t8n.title}</h1>
      <p>{t8n.body}</p>
      <Button value={leave} onClick={() => d({ t: "GoLobbies" })} />
    </div>
  );
};
