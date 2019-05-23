import React from "react";
import t8ns from "../../translations";
import { Lang, D } from "../../etc";
import Button from "../basic/Button";

type Props = {
  lang: Lang;
  d: D;
};

export default ({ lang, d }: Props) => (
  <div className="Disbanded">
    {t8ns[lang].Disbanded.title}
    {t8ns[lang].Disbanded.body}
    <Button value={t8ns[lang].leave} onClick={() => d({ t: "GoLobbies" })} />
  </div>
);
