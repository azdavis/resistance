import React from "react";
import { Translation, D } from "../../etc";
import Button from "../basic/Button";

type Props = {
  t: Translation;
  d: D;
};

export default ({ t, d }: Props) => {
  return (
    <div className="Disbanded">
      <h1>{t.Disbanded.title}</h1>
      <p>{t.Disbanded.body}</p>
      <Button value={t.leave} onClick={() => d({ t: "GoLobbies" })} />
    </div>
  );
};
